package ac006

import (
	"context"

	"github.com/giantswarm/microerror"
)

type ExplainerConfig struct {
}

type Explainer struct {
}

func NewExplainer(config ExplainerConfig) (*Explainer, error) {
	e := &Explainer{}

	return e, nil
}

func (e *Explainer) Explain(ctx context.Context) (string, error) {
	s, err := e.explain(ctx)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return s, nil
}
