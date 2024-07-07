package add

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/weaviate/weaviate/entities/models"

	"github.com/cristure/weaviate-benchmark-tool/client"
)

var (
	objectFilePath string
)

func init() {
	objectCmd.PersistentFlags().StringVarP(&objectFilePath, "file-path", "f", "", "Schema file")
	objectCmd.MarkPersistentFlagRequired("file-path")
}

var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Add a new object",
	Long:  `Add a new object`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		if objectFilePath == "" {
			return fmt.Errorf("must specify file path")
		}

		if _, err := os.Stat(objectFilePath); os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", objectFilePath)
		}

		// Read the file contents
		content, err := os.ReadFile(objectFilePath)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		var obj models.Object
		err = json.Unmarshal(content, &obj)
		if err != nil {
			return fmt.Errorf("failed to unmarshal object: %w", err)
		}

		before := time.Now()
		w, err := c.Data().Creator().
			WithClassName(obj.Class).
			WithProperties(obj.Properties).
			Do(context.Background())

		took := time.Since(before)

		if err != nil {
			return fmt.Errorf("failed to create object: %w", err)
		}

		fmt.Fprintf(os.Stdout, "successfully created object (%s) in %s \n", w.Object.Class, took)
		return nil
	},
}
