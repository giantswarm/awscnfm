package plan

type Step struct {
	// Action is the action/command name, e.g. ac013.
	Action StepAction
	// Backoff is the tolerance within the associated action must be
	// successfully executed without returning errors.
	Backoff *Backoff
	// Condition defines when to execute this step of a plan. By default all
	// steps are executed until an error occurs. This behaviour can be
	// overwritten to always execute certain steps e.g. for cluster deletion.
	Condition StepCondition
}
