package action

var ExplainerCustomBase = "explainer.go"

var ExplainerCustomContent = `package {{ .Action }}

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	return "", nil
}
`
