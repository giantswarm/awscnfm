package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/giantswarm/apiextensions/v6/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/k8smetadata/pkg/annotation"
	"github.com/giantswarm/k8smetadata/pkg/label"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

const (
	defaultMasterInstanceType = "m5.xlarge"
	kindAWSCluster            = "AWSCluster"
	kindAWSControlPlane       = "AWSControlPlane"
	kindG8sControlPlane       = "G8sControlPlane"

	// NameChars represents the character set used to generate resource names.
	// (does not contain 1 and l, to avoid confusion)
	NameChars = "023456789abcdefghijkmnopqrstuvwxyz"
	// NameLengthLong represents the number of characters used to create a resource name when --enable-long-names feature flag is used.
	NameLengthLong = 10
	// NameLengthShort represents the number of characters used to create a resource name.
	NameLengthShort = 5
)

// +k8s:deepcopy-gen=false

type ClusterCRsConfig struct {
	ClusterID         string
	ClusterName       string
	ControlPlaneName  string
	Credential        string
	Domain            string
	EnableLongNames   bool
	ExternalSNAT      bool
	ControlPlaneAZ    []string
	Description       string
	MasterAZ          []string
	PodsCIDR          string
	Owner             string
	Region            string
	ReleaseComponents map[string]string
	ReleaseVersion    string
	Labels            map[string]string
	NetworkPool       string
}

// +k8s:deepcopy-gen=false

type ClusterCRs struct {
	Cluster         *apiv1alpha3.Cluster
	AWSCluster      *v1alpha3.AWSCluster
	G8sControlPlane *v1alpha3.G8sControlPlane
	AWSControlPlane *v1alpha3.AWSControlPlane
}

func NewClusterCRs(config ClusterCRsConfig) (ClusterCRs, error) {
	// Default some essentials in case certain information are not given. E.g.
	// the workload cluster name may be provided by the user.
	{
		if config.ClusterName == "" {
			generatedName, err := GenerateName(config.EnableLongNames)
			if err != nil {
				return ClusterCRs{}, microerror.Mask(err)
			}

			config.ClusterName = generatedName
		}

		if config.ControlPlaneName == "" {
			generatedName, err := GenerateName(config.EnableLongNames)
			if err != nil {
				return ClusterCRs{}, microerror.Mask(err)
			}

			config.ControlPlaneName = generatedName
		}
	}

	awsClusterCR := newAWSClusterCR(config)
	clusterCR := newClusterCR(awsClusterCR, config)
	awsControlPlaneCR := newAWSControlPlaneCR(config)
	g8sControlPlaneCR := newG8sControlPlaneCR(awsControlPlaneCR, config)

	crs := ClusterCRs{
		Cluster:         clusterCR,
		AWSCluster:      awsClusterCR,
		G8sControlPlane: g8sControlPlaneCR,
		AWSControlPlane: awsControlPlaneCR,
	}

	return crs, nil
}

func newAWSClusterCR(c ClusterCRsConfig) *v1alpha3.AWSCluster {
	awsClusterCR := &v1alpha3.AWSCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       kindAWSCluster,
			APIVersion: v1alpha3.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.ClusterName,
			Namespace: metav1.NamespaceDefault,
			Annotations: map[string]string{
				annotation.Docs: "https://docs.giantswarm.io/ui-api/management-api/crd/awsclusters.infrastructure.giantswarm.io/",
			},
			Labels: map[string]string{
				label.AWSOperatorVersion:     c.ReleaseComponents["aws-operator"],
				label.Cluster:                c.ClusterName,
				label.Organization:           c.Owner,
				label.ReleaseVersion:         c.ReleaseVersion,
				apiv1alpha3.ClusterLabelName: c.ClusterName,
			},
		},
		Spec: v1alpha3.AWSClusterSpec{
			Cluster: v1alpha3.AWSClusterSpecCluster{
				Description: c.Description,
				DNS: v1alpha3.AWSClusterSpecClusterDNS{
					Domain: c.Domain,
				},
				OIDC: v1alpha3.AWSClusterSpecClusterOIDC{},
			},
			Provider: v1alpha3.AWSClusterSpecProvider{
				CredentialSecret: v1alpha3.AWSClusterSpecProviderCredentialSecret{
					Name:      c.Credential,
					Namespace: "giantswarm",
				},
				Pods: v1alpha3.AWSClusterSpecProviderPods{
					CIDRBlock:    c.PodsCIDR,
					ExternalSNAT: &c.ExternalSNAT,
				},
				Nodes: v1alpha3.AWSClusterSpecProviderNodes{
					NetworkPool: c.NetworkPool,
				},
				Region: c.Region,
			},
		},
	}

	// Single master node
	if len(c.ControlPlaneAZ) == 1 {
		awsClusterCR.Spec.Provider.Master = v1alpha3.AWSClusterSpecProviderMaster{
			AvailabilityZone: c.ControlPlaneAZ[0],
			InstanceType:     defaultMasterInstanceType,
		}
	}

	return awsClusterCR
}

