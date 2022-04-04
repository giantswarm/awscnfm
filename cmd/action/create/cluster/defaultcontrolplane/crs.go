package defaultcontrolplane

import (
	"github.com/giantswarm/apiextensions/v6/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/release-operator/v3/api/v1alpha1"

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
			crs = key.MoveClusterCRsToOrgNamespace(crs, key.Organization)
		}
	}

	return crs, nil
}
