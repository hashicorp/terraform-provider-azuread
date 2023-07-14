package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// RoleEligibilityScheduleRequestClient performs operations on RoleEligibilityScheduleRequests.
type RoleEligibilityScheduleRequestClient struct {
	BaseClient Client
}

// NewRoleEligibilityScheduleRequest returns a new RoleEligibilityScheduleRequestClient
func NewRoleEligibilityScheduleRequestClient() *RoleEligibilityScheduleRequestClient {
	return &RoleEligibilityScheduleRequestClient{
		BaseClient: NewClient(Version10),
	}
}

// Get retrieves a UnifiedRoleEligibilityScheduleRequest
func (c *RoleEligibilityScheduleRequestClient) Get(ctx context.Context, id string, query odata.Query) (*UnifiedRoleEligibilityScheduleRequest, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		OData:                  query,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/roleManagement/directory/roleEligibilityScheduleRequests/%s", id),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("RoleEligibilityScheduleRequestClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var roleEligibilityScheduleRequest UnifiedRoleEligibilityScheduleRequest
	if err := json.Unmarshal(respBody, &roleEligibilityScheduleRequest); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &roleEligibilityScheduleRequest, status, nil
}

// List retrieves all UnifiedRoleEligibilityScheduleRequests.
func (c *RoleEligibilityScheduleRequestClient) List(ctx context.Context) (*[]UnifiedRoleEligibilityScheduleRequest, int, error) {
	var status int

	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/roleManagement/directory/roleEligibilityScheduleRequests",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("RoleEligibilityScheduleRequestClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		Value []UnifiedRoleEligibilityScheduleRequest `json:"value"`
	}

	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.Value, status, nil
}

// Create creates a new UnifiedRoleEligibilityScheduleRequest.
func (c *RoleEligibilityScheduleRequestClient) Create(ctx context.Context, resr UnifiedRoleEligibilityScheduleRequest) (*UnifiedRoleEligibilityScheduleRequest, int, error) {
	var status int

	body, err := json.Marshal(resr)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		Body:                   body,
		ValidStatusCodes:       []int{http.StatusCreated},
		Uri: Uri{
			Entity: "/roleManagement/directory/roleEligibilityScheduleRequests",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("RoleEligibilityScheduleRequestClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var newEligibilityScheduleRequest UnifiedRoleEligibilityScheduleRequest
	if err := json.Unmarshal(respBody, &newEligibilityScheduleRequest); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &newEligibilityScheduleRequest, status, nil
}

// Cancel revokes a granted UnifiedRoleEligibilityScheduleRequest
func (c *RoleEligibilityScheduleRequestClient) Cancel(ctx context.Context, id string, query odata.Query) (int, error) {
	_, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		OData:            query,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity: fmt.Sprintf("/roleManagement/directory/roleEligibilityScheduleRequests/%s/cancel", id),
		},
	})
	if err != nil {
		return status, fmt.Errorf("RoleEligibilityScheduleRequestClient.BaseClient.Post(): %v", err)
	}

	return status, nil
}
