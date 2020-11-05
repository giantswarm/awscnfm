package customnetworkpool

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/valid"
)

type flag struct {
	// TenantCluster is optional for cluster creations. A random cluster ID will
	// be generated when none is provided.
	TenantCluster string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.TenantCluster, "tenant-cluster", "c", env.TenantCluster(), "Tenant Cluster ID to use for this particular action.")
}

func (f *flag) Validate() error {
	if f.TenantCluster != "" {
		err := valid.ID(f.TenantCluster)
		if err != nil {
			return microerror.Maskf(invalidFlagsError, "-c/--tenant-cluster %s", err.Error())
		}
	}

	return nil
}
