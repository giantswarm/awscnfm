package release

import (
	"strconv"
	"testing"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/google/go-cmp/cmp"
)

func Test_Release_mustFind(t *testing.T) {
	testCases := []struct {
		name            string
		fromEnv         string
		fromProject     string
		releases        []v1alpha1.Release
		expectedRelease v1alpha1.Release
	}{
		{
			name:        "case 0",
			fromEnv:     "",
			fromProject: "",
			releases: []v1alpha1.Release{
				{},
			},
			release: v1alpha1.Release{},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			release := mustFind(tc.fromEnv, tc.fromProject, tc.releases)

			if !cmp.Equal(release, tc.expectedRelease) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedRelease, release))
			}
		})
	}
}
