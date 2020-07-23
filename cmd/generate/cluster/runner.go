package cluster

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	clustertemplate "github.com/giantswarm/awscnfm/cmd/generate/cluster/template"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
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
	var actions []string
	{
		path, err := filepath.Abs(fmt.Sprintf("cmd/%s", r.flag.Cluster))
		if err != nil {
			return microerror.Mask(err)
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			return microerror.Mask(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				continue
			}
			if !strings.HasPrefix(file.Name(), "ac") {
				continue
			}

			actions = append(actions, file.Name())
		}
	}

	data := struct {
		Actions []string
		Cluster string
	}{
		Actions: actions,
		Cluster: r.flag.Cluster,
	}

	// templates is a predefined list of lists for debugging reasons. When
	// defining map[string]string for the key-value pairs the order of items
	// changes since go maps are not deterministic.
	templates := [][]string{
		{clustertemplate.CommandBase, clustertemplate.CommandContent},
		{clustertemplate.ErrorBase, clustertemplate.ErrorContent},
		{clustertemplate.FlagBase, clustertemplate.FlagContent},
		{clustertemplate.RunnerBase, clustertemplate.RunnerContent},
	}

	for _, l := range templates {
		base := l[0]
		cont := l[1]

		path, err := filepath.Abs(filepath.Join(fmt.Sprintf("cmd/%s", data.Cluster), base))
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

		err = ioutil.WriteFile(path, buff.Bytes(), 0644) // nolint:gosec
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
