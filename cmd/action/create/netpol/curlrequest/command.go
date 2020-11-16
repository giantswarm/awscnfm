package curlrequest

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "curlrequest"
	short = "Create a resources to test the network policy 'deny-from-all-namespaces'."
	long  = `Create a curl request job in the 'test' namespace and same job in the default namespace and test the network policy 'deny-from-all-namespaces' by connecting to the nginx test pod. The job in the 'test' namespace should successfully connect to the nginx test pod and the job in 'default' namespace should fail to connect to the nginx test pod.
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
