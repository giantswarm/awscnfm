package action

var ExecutorCustomBase = "executor.go"

var ExecutorCustomContent = `package {{ .Action }}

import (
	"context"{{ if eq .Action "ac001" }}

	"github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"{{ end }}
)

{{ if eq .Action "ac001" }}func (e *Executor) execute(ctx context.Context) (v1alpha2.ClusterCRs, error) {
	return v1alpha2.ClusterCRs{}, nil
}{{ else }}func (e *Executor) execute(ctx context.Context) error {
	return nil
}{{ end }}
`
