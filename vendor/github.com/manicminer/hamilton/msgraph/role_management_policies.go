package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type RoleManagementPolicyClient struct {
	BaseClient Client
}

func NewRoleManagementPolicyClient() *RoleManagementPolicyClient {
	return &RoleManagementPolicyClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List retrieves a list of Role Management Policies
func (c *RoleManagementPolicyClient) List(ctx context.Context, query odata.Query) (*[]UnifiedRoleManagementPolicy, int, error) {
	query.Expand = odata.Expand{Relationship: "*"}
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		OData:                  query,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity: "/policies/roleManagementPolicies",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("RoleManagementPolicyClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		UnifiedRoleManagementPolicy []UnifiedRoleManagementPolicy `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.UnifiedRoleManagementPolicy, status, nil
}

// Get retrieves a UnifiedRoleManagementPolicy
func (c *RoleManagementPolicyClient) Get(ctx context.Context, id string) (*UnifiedRoleManagementPolicy, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		OData: odata.Query{
			Expand: odata.Expand{Relationship: "*"},
		},
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/roleManagementPolicies/%s", id),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("RoleDefinitionsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var policy UnifiedRoleManagementPolicy
	if err := json.Unmarshal(respBody, &policy); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &policy, status, nil
}

// Update amends an existing UnifiedRoleManagementPolicy.
func (c *RoleManagementPolicyClient) Update(ctx context.Context, policy UnifiedRoleManagementPolicy) (int, error) {
	var status int

	body, err := json.Marshal(policy)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}

	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/roleManagementPolicies/%s", *policy.ID),
		},
	})
	if err != nil {
		return status, fmt.Errorf("RoleDefinitionsClient.BaseClient.Patch(): %v", err)
	}

	return status, nil
}
