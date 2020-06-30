package execute

import (
	"path/filepath"

	"github.com/giantswarm/awscnfm/pkg/key"
)

var FlagBase = filepath.Join("execute", key.GeneratedWithPrefix("flag.go"))

var FlagContent = `
package execute

import "github.com/spf13/cobra"

type flag struct {
}

func (f *flag) Init(cmd *cobra.Command) {
}

func (f *flag) Validate() error {
	return nil
}
`
