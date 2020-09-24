package ac008

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Test Kiam app deployed on the tenant cluster. This test should do following:

    * Ensure tls certs for Kiam are created.
    * Ensure kiam-server and kiam-agent pods are running without errors on the tenant cluster.
    * Test AWS API calls are working as expected. This part will create a pod (via job) 
      and try assume role "$cluster_id-Route53-Manager" and call "aws route53 list-domains" to ensure kiam works as expected.
`

	return s, nil
}
