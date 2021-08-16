package pl003

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/generate"
)

type flag struct {
	TenantCluster string
	Output        string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.TenantCluster, "tenant-cluster", "c", env.TenantCluster(), "Tenant Cluster ID to use for this particular action.")
	cmd.Flags().StringVar(&f.Output, "output", "", `The directory in which to store the cluster ID, kubeconfig, and provider of the created cluster.`)
}

func (f *flag) Validate() error {
	if f.TenantCluster == "" {
		f.TenantCluster = generate.ID()
	}

	// Write cluster ID to filesystem
	{
		clusterIDPath := filepath.Join(f.Output, "cluster-id")
		err := ioutil.WriteFile(clusterIDPath, []byte(f.TenantCluster), 0644) //#nosec
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
