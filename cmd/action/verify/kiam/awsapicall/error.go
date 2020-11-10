package awsapicall

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

// jobNotCompleted is an error indicating that the job in kiam test is not jet finished
var jobNotCompleted = &microerror.Error{
	Kind: "jobNotCompleted",
	Desc: "AWS API call job for testing kiam is not completed",
}

// IsJobNotCompleted asserts jobNotCompleted error.
func IsJobNotCompleted(err error) bool {
	return microerror.Cause(err) == jobNotCompleted
}
