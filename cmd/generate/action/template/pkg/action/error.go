package action

import (
	"github.com/giantswarm/awscnfm/v12/pkg/key"
)

var ErrorBase = key.GeneratedWithPrefix("error.go")

var ErrorContent = `package {{ .Action }}

import (
	"github.com/giantswarm/microerror"
)

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}
`
