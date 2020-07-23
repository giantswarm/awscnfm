package action

import (
	"regexp"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

const (
	allActions = "add"
)

type flag struct {
	Action  string
	Cluster string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Action, "action", "a", "", "The name of the action to generate the subcommand for, e.g. ac013 or all to generate all actions.")
	cmd.Flags().StringVarP(&f.Cluster, "cluster", "c", "", "The name of the cluster to generate the subcommand for, e.g. cl001.")
}

func (f *flag) Validate() error {
	if f.Action == "" {
		return microerror.Maskf(invalidFlagError, "-a/--action must not be empty")
	}
	if f.Action != allActions && !strings.HasPrefix(f.Action, "ac") {
		return microerror.Maskf(invalidFlagError, "-a/--action must have ac prefix, e.g. ac013, got %#q", f.Action)
	}
	if f.Action != allActions && !regexp.MustCompile(`[0-9]{3}$`).MatchString(f.Action) {
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
