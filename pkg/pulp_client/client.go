package pulp_client

import (
	"context"
	"net/http"
	"time"

	"github.com/content-services/content-sources-backend/pkg/config"
	zest "github.com/content-services/zest/release/v3"
)

type pulpDaoImpl struct {
	client *zest.APIClient
	ctx    context.Context
}

func GetPulpClient() PulpClient {
	ctx := context.WithValue(context.Background(), zest.ContextServerIndex, 0)

	timeout := 60 * time.Second
	transport := &http.Transport{ResponseHeaderTimeout: timeout}
	httpClient := http.Client{Transport: transport, Timeout: timeout}

	pulpConfig := zest.NewConfiguration()
	pulpConfig.HTTPClient = &httpClient
	pulpConfig.Servers = zest.ServerConfigurations{zest.ServerConfiguration{
		URL: config.Get().Clients.Pulp.Server,
	}}
	client := zest.NewAPIClient(pulpConfig)

	auth := context.WithValue(ctx, zest.ContextBasicAuth, zest.BasicAuth{
		UserName: config.Get().Clients.Pulp.Username,
		Password: config.Get().Clients.Pulp.Password,
	})

	// Return DAO instance
	return pulpDaoImpl{
		client: client,
		ctx:    auth,
	}
}
