package adapters

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/magneticio/vamp-cloud-cli/client"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
)

func createRoundTripper(baseTransport http.RoundTripper, apiVersion, apikey string) *customTransport {

	return &customTransport{
		//originalTransport: http.DefaultTransport,
		originalTransport: baseTransport,
	}
}

type customTransport struct {
	originalTransport http.RoundTripper
	apiVersion        string
	apiKey            string
}

func (c *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {

	r.Header.Add("Accept", fmt.Sprintf("application/vnd.vamp.%v+json", c.apiVersion))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Vamp-Token", c.apiKey)

	logging.Info("Receiveived new request", logging.NewPair("method", r.Method), logging.NewPair("url", r.URL))

	resp, err := c.originalTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewApiClient(host, basePath, apiVersion, apikey string) *client.Anansi {

	transport := httptransport.New(host, basePath, []string{"http", "https"})
	transport.Producers[fmt.Sprintf("application/vnd.vamp.%v+json", apiVersion)] = runtime.JSONProducer()

	customRoundTripper := createRoundTripper(transport.Transport, apiVersion, apikey)

	transport.Transport = customRoundTripper
	return client.New(transport, nil)

}
