package release

import (
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
)

type Config struct {
	// FromEnv is the release version we get from the environment. This might be
	// a custom release version created for testing, e.g. 100.0.0-xh3b4sd.
	FromEnv string
	// Releases are all releases to draw from using the configured release
	// version. See also FromEnv and FromProject.
	Releases []v1alpha1.Release
}

type Release struct {
	fromEnv  string
	releases []v1alpha1.Release
}

func New(config Config) (*Release, error) {
	if config.FromEnv == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.FromEnv must not be empty", config)
	}
	if len(config.Releases) == 0 {
		return nil, microerror.Maskf(invalidConfigError, "%T.Releases must not be empty", config)
	}

	r := &Release{
		fromEnv:  config.FromEnv,
		releases: config.Releases,
	}

	return r, nil
}

func (r *Release) Components() map[string]string {
	m := map[string]string{}
	for _, r := range r.releases {
		for _, c := range r.Spec.Components {
			m[c.Name] = c.Version
		}
	}

	return m
}

func (r *Release) Version() string {
	return r.fromEnv
}
