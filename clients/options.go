package clients

import (
	"errors"
	"net/http"

	"github.cbhq.net/cloud/waas-client-library-go/auth"
	"google.golang.org/api/option"
)

// WaaSClientConfig stores configuration information about a WaaS client.
type WaaSClientConfig struct {
	APIKey        *auth.APIKey
	HTTPClient    *http.Client
	ClientOptions []option.ClientOption
}

// WaaSClientOption is a function that applies changes to a clientConfig.
type WaaSClientOption func(c *WaaSClientConfig)

// WithAPIKey returns an option to set the Coinbase Cloud API Key.
func WithAPIKey(apiKey *auth.APIKey) WaaSClientOption {
	return func(c *WaaSClientConfig) {
		c.APIKey = apiKey
	}
}

// WithHTTPClient returns an option to set the HTTP Client.
func WithHTTPClient(httpClient *http.Client) WaaSClientOption {
	return func(c *WaaSClientConfig) {
		c.HTTPClient = httpClient
	}
}

// WithClientOptions returns an option to set the inner set of ClientOptions.
func WithClientOptions(clientOptions ...option.ClientOption) WaaSClientOption {
	return func(c *WaaSClientConfig) {
		c.ClientOptions = clientOptions
	}
}

// GetConfig returns a WaaSClientConfig with all of the WaaSClientOptions applied.
func GetConfig(waasOpts ...WaaSClientOption) *WaaSClientConfig {
	config := &WaaSClientConfig{}

	for _, waasOpt := range waasOpts {
		waasOpt(config)
	}

	return config
}

// GetClientOptions returns the list of ClientOptions based on the given endpoint, service name,
// and config.
func GetClientOptions(
	endpoint string,
	serviceName string,
	config *WaaSClientConfig,
) ([]option.ClientOption, error) {
	clientOptions := []option.ClientOption{}

	// Start with explicitly indicated client options.
	if len(config.ClientOptions) > 0 {
		clientOptions = append(clientOptions, config.ClientOptions...)
	}

	if endpoint == "" {
		return nil, errors.New("endpoint must be provided")
	}

	clientOptions = append(clientOptions, option.WithEndpoint(endpoint))

	httpClient, err := GetHTTPClient(serviceName, config)
	if err != nil {
		return nil, err
	}

	clientOptions = append(clientOptions, option.WithHTTPClient(httpClient))

	return clientOptions, nil
}

// GetHTTPClient returns the HTTPClient based on the given service name and config.
func GetHTTPClient(serviceName string, config *WaaSClientConfig) (*http.Client, error) {
	var httpClient *http.Client
	if config.HTTPClient != nil {
		httpClient = config.HTTPClient
	} else {
		httpClient = &http.Client{}
	}

	if config.APIKey != nil {
		if httpClient.Transport == nil {
			httpClient.Transport = http.DefaultTransport
		}

		transport, err := NewTransport(
			httpClient.Transport,
			serviceName,
			config.APIKey,
		)
		if err != nil {
			return nil, err
		}

		httpClient.Transport = transport
	}

	return httpClient, nil
}
