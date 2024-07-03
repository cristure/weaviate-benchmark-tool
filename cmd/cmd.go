package cmd

import (
	"github.com/jessevdk/go-flags"

	"github.com/cristure/weaviate-benchmark-tool/config"
)

var (
	opts   config.Options
	parser = flags.NewParser(&opts, flags.Default)

	client client.C
)

func RunCommand() {
	_, err := parser.Parse()
	if err != nil {
		return
	}
}
