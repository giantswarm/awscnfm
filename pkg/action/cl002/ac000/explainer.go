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

	s := "Execute the conformance test plan of this cluster scope. Actions below are\n"
	s += "executed in order. A tenant cluster is conform if the plan executes without\n"
	s += "errors. Plan execution might take up to " + d.String() + ".\n\n"

	t := [][]string{{"ACTION", "RETRY", "WAIT", "COMMENT"}}

	for _, s := range Plan {
		t = append(t, []string{s.Action, s.Backoff.Interval().String(), s.Backoff.Wait().String(), s.Comment})
	}

	colourized := table.Colourize(t)
	s += table.Format(colourized)

	return s, nil
}
