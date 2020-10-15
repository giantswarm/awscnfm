package ac010

import "github.com/giantswarm/microerror"

// jobNotCompleted is an error indicating that the job in kiam test is not jet finished
var jobNotCompleted = &microerror.Error{
	Kind: "jobNotCompleted",
	Desc: "AWS API call job for testing kiam is not completed",
}

// IsJobNotCompleted asserts jobNotCompleted error.
func IsJobNotCompleted(err error) bool {
	return microerror.Cause(err) == jobNotCompleted
}
