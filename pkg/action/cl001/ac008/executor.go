package ac008

import (
	"context"
	"fmt"
	"time"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/backoff"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	valuemodifierpath "github.com/giantswarm/valuemodifier/path"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

const (
	kubeSystemNamespace = "kube-system"

	kiamServerLabelSelector = "app=kiam,component=kiam-server"
	kiamAgentLabelSelector  = "app=kiam,component=kiam-agent"

	masterNodeLabelSelector = "kubernetes.io/role=master"
	workerNodeLabelSelector = "kubernetes.io/role=worker"

	checkJobRetryLimit = 15
	checkJobMaxWait    = time.Second * 10

	draughtsmanNamespace                  = "draughtsman"
	draughtsmanConfigMapName              = "draughtsman-values-configmap"
	draughtsmanConfigMapDataKey           = "values"
	draughtsmanConfigMapDockerRegistryKey = "Installation.V1.Registry.Domain"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: e.logger,

			KubeConfig: env.KubeConfig(),
		}

		cpClients, err = client.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var tcClients k8sclient.Interface
	{
		c := client.TenantClusterConfig{
			ControlPlane: cpClients,
			Logger:       e.logger,

			Scope: e.scope,
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	// awsRegion is necessary to execute aws-cli call
	var awsRegion string
	{
		var cr infrastructurev1alpha2.AWSCluster
		{
			var list infrastructurev1alpha2.AWSClusterList
			err := cpClients.CtrlClient().List(
				ctx,
				&list,
				k8sruntimeclient.InNamespace(cr.GetNamespace()),
				k8sruntimeclient.MatchingLabels{label.Cluster: e.tenantCluster},
			)
			if err != nil {
				return microerror.Mask(err)
			}

			if len(list.Items) == 0 {
				return microerror.Mask(notFoundError)
			}
			if len(list.Items) > 1 {
				return microerror.Mask(tooManyCRsError)
			}

			cr = list.Items[0]
		}

		awsRegion = cr.Spec.Provider.Region
	}

	// dockerRegistry is needed in order to spawn pod with proper docker image that will execute aws-cli call
	var dockerRegistry string
	{
		dockerRegistry, err = e.fetchDockerRegistry(ctx, cpClients.K8sClient())
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = e.checkTLSCerts(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	err = e.checkKiamPods(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	err = e.testAWSApiCalls(ctx, tcClients.K8sClient(), awsRegion, dockerRegistry)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// checkTLSCerts Ensures that kiam  tls certs are created.
var kiamTlSCertSecretNames = []string{"kiam-agent-tls", "kiam-ca-tls", "kiam-server-tls"}

func (e *Executor) checkTLSCerts(ctx context.Context, tcClient kubernetes.Interface) error {
	for _, secret := range kiamTlSCertSecretNames {
		_, err := tcClient.CoreV1().Secrets(kubeSystemNamespace).Get(ctx, secret, metav1.GetOptions{})

		if err != nil {
			e.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("Could not find secret %s in namespace %s", secret, kubeSystemNamespace))
			return microerror.Mask(err)
		}
	}

	return nil
}

// checkKiamPods ensures kiam-agent and kiam-server pods are alive and running
func (e *Executor) checkKiamPods(ctx context.Context, tcClient kubernetes.Interface) error {
	// count expected kiam-server and kiam-agent pods
	var expectedKiamServerPodCount, expectedKiamAgentPodCount int
	{
		masterNodes, err := tcClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{LabelSelector: masterNodeLabelSelector})
		if err != nil {
			return microerror.Mask(err)
		}
		expectedKiamServerPodCount = len(masterNodes.Items)

		workerNodes, err := tcClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{LabelSelector: workerNodeLabelSelector})
		if err != nil {
			return microerror.Mask(err)
		}
		expectedKiamAgentPodCount = len(workerNodes.Items)
	}

	// kiam server
	{
		kiamServerPods, err := tcClient.CoreV1().Pods(kubeSystemNamespace).List(ctx, metav1.ListOptions{LabelSelector: kiamServerLabelSelector})
		if err != nil {
			return microerror.Mask(err)
		}

		if len(kiamServerPods.Items) != expectedKiamServerPodCount {
			e.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("Wrong kiam-server pod count, expected %d but got %d", expectedKiamServerPodCount, len(kiamServerPods.Items)))
			return microerror.Mask(executionFailedError)
		}

		for _, kiamServerPod := range kiamServerPods.Items {
			if kiamServerPod.Status.Phase != apiv1.PodRunning {
				e.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("Pod %s in namespace %s is not running.", kiamServerPod.Name, kiamServerPod.Namespace))
				return microerror.Mask(executionFailedError)
			}
		}
	}

	// kiam agent
	{
		kiamAgentPods, err := tcClient.CoreV1().Pods(kubeSystemNamespace).List(ctx, metav1.ListOptions{LabelSelector: kiamAgentLabelSelector})
		if err != nil {
			return microerror.Mask(err)
		}

		if len(kiamAgentPods.Items) != expectedKiamAgentPodCount {
			e.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("Wrong kiam-agent pod count, expected %d but got %d", expectedKiamAgentPodCount, len(kiamAgentPods.Items)))
			return microerror.Mask(executionFailedError)
		}

		for _, kiamAgentPod := range kiamAgentPods.Items {
			if kiamAgentPod.Status.Phase != apiv1.PodRunning {
				e.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("Pod %s in namespace %s is not running.", kiamAgentPod.Name, kiamAgentPod.Namespace))
				return microerror.Mask(executionFailedError)
			}
		}
	}

	return nil
}

