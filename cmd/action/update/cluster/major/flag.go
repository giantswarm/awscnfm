package major

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/env"
)

type flag struct {
	TenantCluster  string
	ReleaseVersion string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.TenantCluster, "tenant-cluster", "c", env.TenantCluster(), "Tenant Cluster ID to use for this particular action.")
	cmd.Flags().StringVarP(&f.ReleaseVersion, "release-version", "r", env.UpdateReleaseVersion(), "Release Version to use for updating a Tenant Cluster.")
}

func (f *flag) Validate() error {
	if f.TenantCluster == "" {
		return microerror.Maskf(invalidFlagsError, "-c/--tenant-cluster must not be empty")
	}

	if f.ReleaseVersion == "" {
		return microerror.Maskf(invalidFlagsError, "-r/--release-version must not be empty")
	}

	return nil
}
