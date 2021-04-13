package msgraph

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ConditionalAccessPolicyClient performs operations on ConditionalAccessPolicy.
type ConditionalAccessPolicyClient struct {
	BaseClient Client
}

// NewConditionalAccessPolicyClient returns a new ConditionalAccessPolicyClient
func NewConditionalAccessPolicyClient(tenantId string) *ConditionalAccessPolicyClient {
	return &ConditionalAccessPolicyClient{
		BaseClient: NewClient(VersionBeta, tenantId),
	}
}

// List returns a list of ConditionalAccessPolicys, optionally filtered using OData.
func (c *ConditionalAccessPolicyClient) List(ctx context.Context, filter string) (*[]ConditionalAccessPolicy, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/identity/conditionalAccess/policies",
			Params:      params,
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ConditionalAccessPolicyClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var data struct {
		ConditionalAccessPolicys []ConditionalAccessPolicy `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &data.ConditionalAccessPolicys, status, nil
}

// Create creates a new ConditionalAccessPolicy.
func (c *ConditionalAccessPolicyClient) Create(ctx context.Context, conditionalAccessPolicy ConditionalAccessPolicy) (*ConditionalAccessPolicy, int, error) {
	var status int
	body, err := json.Marshal(conditionalAccessPolicy)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}
	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity:      "/identity/conditionalAccess/policies",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ConditionalAccessPolicyClient.BaseClient.Post(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var newConditionalAccessPolicy ConditionalAccessPolicy
	if err := json.Unmarshal(respBody, &newConditionalAccessPolicy); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &newConditionalAccessPolicy, status, nil
}

// Get retrieves an ConditionalAccessPolicy.
func (c *ConditionalAccessPolicyClient) Get(ctx context.Context, id string) (*ConditionalAccessPolicy, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identity/conditionalAccess/policies/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ConditionalAccessPolicyClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var conditionalAccessPolicy ConditionalAccessPolicy
	if err := json.Unmarshal(respBody, &conditionalAccessPolicy); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &conditionalAccessPolicy, status, nil
}

// Update amends an existing ConditionalAccessPolicy.
func (c *ConditionalAccessPolicyClient) Update(ctx context.Context, conditionalAccessPolicy ConditionalAccessPolicy) (int, error) {
	var status int
	if conditionalAccessPolicy.ID == nil {
		return status, errors.New("cannot update conditionalAccessPolicy with nil ID")
	}

	body, err := json.Marshal(conditionalAccessPolicy)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}
	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identity/conditionalAccess/policies/%s", *conditionalAccessPolicy.ID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("ConditionalAccessPolicyClient.BaseClient.Patch(): %v", err)
	}
	return status, nil
}

// Delete removes a ConditionalAccessPolicy.
func (c *ConditionalAccessPolicyClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identity/conditionalAccess/policies/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("ConditionalAccessPolicyClient.BaseClient.Delete(): %v", err)
	}
	return status, nil
}
