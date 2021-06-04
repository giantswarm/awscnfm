package pl002

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/generate"
)

type flag struct {
	TenantCluster string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.TenantCluster, "tenant-cluster", "c", env.TenantCluster(), "Tenant Cluster ID to use for this particular action.")
}

func (f *flag) Validate() error {
	if f.TenantCluster == "" {
		f.TenantCluster = generate.ID()
	}

	return nil
}
