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
	classFilePath string
)

func init() {
	classCmd.PersistentFlags().StringVarP(&classFilePath, "file-path", "f", "", "Schema file")
	classCmd.MarkPersistentFlagRequired("file-path")

}

var classCmd = &cobra.Command{
	Use:   "class",
	Short: "Add a new class",
	Long:  `Add a new class`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New()
		if err != nil {
			return fmt.Errorf("failed to init client: %w", err)
		}

		if classFilePath == "" {
			return fmt.Errorf("must specify file path")
		}

		if _, err := os.Stat(classFilePath); os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", classFilePath)
		}

		// Read the file contents
		content, err := os.ReadFile(classFilePath)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		var cls models.Class
		err = json.Unmarshal(content, &cls)
		if err != nil {
			return fmt.Errorf("failed to unmarshal class: %w", err)
		}

		before := time.Now()
		err = c.Schema().ClassCreator().
			WithClass(&cls).
			Do(context.Background())
		took := time.Since(before)

		if err != nil {
			return fmt.Errorf("failed to create class: %w", err)
		}

		fmt.Fprintf(os.Stdout, "successfully created class (%s) in %s \n", cls.Class, took)
		return nil
	},
}
