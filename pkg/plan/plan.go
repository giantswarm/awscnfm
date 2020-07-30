package plan

import "github.com/giantswarm/backoff"

type Plan struct {
	// Action is the action/command name, e.g. ac013.
	Action string
	// Backoff is the tolerance within the associated action must be
	// successfully executed without returning errors.
	Backoff backoff.Interface
	// Comment is only informational to provide developers hints about the
	// purpose of a given action. This should help understanding what ac013 is
	// doing so that one can reason about why ac027 should run next instead of
	// ac014.
	Comment string
}
