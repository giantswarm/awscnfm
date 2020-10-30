package valid

import (
	"regexp"

	"github.com/giantswarm/microerror"
)

var (
	containsLetter     = regexp.MustCompile(`[a-z]`)
	containsNumber     = regexp.MustCompile(`[0-9]`)
	containsWhitespace = regexp.MustCompile(`[\s]`)
)

func ID(id string) error {
	if !containsLetter.MatchString(id) {
		return microerror.Maskf(invalidIDError, "must contain letter")
	}

	if !containsNumber.MatchString(id) {
		return microerror.Maskf(invalidIDError, "must contain number")
	}

	if containsWhitespace.MatchString(id) {
		return microerror.Maskf(invalidIDError, "must not contain whitespace")
	}

	if len(id) != 5 {
		return microerror.Maskf(invalidIDError, "must be have length of 5")
	}

	return nil
}
