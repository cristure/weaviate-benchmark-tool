package list

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/cristure/weaviate-benchmark-tool/client"
	"github.com/cristure/weaviate-benchmark-tool/cmd/stats"
)

func init() {
	objectsCmd.PersistentFlags().StringVarP(&className, "class-name", "c", "", "Name of the class to add the tenant to")
	objectsCmd.MarkPersistentFlagRequired("class-name")
}

var objectsCmd = &cobra.Command{
	Use:   "object",
	Short: "List objects",
	Long:  `List the objects from a class`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		tenants, err := c.Schema().TenantsGetter().
			WithClassName(className).
			Do(context.Background())

		fmt.Fprintf(os.Stdout, "class (%s) has %d tenants", className, len(tenants))

		latencies := make(map[string]time.Duration)
		for _, t := range tenants {
			before := time.Now()
			objs, err := c.Data().ObjectsGetter().
				WithClassName(className).
				WithTenant(t.Name).
				Do(context.Background())
			if err != nil {
				return err
			}
			took := time.Since(before)

			fmt.Fprintf(os.Stdout, "tenant (%s) has %d objects. List took %s\n", t.Name, len(objs), took)
			latencies[t.Name] = took
		}
		meanLatency := stats.Mean(latencies)
		fmt.Fprintf(os.Stdout, "Average latency: %d\n", meanLatency)

		p50 := stats.Percentile(latencies, 50)
		p90 := stats.Percentile(latencies, 90)
		p99 := stats.Percentile(latencies, 99)

		fmt.Printf("50th percentile (median) duration: %s\n", p50)
		fmt.Printf("90th percentile duration: %s\n", p90)
		fmt.Printf("99th percentile duration: %s\n", p99)

		return nil
	},
}
