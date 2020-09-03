package release

import (
	"strconv"
	"strings"
	"testing"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_Release_mustFindPatch(t *testing.T) {
	testCases := []struct {
		name             string
		version          string
		release          v1alpha1.Release
		releases         []v1alpha1.Release
		expectedPrevious v1alpha1.Release
		expectedLatest   v1alpha1.Release
	}{
		{
			name:    "case 0",
			version: "v12.0.0",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
			},
			expectedPrevious: v1alpha1.Release{},
			expectedLatest: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"},
			},
		},
		{
			name:    "case 1",
			version: "v12.0.0",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.1"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.2"}},
			},
			expectedPrevious: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.1"},
			},
			expectedLatest: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.2"},
			},
		},
		{
			name:    "case 2",
			version: "v12.0.0",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.2.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedPrevious: v1alpha1.Release{},
			expectedLatest: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"},
			},
		},
		{
			name:    "case 3",
			version: "v12.0.0",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedPrevious: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"},
			},
			expectedLatest: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"},
			},
		},
		{
			name:    "case 4",
			version: "v12.0.9",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedPrevious: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"},
			},
			expectedLatest: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"},
			},
		},
		{
			name:    "case 5",
			version: "v12.0.0-dev",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedPrevious: v1alpha1.Release{},
			expectedLatest: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"},
			},
		},
		{
			name:    "case 6",
			version: "v12.0.0-dev",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.0-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.3-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedPrevious: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"},
			},
			expectedLatest: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.3-dev"},
			},
		},
		{
			name:    "case 7",
			version: "v100.0.0-xh3b4sd",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v100.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v100.0.0-xh3b4sd"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v100.0.3"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedPrevious: v1alpha1.Release{},
			expectedLatest: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v100.0.0-xh3b4sd"},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var p Resolver
			{
				c := PatchConfig{
					FromProject: tc.version,
					Releases:    tc.releases,
				}

				p, err = NewPatch(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			v := p.Version()

			expectedPrevious := strings.Replace(tc.expectedPrevious.GetName(), "v", "", 1)
			expectedLatest := strings.Replace(tc.expectedLatest.GetName(), "v", "", 1)

			if !cmp.Equal(v.Previous(), expectedPrevious) {
				t.Fatalf("\n\n%s\n", cmp.Diff(expectedPrevious, v.Previous()))
			}
			if !cmp.Equal(v.Latest(), expectedLatest) {
				t.Fatalf("\n\n%s\n", cmp.Diff(expectedLatest, v.Latest()))
			}
		})
	}
}
