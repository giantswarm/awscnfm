package ac001

import (
	"fmt"
	"strings"

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

	var r *release.Release
	{
		c := release.Config{
			FromEnv:     env.ReleaseVersion(),
			FromProject: project.Version(),
			Releases:    releases,
		}

		r, err = release.New(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}
	}

	var crs v1alpha2.ClusterCRs
	fmt.Println(r.Version())
	{
		c := v1alpha2.ClusterCRsConfig{
			Credential:        key.Credential,
			Domain:            key.DomainFromHost(host),
			Description:       explainerCommand,
			Owner:             key.Organization,
			Region:            key.RegionFromHost(host),
			ReleaseComponents: r.Components(),
			ReleaseVersion:    strings.Replace(r.Version(), "v", "", -1),
		}

		crs, err = v1alpha2.NewClusterCRs(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}
	}

	return crs, nil
}
