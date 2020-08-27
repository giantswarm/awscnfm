package release

import (
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
)

type MajorConfig struct {
	// FromEnv is the release version we get from the environment. This might be
	// a custom release version created for testing, e.g. 100.0.0-xh3b4sd. If
	// this information is given it is preferred over the value given by
	// FromProject.
	FromEnv string
	// FromProject is the awscnfm project version. If only this one is given we
	// try to derive the actual release version from it which we want to use for
	// testing. The project version might be v12.0.1-dev and there might be the
	// latest v12.0.x release on the Control Plane, e.g. v12.0.16, which we then
	// use for conformance testing.
	FromProject string
	// Releases are all releases to draw from using the configured release
	// version. See also FromEnv and FromProject.
	Releases []v1alpha1.Release
}

type Major struct {
	fromEnv     string
	fromProject string
	releases    []v1alpha1.Release
}

func NewMajor(config MajorConfig) (*Major, error) {
	if config.FromProject == "" && config.FromEnv == "" {
		return nil, microerror.Maskf(invalidConfigError, "either %T.FromProject or %T.FromEnv must be given", config, config)
	}
	if len(config.Releases) == 0 {
		return nil, microerror.Maskf(invalidConfigError, "%T.Releases must not be empty", config)
	}

	m := &Major{
		fromEnv:     config.FromEnv,
		fromProject: config.FromProject,
		releases:    config.Releases,
	}

	return m, nil
}

func (m *Major) Components() map[string]string {
	// Collecting the components of the release we found based on the input
	// configuration.
	releaseComponents := map[string]string{}
	{
		for _, c := range m.release().Spec.Components {
			releaseComponents[c.Name] = c.Version
		}
	}

	return releaseComponents
}

func (m *Major) Version() string {
	return strings.Replace(m.release().GetName(), "v", "", 1)
}

func (m *Major) release() *v1alpha1.Release {
	version := findVersion(m.fromEnv, m.fromProject)

	release := findRelease(version, m.releases)
	if release.GetName() != "" {
		return &release
	}

	major := mustFindMajor(version, m.releases)
	return &major
}

func mustFindMajor(version string, releases []v1alpha1.Release) v1alpha1.Release {
	vv := mustToSemver(version)

	var versions semver.Versions
	for _, r := range releases {
		rv := mustToSemver(r.GetName())

		if vv.Major <= rv.Major {
			continue
		}
		if vv.PreRelease != rv.PreRelease {
			continue
		}

		versions = append(versions, rv)
	}

	if len(versions) == 0 {
		return v1alpha1.Release{}
	}

	sort.Sort(sort.Reverse(versions))

	return findRelease("v"+versions[0].String(), releases)
}
