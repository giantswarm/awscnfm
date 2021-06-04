package curlrequest

import (
	"context"

	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	batchapiv1 "k8s.io/api/batch/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	err = r.cleanupCurlRequestJobs(ctx, tcClients.CtrlClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// cleanupCurlRequestJobs will cleanup curl request test resources
func (r *runner) cleanupCurlRequestJobs(ctx context.Context, tcClient k8sruntimeclient.Client) error {
	successJob := &batchapiv1.Job{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "Job",
			APIVersion: batchapiv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      key.NetPolTestJobName(key.NetPolTestNamespaceName),
			Namespace: key.NetPolTestNamespaceName,
		},
	}

	err := tcClient.Delete(ctx, successJob)
	if err != nil {
		return microerror.Mask(err)
	}

	failJob := &batchapiv1.Job{
		TypeMeta: apismetav1.TypeMeta{
			Kind:       "Job",
			APIVersion: batchapiv1.GroupName,
		},
		ObjectMeta: apismetav1.ObjectMeta{
			Name:      key.NetPolTestJobName(key.NetPolDefaultNamespaceName),
			Namespace: key.NetPolDefaultNamespaceName,
		},
	}

	err = tcClient.Delete(ctx, failJob)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
