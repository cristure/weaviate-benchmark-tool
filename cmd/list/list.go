package list

import (
	"github.com/spf13/cobra"
)

func init() {
	ListCmd.AddCommand(tenantCmd)
	ListCmd.AddCommand(shardCmd)
	ListCmd.AddCommand(objectsCmd)
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources",
	Long: `List resources and latencies

Add requires a subcommand, e.g. list tenant`,
	Run: nil,
}
