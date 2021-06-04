package pl002

import (
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/pkg/table"
)

const (
	name  = "pl002"
	short = "Execute plan pl002 automatically."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
	}

	c := &cobra.Command{
		Use:   name,
		Short: short,
		Long:  mustLong(),
		RunE:  r.Run,
	}

	f.Init(c)

	return c, nil
}

func mustLong() string {
	var d time.Duration
	for _, s := range Plan {
		d += s.Backoff.Wait() + s.CoolDown
	}

	s := "Test plan pl002 launches a basic Tenant Cluster and upgrades it according\n"
	s += "to the provided release versions. With this plan Tenant Clusters may be\n"
	s += "up or downgraded using major, minor or patch version changes. Plan\n"
	s += "execution might take up to " + d.String() + ".\n\n"

	t := [][]string{{"ACTION", "RETRY", "WAIT", "COOLDOWN"}}

	var h bool
	for _, s := range Plan {
		a := s.Action.String()
		i := s.Backoff.Interval().String()
		w := s.Backoff.Wait().String()
		c := ""
		if s.CoolDown != 0 {
			c = s.CoolDown.String()
			h = true
		}

		t = append(t, []string{a, i, w, c})
	}

	if !h {
		t[0][3] = ""
	}

	colourized := table.Colourize(t)
	s += table.Format(colourized)

	return s
}
