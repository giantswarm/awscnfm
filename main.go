package main

import (
	"context"
	"fmt"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/cmd"
	"github.com/giantswarm/awscnfm/pkg/project"
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		mErr, ok := microerror.Cause(err).(*microerror.Error)
		if ok && mErr.Desc != "" {
			fmt.Println(mErr.Desc)
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
		c := micrologger.Config{
			//IOWriter: ioutil.Discard,
		}

		logger, err = micrologger.New(c)
		if err != nil {
			return microerror.Mask(err)
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
