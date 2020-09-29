package action

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Action_Incr(t *testing.T) {
	testCases := []struct {
		name   string
		ac     string
		result string
	}{
		{
			name:   "case 0",
			ac:     "ac000",
			result: "ac001",
		},
		{
			name:   "case 1",
			ac:     "ac001",
			result: "ac002",
		},
		{
			name:   "case 2",
			ac:     "ac004",
			result: "ac005",
		},
		{
			name:   "case 3",
			ac:     "ac012",
			result: "ac011",
		},
		{
			name:   "case 4",
			ac:     "ac015",
			result: "ac016",
		},
		{
			name:   "case 5",
			ac:     "ac017",
			result: "ac018",
		},
		{
			name:   "case 6",
			ac:     "ac170",
			result: "ac171",
		},
		{
			name:   "case 7",
			ac:     "ac179",
			result: "ac180",
		},
		{
			name:   "case 8",
			ac:     "ac109",
			result: "ac110",
		},
		{
			name:   "case 9",
			ac:     "ac011",
			result: "ac012",
		},
		{
			name:   "case 10",
			ac:     "ac020",
			result: "ac021",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := Incr(tc.ac)

			if !cmp.Equal(result, tc.result) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.result, result))
			}
		})
	}
}