// testAWSApiCalls will spawn a job in k8s tenant cluster to test calling AWS API to ensure kiam works as expected
func (e *Executor) testAWSApiCalls(ctx context.Context, tcClient kubernetes.Interface, awsRegion string, dockerRegistry string) error {
	networkPolicy := jobNetworkPolicy()
	_, err := tcClient.NetworkingV1().NetworkPolicies(kubeSystemNamespace).Create(ctx, networkPolicy, metav1.CreateOptions{})
	if err != nil {
		return microerror.Mask(err)
	}
	// delete the network policies when function ends
	defer func() {
		_ = tcClient.NetworkingV1().NetworkPolicies(kubeSystemNamespace).Delete(ctx, networkPolicy.Name, metav1.DeleteOptions{})
	}()

	job := awsApiCallJob(dockerRegistry, awsRegion, e.tenantCluster)
	_, err = tcClient.BatchV1().Jobs(kubeSystemNamespace).Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		return microerror.Mask(err)
	}
	// delete the job when function ends
	defer func() {
		_ = tcClient.BatchV1().Jobs(kubeSystemNamespace).Delete(ctx, job.Name, metav1.DeleteOptions{})
	}()

	// check if job is completed
	o := func() error {
		job, err := tcClient.BatchV1().Jobs(kubeSystemNamespace).Get(ctx, job.Name, metav1.GetOptions{})
		if err != nil {
			return microerror.Mask(err)
		}

		if !isJobCompleted(job) {
			return microerror.Mask(jobNotCompleted)
		}
		return nil
	}
	b := backoff.NewMaxRetries(checkJobRetryLimit, checkJobMaxWait)
	err = backoff.Retry(o, b)
	if err != nil {
		e.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("Job %s in namespace %s for testing AWS api call for kiam test did not succeed. Check pods logs for more details.", job.Name, job.Namespace))
		return microerror.Mask(err)
	}

	return nil
}

func (e *Executor) fetchDockerRegistry(ctx context.Context, cpClient kubernetes.Interface) (string, error) {
	var dockerRegistry string
	cm, err := cpClient.CoreV1().ConfigMaps(draughtsmanNamespace).Get(ctx, draughtsmanConfigMapName, metav1.GetOptions{})
	if err != nil {
		return "", microerror.Mask(err)
	}

	var valueReaderService *valuemodifierpath.Service
	{
		valueReaderConfig := valuemodifierpath.DefaultConfig()
		valueReaderConfig.InputBytes = []byte(cm.Data[draughtsmanConfigMapDataKey])
		valueReaderService, err = valuemodifierpath.New(valueReaderConfig)
		if err != nil {
			return "", microerror.Mask(err)
		}

		value, err := valueReaderService.Get(draughtsmanConfigMapDockerRegistryKey)
		if err != nil {
			return "", microerror.Mask(err)
		}

		var ok bool
		dockerRegistry, ok = value.(string)
		if !ok {
			e.logger.LogCtx(ctx, "level", "error", "message", "Failed to parse DockerRegistry value from draughtsman configmap on CP.")
			return "", microerror.Mask(executionFailedError)
		}
	}

	return dockerRegistry, nil
}
