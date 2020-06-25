package action

import (
	"context"
)

type Executor interface {
	Execute(ctx context.Context) error
}

type Explainer interface {
	Explain() string
}
