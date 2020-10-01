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
	{
		Action:  "ac002",
		Backoff: plan.NewBackoff(30*time.Minute, 3*time.Minute),
		Comment: "check cluster access",
	},
	{
		Action:  "ac003",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "check master count",
	},
	{
		Action:  "ac004",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "check worker count",
	},
	{
		Action:  "ac005",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "create node pool",
	},
	{
		Action:  "ac004",
		Backoff: plan.NewBackoff(30*time.Minute, 3*time.Minute),
		Comment: "check worker count",
	},
	{
		Action:  "ac006",
		Backoff: plan.NewBackoff(15*time.Minute, 1*time.Minute),
		Comment: "check master with host network set",
	},
	{
		Action:  "ac007",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "check worker with host network set",
	},
	{
		Action:  "ac008",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "check if kiam's pods and secrets are present",
	},
	{
		Action:  "ac009",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "create job to test kiam aws api call",
	},
	{
		Action:  "ac010",
		Backoff: plan.NewBackoff(60*time.Second, 2*time.Second),
		Comment: "check job to test kiam aws api call",
	},
	{
		Action:  "ac011",
		Backoff: plan.NewBackoff(60*time.Second, 2*time.Second),
		Comment: "cleanup kiam test resources",
	},
	{
		Action:  "ac012",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "delete cluster CRs",
	},
	{
		Action:  "ac013",
		Backoff: plan.NewBackoff(90*time.Minute, 9*time.Minute),
		Comment: "check CRs deleted",
	},
}
