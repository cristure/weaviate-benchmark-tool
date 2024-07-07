package generate

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate/entities/models"

	"github.com/cristure/weaviate-benchmark-tool/client"
	"github.com/cristure/weaviate-benchmark-tool/cmd/stats"
)

var (
	objectClassName string
	numberOfObjects int
	numberOfTenants int
	vectorLength    int

	namePrefixObject string
)

func init() {
	objectCmd.PersistentFlags().StringVarP(&objectClassName, "class-name", "c", "", "Class definition name")
	objectCmd.MarkPersistentFlagRequired("class-name")

	objectCmd.PersistentFlags().IntVarP(&numberOfObjects, "objects-number", "o", 0, "Number of objects to generate")
	objectCmd.PersistentFlags().IntVarP(&numberOfTenants, "tenants-number", "t", 0, "Number of tenants to generate")
	objectCmd.PersistentFlags().IntVarP(&vectorLength, "vector-length", "v", 0, "Length of vector to generate")

	objectCmd.PersistentFlags().StringVarP(&namePrefixObject, "name-prefix", "p", "tenant", "Prefix for tenant names")
}

var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Generate object(s) from a class",
	Long:  `Generate a nubmer of object(s) from a class with tenants`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		if objectClassName == "" {
			return fmt.Errorf("must specify file path")
		}

		cls, err := c.Schema().ClassGetter().WithClassName(objectClassName).Do(context.Background())
		if err != nil {
			return fmt.Errorf("failed to get class: %w", err)
		}

		latencies := make(map[string]time.Duration)
		for i := 0; i < numberOfTenants; i++ {
			tenantName := fmt.Sprintf("%s-%d", namePrefixObject, i)

			fmt.Fprintf(os.Stdout, "creating objects in tenant (%s)\n", tenantName)
			for j := 0; j < numberOfObjects; j++ {
				id := uuid.Must(uuid.NewUUID()).String()

				props, err := generatePropertiesForObject(*cls)
				if err != nil {
					return fmt.Errorf("failed to generate object properties: %w", err)
				}
				before := time.Now()
				creator := c.Data().Creator()

				w, err := creator.
					WithClassName(cls.Class).
					WithProperties(props).
					WithID(id).
					WithTenant(tenantName).
					WithVector(generateVector(vectorLength)).
					Do(context.Background())

				if err != nil {
					return fmt.Errorf("failed to create object: %w", err)
				}
				took := time.Since(before)

				fmt.Fprintf(os.Stdout, "successfully created object (%s) in %s\n", w.Object.ID, took)
				latencies[id] = took
			}
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

func generatePropertiesForObject(class models.Class) (map[string]interface{}, error) {
	props := make(map[string]interface{}, 0)
	for _, p := range class.Properties {
		if len(p.DataType) > 1 {
			return nil, errors.New("cannot generate cross-reference types")
		}

		switch p.DataType[0] {
		case "text":
			props[p.Name] = uuid.Must(uuid.NewUUID()).String()
		default:
			panic("unknown data type")
		}
	}
	return props, nil
}

func generateVector(vectorLength int) []float32 {
	vector := []float32{}

	for i := 0; i < vectorLength; i++ {
		vector = append(vector, rand.Float32())
	}

	return vector
}
