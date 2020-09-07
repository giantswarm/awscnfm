package ac000

import (
	"time"

	"github.com/giantswarm/awscnfm/v12/pkg/plan"
)

// Plan describes in which order and with which tolerance to execute actions of
// this cluster scope.
var Plan = []plan.Step{
	{
		Action:  "ac001",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "create cluster CRs",
	},
}
