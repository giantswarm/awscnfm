package action

import (
	"context"
)

type Executor interface {
	Execute(ctx context.Context) error
}

type Explainer interface {
	Explain(ctx context.Context) (string, error)
}
