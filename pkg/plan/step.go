package plan

import "strings"

type Step struct {
	// Action is the action/command name, e.g. ac013.
	Action StepAction
	// Backoff is the tolerance within the associated action must be
	// successfully executed without returning errors.
	Backoff *Backoff
}

type StepAction string

func (a StepAction) Split() []string {
	return strings.Split(string(a), "/")
}

func (a StepAction) String() string {
	return strings.ReplaceAll(string(a), "/", " ")
}
