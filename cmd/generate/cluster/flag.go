package cluster

import (
	"regexp"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

type flag struct {
	Cluster string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Cluster, "cluster", "c", "", "The name of the cluster to generate the sub command for, e.g. cl001.")
}

func (f *flag) Validate() error {
	if f.Cluster == "" {
		return microerror.Maskf(invalidFlagError, "-c/--cluster must not be empty")
	}
	if !strings.HasPrefix(f.Cluster, "cl") {
		return microerror.Maskf(invalidFlagError, "-c/--cluster must have cl prefix, e.g. cl001, got %#q", f.Cluster)
	}
	if !regexp.MustCompile(`[0-9]{3}$`).MatchString(f.Cluster) {
		return microerror.Maskf(invalidFlagError, "-c/--cluster must have numbered suffix, e.g. cl001, got %#q", f.Cluster)
	}

	return nil
}