func newAWSControlPlaneCR(c ClusterCRsConfig) *v1alpha3.AWSControlPlane {
	return &v1alpha3.AWSControlPlane{
		TypeMeta: metav1.TypeMeta{
			Kind:       kindAWSControlPlane,
			APIVersion: v1alpha3.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.ControlPlaneName,
			Namespace: metav1.NamespaceDefault,
			Annotations: map[string]string{
				annotation.Docs: "https://docs.giantswarm.io/ui-api/management-api/crd/awscontrolplanes.infrastructure.giantswarm.io/",
			},
			Labels: map[string]string{
				label.AWSOperatorVersion:     c.ReleaseComponents["aws-operator"],
				label.Cluster:                c.ClusterName,
				label.ControlPlane:           c.ControlPlaneName,
				label.Organization:           c.Owner,
				label.ReleaseVersion:         c.ReleaseVersion,
				apiv1alpha3.ClusterLabelName: c.ClusterName,
			},
		},
		Spec: v1alpha3.AWSControlPlaneSpec{
			AvailabilityZones: c.ControlPlaneAZ,
			InstanceType:      defaultMasterInstanceType,
		},
	}
}

func newClusterCR(obj *v1alpha3.AWSCluster, c ClusterCRsConfig) *apiv1alpha3.Cluster {
	clusterLabels := map[string]string{}
	{
		for key, value := range c.Labels {
			clusterLabels[key] = value
		}

		gsLabels := map[string]string{
			label.ClusterOperatorVersion: c.ReleaseComponents["cluster-operator"],
			label.Cluster:                c.ClusterName,
			apiv1alpha3.ClusterLabelName: c.ClusterName,
			label.Organization:           c.Owner,
			label.ReleaseVersion:         c.ReleaseVersion,
		}

		for key, value := range gsLabels {
			clusterLabels[key] = value
		}
	}

	clusterCR := &apiv1alpha3.Cluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Cluster",
			APIVersion: "cluster.x-k8s.io/v1alpha3",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.ClusterName,
			Namespace: metav1.NamespaceDefault,
			Annotations: map[string]string{
				annotation.Docs: "https://docs.giantswarm.io/ui-api/management-api/crd/clusters.cluster.x-k8s.io/",
			},
			Labels: clusterLabels,
		},
		Spec: apiv1alpha3.ClusterSpec{
			InfrastructureRef: &corev1.ObjectReference{
				APIVersion: obj.TypeMeta.APIVersion,
				Kind:       obj.TypeMeta.Kind,
				Name:       obj.GetName(),
				Namespace:  obj.GetNamespace(),
			},
		},
	}

	return clusterCR
}

func newG8sControlPlaneCR(obj *v1alpha3.AWSControlPlane, c ClusterCRsConfig) *v1alpha3.G8sControlPlane {
	return &v1alpha3.G8sControlPlane{
		TypeMeta: metav1.TypeMeta{
			Kind:       kindG8sControlPlane,
			APIVersion: v1alpha3.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.ControlPlaneName,
			Namespace: metav1.NamespaceDefault,
			Annotations: map[string]string{
				annotation.Docs: "https://docs.giantswarm.io/ui-api/management-api/crd/g8scontrolplanes.infrastructure.giantswarm.io/",
			},
			Labels: map[string]string{
				label.ClusterOperatorVersion: c.ReleaseComponents["cluster-operator"],
				label.Cluster:                c.ClusterName,
				label.ControlPlane:           c.ControlPlaneName,
				label.Organization:           c.Owner,
				label.ReleaseVersion:         c.ReleaseVersion,
				apiv1alpha3.ClusterLabelName: c.ClusterName,
			},
		},
		Spec: v1alpha3.G8sControlPlaneSpec{
			Replicas: len(c.ControlPlaneAZ),
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: obj.TypeMeta.APIVersion,
				Kind:       obj.TypeMeta.Kind,
				Name:       obj.GetName(),
				Namespace:  obj.GetNamespace(),
			},
		},
	}
}

func GenerateName(enableLongNames bool) (string, error) {
	for {
		letterRunes := []rune(NameChars)
		length := NameLengthShort
		if enableLongNames {
			length = NameLengthLong
		}
		characters := make([]rune, length)
		rand.Seed(time.Now().UnixNano())
		for i := range characters {
			characters[i] = letterRunes[rand.Intn(len(letterRunes))] //nolint:gosec
		}

		generatedName := string(characters)

		if valid, err := ValidateName(generatedName, enableLongNames); err != nil {
			return "", microerror.Mask(err)
		} else if !valid {
			continue
		}

		return generatedName, nil
	}
}

func ValidateName(name string, enableLongNames bool) (bool, error) {
	maxLength := NameLengthShort
	if enableLongNames {
		maxLength = NameLengthLong
	}

	pattern := fmt.Sprintf("^[a-z][a-z0-9]{0,%d}$", maxLength-1)
	matched, err := regexp.MatchString(pattern, name)
	return matched, microerror.Mask(err)
}
