package action

var ExecutorCustomBase = "executor.go"

var ExecutorCustomContent = `package {{ .Action }}

import (
	"context"
)

func (e *Executor) execute(ctx context.Context) error {
	return nil
}
`
