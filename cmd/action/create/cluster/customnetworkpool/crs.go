package customnetworkpool

import (
	"github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/apiextensions/v3/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v15/pkg/key"
	"github.com/giantswarm/awscnfm/v15/pkg/release"
)

func (r *runner) newCRs(releases []v1alpha1.Release, host string) (v1alpha3.ClusterCRs, v1alpha3.NetworkPoolCRs, error) {
	var err error

	var re *release.Release
	{
		c := release.Config{
			FromEnv:  r.flag.ReleaseVersion,
			Releases: releases,
		}

		re, err = release.New(c)
		if err != nil {
			return v1alpha3.ClusterCRs{}, v1alpha3.NetworkPoolCRs{}, microerror.Mask(err)
		}
	}

	var crs v1alpha3.ClusterCRs
	{
		c := v1alpha3.ClusterCRsConfig{
			ClusterID:         r.flag.TenantCluster,
			Credential:        key.Credential,
			Domain:            key.DomainFromHost(host),
			Description:       "awscnfm action create cluster onenodepool",
			Owner:             key.Organization,
			Region:            key.RegionFromHost(host),
			ReleaseComponents: re.Components(),
			ReleaseVersion:    re.Version(),
			NetworkPool:       "custom",
			PodsCIDR:          "192.168.0.0/18",
		}

		crs, err = v1alpha3.NewClusterCRs(c)
		if err != nil {
			return v1alpha3.ClusterCRs{}, v1alpha3.NetworkPoolCRs{}, microerror.Mask(err)
		}

		if key.IsOrgNamespaceVersion(c.ReleaseVersion) {
			crs = moveCRsToOrgNamespace(crs, key.Organization)
		}

	}

	var npcrs v1alpha3.NetworkPoolCRs
	{
		npcrs, err = v1alpha3.NewNetworkPoolCRs(v1alpha3.NetworkPoolCRsConfig{
			CIDRBlock:     "192.168.64.0/18",
			NetworkPoolID: "custom",
			Owner:         "giantswarm",
		})
		if err != nil {
			return v1alpha3.ClusterCRs{}, v1alpha3.NetworkPoolCRs{}, microerror.Mask(err)
		}
	}

	return crs, npcrs, nil
}

func moveCRsToOrgNamespace(crs v1alpha3.ClusterCRs, organization string) v1alpha3.ClusterCRs {
	crs.Cluster.SetNamespace(key.OrganizationNamespaceFromName(organization))
	crs.Cluster.Spec.InfrastructureRef.Namespace = key.OrganizationNamespaceFromName(organization)
	crs.AWSCluster.SetNamespace(key.OrganizationNamespaceFromName(organization))
	crs.G8sControlPlane.SetNamespace(key.OrganizationNamespaceFromName(organization))
	crs.G8sControlPlane.Spec.InfrastructureRef.Namespace = key.OrganizationNamespaceFromName(organization)
	crs.AWSControlPlane.SetNamespace(key.OrganizationNamespaceFromName(organization))
	return crs
}
