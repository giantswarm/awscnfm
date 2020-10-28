package ac000

import (
	"context"
	"time"

	"github.com/giantswarm/awscnfm/v12/pkg/table"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	var d time.Duration
	for _, s := range Plan {
		d += s.Backoff.Wait()
	}

	s := "Test plan for " + e.scope + " launches a basic Tenant Cluster with a single master\n"
	s += "and upgrades the Tenant Cluster to HA once it is up. Plan\n"
	s += "execution might take up to " + d.String() + ".\n\n"

	t := [][]string{{"ACTION", "RETRY", "WAIT", "COMMENT"}}

	for _, s := range Plan {
		t = append(t, []string{s.Action, s.Backoff.Interval().String(), s.Backoff.Wait().String(), s.Comment})
	}

	colourized := table.Colourize(t)
	s += table.Format(colourized)

	return s, nil
}
