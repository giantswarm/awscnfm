package release

import (
	"sort"

	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
)

func mustFindLatest(version string, releases []v1alpha1.Release) v1alpha1.Release {
	vv := mustToSemver(version)

	var versions semver.Versions
	for _, r := range releases {
		rv := mustToSemver(r.GetName())

		if vv.PreRelease != rv.PreRelease {
			continue
		}

		versions = append(versions, rv)
	}

	if len(versions) == 0 {
		return v1alpha1.Release{}
	}

	sort.Sort(sort.Reverse(versions))

	return findRelease(versions[0].String(), releases)
}
