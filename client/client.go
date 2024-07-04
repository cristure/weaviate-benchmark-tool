package client

import (
	"fmt"
	"os"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/auth"

	"github.com/cristure/weaviate-benchmark-tool/cmd/config"
)

func New() (*weaviate.Client, error) {
	var (
		err        error
		authConfig auth.Config
	)

	apiKey, ok := os.LookupEnv("WEAVIATE_API_KEY")
	if ok {
		authConfig = auth.ApiKey{Value: apiKey}
	}

	client, err := weaviate.NewClient(weaviate.Config{
		Host:       config.GlobalConfig.Host,
		Scheme:     config.GlobalConfig.Scheme,
		AuthConfig: authConfig,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init client: %v", err)
	}

	return client, nil
}
