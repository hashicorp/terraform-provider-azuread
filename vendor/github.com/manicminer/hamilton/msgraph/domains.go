package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DomainsClient performs operations on Domains.
type DomainsClient struct {
	BaseClient Client
}

// NewDomainsClient returns a new DomainsClient.
func NewDomainsClient(tenantId string) *DomainsClient {
	return &DomainsClient{
		BaseClient: NewClient(Version10, tenantId),
	}
}

// List returns a list of Domains.
func (c *DomainsClient) List(ctx context.Context) (*[]Domain, int, error) {
	var status int
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/domains",
			HasTenantId: true,
		},
	})

	if err != nil {
		return nil, status, err
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	var data struct {
		Domains []Domain `json:"value"`
	}

	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}

	return &data.Domains, status, nil
}

// Get retrieves a Domain.
func (c *DomainsClient) Get(ctx context.Context, id string) (*Domain, int, error) {
	var status int
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/domains/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var domain Domain
	if err := json.Unmarshal(respBody, &domain); err != nil {
		return nil, status, err
	}
	return &domain, status, nil
}
