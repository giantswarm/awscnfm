package defaultnetpol

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "defaultnetpol"
	short = "Create a default network policy 'deny-from-all' and test pod which will be used to test the network policy."
	long  = `Create a test resources for a default network policy that denies traffic from other namespaces.
Created resources:
* namespace 'test'
* network policy 'deny-from-other-namespaces' in namespace 'test'
* nginx pod running in 'test' namespace
* service for the nginx to test the connection to the pod
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
