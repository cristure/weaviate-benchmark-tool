package add

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate/entities/models"

	"github.com/cristure/weaviate-benchmark-tool/client"
)

var (
	collectionName string
	tenantsNo      int
)

func init() {
	tenantCmd.PersistentFlags().StringVarP(&collectionName, "name", "N", "", "Name of the collection to add to")
	tenantCmd.PersistentFlags().IntVarP(&tenantsNo, "tenants-number", "T", 1, "Number of the tenants to add")
}

var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Add tenants",
	Long:  `Add a specified amount of tenants to a collection`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		if collectionName == "" {
			return fmt.Errorf("must specify collection name")
		}

		// Add the class to the schema
		err = c.Schema().TenantsCreator().
			WithClassName(collectionName).
			WithTenants(generateTenants()...).
			Do(context.Background())

		fmt.Fprintf(os.Stdout, "Tenants have been added successfully to collection (%s)", collectionName)
		return nil
	},
}

func generateTenants() []models.Tenant {
	tenants := make([]models.Tenant, tenantsNo)
	for i := 0; i < tenantsNo; i++ {
		tenants[i] = models.Tenant{
			Name: fmt.Sprintf("tenant%d", i),
		}
	}

	return tenants
}
