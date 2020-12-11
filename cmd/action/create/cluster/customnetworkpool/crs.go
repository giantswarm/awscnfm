package customnetworkpool

import (
	"github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/release"
)

func (r *runner) newCRs(releases []v1alpha1.Release, host string) (v1alpha2.ClusterCRs, v1alpha2.NetworkPoolCRs, error) {
	var err error

	var re *release.Release
	{
		c := release.Config{
			FromEnv:  r.flag.ReleaseVersion,
			Releases: releases,
		}

		re, err = release.New(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NetworkPoolCRs{}, microerror.Mask(err)
		}
	}

	var crs v1alpha2.ClusterCRs
	{
		c := v1alpha2.ClusterCRsConfig{
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

		crs, err = v1alpha2.NewClusterCRs(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NetworkPoolCRs{}, microerror.Mask(err)
		}
	}

	var npcrs v1alpha2.NetworkPoolCRs
	{
		npcrs, err = v1alpha2.NewNetworkPoolCRs(v1alpha2.NetworkPoolCRsConfig{
			CIDRBlock:     "192.168.64.0/18",
			NetworkPoolID: "custom",
			Owner:         "giantswarm",
		})
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NetworkPoolCRs{}, microerror.Mask(err)
		}
	}

	return crs, npcrs, nil
}
