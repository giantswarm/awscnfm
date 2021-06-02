package ready

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v14/pkg/env"
)

type flag struct {
	TenantCluster string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.TenantCluster, "tenant-cluster", "c", env.TenantCluster(), "Tenant Cluster ID to use for this particular action.")
}

func (f *flag) Validate() error {
	if f.TenantCluster == "" {
		return microerror.Maskf(invalidFlagsError, "-c/--tenant-cluster must not be empty")
	}

	return nil
}
