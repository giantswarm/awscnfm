package ac009

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Check if Kiam's pods and secrets are present on the tenant cluster. This test should do following:

    * Ensure tls certs for Kiam are created.
    * Ensure kiam-server and kiam-agent pods are running without errors on the tenant cluster.

`

	return s, nil
}
