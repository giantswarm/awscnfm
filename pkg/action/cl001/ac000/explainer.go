package ac000

import (
	"context"

	"github.com/giantswarm/awscnfm/pkg/table"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := "Execute the conformance test plan of this cluster scope. Actions below are\n"
	s += "executed in order. A tenant cluster is conform if the plan executes without\n"
	s += "errors.\n\n"

	t := [][]string{{"ACTION", "COMMENT"}}

	for _, s := range Plan {
		t = append(t, []string{s.Action, s.Comment})
	}

	colourized := table.Colourize(t)
	s += table.Format(colourized)

	return s, nil
}
