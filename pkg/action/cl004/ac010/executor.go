package ac010

import (
	"context"
	"fmt"

	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	batchapiv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
)

const (
	kubeSystemNamespace = "kube-system"
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

	err = e.checkKiamAWScallPod(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// checkKiamAWScallPod will check if the AWS API call job was successful
func (e *Executor) checkKiamAWScallPod(ctx context.Context, tcClient kubernetes.Interface) error {

	name := fmt.Sprintf("%s-kiam-test", project.Name())

	job, err := tcClient.BatchV1().Jobs(kubeSystemNamespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	if !isJobCompleted(job) {
		return microerror.Mask(jobNotCompleted)
	}

	return nil
}

func isJobCompleted(j *batchapiv1.Job) bool {
	for _, c := range j.Status.Conditions {
		if c.Type == "Complete" && c.Status == apiv1.ConditionTrue {
			return true
		}
	}
	return false
}
