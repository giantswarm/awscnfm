package ac001

import (
	"github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
	"github.com/giantswarm/awscnfm/v12/pkg/release"
)

func newCRs(releases []v1alpha1.Release, host string) (v1alpha2.ClusterCRs, v1alpha2.NetworkPoolCRs, error) {
	var err error

	var p *release.Patch
	{
		c := release.PatchConfig{
			FromEnv:     env.ReleaseVersion(),
			FromProject: project.Version(),
			Releases:    releases,
		}

		p, err = release.NewPatch(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NetworkPoolCRs{}, microerror.Mask(err)
		}
	}

	var crs v1alpha2.ClusterCRs
	{
		c := v1alpha2.ClusterCRsConfig{
			Credential:        key.Credential,
			Domain:            key.DomainFromHost(host),
			Description:       explainerCommand,
			Owner:             key.Organization,
			Region:            key.RegionFromHost(host),
			ReleaseComponents: p.Components().Latest(),
			ReleaseVersion:    p.Version().Latest(),
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
