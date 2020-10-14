package ac010

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Test AWS API calls are working as expected. This part will create a pod (via job)
and try assume role "$cluster_id-Route53-Manager" and call "aws route53 list-domains"
to ensure kiam works as expected.
`
	return s, nil
}
