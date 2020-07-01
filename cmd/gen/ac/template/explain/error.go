package explain

import (
	"path/filepath"

	"github.com/giantswarm/awscnfm/pkg/key"
)

var ErrorBase = filepath.Join("explain", key.GeneratedWithPrefix("error.go"))

var ErrorContent = `package explain

import "github.com/giantswarm/microerror"

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidFlagsError = &microerror.Error{
	Kind: "invalidFlagsError",
}

// IsInvalidFlags asserts invalidFlagsError.
func IsInvalidFlags(err error) bool {
	return microerror.Cause(err) == invalidFlagsError
}
`
