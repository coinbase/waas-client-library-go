package auth

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math"
	"math/big"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// Authenticator builds a JWT based on the APIKey.
type Authenticator struct {
	apiKey *APIKey
}

// APIKeyClaims holds public claim values for a JWT, as well as a URI.
type APIKeyClaims struct {
	*jwt.Claims
	URI string `json:"uri"`
}

// NewAuthenticator returns a new Authenticator.
func NewAuthenticator(apiKey *APIKey) *Authenticator {
	return &Authenticator{
		apiKey: apiKey,
	}
}

// BuildJWT constructs and returns the JWT as a string along with an error, if any.
func (a *Authenticator) BuildJWT(service, uri string) (string, error) {
	block, _ := pem.Decode([]byte(a.apiKey.PrivateKey))
	if block == nil {
		return "", fmt.Errorf("jwt: Could not decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.ES256, Key: key},
		(&jose.SignerOptions{NonceSource: nonceSource{}}).WithType("JWT").WithHeader("kid", a.apiKey.Name),
	)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	cl := &APIKeyClaims{
		Claims: &jwt.Claims{
			Subject:   a.apiKey.Name,
			Issuer:    "coinbase-cloud",
			NotBefore: jwt.NewNumericDate(time.Now()),
			Expiry:    jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			Audience:  jwt.Audience{service},
		},
		URI: uri,
	}
	jwtString, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}
	return jwtString, nil
}

// nonceSource implements the jose.NonceSource interface. It is used for building
// the JWT.
type nonceSource struct{}

// Nonce calculates and returns nonce as a string and an error, if any.
func (n nonceSource) Nonce() (string, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return "", err
	}
	return r.String(), nil
}
