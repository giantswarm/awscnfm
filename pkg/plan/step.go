package plan

type Step struct {
	// Action is the action/command name, e.g. ac013.
	Action string
	// Backoff is the tolerance within the associated action must be
	// successfully executed without returning errors.
	Backoff *Backoff
}
