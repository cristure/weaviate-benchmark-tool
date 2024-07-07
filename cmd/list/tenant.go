package list

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/cristure/weaviate-benchmark-tool/client"
)

var (
	className string
)

func init() {
	tenantCmd.PersistentFlags().StringVarP(&className, "class-name", "c", "", "Name of the class to add the tenant to")
	tenantCmd.MarkPersistentFlagRequired("class-name")
}

var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "List tenants",
	Long:  `List the tenants from a class`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		before := time.Now()
		tenants, err := c.Schema().TenantsGetter().
			WithClassName(className).
			Do(context.Background())

		if err != nil {
			return err
		}

		took := time.Since(before)
		fmt.Fprintf(os.Stdout, "class (%s) has %d tenants. List took %s\n", className, len(tenants), took)

		return nil
	},
}
