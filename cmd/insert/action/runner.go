package action

import (
	"context"
	"fmt"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/pkg/action"
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
	// We compute the mapping of actions we have to move based on the given
	// action we want to insert. Consider having a list of actions from ac000 to
	// ac009 and we want to insert an empty action ac004. The resulting mapping
	// below would look like this.
	//
	//     map[string]string{
	//         "ac004": "ac005",
	//         "ac005": "ac006",
	//         "ac006": "ac007",
	//         "ac007": "ac008",
	//         "ac008": "ac009",
	//         "ac009": "ac010",
	//     }
	//
	actions := map[string]string{}
	{
		all, err := action.All(r.flag.Cluster)
		if err != nil {
			return microerror.Mask(err)
		}

		for _, a := range all {
			if action.Lower(a, r.flag.Action) {
				continue
			}

			actions[a] = action.Incr(a)
		}
	}

	fmt.Printf("%#v\n", actions)

	return nil
}
