package pl002

import (
	"time"

	"github.com/giantswarm/awscnfm/v12/pkg/plan"
)

// Plan describes in which order and with which tolerance to execute actions of
// this test plan.
var Plan = []plan.Step{
	{
		Action:  "create/cluster/defaultcontrolplane",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "create/nodepool/defaultdataplane",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "verify/cluster/created",
		Backoff: plan.NewBackoff(30*time.Minute, 3*time.Minute),
	},
	{
		Action:  "verify/master/ready",
		Backoff: plan.NewBackoff(10*time.Minute, 30*time.Second),
	},
	{
		Action:  "verify/worker/ready",
		Backoff: plan.NewBackoff(10*time.Minute, 30*time.Second),
	},
	{
		Action:  "update/cluster/patch",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "verify/cluster/updated",
		Backoff: plan.NewBackoff(120*time.Minute, 10*time.Minute),
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
