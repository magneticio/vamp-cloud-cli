package adapters

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/magneticio/vamp-cloud-cli/client"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
)

func createRoundTripper(baseTransport http.RoundTripper, apiVersion, apiKey string) *customTransport {

	return &customTransport{

		originalTransport: baseTransport,
		apiVersion:        apiVersion,
		apiKey:            apiKey,
	}
}

type customTransport struct {
	originalTransport http.RoundTripper
	apiVersion        string
	apiKey            string
}

func (c *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {

	r.Header.Add("Accept", fmt.Sprintf("application/vnd.vamp.%v+json", c.apiVersion))
	r.Header.Add("X-Vamp-Token", c.apiKey)

	logging.Info("Received new request", logging.NewPair("method", r.Method), logging.NewPair("url", r.URL))

	resp, err := c.originalTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	logging.Info("Received new response", logging.NewPair("response-status", resp.Status))

	return resp, nil
}

func NewApiClient(apiUrl, apiVersion, apikey string) (*client.Anansi, error) {

	parsedUrl, err := url.Parse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}

	transport := httptransport.New(parsedUrl.Host, parsedUrl.Path, []string{parsedUrl.Scheme})
	transport.Producers[fmt.Sprintf("application/vnd.vamp.%v+json", apiVersion)] = runtime.JSONProducer()
	transport.Consumers[fmt.Sprintf("application/vnd.vamp.%v+json", apiVersion)] = runtime.JSONConsumer()

	customRoundTripper := createRoundTripper(transport.Transport, apiVersion, apikey)

	transport.Transport = customRoundTripper
	return client.New(transport, nil), nil

}
