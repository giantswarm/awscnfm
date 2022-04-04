package key

import (
	"context"
	"fmt"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/giantswarm/apiextensions/v6/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/microerror"
	valuemodifierpath "github.com/giantswarm/valuemodifier/path"
	apiv1 "k8s.io/api/core/v1"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v15/pkg/normalize"
	"github.com/giantswarm/awscnfm/v15/pkg/project"
)

const (
	// Organization is the Giant Swarm specific organization we create our
	// conformance test clusters in.
	Organization = "conformance-testing"
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

const (
	organizationNamespaceFormat = "org-%s"
)

const (
	// FirstOrgNamespaceRelease is the first GS release that creates Clusters in Org Namespaces by default
	FirstAWSOrgNamespaceRelease = "16.0.0"
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

func EBSTestJobName() string {
	return fmt.Sprintf("%s-ebs-volume-test", project.Name())
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

func OrganizationNamespaceFromName(name string) string {
	name = normalize.AsDNSLabelName(fmt.Sprintf(organizationNamespaceFormat, name))

	return name
}

// IsOrgNamespaceVersion returns whether a given AWS GS Release Version is based on clusters in Org Namespace
func IsOrgNamespaceVersion(version string) bool {
	// TODO: this has to return true as soon as v16.0.0 is the newest version
	// Background: in case the release version is not set, aws-admission-controller mutates to the the latest AWS version,
	// see https://github.com/giantswarm/aws-admission-controller/blob/ef83d90fc856fbc0484bec967064834c0b8d2c1e/pkg/aws/v1alpha3/cluster/mutate_cluster.go#L191-L202
	// so as soon as the latest version is >=16.0.0 we are going to need the org-namespace as default here.
	if version == "" {
		return true
	}
	OrgNamespaceVersion, _ := semver.New(FirstAWSOrgNamespaceRelease)
	releaseVersion, _ := semver.New(version)
	return releaseVersion.GE(*OrgNamespaceVersion)
}

func MoveClusterCRsToOrgNamespace(crs v1alpha3.ClusterCRs, organization string) v1alpha3.ClusterCRs {
	crs.Cluster.SetNamespace(OrganizationNamespaceFromName(organization))
	crs.Cluster.Spec.InfrastructureRef.Namespace = OrganizationNamespaceFromName(organization)
	crs.AWSCluster.SetNamespace(OrganizationNamespaceFromName(organization))
	crs.G8sControlPlane.SetNamespace(OrganizationNamespaceFromName(organization))
	crs.G8sControlPlane.Spec.InfrastructureRef.Namespace = OrganizationNamespaceFromName(organization)
	crs.AWSControlPlane.SetNamespace(OrganizationNamespaceFromName(organization))
	return crs
}

func MoveNodePoolCRsToOrgNamespace(crs v1alpha3.NodePoolCRs, namespace string) v1alpha3.NodePoolCRs {
	crs.MachineDeployment.SetNamespace(OrganizationNamespaceFromName(namespace))
	crs.MachineDeployment.Spec.Template.Spec.InfrastructureRef.Namespace = OrganizationNamespaceFromName(namespace)
	crs.AWSMachineDeployment.SetNamespace(OrganizationNamespaceFromName(namespace))
	return crs
}
