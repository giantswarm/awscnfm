package action

import (
	"github.com/giantswarm/awscnfm/v12/pkg/key"
)

var ExplainerBase = key.GeneratedWithPrefix("explainer.go")

var ExplainerContent = `package {{ .Action }}

import (
	"context"

	"github.com/giantswarm/microerror"
)

{{ if eq .Action "ac001" -}}
const (
	// explainerCommand is for internal documentation purposes only so that
	// commands can self describe and explain themselves better. This
	// information might be used in different creative ways.
	explainerCommand = "awscnfm {{ .Cluster }} {{ .Action }} explain"
)

{{ end -}}

type ExplainerConfig struct {
	TenantCluster string
}

type Explainer struct {
	tenantCluster string
}

func NewExplainer(config ExplainerConfig) (*Explainer, error) {
	e := &Explainer{
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
`
