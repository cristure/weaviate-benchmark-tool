package add

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate/entities/models"

	"github.com/cristure/weaviate-benchmark-tool/client"
)

var (
	className  string
	tenantName string
)

func init() {
	tenantCmd.PersistentFlags().StringVarP(&className, "class", "c", "", "Name of the class to add the tenant to")
	tenantCmd.PersistentFlags().StringVarP(&tenantName, "name", "n", "", "Name of the tenant to add")
	tenantCmd.MarkPersistentFlagRequired("class-name")
	tenantCmd.MarkPersistentFlagRequired("name")
}

var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Add tenants",
	Long:  `Add only a tenant to a given class`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		before := time.Now()
		err = c.Schema().TenantsCreator().
			WithClassName(className).
			WithTenants(models.Tenant{Name: tenantName}).
			Do(context.Background())

		took := time.Since(before)

		fmt.Fprintf(os.Stdout, "Tenant (%s) has been created in %s", tenantName, took)
		return nil
	},
}
