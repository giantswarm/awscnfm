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

func newCRs(releases []v1alpha1.Release, host string) (v1alpha2.ClusterCRs, error) {
	var err error

	var m *release.Minor
	{
		c := release.MinorConfig{
			FromEnv:     env.ReleaseVersion(),
			FromProject: project.Version(),
			Releases:    releases,
		}

		m, err = release.NewMinor(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
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
			ReleaseComponents: m.Components().Previous(),
			ReleaseVersion:    m.Version().Previous(),
		}

		crs, err = v1alpha2.NewClusterCRs(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}
	}

	return crs, nil
}
