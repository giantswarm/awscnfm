package customnetworkpool

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/valid"
)

type flag struct {
	// TenantCluster is optional for cluster creations. A random cluster ID will
	// be generated when none is provided.
	TenantCluster  string
	ReleaseVersion string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.TenantCluster, "tenant-cluster", "c", env.TenantCluster(), "Tenant Cluster ID to use for this particular action.")
	cmd.Flags().StringVarP(&f.ReleaseVersion, "release-version", "r", env.CreateReleaseVersion(), "Release Version to use for creating a Tenant Cluster.")
}

func (f *flag) Validate() error {
	if f.TenantCluster != "" {
		err := valid.ID(f.TenantCluster)
		if err != nil {
			return microerror.Maskf(invalidFlagsError, "-c/--tenant-cluster %s", err.Error())
		}
	}

	if f.ReleaseVersion == "" {
		return microerror.Maskf(invalidFlagsError, "-r/--release-version must not be empty")
	}

	return nil
}
