package ac008

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	return "Check that masters and nodepools all belong to the NetworkPool CIDR based on the reported subet used in the CRs", nil
}
