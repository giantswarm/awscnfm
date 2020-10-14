package ac001

import (
	"context"

	"github.com/giantswarm/microerror"
)

const (
	// explainerCommand is for internal documentation purposes only so that
	// commands can self describe and explain themselves better. This
	// information might be used in different creative ways.
	explainerCommand = "awscnfm cl005 ac001 explain"
)

const (
	// This is a hack to make initially generated code compile because there is
	// nothing making use of the constant when starting out.
	_ = explainerCommand
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
