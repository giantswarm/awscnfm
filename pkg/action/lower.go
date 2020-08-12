package action

import (
	"strconv"
	"strings"
)

func Lower(a string, b string) bool {
	na, err := strconv.Atoi(strings.Replace(a, "ac", "", 1))
	if err != nil {
		panic(err)
	}
	nb, err := strconv.Atoi(strings.Replace(b, "ac", "", 1))
	if err != nil {
		panic(err)
	}

	return na < nb
}
