package ac011

import (
	"context"

	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/key"
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

	err = e.cleanupKiamTestResources(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// cleanupKiamTestResources will cleanup kiam test resources
func (e *Executor) cleanupKiamTestResources(ctx context.Context, tcClient kubernetes.Interface) error {
	err := tcClient.BatchV1().Jobs(key.KiamTestNamespace).Delete(ctx, key.KiamTestJobName(), metav1.DeleteOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	err = tcClient.NetworkingV1().NetworkPolicies(key.KiamTestNamespace).Delete(ctx, key.KiamTestNetPolName(), metav1.DeleteOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
