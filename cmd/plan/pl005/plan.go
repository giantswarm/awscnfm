package pl005

import (
	"time"

	"github.com/giantswarm/awscnfm/v12/pkg/plan"
)

// Plan describes in which order and with which tolerance to execute actions of
// this test plan.
var Plan = []plan.Step{
	{
		Action:  "create/cluster/singlecontrolplane",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "verify/cluster/created",
		Backoff: plan.NewBackoff(30*time.Minute, 1*time.Minute),
	},
	{
		Action:  "update/cluster/major",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "verify/cluster/updated",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:    "delete/cluster",
		Backoff:   plan.NewBackoff(10*time.Second, 2*time.Second),
		Condition: plan.ConditionAlwaysExecute,
	},
	{
		Action:    "verify/cluster/deleted",
		Backoff:   plan.NewBackoff(90*time.Minute, 9*time.Minute),
		Condition: plan.ConditionAlwaysExecute,
	},
}
