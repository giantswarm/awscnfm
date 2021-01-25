package plan

import "time"

type Step struct {
	// Action is the action/command name, e.g. ac013.
	Action StepAction
	// Backoff is the tolerance within the associated action must be
	// successfully executed without returning errors.
	Backoff *Backoff
	// CoolDown is the time to wait after the action was executed, before moving
	// on to execute the following action. This is a delay mechanism in order to
	// defer the execution of actions within a test plan.
	CoolDown time.Duration
	// Condition defines when to execute this step of a plan. By default all
	// steps are executed until an error occurs. This behaviour can be
	// overwritten to always execute certain steps e.g. for cluster deletion.
	Condition StepCondition
}
