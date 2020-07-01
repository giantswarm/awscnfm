package ac

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

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
	data := struct {
		Action  string
		Cluster string
	}{
		Action:  r.flag.Action,
		Cluster: r.flag.Cluster,
	}

	// templates is a predefined list of lists for debugging reasons. When
	// defining map[string]string for the key-value pairs the order of items
	// changes since go maps are not deterministic.
	templates := [][]string{
		{actemplateexecute.CommandBase, actemplateexecute.CommandContent},
		{actemplateexecute.ErrorBase, actemplateexecute.ErrorContent},
		{actemplateexecute.FlagBase, actemplateexecute.FlagContent},
		{actemplateexecute.RunnerBase, actemplateexecute.RunnerContent},

		{actemplateexplain.CommandBase, actemplateexplain.CommandContent},
		{actemplateexplain.ErrorBase, actemplateexplain.ErrorContent},
		{actemplateexplain.FlagBase, actemplateexplain.FlagContent},
		{actemplateexplain.RunnerBase, actemplateexplain.RunnerContent},

		{actemplate.CommandBase, actemplate.CommandContent},
		{actemplate.ErrorBase, actemplate.ErrorContent},
		{actemplate.FlagBase, actemplate.FlagContent},
		{actemplate.RunnerBase, actemplate.RunnerContent},
	}

	for _, l := range templates {
		base := l[0]
		cont := l[1]

		path, err := filepath.Abs(filepath.Join(fmt.Sprintf("cmd/%s/%s", data.Cluster, data.Action), base))
		if err != nil {
			return microerror.Mask(err)
		}

		t, err := template.New(path).Parse(cont)
		if err != nil {
			return microerror.Mask(err)
		}

		var buff bytes.Buffer
		err = t.ExecuteTemplate(&buff, path, data)
		if err != nil {
			return microerror.Mask(err)
		}

		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return microerror.Mask(err)
		}

		err = ioutil.WriteFile(path, buff.Bytes(), 0644)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
