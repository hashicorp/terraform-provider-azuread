package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// EntitlementRoleAssignmentsClient performs operations on RoleAssignments.
type EntitlementRoleAssignmentsClient struct {
	BaseClient Client
}

// NewEntitlementRoleAssignmentsClient returns a new EntitlementRoleAssignmentsClient
func NewEntitlementRoleAssignmentsClient(tenantId string) *EntitlementRoleAssignmentsClient {
	return &EntitlementRoleAssignmentsClient{
		BaseClient: NewClient(Version10, tenantId),
	}
}

// List returns a list of RoleAssignments
func (c *EntitlementRoleAssignmentsClient) List(ctx context.Context, query odata.Query) (*[]UnifiedRoleAssignment, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/roleManagement/entitlementManagement/roleAssignments",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("EntitlementRoleAssignmentsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		RoleAssignments []UnifiedRoleAssignment `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.RoleAssignments, status, nil
}

// Get retrieves a UnifiedRoleAssignment
func (c *EntitlementRoleAssignmentsClient) Get(ctx context.Context, id string, query odata.Query) (*UnifiedRoleAssignment, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		OData:                  query,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/roleManagement/entitlementManagement/roleAssignments/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("EntitlementRoleAssignmentsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var dirRole UnifiedRoleAssignment
	if err := json.Unmarshal(respBody, &dirRole); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &dirRole, status, nil
}

// Create creates a new UnifiedRoleAssignment.
func (c *EntitlementRoleAssignmentsClient) Create(ctx context.Context, roleAssignment UnifiedRoleAssignment) (*UnifiedRoleAssignment, int, error) {
	var status int

	body, err := json.Marshal(roleAssignment)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/roleManagement/entitlementManagement/roleAssignments",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("EntitlementRoleAssignmentsClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var newRoleAssignment UnifiedRoleAssignment
	if err := json.Unmarshal(respBody, &newRoleAssignment); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &newRoleAssignment, status, nil
}

// Delete removes a UnifiedRoleAssignment.
func (c *EntitlementRoleAssignmentsClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/roleManagement/entitlementManagement/roleAssignments/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("RoleAssignments.BaseClient.Get(): %v", err)
	}

	return status, nil
}
