package ac001

import (
	"fmt"

	"github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/pkg/key"
	"github.com/giantswarm/awscnfm/pkg/release"
)

func newCRs(host string) (v1alpha2.ClusterCRs, v1alpha2.NodePoolCRs, error) {
	var err error

	var releaseComponents map[string]string
	{
		c := release.Config{
			Branch: "aws-release-12-0-0-2",
		}

		releaseCollection, err := release.New(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}

		releaseComponents = releaseCollection.ReleaseComponents(key.Release)
	}

	var crs v1alpha2.ClusterCRs
	{
		c := v1alpha2.ClusterCRsConfig{
			Domain:            key.DomainFromHost(host),
			Description:       explainerCommand,
			Owner:             key.Organization,
			Region:            key.RegionFromHost(host),
			ReleaseComponents: releaseComponents,
			ReleaseVersion:    key.Release,
		}

		crs, err = v1alpha2.NewClusterCRs(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}
	}

	var azs []string
	{
		fmt.Println(host)
		azs, err = key.AZsFromRegion(key.RegionFromHost(host))
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}
	}

	var npCrs v1alpha2.NodePoolCRs
	{
		c := v1alpha2.NodePoolCRsConfig{
			AvailabilityZones:                   []string{azs[0]},
			AWSInstanceType:                     "m5.xlarge",
			ClusterID:                           crs.AWSCluster.GetName(),
			Description:                         nodePoolDescription,
			NodesMax:                            2,
			NodesMin:                            1,
			OnDemandBaseCapacity:                0,
			OnDemandPercentageAboveBaseCapacity: 0,
			Owner:                               "giantswarm",
			ReleaseComponents:                   releaseComponents,
			ReleaseVersion:                      key.Release,
			UseAlikeInstanceTypes:               true,
		}
		npCrs, err = v1alpha2.NewNodePoolCRs(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}
	}

	return crs, npCrs, nil
}
