package webclient

import (
	"io"
	"net/http"
	"time"
)

type (
	WebClient interface {
		Create(timeout time.Duration) Client
		CreateWithRetry(timeout time.Duration, retryCount int) Client
	}

	Client interface {
		Get(url string, headers http.Header, queryString map[string]string) (*http.Response, error)
		Post(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Put(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Patch(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Delete(url string, headers http.Header) (*http.Response, error)
		Do(r *http.Request) (*http.Response, error)
	}
)
