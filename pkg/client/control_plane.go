package client

import (
	"os/user"
	"path/filepath"
	"strings"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha2"
	releasev1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"

	"github.com/giantswarm/awscnfm/v12/pkg/env"
)

type ControlPlaneConfig struct {
	Logger micrologger.Logger

	KubeConfig string
}

func NewControlPlane(config ControlPlaneConfig) (k8sclient.Interface, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	// We want to default the local kube config path in case it is not given.
	// The default kube config path is usually prefixed by "~/", which does not
	// properly resolve on all systems. That is why we trim this particular
	// prefix before joining the current users' home dir.
	if config.KubeConfig == "" {
		u, err := user.Current()
		if err != nil {
			return nil, microerror.Mask(err)
		}

		config.KubeConfig = filepath.Join(u.HomeDir, strings.TrimPrefix(env.DefaultControlPlaneKubeConfig, "~/"))
	}

	var err error

	var clients *k8sclient.Clients
	{
		c := k8sclient.ClientsConfig{
			SchemeBuilder: k8sclient.SchemeBuilder{
				apiv1alpha2.AddToScheme,
				infrastructurev1alpha2.AddToScheme,
				releasev1alpha1.AddToScheme,
			},
			Logger: config.Logger,

			KubeConfigPath: config.KubeConfig,
		}

		clients, err = k8sclient.NewClients(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return clients, nil
}
