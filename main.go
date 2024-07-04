package main

import (
	"github.com/cristure/weaviate-benchmark-tool/cmd"
)

//TODO: refactor add collection, to be able to do the schema from a file
//TODO: add latencies for add
//TODO: implement list command for objects, schema, ?tenants
//TODO: add setup script on kind with prometheus and grafana dashboard

func main() {
	cmd.Execute()
}
