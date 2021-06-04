package networkpool

import (
	"context"

	"github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	pkgclient "github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
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
		c := pkgclient.ControlPlaneConfig{
			Logger: r.logger,

			KubeConfig: env.ControlPlaneKubeConfig(),
		}

		cpClients, err = pkgclient.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		np := &v1alpha2.NetworkPool{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: metav1.NamespaceDefault,
				Name:      "custom",
			},
		}
		err = cpClients.CtrlClient().Delete(context.Background(), np)

		if errors.IsNotFound(err) {
			// fall through
		} else if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
