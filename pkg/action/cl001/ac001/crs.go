package ac001

import (
	"github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v12/pkg/key"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
	"github.com/giantswarm/awscnfm/v12/pkg/release"
)

func newCRs(host string) (v1alpha2.ClusterCRs, error) {
	var err error

	var releaseComponents map[string]string
	{
		c := release.Config{}

		releaseCollection, err := release.New(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}

		releaseComponents = releaseCollection.ReleaseComponents(project.Version())
	}

	var crs v1alpha2.ClusterCRs
	{
		c := v1alpha2.ClusterCRsConfig{
			Credential:        key.Credential,
			Domain:            key.DomainFromHost(host),
			Description:       explainerCommand,
			Owner:             key.Organization,
			Region:            key.RegionFromHost(host),
			ReleaseComponents: releaseComponents,
			ReleaseVersion:    project.Version(),
		}

		crs, err = v1alpha2.NewClusterCRs(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}
	}

	return crs, nil
}
