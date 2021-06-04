package volume

import (
	"context"

	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

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

	err = r.cleanupEBSTestResources(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// cleanupEBSTestResources will cleanup EBS volume test resources
func (r *runner) cleanupEBSTestResources(ctx context.Context, tcClient kubernetes.Interface) error {
	// in case we have errors, lets collect them to ensure other resources get deleted in case they exists
	var errSlice []error

	err := tcClient.BatchV1().Jobs("default").Delete(ctx, key.EBSTestJobName(), metav1.DeleteOptions{})
	if err != nil {
		errSlice = append(errSlice, err)
	}

	err = tcClient.RbacV1().ClusterRoles().Delete(ctx, "enable-ebs-psp", metav1.DeleteOptions{})
	if err != nil {
		errSlice = append(errSlice, err)
	}

	err = tcClient.RbacV1().RoleBindings("default").Delete(ctx, "ebs-rolebinding", metav1.DeleteOptions{})
	if err != nil {
		errSlice = append(errSlice, err)
	}

	err = tcClient.CoreV1().PersistentVolumeClaims("default").Delete(ctx, "ebs-claim", metav1.DeleteOptions{})
	if err != nil {
		errSlice = append(errSlice, err)
	}

	pods, err := tcClient.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
		LabelSelector: "app=awscnfm-ebs-volume-test",
	})
	if err != nil {
		errSlice = append(errSlice, err)
	}

	err = tcClient.CoreV1().Pods("default").Delete(ctx, pods.Items[0].GetName(), metav1.DeleteOptions{})
	if err != nil {
		errSlice = append(errSlice, err)
	}

	// return the first error
	if len(errSlice) > 0 {
		return errSlice[0]
	}

	return nil
}
