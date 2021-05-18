package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AppRoleAssignmentsClient performs operations on AppRoleAssignments.
type AppRoleAssignmentsClient struct {
	BaseClient Client
}

// NewAppRoleAssignmentsClient returns a new AppRoleAssignmentsClient
func NewAppRoleAssignmentsClient(tenantId string) *AppRoleAssignmentsClient {
	return &AppRoleAssignmentsClient{
		BaseClient: NewClient(Version10, tenantId),
	}
}

// List returns a list of app role assignments.
func (c *AppRoleAssignmentsClient) List(ctx context.Context, groupId string) (*[]AppRoleAssignment, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/groups/%s/appRoleAssignments", groupId),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AppRoleAssignmentsClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var data struct {
		AppRoleAssignments []AppRoleAssignment `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &data.AppRoleAssignments, status, nil
}

// Remove removes a app role assignment.
func (c *AppRoleAssignmentsClient) Remove(ctx context.Context, groupId, appRoleAssignmentId string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/groups/%s/appRoleAssignments/%s", groupId, appRoleAssignmentId),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("AppRoleAssignmentsClient.BaseClient.Delete(): %v", err)
	}
	return status, nil
}

// Assign assigns an app role to a group.
func (c *AppRoleAssignmentsClient) Assign(ctx context.Context, groupId, resourceId, appRoleId string) (*AppRoleAssignment, int, error) {
	var status int
	data := struct {
		PrincipalId string `json:"principalId"`
		ResourceId  string `json:"resourceId"`
		AppRoleId   string `json:"appRoleId"`
	}{
		PrincipalId: groupId,
		ResourceId:  resourceId,
		AppRoleId:   appRoleId,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}
	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity:      fmt.Sprintf("/groups/%s/appRoleAssignments", groupId),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AppRoleAssignmentsClient.BaseClient.Post(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var appRoleAssignment AppRoleAssignment
	if err := json.Unmarshal(respBody, &appRoleAssignment); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &appRoleAssignment, status, nil
}
