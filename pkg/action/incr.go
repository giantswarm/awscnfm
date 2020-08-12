package action

import (
	"fmt"
	"strconv"
	"strings"
)

func Incr(a string) string {
	na, err := strconv.Atoi(strings.Replace(a, "ac", "", 1))
	if err != nil {
		panic(err)
	}

	na++

	return fmt.Sprintf("ac%03d", na)
}
