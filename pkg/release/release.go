package release

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
)

type Config struct {
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

type Release struct {
	fromEnv     string
	fromProject string
	releases    []v1alpha1.Release
}

func New(config Config) (*Release, error) {
	if config.FromProject == "" && config.FromEnv == "" {
		return nil, microerror.Maskf(invalidConfigError, "either %T.FromProject or %T.FromEnv must be given", config, config)
	}
	if len(config.Releases) == 0 {
		return nil, microerror.Maskf(invalidConfigError, "%T.Releases must not be empty", config)
	}

	r := &Release{
		fromEnv:     config.FromEnv,
		fromProject: config.FromProject,
		releases:    config.Releases,
	}

	return r, nil
}

func (r *Release) Components() map[string]string {
	// Collecting the components of the release we found based on the input
	// configuration.
	releaseComponents := map[string]string{}
	{
		release := mustFind(r.fromEnv, r.fromProject, r.releases)

		for _, c := range release.Spec.Components {
			releaseComponents[c.Name] = c.Version
		}
	}

	return releaseComponents
}

func (r *Release) Version() string {
	release := mustFind(r.fromEnv, r.fromProject, r.releases)
	return release.GetName()
}

func mustFind(fromEnv string, fromProject string, releases []v1alpha1.Release) v1alpha1.Release {
	var version string
	{
		if fromEnv == "" {
			version = fromProject
		} else {
			version = fromEnv
		}

		if !strings.HasPrefix(version, "v") {
			version = fmt.Sprintf("v%s", version)
		}
	}

	var release v1alpha1.Release
	{
		// First we check if we have a direct match. This might be the case if
		// somebody specifies the exact test release they want to test for
		// conformity.
		//
		//     v100.0.0-xh3b4sd
		//     v24.6.8-dev
		//
		for _, o := range releases {
			if o.GetName() == version {
				release = o
				break
			}
		}

		// We might not have an exact match. Then we want to check for the
		// latest release that aligns with our major and minor version. Such a
		// scenario might be if somebody wants to test conformity of a release
		// we want to publish. Note that in case of a test release we fall back
		// to the exact match again. The fuzzy search on using the latest patch
		// release does only work with published releases.
		//
		//     v13.4.3
		//     v18.6.8
		//
		rv := mustToSemver(version)

		for _, o := range releases {
			ov := mustToSemver(o.GetName())

			// Major and minor must match. We ignore the rest.
			if rv.Major != ov.Major {
				continue
			}
			if rv.Minor != ov.Minor {
				continue
			}
			if rv.PreRelease != ov.PreRelease {
				continue
			}

			if release.GetName() == "" {
				// Remembering the first release we find. We might find another
				// one with a bigger patch version during the next iterations.
				release = o
			} else {
				pv := mustToSemver(release.GetName())

				if ov.Patch > pv.Patch {
					// We found a release with a bigger patch version than we
					// already kept track of.
					release = o
				}
			}
		}
	}

	return release
}

func mustToSemver(s string) *semver.Version {
	v, err := semver.NewVersion(strings.Replace(s, "v", "", 1))
	if err != nil {
		panic(err)
	}

	return v
}
