package auth

import (
	"errors"
	"os"
)

const (
	// nameEnvVar is the environment variable for the Coinbase Cloud API Key name.
	nameEnvVar = "COINBASE_CLOUD_API_KEY_NAME"

	// privateKeyEnvVar is the environment variable for the Coinbase Cloud API Key private key.
	privateKeyEnvVar = "COINBASE_CLOUD_API_KEY_PRIVATE_KEY"
)

// APIKey represents a Coinbase Cloud API Key.
type APIKey struct {
	Name       string
	PrivateKey string
}

// apiKeyConfig contains configuration information for constructing the API Key.
type apiKeyConfig struct {
	loadApiKeyFromEnv bool
	apiKeyName        string
	apiKeyPrivateKey  string
}

// apiKeyOption is a function that applies changes to a apiKeyConfig.
type apiKeyOption func(t *apiKeyConfig)

// WithAPIKeyName returns an option to set the API Key.
func WithAPIKeyName(apiKeyName, apiKeyPrivateKey string) apiKeyOption {
	return func(t *apiKeyConfig) {
		t.apiKeyName = apiKeyName
		t.apiKeyPrivateKey = apiKeyPrivateKey
	}
}

// WithLoadAPIKeyFromEnv returns an option to set whether or not to load the API
// Key from environment variables. If the API Key name and private key are both set,
// they take precedence over the environment variables.
func WithLoadAPIKeyFromEnv(loadAPIKeyFromEnv bool) apiKeyOption {
	return func(t *apiKeyConfig) {
		t.loadApiKeyFromEnv = loadAPIKeyFromEnv
	}
}

// NewAPIKey creates a new Coinbase Cloud API Key based on the provided options.
func NewAPIKey(opts ...apiKeyOption) (*APIKey, error) {
	config := &apiKeyConfig{}

	for _, opt := range opts {
		opt(config)
	}

	if config.apiKeyName != "" && config.apiKeyPrivateKey != "" {
		return &APIKey{
			Name:       config.apiKeyName,
			PrivateKey: config.apiKeyPrivateKey,
		}, nil
	}

	return loadAPIKeyFromEnv()
}

// loadAPIKeyFromEnv loads a new Coinbase Cloud API Key from environment variables.
func loadAPIKeyFromEnv() (*APIKey, error) {
	name := os.Getenv(nameEnvVar)
	if name == "" {
		return nil, errors.New("environment variable COINBASE_CLOUD_API_KEY_NAME must be set")
	}

	privateKey := os.Getenv(privateKeyEnvVar)
	if privateKey == "" {
		return nil, errors.New("environment variable COINBASE_CLOUD_API_KEY_PRIVATE_KEY must be set")
	}

	return &APIKey{
		Name:       name,
		PrivateKey: privateKey,
	}, nil
}
