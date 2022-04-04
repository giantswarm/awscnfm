package singlecontrolplane

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/release-operator/v3/api/v1alpha1"

	"github.com/giantswarm/awscnfm/v15/cmd/action/create/cluster/util"
	"github.com/giantswarm/awscnfm/v15/pkg/key"
	"github.com/giantswarm/awscnfm/v15/pkg/release"
)

func (r *runner) newCRs(releases []v1alpha1.Release, host string) (util.ClusterCRs, error) {
	var err error

	var re *release.Release
	{
		c := release.Config{
			FromEnv:  r.flag.ReleaseVersion,
			Releases: releases,
		}

		re, err = release.New(c)
		if err != nil {
			return util.ClusterCRs{}, microerror.Mask(err)
		}
	}

	var crs util.ClusterCRs
	{
		c := util.ClusterCRsConfig{
			ClusterID:         r.flag.TenantCluster,
			Domain:            key.DomainFromHost(host),
			Description:       "awscnfm action create cluster single-master onenodepool",
			Owner:             key.Organization,
			Region:            key.RegionFromHost(host),
			MasterAZ:          []string{key.RegionFromHost(host) + "a"},
			ReleaseComponents: re.Components(),
			ReleaseVersion:    re.Version(),
		}

		crs, err = util.NewClusterCRs(c)
		if err != nil {
			return util.ClusterCRs{}, microerror.Mask(err)
		}

		if key.IsOrgNamespaceVersion(c.ReleaseVersion) {
			crs = key.MoveClusterCRsToOrgNamespace(crs, key.Organization)
		}
	}

	return crs, nil
}
