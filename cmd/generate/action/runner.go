package action

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

	actiontemplate "github.com/giantswarm/awscnfm/cmd/generate/action/template"
	actiontemplateexecute "github.com/giantswarm/awscnfm/cmd/generate/action/template/execute"
	actiontemplateexplain "github.com/giantswarm/awscnfm/cmd/generate/action/template/explain"
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
		{actiontemplateexecute.CommandBase, actiontemplateexecute.CommandContent},
		{actiontemplateexecute.ErrorBase, actiontemplateexecute.ErrorContent},
		{actiontemplateexecute.FlagBase, actiontemplateexecute.FlagContent},
		{actiontemplateexecute.RunnerBase, actiontemplateexecute.RunnerContent},

		{actiontemplateexplain.CommandBase, actiontemplateexplain.CommandContent},
		{actiontemplateexplain.ErrorBase, actiontemplateexplain.ErrorContent},
		{actiontemplateexplain.FlagBase, actiontemplateexplain.FlagContent},
		{actiontemplateexplain.RunnerBase, actiontemplateexplain.RunnerContent},

		{actiontemplate.CommandBase, actiontemplate.CommandContent},
		{actiontemplate.ErrorBase, actiontemplate.ErrorContent},
		{actiontemplate.FlagBase, actiontemplate.FlagContent},
		{actiontemplate.RunnerBase, actiontemplate.RunnerContent},
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
