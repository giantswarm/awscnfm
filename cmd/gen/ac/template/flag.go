package template

import "github.com/giantswarm/awscnfm/pkg/key"

var FlagBase = key.GeneratedWithPrefix("flag.go")

var FlagContent = `package {{ .Action }}

import "github.com/spf13/cobra"

type flag struct {
}

func (f *flag) Init(cmd *cobra.Command) {
}

func (f *flag) Validate() error {
	return nil
}
`
