package ac000

import (
	"context"
	"fmt"
)

func (e *Executor) execute(ctx context.Context) error {
	for _, c := range e.command.Parent().Parent().Commands() {
		fmt.Printf("%#v\n", c.Name())
	}

	return nil
}
