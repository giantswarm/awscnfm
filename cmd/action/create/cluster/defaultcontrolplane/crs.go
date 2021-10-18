package defaultcontrolplane

import (
	"github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/apiextensions/v3/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v15/pkg/key"
	"github.com/giantswarm/awscnfm/v15/pkg/release"
)

func (r *runner) newCRs(releases []v1alpha1.Release, host string) (v1alpha3.ClusterCRs, error) {
	var err error

	var re *release.Release
	{
		c := release.Config{
			FromEnv:  r.flag.ReleaseVersion,
			Releases: releases,
		}

		re, err = release.New(c)
		if err != nil {
			return v1alpha3.ClusterCRs{}, microerror.Mask(err)
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
		}

		crs, err = v1alpha3.NewClusterCRs(c)
		if err != nil {
			return v1alpha3.ClusterCRs{}, microerror.Mask(err)
		}

		if key.IsOrgNamespaceVersion(c.ReleaseVersion) {
			crs = moveCRsToOrgNamespace(crs, key.Organization)
		}
	}

	return crs, nil
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
