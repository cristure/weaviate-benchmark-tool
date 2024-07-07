package list

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/cristure/weaviate-benchmark-tool/client"
)

func init() {
	shardCmd.PersistentFlags().StringVarP(&className, "class-name", "c", "", "Name of the class to add the tenant to")
	shardCmd.MarkPersistentFlagRequired("class-name")
}

var shardCmd = &cobra.Command{
	Use:   "shard",
	Short: "List shards",
	Long:  `List the shards from a class`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		before := time.Now()
		shards, err := c.Schema().ShardsGetter().
			WithClassName(className).
			Do(context.Background())

		if err != nil {
			return fmt.Errorf("failed to list shards: %w", err)
		}

		took := time.Since(before)

		fmt.Fprintf(os.Stdout, "class (%s) has %d shards. List took %s\n", className, len(shards), took)
		var sum int64
		for _, shard := range shards {
			sum += shard.VectorQueueSize
		}

		fmt.Fprintf(os.Stdout, "Average vector queue size: %d\n", sum/int64(len(shards)))
		return nil
	},
}
