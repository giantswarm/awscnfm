package pl003

import (
	"time"

	"github.com/giantswarm/awscnfm/v14/pkg/plan"
)

// Plan describes in which order and with which tolerance to execute actions of
// this test plan.
var Plan = []plan.Step{
	{
		Action:  "create/cluster/customnetworkpool",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "create/nodepool/defaultdataplane",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "verify/cluster/created",
		Backoff: plan.NewBackoff(30*time.Minute, 1*time.Minute),
	},
	{
		Action:  "verify/cluster/networkpool",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
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
		Action:  "verify/apps/installed",
		Backoff: plan.NewBackoff(20*time.Minute, 1*time.Minute),
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
		Action:  "create/netpol/defaultnetpol",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "create/netpol/curlrequest",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "verify/netpol/curlrequest",
		Backoff: plan.NewBackoff(15*time.Minute, 1*time.Minute),
	},
	{
		Action:  "delete/netpol/curlrequest",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
	},
	{
		Action:  "delete/netpol/defaultnetpol",
		Backoff: plan.NewBackoff(30*time.Second, 4*time.Second),
	},
	{
		Action:    "delete/cluster",
		Backoff:   plan.NewBackoff(10*time.Second, 2*time.Second),
		Condition: plan.ConditionAlwaysExecute,
	},
	{
		Action:    "delete/networkpool",
		Backoff:   plan.NewBackoff(10*time.Second, 2*time.Second),
		Condition: plan.ConditionAlwaysExecute,
	},
	{
		Action:    "verify/cluster/deleted",
		Backoff:   plan.NewBackoff(90*time.Minute, 9*time.Minute),
		Condition: plan.ConditionAlwaysExecute,
	},
}
