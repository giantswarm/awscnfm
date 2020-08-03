package ac000

import (
	"time"

	"github.com/giantswarm/backoff"

	"github.com/giantswarm/awscnfm/pkg/plan"
)

// Plan describes in which order and with which tolerance to execute actions of
// this cluster scope.
var Plan = []plan.Step{
	{
		Action:  "ac001",
		Backoff: backoff.NewConstant(24*time.Minute, 3*time.Minute),
		Comment: "create cluster",
	},
	{
		Action:  "ac002",
		Backoff: backoff.NewConstant(10*time.Second, 2*time.Second),
		Comment: "check master count",
	},
	{
		Action:  "ac003",
		Backoff: backoff.NewConstant(10*time.Second, 2*time.Second),
		Comment: "check worker count",
	},
	{
		Action:  "ac004",
		Backoff: backoff.NewConstant(90*time.Minute, 9*time.Minute),
		Comment: "delete cluster",
	},
}
