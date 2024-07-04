package add

import (
	"github.com/spf13/cobra"
)

func init() {
	AddCmd.AddCommand(collectionCmd)
	AddCmd.AddCommand(tenantCmd)
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new object",
	Long: `Add a new object, it can be either a collection, tenant or plain object.

Add requires a subcommand, e.g. add collection`,
	Run: nil,
}
