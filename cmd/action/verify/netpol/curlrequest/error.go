package curlrequest

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

// jobNotCompleted is an error indicating that the job in netpol test is not yet finished
var jobNotCompleted = &microerror.Error{
	Kind: "jobNotCompleted",
	Desc: "curl request job for testing netpol is not successfully completed",
}

// IsJobNotCompleted asserts jobNotCompleted error.
func IsJobNotCompleted(err error) bool {
	return microerror.Cause(err) == jobNotCompleted
}

// jobNotFailed is an error indicating that the job in netpol test have not failed even tho its expected
var jobNotFailed = &microerror.Error{
	Kind: "jobNotFailed",
	Desc: "curl request job for testing netpol is not failed",
}

// IsJobNotFailed asserts jobNotCompleted error.
func IsJobNotFailed(err error) bool {
	return microerror.Cause(err) == jobNotFailed
}
