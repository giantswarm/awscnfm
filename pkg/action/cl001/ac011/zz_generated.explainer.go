package ac011

import (
	"context"

	"github.com/giantswarm/microerror"
)

type ExplainerConfig struct {
	Scope         string
	TenantCluster string
}

type Explainer struct {
	scope         string
	tenantCluster string
}

func NewExplainer(config ExplainerConfig) (*Explainer, error) {
	e := &Explainer{
		scope:         config.Scope,
		tenantCluster: config.TenantCluster,
	}

	return e, nil
}

func (e *Explainer) Explain(ctx context.Context) (string, error) {
	s, err := e.explain(ctx)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return s, nil
}
