package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cristure/weaviate-benchmark-tool/cmd/add"
	"github.com/cristure/weaviate-benchmark-tool/cmd/config"
)

func init() {
	initRootCommand()
	//initRandomText()
	//initDataset()
}

func initRootCommand() {
	RootCmd.PersistentFlags().StringVarP(&config.GlobalConfig.Host, "host", "H", "localhost", "Host to connect to")
	RootCmd.PersistentFlags().StringVarP(&config.GlobalConfig.Scheme, "scheme", "S", "http", "Schme to connect to")
	RootCmd.AddCommand(add.AddCmd)
}

var RootCmd = &cobra.Command{
	Use:   "weaviate-benchmark-tool",
	Short: "Weaviate Benchmarker",
	Long:  `A Weaviate Benchmarker`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("running the root command, see help or -h for available commands\n")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}