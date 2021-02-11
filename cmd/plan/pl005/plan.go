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
		Action:   "verify/cluster/created",
		Backoff:  plan.NewBackoff(30*time.Minute, 3*time.Minute),
		CoolDown: 30 * time.Minute,
	},
	{
		Action:  "verify/master/ready",
		Backoff: plan.NewBackoff(10*time.Minute, 30*time.Second),
	},
	{
		Action:  "create/nodepool/defaultdataplane",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "verify/worker/ready",
		Backoff: plan.NewBackoff(10*time.Minute, 30*time.Second),
	},
	{
		Action:  "verify/master/hostnetworkpod",
		Backoff: plan.NewBackoff(15*time.Minute, 1*time.Minute),
	},
	{
		Action:  "verify/worker/hostnetworkpod",
		Backoff: plan.NewBackoff(15*time.Minute, 1*time.Minute),
	},
	{
		Action:  "verify/kiam/podandsecret",
		Backoff: plan.NewBackoff(10*time.Minute, 30*time.Second),
	},
	{
		Action:  "create/kiam/awsapicall",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "verify/kiam/awsapicall",
		Backoff: plan.NewBackoff(2*time.Minute, 5*time.Second),
	},
	{
		Action:  "delete/kiam/awsapicall",
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
