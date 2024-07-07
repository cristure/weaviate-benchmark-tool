package generate

import (
	"github.com/spf13/cobra"
)

func init() {
	GenerateCmd.AddCommand(objectCmd)
	GenerateCmd.AddCommand(tenantCmd)
}

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates resources for a class",
	Long: `Generates resources for a class. It can be either tenants or objects

Add requires a subcommand, e.g. generate tenant`,
	Run: nil,
}
