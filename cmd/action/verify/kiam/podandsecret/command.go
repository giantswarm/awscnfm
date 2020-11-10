package podandsecret

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "podandsecret"
	short = "Check if Kiam's pods and secrets are present on the tenant cluster"
	long  = `Check if Kiam's pods and secrets are present on the tenant cluster. This test
should do following:

    * Ensure tls certs for Kiam are created.
    * Ensure kiam-server and kiam-agent pods are running without errors on the tenant cluster.
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
