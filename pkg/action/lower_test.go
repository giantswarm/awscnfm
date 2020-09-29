package action

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Action_Lower(t *testing.T) {
	testCases := []struct {
		name   string
		aca    string
		acb    string
		result bool
	}{
		{
			name:   "case 0",
			aca:    "ac000",
			acb:    "ac000",
			result: false,
		},
		{
			name:   "case 1",
			aca:    "ac001",
			acb:    "ac000",
			result: false,
		},
		{
			name:   "case 2",
			aca:    "ac001",
			acb:    "ac001",
			result: false,
		},
		{
			name:   "case 3",
			aca:    "ac012",
			acb:    "ac012",
			result: false,
		},
		{
			name:   "case 4",
			aca:    "ac015",
			acb:    "ac012",
			result: false,
		},
		{
			name:   "case 5",
			aca:    "ac017",
			acb:    "ac017",
			result: false,
		},
		{
			name:   "case 6",
			aca:    "ac170",
			acb:    "ac170",
			result: false,
		},
		{
			name:   "case 7",
			aca:    "ac179",
			acb:    "ac179",
			result: false,
		},
		{
			name:   "case 8",
			aca:    "ac109",
			acb:    "ac179",
			result: true,
		},
		{
			name:   "case 9",
			aca:    "ac011",
			acb:    "ac029",
			result: true,
		},
		{
			name:   "case 10",
			aca:    "ac020",
			acb:    "ac021",
			result: true,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := Lower(tc.aca, tc.acb)

			if !cmp.Equal(result, tc.result) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.result, result))
			}
		})
	}
}
