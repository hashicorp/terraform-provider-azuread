package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// EntitlementRoleDefinitionsClient performs operations on RoleDefinitions.
type EntitlementRoleDefinitionsClient struct {
	BaseClient Client
}

// NewEntitlementRoleDefinitionsClient returns a new EntitlementRoleDefinitionsClient
func NewEntitlementRoleDefinitionsClient() *EntitlementRoleDefinitionsClient {
	return &EntitlementRoleDefinitionsClient{
		BaseClient: NewClient(Version10),
	}
}

// List returns a list of RoleDefinitions
func (c *EntitlementRoleDefinitionsClient) List(ctx context.Context, query odata.Query) (*[]UnifiedRoleDefinition, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/roleManagement/entitlementManagement/roleDefinitions",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("EntitlementRoleDefinitionsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		RoleDefinitions []UnifiedRoleDefinition `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.RoleDefinitions, status, nil
}

// Get retrieves a UnifiedRoleDefinition
func (c *EntitlementRoleDefinitionsClient) Get(ctx context.Context, id string, query odata.Query) (*UnifiedRoleDefinition, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/roleManagement/entitlementManagement/roleDefinitions/%s", id),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("EntitlementRoleDefinitionsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var dirRole UnifiedRoleDefinition
	if err := json.Unmarshal(respBody, &dirRole); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &dirRole, status, nil
}
