package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type RoleManagementPolicyAssignmentClient struct {
	BaseClient Client
}

func NewRoleManagementPolicyAssignmentClient() *RoleManagementPolicyAssignmentClient {
	return &RoleManagementPolicyAssignmentClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List retrieves a list of Role Management Policies
func (c *RoleManagementPolicyAssignmentClient) List(ctx context.Context, query odata.Query) (*[]UnifiedRoleManagementPolicyAssignment, int, error) {
	query.Expand = odata.Expand{Relationship: "*"}
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/policies/roleManagementPolicyAssignments",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("RoleManagementPolicyAssignmentClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		UnifiedRoleManagementPolicyAssignment []UnifiedRoleManagementPolicyAssignment `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.UnifiedRoleManagementPolicyAssignment, status, nil
}

// Get retrieves a UnifiedRoleManagementPolicy
func (c *RoleManagementPolicyAssignmentClient) Get(ctx context.Context, id string) (*UnifiedRoleManagementPolicyAssignment, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		OData: odata.Query{
			Expand: odata.Expand{Relationship: "*"},
		},
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/roleManagementPolicyAssignments/%s", id),
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

	var assign UnifiedRoleManagementPolicyAssignment
	if err := json.Unmarshal(respBody, &assign); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &assign, status, nil
}
