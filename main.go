package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v14/cmd"
	"github.com/giantswarm/awscnfm/v14/pkg/project"
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		mErr, ok := microerror.Cause(err).(*microerror.Error)
		if ok && mErr.Desc != "" {
			fmt.Println(strings.Title(err.Error()))
			fmt.Println()
			fmt.Println("    " + mErr.Desc)
			fmt.Println()
			os.Exit(1)
		} else {
			panic(microerror.JSON(err))
		}
	}
}

func mainE(ctx context.Context) error {
	var err error

	var logger micrologger.Logger
	{
		c := micrologger.Config{}

		logger, err = micrologger.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		c := micrologger.ActivationLoggerConfig{
			Underlying: logger,

			Activations: map[string]interface{}{
				micrologger.KeyLevel: "info",
			},
		}

		logger, err = micrologger.NewActivation(c)
		if err != nil {
			panic(err)
		}
	}

	var rootCommand *cobra.Command
	{
		c := cmd.Config{
			Logger: logger,

			GitCommit: project.GitSHA(),
			Source:    project.Source(),
		}

		rootCommand, err = cmd.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = rootCommand.Execute()
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
