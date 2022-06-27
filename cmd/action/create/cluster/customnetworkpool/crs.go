package customnetworkpool

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/release-operator/v3/api/v1alpha1"

	"github.com/giantswarm/awscnfm/v15/cmd/action/create/cluster/util"
	"github.com/giantswarm/awscnfm/v15/pkg/key"
	"github.com/giantswarm/awscnfm/v15/pkg/release"
)

func (r *runner) newCRs(releases []v1alpha1.Release, host string) (util.ClusterCRs, util.NetworkPoolCRs, error) {
	var err error

	var re *release.Release
	{
		c := release.Config{
			FromEnv:  r.flag.ReleaseVersion,
			Releases: releases,
		}

		re, err = release.New(c)
		if err != nil {
			return util.ClusterCRs{}, util.NetworkPoolCRs{}, microerror.Mask(err)
		}
	}

	var crs util.ClusterCRs
	{
		c := util.ClusterCRsConfig{
			ClusterID:         r.flag.TenantCluster,
			Domain:            key.DomainFromHost(host),
			Description:       "awscnfm action create cluster onenodepool",
			Owner:             key.Organization,
			Region:            key.RegionFromHost(host),
			ReleaseComponents: re.Components(),
			ReleaseVersion:    re.Version(),
			NetworkPool:       "custom",
			PodsCIDR:          "192.168.0.0/18",
		}

		crs, err = util.NewClusterCRs(c)
		if err != nil {
			return util.ClusterCRs{}, util.NetworkPoolCRs{}, microerror.Mask(err)
		}

		if key.IsOrgNamespaceVersion(c.ReleaseVersion) {
			crs = key.MoveClusterCRsToOrgNamespace(crs, key.Organization)
		}

	}

	var npcrs util.NetworkPoolCRs
	{
		npcrs, err = util.NewNetworkPoolCRs(util.NetworkPoolCRsConfig{
			CIDRBlock:     "192.168.64.0/18",
			NetworkPoolID: "custom",
			Owner:         "giantswarm",
		})
		if err != nil {
			return util.ClusterCRs{}, util.NetworkPoolCRs{}, microerror.Mask(err)
		}
	}

	return crs, npcrs, nil
}
