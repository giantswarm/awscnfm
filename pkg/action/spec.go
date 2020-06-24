package action

import (
	"context"

	"github.com/giantswarm/awscnfm/pkg/client"
)

type Interface interface {
	Execute(ctx context.Context, cli *client.Client) error
	Explain() string
}
