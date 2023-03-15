package clients

import (
	"fmt"
	"net/http"

	"github.com/coinbase/waas-client-library-go/auth"
)

// transport implements the http.RoundTripper interface for use by WaaS HTTP clients.
type transport struct {
	roundTripper  http.RoundTripper
	authenticator *auth.Authenticator
	serviceName   string
}

// NewTransport returns a new transport based on the given inputs.
func NewTransport(
	roundTripper http.RoundTripper,
	serviceName string,
	apiKey *auth.APIKey,
) (*transport, error) {
	return &transport{
		roundTripper:  roundTripper,
		authenticator: auth.NewAuthenticator(apiKey),
		serviceName:   serviceName,
	}, nil
}

// RoundTrip implements the http.RoundTripper interface and wraps
// the base round tripper with logic to inject the API key auth-based HTTP headers
// into the request. Reference: https://pkg.go.dev/net/http#RoundTripper
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	jwt, err := t.authenticator.BuildJWT(
		t.serviceName,
		fmt.Sprintf("%s %s%s", req.Method, req.URL.Host, req.URL.Path),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))
	return t.roundTripper.RoundTrip(req)
}
