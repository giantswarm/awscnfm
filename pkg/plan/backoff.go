package plan

import (
	"time"

	"github.com/giantswarm/backoff"
)

// Backoff is a simple wrapper that allows us to better explain how the backoff
// got configured.
type Backoff struct {
	// interval is the time after which we retry the configured action again.
	interval time.Duration
	// wait is the time after which the configured action must have succeeded.
	// If we could not execute the configured action successfully after this
	// timespan we give up and return the most recent error.
	wait time.Duration
	// wrapped is the backoff we transparently execute.
	wrapped backoff.Interface
}

func NewBackoff(wait time.Duration, interval time.Duration) *Backoff {
	return &Backoff{
		interval: interval,
		wait:     wait,
		wrapped:  backoff.NewConstant(wait, interval),
	}
}

func (b *Backoff) Interval() time.Duration {
	return b.interval
}

func (b *Backoff) Wait() time.Duration {
	return b.wait
}

func (b *Backoff) Wrapped() backoff.Interface {
	return b.wrapped
}
