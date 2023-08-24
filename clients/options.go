package clients

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/coinbase/waas-client-library-go/auth"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
)

// WaaSClientConfig stores configuration information about a WaaS client.
type WaaSClientConfig struct {
	Endpoint      string
	ServiceName   string
	APIKey        *auth.APIKey
	HTTPClient    *http.Client
	ClientOptions []option.ClientOption
	Insecure      bool
}

// WaaSClientOption is a function that applies changes to a clientConfig.
type WaaSClientOption func(c *WaaSClientConfig)

// WithEndpoint returns an option to set the service endpoint.
func WithEndpoint(endpoint string) WaaSClientOption {
	return func(c *WaaSClientConfig) {
		c.Endpoint = endpoint
	}
}

// WithAPIKey returns an option to set the Coinbase Cloud API Key.
func WithAPIKey(apiKey *auth.APIKey) WaaSClientOption {
	return func(c *WaaSClientConfig) {
		c.APIKey = apiKey
	}
}

// WithInsecure returns an option to set the insecure flag.
// If insecure is set to true, the client will not use transport authentication.
func WithInsecure(insecure bool) WaaSClientOption {
	return func(c *WaaSClientConfig) {
		c.Insecure = insecure
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
// It uses the given default endpoint if none of the options set it.
func GetConfig(
	serviceName string,
	defaultEndpoint string,
	waasOpts ...WaaSClientOption) (*WaaSClientConfig, error) {
	config := &WaaSClientConfig{}

	for _, waasOpt := range waasOpts {
		waasOpt(config)
	}

	if config.Endpoint == "" {
		config.Endpoint = defaultEndpoint
	}

	if serviceName == "" {
		return nil, errors.New("service name must be provided")
	}

	config.ServiceName = serviceName

	return config, nil
}

// GetClientOptions returns the list of ClientOptions based on the given endpoint, service name,
// and config.
func GetClientOptions(config *WaaSClientConfig) ([]option.ClientOption, error) {
	clientOptions := []option.ClientOption{}

	// Start with explicitly indicated client options.
	if len(config.ClientOptions) > 0 {
		clientOptions = append(clientOptions, config.ClientOptions...)
	}

	if config.Endpoint == "" {
		return nil, errors.New("endpoint must be provided")
	}

	clientOptions = append(clientOptions, option.WithEndpoint(config.Endpoint))

	if config.ServiceName == "" {
		return nil, errors.New("service name must be provided")
	}

	httpClient, err := GetHTTPClient(config.ServiceName, config)
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

	if config.Insecure == false {
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

// LongRunningOperation is the interface for long-running operations that is
// used to create the gax call options for interacting with LROs.
type LongRunningOperation interface {
	Name() string
	PathPrefix() string
}

// LROOptions returns the call options for long-running operations.
// This overrides the gapic generated client `WithPath` call option that ignores the
// path prefix, with a call option that includes the path prefix.
func LROOptions(op LongRunningOperation, version string, opts []gax.CallOption) []gax.CallOption {
	if op.PathPrefix() == "" {
		return opts
	}

	return append(opts, gax.WithPath(fmt.Sprintf("%s/%s/%s", op.PathPrefix(), version, op.Name())))
}
