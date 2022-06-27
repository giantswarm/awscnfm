package volume

import (
	"context"

	"github.com/giantswarm/k8sclient/v7/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	batchapiv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/key"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: r.logger,

			KubeConfig: env.ControlPlaneKubeConfig(),
		}

		cpClients, err = client.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var tcClients k8sclient.Interface
	{
		c := client.TenantClusterConfig{
			ControlPlane: cpClients,
			Logger:       r.logger,

			TenantCluster: r.flag.TenantCluster,
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = r.checkEBSVolumePod(ctx, tcClients.CtrlClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// checkEBSVolumePod will check if the EBS volume job was successful
func (r *runner) checkEBSVolumePod(ctx context.Context, tcClient k8sruntimeclient.Client) error {
	job := &batchapiv1.Job{}

	err := tcClient.Get(
		ctx,
		k8sruntimeclient.ObjectKey{
			Namespace: "default",
			Name:      key.EBSTestJobName(),
		},
		job,
	)
	if err != nil {
		return microerror.Mask(err)
	}

	if !isJobCompleted(job) {
		return microerror.Mask(jobNotCompleted)
	}

	return nil
}

func isJobCompleted(j *batchapiv1.Job) bool {
	for _, c := range j.Status.Conditions {
		if c.Type == "Complete" && c.Status == apiv1.ConditionTrue {
			return true
		}
	}
	return false
}
