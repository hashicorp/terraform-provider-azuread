package utils

import (
	"net/http"

	"github.com/Azure/go-autorest/autorest"
)

func ResponseWasNotFound(resp autorest.Response) bool {
	return ResponseWasStatusCode(resp, http.StatusNotFound)
}

func ResponseWasStatusCode(resp autorest.Response, statusCode int) bool { // nolint: unparam
	if r := resp.Response; r != nil {
		if r.StatusCode == statusCode {
			return true
		}
	}

	return false
}
