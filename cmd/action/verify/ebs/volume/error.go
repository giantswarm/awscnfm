package volume

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

// jobNotCompleted is an error indicating that the job in EBS volume test is not jet finished
var jobNotCompleted = &microerror.Error{
	Kind: "jobNotCompleted",
	Desc: "EBS volume job for testing CSI driver is not completed",
}

// IsJobNotCompleted asserts jobNotCompleted error.
func IsJobNotCompleted(err error) bool {
	return microerror.Cause(err) == jobNotCompleted
}
