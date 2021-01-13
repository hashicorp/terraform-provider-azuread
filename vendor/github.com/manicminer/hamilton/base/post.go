package base

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/manicminer/hamilton/base/odata"
)

// PostHttpRequestInput configures a POST request.
type PostHttpRequestInput struct {
	Body             []byte
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
}

// GetValidStatusCodes returns a []int of status codes considered valid for a POST request.
func (i PostHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a POST request is considered valid.
func (i PostHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Post performs a POST request.
func (c Client) Post(ctx context.Context, input PostHttpRequestInput) (*http.Response, int, *odata.OData, error) {
	var status int
	url, err := c.buildUri(input.Uri)
	if err != nil {
		return nil, status, nil, fmt.Errorf("unable to make request: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(input.Body))
	if err != nil {
		return nil, status, nil, err
	}
	resp, status, o, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, o, err
	}
	return resp, status, o, nil
}
