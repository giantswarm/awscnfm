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
		Comment: "upgrade tenant cluster",
	},
	{
		Action:  "ac004",
		Backoff: plan.NewBackoff(120*time.Minute, 10*time.Minute),
		Comment: "check upgrade successful",
	},
	{
		Action:  "ac005",
		Backoff: plan.NewBackoff(10*time.Second, 2*time.Second),
		Comment: "delete cluster CRs",
	},
	{
		Action:  "ac006",
		Backoff: plan.NewBackoff(90*time.Minute, 9*time.Minute),
		Comment: "check CRs deleted",
	},
}
