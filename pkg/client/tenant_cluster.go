package client

import (
	"context"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	releasev1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/certs/v3/pkg/certs"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/tenantcluster/v3/pkg/tenantcluster"
	"k8s.io/client-go/rest"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pkgconfig "github.com/giantswarm/awscnfm/v12/pkg/config"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

type TenantClusterConfig struct {
	ControlPlane k8sclient.Interface
	Logger       micrologger.Logger

	KubeConfig string
	Scope      string
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

	// If a kube config is provided we prefer it. Usually it should be given via
	// environment variables. That might be most often the case when running in
	// some CI system.
	if config.KubeConfig != "" {
		clients, err := clientsFromKubeConfig(config)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		return clients, nil
	}

	// If there is no kube config provided we try to look it up via Control
	// Plane resources. Therefore the Control Plane client is necessary as well
	// as the cluster scope so that we can lookup the actual Tenant Cluster ID.
	clients, err := clientsFromAPISecret(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return clients, nil
}

func clientsFromAPISecret(config TenantClusterConfig) (k8sclient.Interface, error) {
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

func clientsFromKubeConfig(config TenantClusterConfig) (k8sclient.Interface, error) {
	var err error

	var clients *k8sclient.Clients
	{
		c := k8sclient.ClientsConfig{
			SchemeBuilder: k8sclient.SchemeBuilder{
				apiv1alpha2.AddToScheme,
				infrastructurev1alpha2.AddToScheme,
				releasev1alpha1.AddToScheme,
			},
			Logger: config.Logger,

			KubeConfigPath: config.KubeConfig,
		}

		clients, err = k8sclient.NewClients(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return clients, nil
}
