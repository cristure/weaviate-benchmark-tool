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
	name    string
	propsNo int
)

func init() {
	collectionCmd.PersistentFlags().StringVarP(&name, "name", "N", "", "name of the collection")
	collectionCmd.PersistentFlags().IntVarP(&propsNo, "properties-number", "P", 1, "Properties number")
	//collectionCmd.PersistentFlags().IntVarP(&tenantsNo, "tenants-number", "T", 1, "Tenants number")
}

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Add a new collection",
	Long:  `Add a new collection`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		if name == "" {
			return fmt.Errorf("must specify collection name")
		}

		class := &models.Class{
			Class:      name,
			Properties: generateProperties(),
			MultiTenancyConfig: &models.MultiTenancyConfig{
				Enabled:            true,
				AutoTenantCreation: true,
			},
		}

		// Add the class to the schema
		err = c.Schema().ClassCreator().
			WithClass(class).
			Do(context.Background())
		if err != nil {
			return fmt.Errorf("error while creating new collection: %w", err)
		}

		fmt.Fprintf(os.Stdout, "Collection (%s) created successfully!\n", name)
		return nil
	},
}

func generateProperties() []*models.Property {
	properties := make([]*models.Property, propsNo)
	for i := 0; i < propsNo; i++ {
		properties[i] = &models.Property{
			Name:     fmt.Sprintf("prop%d", i),
			DataType: []string{"text"},
		}
	}

	return properties
}
