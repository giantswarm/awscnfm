package env

import (
	"os"

	"github.com/giantswarm/microerror"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	EnvVarKubeConfig = "AWSCNFM_KUBECONFIG"
)

// KubeConfig is the local path to the kube config file used to create the
// Control Plane specific rest config.
func KubeConfig() string {
	return os.Getenv(EnvVarKubeConfig)
}

// CurrentServer returns the server from the current context
// e.g. https://my.installation.region.provider.company.domain:443
func CurrentServer(kubeConfigPath string) (string, error) {
	kubeConfig, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return "", microerror.Mask(err)
	}
	context := kubeConfig.Contexts[kubeConfig.CurrentContext]
	return kubeConfig.Clusters[context.Cluster].Server, nil
}
