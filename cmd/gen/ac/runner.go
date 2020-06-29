package ac

import (
	"bytes"
	"context"
	"html/template"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	actemplate "github.com/giantswarm/awscnfm/cmd/gen/ac/template"
	actemplateexecute "github.com/giantswarm/awscnfm/cmd/gen/ac/template/execute"
	actemplateexplain "github.com/giantswarm/awscnfm/cmd/gen/ac/template/explain"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	var err error

	ctx := context.Background()

	err = r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	data := struct {
		Action  string
		Cluster string
	}{
		Action:  r.flag.Action,
		Cluster: r.flag.Cluster,
	}

	templates := map[string]string{
		"command.go":         actemplate.CommandGo,
		"error.go":           actemplate.ErrorGo,
		"flag.go":            actemplate.FlagGo,
		"runner.go":          actemplate.RunnerGo,
		"execute/command.go": actemplateexecute.CommandGo,
		"execute/error.go":   actemplateexecute.ErrorGo,
		"execute/flag.go":    actemplateexecute.FlagGo,
		"execute/runner.go":  actemplateexecute.RunnerGo,
		"explain/command.go": actemplateexplain.CommandGo,
		"explain/error.go":   actemplateexplain.ErrorGo,
		"explain/flag.go":    actemplateexplain.FlagGo,
		"explain/runner.go":  actemplateexplain.RunnerGo,
	}

	for p, k := range templates {
		main := template.New(p)

		main, err = main.Parse(t)
		if err != nil {
			return microerror.Mask(err)
		}

		var b bytes.Buffer
		err = main.ExecuteTemplate(&b, "main", data)
		if err != nil {
			return microerror.Mask(err)
		}

		return b.String()
	}

	return nil
}
