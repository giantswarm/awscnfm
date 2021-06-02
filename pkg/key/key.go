package key

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/microerror"
	valuemodifierpath "github.com/giantswarm/valuemodifier/path"
	apiv1 "k8s.io/api/core/v1"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v14/pkg/project"
)

const (
	// Credential is the default credential we use for most of our conformance
	// test clusters. These credentials define which AWS Account to use.
	Credential = "credential-default"
	// Organization is the Giant Swarm specific organization we create our
	// conformance test clusters in.
	Organization = "giantswarm"
)

const (
	KiamTestNamespace = "kube-system"
)

const (
	draughtsmanNamespace                  = "draughtsman"
	draughtsmanConfigMapName              = "draughtsman-values-configmap"
	draughtsmanConfigMapDataKey           = "values"
	draughtsmanConfigMapDockerRegistryKey = "Installation.V1.Registry.Domain"
)

const (
	NetPolDefaultNamespaceName = "default"
	NetPolTestNamespaceName    = "test"

	NetPolName                 = "deny-from-all-namespaces"
	NetPolNginxTestPodName     = "netpol-nginx-test-pod"
	NetPolNginxTestPodAppLabel = "nginx-test-pod"
	NetPolNginxSvcName         = "netpol-nginx-test-svc"

	LabelApp       = "app"
	LabelManagedBy = "managed-by"
)

func APIEndpoint(id string, base string) string {
	return fmt.Sprintf("api.%s.k8s.%s", id, base)
}

func DomainFromHost(h string) string {
	h = strings.Replace(h, "https://", "", 1)
	h = strings.Replace(h, "g8s.", "", 1)
	h = strings.Replace(h, ":443", "", 1)
	return h
}

func KiamTestJobName() string {
	return fmt.Sprintf("%s-kiam-test", project.Name())
}

func KiamTestNetPolName() string {
	return fmt.Sprintf("%s-kiam-test", project.Name())
}

func NetPolCurlTestCommand() string {
	return fmt.Sprintf("/usr/bin/curl --connect-timeout 5 %s.%s.svc", NetPolNginxSvcName, NetPolTestNamespaceName)
}

func NetPolTestJobName(namespace string) string {
	return fmt.Sprintf("%s-netpol-%s-job", project.Name(), namespace)
}

func RegionFromHost(h string) string {
	h = strings.Split(DomainFromHost(h), ".")[1]
	// Domain in giraffe is giraffe.pek.aws.k8s.adidas.com.cn which does not follow the convention
	if h == "pek" {
		h = "cn-north-1"
	}
	return h
}

func FetchDockerRegistry(ctx context.Context, cpCtrlClient k8sruntimeclient.Client) (string, error) {
	var dockerRegistry string

	cm := &apiv1.ConfigMap{}

	err := cpCtrlClient.Get(
		ctx,
		k8sruntimeclient.ObjectKey{
			Name:      draughtsmanConfigMapName,
			Namespace: draughtsmanNamespace,
		},
		cm)

	if err != nil {
		return "", microerror.Mask(err)
	}

	var valueReaderService *valuemodifierpath.Service
	{
		valueReaderConfig := valuemodifierpath.DefaultConfig()
		valueReaderConfig.InputBytes = []byte(cm.Data[draughtsmanConfigMapDataKey])
		valueReaderService, err = valuemodifierpath.New(valueReaderConfig)
		if err != nil {
			return "", microerror.Mask(err)
		}

		value, err := valueReaderService.Get(draughtsmanConfigMapDockerRegistryKey)
		if err != nil {
			return "", microerror.Mask(err)
		}

		var ok bool
		dockerRegistry, ok = value.(string)
		if !ok {
			return "", microerror.Maskf(executionFailedError, "Failed to parse DockerRegistry value from draughtsman configmap on CP.")
		}
	}

	return dockerRegistry, nil
}
