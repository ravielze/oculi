package heimdall

import (
	"fmt"
	"io"
	"net/http"
	neturl "net/url"

	"github.com/gojek/heimdall/httpclient"
	errConsts "github.com/ravielze/oculi/constant/errors"
	errUtils "github.com/ravielze/oculi/errors"
)

type (
	client struct {
		Heimdall httpclient.Client
	}

	WebClientErrorDetails struct {
		ResponseStatus     string `json:"response_status"`
		ResponseStatusCode int    `json:"response_status_code"`
	}
)

func (c *client) Get(url string, headers http.Header, queryString map[string]string) (*http.Response, error) {

	if len(queryString) > 0 {
		format := "?%s=%s"
		for key, value := range queryString {
			url += fmt.Sprintf(format, key, neturl.QueryEscape(value))
			format = "&%s=%s"
		}
	}

	response, err := c.Heimdall.Get(url, headers)
	return c.checkClientError(response, err)
}

func (c *client) Post(url string, body io.Reader, headers http.Header) (*http.Response, error) {

	response, err := c.Heimdall.Post(url, body, headers)
	return c.checkClientError(response, err)
}

func (c *client) Put(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	response, err := c.Heimdall.Put(url, body, headers)
	return c.checkClientError(response, err)
}

func (c *client) Patch(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	response, err := c.Heimdall.Patch(url, body, headers)
	return c.checkClientError(response, err)
}

func (c *client) Delete(url string, headers http.Header) (*http.Response, error) {
	response, err := c.Heimdall.Delete(url, headers)
	return c.checkClientError(response, err)
}

func (c *client) Do(r *http.Request) (*http.Response, error) {
	response, err := c.Heimdall.Do(r)
	return c.checkClientError(response, err)
}

func (c *client) checkClientError(response *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return response,
			errUtils.NewDetailedErrors(
				errConsts.WebClientError, err.Error(),
			)
	} else if response.StatusCode >= http.StatusBadRequest {
		err = errUtils.NewDetailedErrors(
			errConsts.WebClientError,
			WebClientErrorDetails{
				ResponseStatus:     response.Status,
				ResponseStatusCode: response.StatusCode,
			},
		)
		return response, err
	}

	return response, nil
}
