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

	s := "Test plan for " + e.scope + " launches a Tenant Cluster with a custom NetworkPool and verifies basic criteria\n"
	s += "like the correct number of nodes, allocated subnets and pods are running. Plan execution might take\n"
	s += "up to " + d.String() + ".\n\n"

	t := [][]string{{"ACTION", "RETRY", "WAIT", "COMMENT"}}

	for _, s := range Plan {
		t = append(t, []string{s.Action, s.Backoff.Interval().String(), s.Backoff.Wait().String(), s.Comment})
	}

	colourized := table.Colourize(t)
	s += table.Format(colourized)

	return s, nil
}
