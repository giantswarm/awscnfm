package awsapicall

import (
	"context"

	infrastructurev1alpha3 "github.com/giantswarm/apiextensions/v6/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/k8sclient/v7/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/key"
	"github.com/giantswarm/awscnfm/v15/pkg/label"
)

const (
	kubeSystemNamespace = "kube-system"
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

	// awsRegion is necessary to execute aws-cli call
	var awsRegion string
	{
		var cr infrastructurev1alpha3.AWSCluster
		{
			var list infrastructurev1alpha3.AWSClusterList
			err := cpClients.CtrlClient().List(
				ctx,
				&list,
				k8sruntimeclient.InNamespace(cr.GetNamespace()),
				k8sruntimeclient.MatchingLabels{label.Cluster: r.flag.TenantCluster},
			)
			if err != nil {
				return microerror.Mask(err)
			}

			if len(list.Items) == 0 {
				return microerror.Mask(notFoundError)
			}
			if len(list.Items) > 1 {
				return microerror.Mask(tooManyCRsError)
			}

			cr = list.Items[0]
		}

		awsRegion = cr.Spec.Provider.Region
	}

	// dockerRegistry is needed in order to spawn pod with proper docker image that will execute aws-cli call
	var dockerRegistry string
	{
		dockerRegistry, err = key.FetchDockerRegistry(ctx, cpClients.CtrlClient())
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = r.createAWSApiCallJob(ctx, tcClients.CtrlClient(), awsRegion, dockerRegistry)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// createAWSApiCallJob will spawn a job in k8s tenant cluster to test calling AWS API to ensure kiam works as expected
func (r *runner) createAWSApiCallJob(ctx context.Context, tcClient k8sruntimeclient.Client, awsRegion string, dockerRegistry string) error {
	networkPolicy := jobNetworkPolicy()

	err := tcClient.Create(ctx, networkPolicy)
	if err != nil {
		return microerror.Mask(err)
	}

	job := awsApiCallJob(dockerRegistry, awsRegion, r.flag.TenantCluster)
	err = tcClient.Create(ctx, job)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
