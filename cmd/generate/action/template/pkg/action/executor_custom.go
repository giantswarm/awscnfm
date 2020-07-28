package action

var ExecutorCustomBase = "executor.go"

var ExecutorCustomContent = `package {{ .Action }}

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
)

func (e *Executor) execute(ctx context.Context) (v1alpha2.ClusterCRs, error) {
	return v1alpha2.ClusterCRs{}, nil
}
`
