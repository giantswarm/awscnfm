package key

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Key_DomainFromHost(t *testing.T) {
	testCases := []struct {
		name           string
		host           string
		expectedDomain string
	}{
		{
			name:           "case 0",
			host:           "https://g8s.ginger.eu-west-1.aws.gigantic.io:443",
			expectedDomain: "ginger.eu-west-1.aws.gigantic.io",
		},
		{
			name:           "case 1",
			host:           "https://g8s.gauss.eu-central-1.aws.gigantic.io:443",
			expectedDomain: "gauss.eu-central-1.aws.gigantic.io",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			domain := DomainFromHost(tc.host)

			if !cmp.Equal(domain, tc.expectedDomain) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedDomain, domain))
			}
		})
	}
}

func Test_Key_RegionFromHost(t *testing.T) {
	testCases := []struct {
		name           string
		host           string
		expectedRegion string
	}{
		{
			name:           "case 0",
			host:           "https://g8s.ginger.eu-west-1.aws.gigantic.io:443",
			expectedRegion: "eu-west-1",
		},
		{
			name:           "case 1",
			host:           "https://g8s.gauss.eu-central-1.aws.gigantic.io:443",
			expectedRegion: "eu-central-1",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			domain := RegionFromHost(tc.host)

			if !cmp.Equal(domain, tc.expectedRegion) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedRegion, domain))
			}
		})
	}
}
