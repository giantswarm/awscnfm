package action

import (
	"github.com/giantswarm/awscnfm/pkg/key"
)

var ExplainerBase = key.GeneratedWithPrefix("explainer.go")

var ExplainerContent = `package {{ .Action }}

import (
	"context"
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
`
