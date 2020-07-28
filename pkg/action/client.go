package action

import (
	"context"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	releasev1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/certs/v2/pkg/certs"
	"github.com/giantswarm/k8sclient/v3/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/tenantcluster/v2/pkg/tenantcluster"
	"k8s.io/client-go/rest"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/pkg/key"
	"github.com/giantswarm/awscnfm/pkg/label"
)

type Config struct {
	Logger micrologger.Logger

	// KubeConfig is the local path to the kube config file used to create the
	// Control Plane specific rest config.
	KubeConfig string
	// TenantCluster is the Tenant Cluster ID to consider, e.g. al9qy.
	TenantCluster string
}

type Clients struct {
	ControlPlane  k8sclient.Interface
	TenantCluster k8sclient.Interface

	logger micrologger.Logger

	kubeConfig    string
	tenantCluster string
}

func NewClients(config Config) (*Clients, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.KubeConfig == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.KubeConfig must not be empty", config)
	}

	c := &Clients{
		ControlPlane:  nil,
		TenantCluster: nil,

		logger: config.Logger,

		kubeConfig:    config.KubeConfig,
		tenantCluster: config.TenantCluster,
	}

	return c, nil
}

func (c *Clients) InitControlPlane(ctx context.Context) error {
	var err error

	var cpClient *k8sclient.Clients
	{
		c := k8sclient.ClientsConfig{
			SchemeBuilder: k8sclient.SchemeBuilder{
				apiv1alpha2.AddToScheme,
				infrastructurev1alpha2.AddToScheme,
				releasev1alpha1.AddToScheme,
			},
			Logger: c.logger,

			KubeConfigPath: c.kubeConfig,
		}

		cpClient, err = k8sclient.NewClients(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	c.ControlPlane = cpClient

	return nil
}

func (c *Clients) InitTenantCluster(ctx context.Context) error {
	var err error

	var tcClient *k8sclient.Clients
	{
		var cr infrastructurev1alpha2.AWSCluster
		{
			var list infrastructurev1alpha2.AWSClusterList
			err := c.ControlPlane.CtrlClient().List(
				ctx,
				&list,
				client.InNamespace(cr.GetNamespace()),
				client.MatchingLabels{label.Cluster: c.tenantCluster},
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

		var certsSearcher *certs.Searcher
		{
			c := certs.Config{
				K8sClient: c.ControlPlane.K8sClient(),
				Logger:    c.logger,
			}

			certsSearcher, err = certs.NewSearcher(c)
			if err != nil {
				return microerror.Mask(err)
			}
		}

		var tenantCluster tenantcluster.Interface
		{
			c := tenantcluster.Config{
				CertsSearcher: certsSearcher,
				Logger:        c.logger,

				CertID: certs.ClusterOperatorAPICert,
			}

			tenantCluster, err = tenantcluster.New(c)
			if err != nil {
				return microerror.Mask(err)
			}
		}

		var restConfig *rest.Config
		{
			restConfig, err = tenantCluster.NewRestConfig(
				ctx,
				c.tenantCluster,
				key.APIEndpoint(c.tenantCluster, cr.Spec.Cluster.DNS.Domain),
			)
			if err != nil {
				return microerror.Mask(err)
			}
		}

		{
			c := k8sclient.ClientsConfig{
				Logger:     c.logger,
				RestConfig: rest.CopyConfig(restConfig),
			}

			tcClient, err = k8sclient.NewClients(c)
			if err != nil {
				return microerror.Mask(err)
			}
		}
	}

	c.TenantCluster = tcClient

	return nil
}
