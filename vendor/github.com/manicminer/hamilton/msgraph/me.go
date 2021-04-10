package msgraph

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// MeClient performs operations on the authenticated user.
type MeClient struct {
	BaseClient Client
}

// NewMeClient returns a new MeClient.
func NewMeClient(tenantId string) *MeClient {
	return &MeClient{
		BaseClient: NewClient(VersionBeta, tenantId),
	}
}

// Get retrieves information about the authenticated user.
func (c *MeClient) Get(ctx context.Context) (*Me, int, error) {
	var status int
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/me",
			HasTenantId: false,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var me Me
	if err := json.Unmarshal(respBody, &me); err != nil {
		return nil, status, err
	}
	return &me, status, nil
}

// GetProfile retrieves the profile of the authenticated user.
func (c *MeClient) GetProfile(ctx context.Context) (*Me, int, error) {
	var status int
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/me/profile",
			HasTenantId: false,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var me Me
	if err := json.Unmarshal(respBody, &me); err != nil {
		return nil, status, err
	}
	return &me, status, nil
}
