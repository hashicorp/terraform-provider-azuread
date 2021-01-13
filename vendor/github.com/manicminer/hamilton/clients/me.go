package clients

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/manicminer/hamilton/base"
	"github.com/manicminer/hamilton/models"
)

// MeClient performs operations on the authenticated user.
type MeClient struct {
	BaseClient base.Client
}

// NewMeClient returns a new MeClient.
func NewMeClient(tenantId string) *MeClient {
	return &MeClient{
		BaseClient: base.NewClient(base.VersionBeta, tenantId),
	}
}

// Get retrieves information about the authenticated user.
func (c *MeClient) Get(ctx context.Context) (*models.Me, int, error) {
	var status int
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      "/me",
			HasTenantId: false,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var me models.Me
	if err := json.Unmarshal(respBody, &me); err != nil {
		return nil, status, err
	}
	return &me, status, nil
}

// GetProfile retrieves the profile of the authenticated user.
func (c *MeClient) GetProfile(ctx context.Context) (*models.Me, int, error) {
	var status int
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      "/me/profile",
			HasTenantId: false,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var me models.Me
	if err := json.Unmarshal(respBody, &me); err != nil {
		return nil, status, err
	}
	return &me, status, nil
}
