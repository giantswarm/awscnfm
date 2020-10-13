package ac010

import (
	"context"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	valuemodifierpath "github.com/giantswarm/valuemodifier/path"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

const (
	kubeSystemNamespace = "kube-system"

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

			KubeConfig: env.ControlPlaneKubeConfig(),
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

	err = e.createAWSApiCallJob(ctx, tcClients.K8sClient(), awsRegion, dockerRegistry)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// createAWSApiCallJob will spawn a job in k8s tenant cluster to test calling AWS API to ensure kiam works as expected
func (e *Executor) createAWSApiCallJob(ctx context.Context, tcClient kubernetes.Interface, awsRegion string, dockerRegistry string) error {
	networkPolicy := jobNetworkPolicy()
	_, err := tcClient.NetworkingV1().NetworkPolicies(kubeSystemNamespace).Create(ctx, networkPolicy, metav1.CreateOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	job := awsApiCallJob(dockerRegistry, awsRegion, e.tenantCluster)
	_, err = tcClient.BatchV1().Jobs(kubeSystemNamespace).Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
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
			return "", microerror.Maskf(executionFailedError, "Failed to parse DockerRegistry value from draughtsman configmap on CP.")
		}
	}

	return dockerRegistry, nil
}
