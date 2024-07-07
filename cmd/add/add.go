package add

import (
	"github.com/spf13/cobra"
)

func init() {
	AddCmd.AddCommand(classCmd)
	AddCmd.AddCommand(tenantCmd)
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new resource",
	Long: `Add a new resource, it can be either a class, tenant or plain object.

Add requires a subcommand, e.g. add collection`,
	Run: nil,
}
