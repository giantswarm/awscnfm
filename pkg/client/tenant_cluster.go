package client

import (
	"context"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/certs/v3/pkg/certs"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/tenantcluster/v3/pkg/tenantcluster"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pkgconfig "github.com/giantswarm/awscnfm/v12/pkg/config"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

type TenantClusterConfig struct {
	ControlPlane k8sclient.Interface
	Logger       micrologger.Logger

	Scope string
}

func NewTenantCluster(config TenantClusterConfig) (k8sclient.Interface, error) {
	if config.ControlPlane == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.ControlPlane must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.Scope == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Scope must not be empty", config)
	}

	var err error

	ctx := context.Background()
	id := pkgconfig.Cluster(config.Scope, env.TenantCluster())

	var clients *k8sclient.Clients
	{
		var cr infrastructurev1alpha2.AWSCluster
		{
			var list infrastructurev1alpha2.AWSClusterList
			err := config.ControlPlane.CtrlClient().List(
				ctx,
				&list,
				client.InNamespace(cr.GetNamespace()),
				client.MatchingLabels{label.Cluster: id},
			)
			if err != nil {
				return nil, microerror.Mask(err)
			}

			if len(list.Items) == 0 {
				return nil, microerror.Mask(notFoundError)
			}
			if len(list.Items) > 1 {
				return nil, microerror.Mask(tooManyCRsError)
			}

			cr = list.Items[0]
		}

		var certsSearcher *certs.Searcher
		{
			c := certs.Config{
				K8sClient: config.ControlPlane.K8sClient(),
				Logger:    config.Logger,
			}

			certsSearcher, err = certs.NewSearcher(c)
			if err != nil {
				return nil, microerror.Mask(err)
			}
		}

		var tenantCluster tenantcluster.Interface
		{
			c := tenantcluster.Config{
				CertsSearcher: certsSearcher,
				Logger:        config.Logger,

				CertID: certs.ClusterOperatorAPICert,
			}

			tenantCluster, err = tenantcluster.New(c)
			if err != nil {
				return nil, microerror.Mask(err)
			}
		}

		var restConfig *rest.Config
		{
			restConfig, err = tenantCluster.NewRestConfig(
				ctx,
				id,
				key.APIEndpoint(id, cr.Spec.Cluster.DNS.Domain),
			)
			if err != nil {
				return nil, microerror.Mask(err)
			}
		}

		{
			c := k8sclient.ClientsConfig{
				Logger:     config.Logger,
				RestConfig: rest.CopyConfig(restConfig),
			}

			clients, err = k8sclient.NewClients(c)
			if err != nil {
				return nil, microerror.Mask(err)
			}
		}
	}

	return clients, nil
}
