package release

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
)

func mustToSemver(s string) *semver.Version {
	v, err := semver.NewVersion(strings.Replace(s, "v", "", 1))
	if err != nil {
		panic(err)
	}

	return v
}

// findRelease looks for the exact release match. We use this to figure out
// which release version to consider for our conformance tests. We might have a
// direct match in case somebody specifies the exact test release they want to
// test for conformity. Examples of such direct matches would be test releases
// like shown below.
//
//     v100.0.0-xh3b4sd
//     v24.6.8-dev
//
func findRelease(version string, releases []v1alpha1.Release) v1alpha1.Release {
	for _, r := range releases {
		if r.GetName() == version {
			return r
		}
	}

	return v1alpha1.Release{}
}

func findVersion(fromEnv string, fromProject string) string {
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

	return version
}
