package heimdall

import (
	"time"

	"github.com/gojek/heimdall/httpclient"
	"github.com/ravielze/oculi/webclient"
)

type (
	clientFactory struct{}
)

func (cf *clientFactory) Create(timeout time.Duration) webclient.Client {
	return &client{
		Heimdall: *httpclient.NewClient(httpclient.WithHTTPTimeout(timeout)),
	}
}

func (cf *clientFactory) CreateWithRetry(timeout time.Duration, retryCount int) webclient.Client {
	return &client{
		Heimdall: *httpclient.NewClient(
			httpclient.WithHTTPTimeout(timeout),
			httpclient.WithRetryCount(retryCount),
		),
	}
}

func New() webclient.WebClient {
	return &clientFactory{}
}
