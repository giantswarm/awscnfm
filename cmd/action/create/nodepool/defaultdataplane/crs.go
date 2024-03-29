package defaultdataplane

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/release-operator/v3/api/v1alpha1"
	regions "github.com/jsonmaur/aws-regions/v2"

	"github.com/giantswarm/awscnfm/v15/cmd/action/create/cluster/util"
	"github.com/giantswarm/awscnfm/v15/pkg/key"
	"github.com/giantswarm/awscnfm/v15/pkg/release"
)

func (r *runner) newCRs(releases []v1alpha1.Release, host string) (util.NodePoolCRs, error) {
	var err error

	var re *release.Release
	{
		c := release.Config{
			FromEnv:  r.flag.ReleaseVersion,
			Releases: releases,
		}

		re, err = release.New(c)
		if err != nil {
			return util.NodePoolCRs{}, microerror.Mask(err)
		}
	}

	var azs []string
	{
		region, err := regions.LookupByCode(key.RegionFromHost(host))
		if err != nil {
			return util.NodePoolCRs{}, microerror.Mask(err)
		}

		azs = region.Zones
	}

	var crs util.NodePoolCRs
	{
		c := util.NodePoolCRsConfig{
			AvailabilityZones:                   []string{azs[0]},
			AWSInstanceType:                     "m5.xlarge",
			ClusterID:                           r.flag.TenantCluster,
			Description:                         "awscnfm cl001 ac005",
			NodesMax:                            2,
			NodesMin:                            1,
			OnDemandBaseCapacity:                0,
			OnDemandPercentageAboveBaseCapacity: 0,
			Owner:                               key.Organization,
			ReleaseComponents:                   re.Components(),
			ReleaseVersion:                      re.Version(),
			UseAlikeInstanceTypes:               true,
		}

		crs, err = util.NewNodePoolCRs(c)
		if err != nil {
			return util.NodePoolCRs{}, microerror.Mask(err)
		}

		if key.IsOrgNamespaceVersion(c.ReleaseVersion) {
			crs = key.MoveNodePoolCRsToOrgNamespace(crs, key.Organization)
		}
	}

	return crs, nil
}
