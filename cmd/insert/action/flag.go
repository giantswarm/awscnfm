package action

import (
	"regexp"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

type flag struct {
	Action  string
	Cluster string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Action, "action", "a", "", "The name of the action to insert, e.g. ac013.")
	cmd.Flags().StringVarP(&f.Cluster, "cluster", "c", "", "The name of the cluster to insert the action for, e.g. cl001.")
}

func (f *flag) Validate() error {
	if f.Action == "" {
		return microerror.Maskf(invalidFlagError, "-a/--action must not be empty")
	}
	if !strings.HasPrefix(f.Action, "ac") {
		return microerror.Maskf(invalidFlagError, "-a/--action must have ac prefix, e.g. ac013, got %#q", f.Action)
	}
	if !regexp.MustCompile(`[0-9]{3}$`).MatchString(f.Action) {
		return microerror.Maskf(invalidFlagError, "-a/--action must have numbered suffix, e.g. ac013, got %#q", f.Action)
	}

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
