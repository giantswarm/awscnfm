package hostnetworkpod

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "hostnetworkpod"
	short = "Verify the number of host network pods on master nodes."
	long  = `Check if the number of host network pods on tenant cluster master nodes
matches the number we expect from k8scloudconfig.

	* Fetch all Tenant Cluster nodes and take the the first master node by label.
	* Compare the current pods with host network set with the expected amount of pods on master node.
	* See also https://github.com/giantswarm/k8scloudconfig/blob/529491d591e039da1ffde03fef070101c8d4a95c/files/conf/setup-kubelet-environment#L21.

The action waits up to 15 minutes to verify all pods with host network set
are being deployed. This ensures pods (e.g. node-exporter) have enough time
to being scheduled.
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
