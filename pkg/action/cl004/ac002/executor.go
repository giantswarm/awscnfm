package ac002

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
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

	// Check for Tenant Cluster API accessibility. If we are able to list the
	// Tenant Cluster nodes we are good to move on.
	_, err = tcClients.K8sClient().CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	var cl v1alpha2.AWSCluster
	{
		err = cpClients.CtrlClient().Get(
			ctx,
			types.NamespacedName{Name: e.tenantCluster, Namespace: corev1.NamespaceDefault},
			&cl,
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	// Checking API availability is one thing, though there is more to the
	// cluster creation as far as the system is concerned. We want to check of
	// the Created status condition is properly set. Only then we assume the
	// Tenant Cluster got successfully created.
	if cl.Status.Cluster.LatestCondition() == v1alpha2.ClusterStatusConditionCreated {
		return nil
	}

	return microerror.Maskf(wrongClusterStatusConditionError, cl.Status.Cluster.LatestCondition())
}
