package generate

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const (
	// idChars represents the character set used to generate cluster IDs.
	// (does not contain 1 and l, to avoid confusion)
	idChars = "023456789abcdefghijkmnopqrstuvwxyz"

	// idLength represents the number of characters used to create a cluster ID.
	idLength = 5
)

var (
	// Use local instance of RNG. Can be overwritten with fixed seed in tests
	// if needed.
	localRng = rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404
)

func ID() string {
	pattern := regexp.MustCompile("^[a-z]+$")
	for {
		letterRunes := []rune(idChars)
		b := make([]rune, idLength)
		for i := range b {
			b[i] = letterRunes[localRng.Intn(len(letterRunes))]
		}

		id := string(b)

		if _, err := strconv.Atoi(id); err == nil {
			// string is numbers only, which we want to avoid
			continue
		}

		if pattern.MatchString(id) {
			// strings is letters only, which we also avoid
			continue
		}

		return id
	}
}
