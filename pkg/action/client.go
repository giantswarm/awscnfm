package action

import (
	"context"

	"github.com/giantswarm/k8sclient/v3/pkg/k8sclient"
	"github.com/giantswarm/micrologger"
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
	return nil, nil
}

func (c *Clients) InitControlPlane(ctx context.Context) error {
	return nil
}

func (c *Clients) InitTenantCluster(ctx context.Context) error {
	return nil
}
