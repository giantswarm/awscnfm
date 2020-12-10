package defaultdataplane

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/jsonmaur/aws-regions/go/regions"

	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/release"
)

func newCRs(ctx context.Context, releases []v1alpha1.Release, id string, host string) (v1alpha2.NodePoolCRs, error) {
	var err error

	var re *release.Release
	{
		c := release.Config{
			FromEnv:  env.CreateReleaseVersion(),
			Releases: releases,
		}

		re, err = release.New(c)
		if err != nil {
			return v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}
	}

	var azs []string
	{
		region, err := regions.LookupByCode(key.RegionFromHost(host))
		if err != nil {
			return v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}

		azs = region.Zones
	}

	var crs v1alpha2.NodePoolCRs
	{
		c := v1alpha2.NodePoolCRsConfig{
			AvailabilityZones:                   []string{azs[0]},
			AWSInstanceType:                     "m5.xlarge",
			ClusterID:                           id,
			Description:                         "awscnfm cl001 ac005",
			NodesMax:                            2,
			NodesMin:                            1,
			OnDemandBaseCapacity:                0,
			OnDemandPercentageAboveBaseCapacity: 0,
			Owner:                               "giantswarm",
			ReleaseComponents:                   re.Components(),
			ReleaseVersion:                      re.Version(),
			UseAlikeInstanceTypes:               true,
		}

		crs, err = v1alpha2.NewNodePoolCRs(c)
		if err != nil {
			return v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}
	}

	return crs, nil
}
