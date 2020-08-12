package explain

import (
	"path/filepath"

	"github.com/giantswarm/awscnfm/v12/pkg/key"
)

var FlagBase = filepath.Join("explain", key.GeneratedWithPrefix("flag.go"))

var FlagContent = `package explain

import "github.com/spf13/cobra"

type flag struct {
}

func (f *flag) Init(cmd *cobra.Command) {
}

func (f *flag) Validate() error {
	return nil
}
`
