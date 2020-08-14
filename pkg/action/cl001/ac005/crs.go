package ac005

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/microerror"
	"github.com/jsonmaur/aws-regions/go/regions"

	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
	"github.com/giantswarm/awscnfm/v12/pkg/release"
)

func newCRs(ctx context.Context, id string, host string) (v1alpha2.NodePoolCRs, error) {
	var err error

	var releaseComponents map[string]string
	{
		c := release.Config{}

		releaseCollection, err := release.New(c)
		if err != nil {
			return v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}

		releaseComponents = releaseCollection.ReleaseComponents(project.Version())
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
			Description:                         "awscnfm cl001 ac005 explain",
			NodesMax:                            2,
			NodesMin:                            1,
			OnDemandBaseCapacity:                0,
			OnDemandPercentageAboveBaseCapacity: 0,
			Owner:                               "giantswarm",
			ReleaseComponents:                   releaseComponents,
			ReleaseVersion:                      project.Version(),
			UseAlikeInstanceTypes:               true,
		}

		crs, err = v1alpha2.NewNodePoolCRs(c)
		if err != nil {
			return v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}
	}

	return crs, nil
}
