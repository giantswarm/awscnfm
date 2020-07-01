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
	templatecmd "github.com/giantswarm/awscnfm/cmd/generate/action/template/cmd"
	templatecmdexecute "github.com/giantswarm/awscnfm/cmd/generate/action/template/cmd/execute"
	templatecmdexplain "github.com/giantswarm/awscnfm/cmd/generate/action/template/cmd/explain"
	templatepkgaction "github.com/giantswarm/awscnfm/cmd/generate/action/template/pkg/action"
	"github.com/giantswarm/awscnfm/pkg/key"
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
	data := actiontemplate.Data{
		Action:  r.flag.Action,
		Cluster: r.flag.Cluster,
	}

	{
		// templates is a predefined list of lists for debugging reasons. When
		// defining map[string]string for the key-value pairs the order of items
		// changes since go maps are not deterministic.
		templates := [][]string{
			{templatecmdexecute.CommandBase, templatecmdexecute.CommandContent},
			{templatecmdexecute.ErrorBase, templatecmdexecute.ErrorContent},
			{templatecmdexecute.FlagBase, templatecmdexecute.FlagContent},
			{templatecmdexecute.RunnerBase, templatecmdexecute.RunnerContent},

			{templatecmdexplain.CommandBase, templatecmdexplain.CommandContent},
			{templatecmdexplain.ErrorBase, templatecmdexplain.ErrorContent},
			{templatecmdexplain.FlagBase, templatecmdexplain.FlagContent},
			{templatecmdexplain.RunnerBase, templatecmdexplain.RunnerContent},

			{templatecmd.CommandBase, templatecmd.CommandContent},
			{templatecmd.ErrorBase, templatecmd.ErrorContent},
			{templatecmd.FlagBase, templatecmd.FlagContent},
			{templatecmd.RunnerBase, templatecmd.RunnerContent},
		}

		err := write(data, templates, "cmd/%s/%s")
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		// templates is a predefined list of lists for debugging reasons. When
		// defining map[string]string for the key-value pairs the order of items
		// changes since go maps are not deterministic.
		templates := [][]string{
			{templatepkgaction.ErrorBase, templatepkgaction.ErrorContent},
			{templatepkgaction.ExecutorBase, templatepkgaction.ExecutorContent},
			{templatepkgaction.ExecutorCustomBase, templatepkgaction.ExecutorCustomContent},
			{templatepkgaction.ExplainerBase, templatepkgaction.ExplainerContent},
			{templatepkgaction.ExplainerCustomBase, templatepkgaction.ExplainerCustomContent},
		}

		err := write(data, templates, "pkg/action/%s/%s")
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

func write(data actiontemplate.Data, templates [][]string, dirfmt string) error {
	for _, l := range templates {
		base := l[0]
		cont := l[1]

		path, err := filepath.Abs(filepath.Join(fmt.Sprintf(dirfmt, data.Cluster, data.Action), base))
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

		var write bool
		{
			var exists bool
			_, err = os.Stat(path)
			if os.IsNotExist(err) {
				// fall through, exists is already false
			} else if err != nil {
				return microerror.Mask(err)
			} else {
				exists = true
			}

			var regenerate bool
			if key.HasGeneratedPrefix(base) {
				regenerate = true
			}

			// This means that r.g. custom business logic in explainer.go is
			// only bootstrapped, but not regenerated once the file exists,
			// because this means whoever generated the action subcommand in the
			// first place got the chance to add their custom code in there
			// already. At the same time we want to regenerate files like
			// zz_generated.explainer.go, since this is where no custom code
			// should go.
			write = !exists || regenerate
		}

		if write {
			err = ioutil.WriteFile(path, buff.Bytes(), 0644)
			if err != nil {
				return microerror.Mask(err)
			}
		}
	}

	return nil
}
