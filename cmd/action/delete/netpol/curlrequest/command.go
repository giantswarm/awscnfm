package curlrequest

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "curlrequest"
	short = "Cleanup curl request resources for testing network policy from the tenant Cluster."
	long  = "Cleanup curl request resources for testing network policy from the tenant Cluster."
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
