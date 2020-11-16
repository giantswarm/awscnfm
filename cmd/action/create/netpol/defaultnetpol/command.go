package defaultnetpol

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "defaultnetpol"
	short = "Create a default network policy 'deny-from-all' and test pod which will be used to test the network policy."
	long  = `Create a default network policy that denies traffic from any other namespaces in default namespace and a simple pod running in default namespace that can be used for testing if the network policy work.
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
