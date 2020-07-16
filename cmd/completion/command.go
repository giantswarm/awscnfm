package completion

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name             = "completion [bash|zsh|fish|powershell]"
	shortDescription = "Generate completion script."
	longDescription  = `To load completions:

	# Bash:

	$ source <(awscnfm completion bash)

	# To load completions for each session, execute once:
	# Linux:
	  $ awscnfm completion bash > /etc/bash_completion.d/awscnfm
	# MacOS:
	  $ awscnfm completion bash > /usr/local/etc/bash_completion.d/awscnfm

	# Zsh:

	$ source <(awscnfm completion zsh)

	# To load completions for each session, execute once:
	$ awscnfm completion zsh > "${fpath[1]}/_awscnfm"

	# Fish:

	$ awscnfm completion fish | source

	# To load completions for each session, execute once:
	$ awscnfm completion fish > ~/.config/fish/completions/awscnfm.fish`
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

	r := &runner{
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:                   name,
		Short:                 shortDescription,
		Long:                  longDescription,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		RunE:                  r.Run,
	}

	return c, nil
}
