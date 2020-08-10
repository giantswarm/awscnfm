package template

import "github.com/giantswarm/awscnfm/pkg/key"

var CommandBase = key.GeneratedWithPrefix("command.go")

var CommandContent = `package {{ .Cluster }}

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"{{ if not .Actions }}
){{ else }}
{{ range $a := .Actions }}
	"github.com/giantswarm/awscnfm/cmd/{{ $.Cluster }}/{{ $a }}"
{{- end }}
){{ end }}

const (
	name        = "{{ .Cluster }}"
	description = "Conformance tests for cluster {{ .Cluster }}."
)

type Config struct {
	Logger micrologger.Logger
	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	var err error
{{ range $a := .Actions }}
	var {{ $a }}Cmd *cobra.Command
	{
		c := {{ $a }}.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		{{ $a }}Cmd, err = {{ $a }}.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}
{{ end }}
	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
		RunE:  r.Run,
	}

	f.Init(c){{ if not .Actions }}
{{- else }}
{{ range $a := .Actions }}
	c.AddCommand({{ $a }}Cmd)
{{- end -}}
{{ end }}

	return c, nil
}
`
