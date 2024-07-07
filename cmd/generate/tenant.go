package generate

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate/entities/models"

	"github.com/cristure/weaviate-benchmark-tool/client"
	"github.com/cristure/weaviate-benchmark-tool/cmd/stats"
)

var (
	className  string
	tenantsNo  int
	namePrefix string
)

func init() {
	tenantCmd.PersistentFlags().StringVarP(&className, "class-name", "c", "", "Name of the class to add the tenant to")
	tenantCmd.PersistentFlags().IntVarP(&tenantsNo, "tenants-number", "t", 0, "Number of tenants to generate")
	tenantCmd.PersistentFlags().StringVarP(&namePrefix, "name-prefix", "p", "tenant", "Prefix for tenant names")
	tenantCmd.MarkPersistentFlagRequired("class-name")
	tenantCmd.MarkPersistentFlagRequired("name")
}

var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Generate tenants",
	Long:  `Add a specified amount of tenants to a collection`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		latencies := make(map[string]time.Duration)
		for i := 0; i < tenantsNo; i++ {
			tenantName := fmt.Sprintf("%s-%d", namePrefix, i)
			fmt.Printf("Creating tenant %d...\n", i)
			before := time.Now()
			err = c.Schema().TenantsCreator().
				WithClassName(className).
				WithTenants(models.Tenant{Name: tenantName}).
				Do(context.Background())

			if err != nil {
				return fmt.Errorf("failed to create tenant %s: %w", tenantName, err)
			}

			took := time.Since(before)
			fmt.Fprintf(os.Stdout, "successfully created tenant (%s) in %s\n", tenantName, took)
			latencies[tenantName] = took
		}

		meanLatency := stats.Mean(latencies)
		fmt.Printf("Mean latency: %s\n", meanLatency)

		p50 := stats.Percentile(latencies, 50)
		p90 := stats.Percentile(latencies, 90)
		p99 := stats.Percentile(latencies, 99)

		fmt.Printf("50th percentile (median) duration: %s\n", p50)
		fmt.Printf("90th percentile duration: %s\n", p90)
		fmt.Printf("99th percentile duration: %s\n", p99)
		return nil
	},
}
