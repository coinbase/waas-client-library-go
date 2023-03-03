package clients

import (
	"net/http"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
)

// HTTPCode attempts to return the HTTP status code for the given error.
// If it cannot be determined, it returns a 500 (Internal Server Error).
func HTTPCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if googleapiErr, ok := err.(*googleapi.Error); ok {
		return googleapiErr.Code
	}

	return http.StatusInternalServerError
}

// UnwrapError attempts to unwrap the passed error and return a googleapi error.
// It is expected that the passed error is of type apierror.APIError.
// If it is not, it returns the passed error unchanged.
func UnwrapError(err error) error {
	if err == nil {
		return nil
	}

	if apiErr, ok := err.(*apierror.APIError); ok { //nolint:errorlint
		// Unwrap the the `apierror` to extract the HTTP status code and original error message.
		unwrappedErr := apiErr.Unwrap()
		if googleapiErr, ok := unwrappedErr.(*googleapi.Error); ok { //nolint:errorlint
			return googleapiErr
		}
	}

	return err
}
