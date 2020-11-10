package updated

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "updated"
	short = "Check if the Tenant Cluster got successfully upgraded."
	long  = `Check if the Tenant Cluster got successfully upgraded. Note that this
particular action is not meant to be reliably used for other purposes than
for the plan exection. Executing this action against a Tenant Cluster that
got already upgraded may lead to wrong results in case you want to assert an
additional Tenant Cluster upgrade.

    * Fetch the AWSCluster CR.
    * Check if the latest cluster status condition is "Updated".
    * Return an error if we see other cluster status conditions than "Updated".

	`
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
	}

	c := &cobra.Command{
		Use:   name,
		Short: short,
		Long:  long,
		RunE:  r.Run,
	}

	f.Init(c)

	return c, nil
}
