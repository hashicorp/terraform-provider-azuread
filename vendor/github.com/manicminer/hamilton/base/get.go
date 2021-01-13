package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/manicminer/hamilton/base/odata"
)

// GetHttpRequestInput configures a GET request.
type GetHttpRequestInput struct {
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
	rawUri           string
}

// GetValidStatusCodes returns a []int of status codes considered valid for a GET request.
func (i GetHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a GET request is considered valid.
func (i GetHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Get performs a GET request.
func (c Client) Get(ctx context.Context, input GetHttpRequestInput) (*http.Response, int, *odata.OData, error) {
	var status int

	// Check for a raw uri, else build one from the Uri field
	url := input.rawUri
	if url == "" {
		var err error
		url, err = c.buildUri(input.Uri)
		if err != nil {
			return nil, status, nil, fmt.Errorf("unable to make request: %v", err)
		}
	}

	// Build a new request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, status, nil, err
	}

	// Perform the request
	resp, status, o, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, o, err
	}

	// Check for json content before handling pagination
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	if strings.HasPrefix(contentType, "application/json") {
		// Read the response body and close it
		respBody, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		// Unmarshall firstOdata
		var firstOdata odata.OData
		if err := json.Unmarshal(respBody, &firstOdata); err != nil {
			return nil, status, o, err
		}

		if firstOdata.NextLink == nil || firstOdata.Value == nil {
			// No more pages, reassign response body and return
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))
			return resp, status, o, nil
		}

		// Get the next page, recursively
		nextInput := input
		nextInput.rawUri = *firstOdata.NextLink
		nextResp, status, o, err := c.Get(ctx, nextInput)
		if err != nil {
			return resp, status, o, err
		}

		// Read the next page response body and close it
		nextRespBody, _ := ioutil.ReadAll(nextResp.Body)
		nextResp.Body.Close()

		// Unmarshall firstOdata from the next page
		var nextOdata odata.OData
		if err := json.Unmarshal(nextRespBody, &nextOdata); err != nil {
			return resp, status, o, err
		}

		if nextOdata.Value != nil {
			// Next page has results, append to current page
			value := append(*firstOdata.Value, *nextOdata.Value...)
			nextOdata.Value = &value
		}

		// Marshal the entire result, along with fields from the final page
		newJson, err := json.Marshal(nextOdata)
		if err != nil {
			return resp, status, o, err
		}

		// Reassign the response body
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(newJson))
	}

	return resp, status, o, nil
}
