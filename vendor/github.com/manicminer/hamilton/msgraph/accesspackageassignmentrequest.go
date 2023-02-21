package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/manicminer/hamilton/odata"
)

type AccessPackageAssignmentRequestClient struct {
	BaseClient Client
}

func NewAccessPackageAssignmentRequestClient(tenantId string) *AccessPackageAssignmentRequestClient {
	return &AccessPackageAssignmentRequestClient{
		BaseClient: NewClient(Version10, tenantId),
	}
}

// List will list all access package assignment requests
func (c *AccessPackageAssignmentRequestClient) List(ctx context.Context, query odata.Query) (*[]AccessPackageAssignmentRequest, int, error) {
	entity := getEntity(c.BaseClient.ApiVersion)

	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      entity,
			Params:      query.Values(),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AccessPackageAssignmentRequestClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		AccessPackageAssignmentRequest []AccessPackageAssignmentRequest `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.AccessPackageAssignmentRequest, status, nil
}

// Get will get an Access Package request
func (c *AccessPackageAssignmentRequestClient) Get(ctx context.Context, id string) (*AccessPackageAssignmentRequest, int, error) {
	entity := getEntity(c.BaseClient.ApiVersion)
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("%s/%s", entity, id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AccessPackageAssignmentRequestClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var accessPackageAssignmentRequest AccessPackageAssignmentRequest
	if err := json.Unmarshal(respBody, &accessPackageAssignmentRequest); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &accessPackageAssignmentRequest, status, nil
}

// Create will create an access package request
func (c *AccessPackageAssignmentRequestClient) Create(ctx context.Context, accessPackageAssignementRequest AccessPackageAssignmentRequest) (*AccessPackageAssignmentRequest, int, error) {
	var status int
	entity := getEntity(c.BaseClient.ApiVersion)
	body, err := json.Marshal(accessPackageAssignementRequest)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      entity,
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AccessPackageAssignmentRequestClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var newAccessPackageAssignmentRequest AccessPackageAssignmentRequest
	if err := json.Unmarshal(respBody, &newAccessPackageAssignmentRequest); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &newAccessPackageAssignmentRequest, status, nil
}

// Delete will delete an access package request
func (c *AccessPackageAssignmentRequestClient) Delete(ctx context.Context, id string) (int, error) {
	entity := getEntity(c.BaseClient.ApiVersion)
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("%s/%s", entity, id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("AccessPackageAssignmentPolicyClient.BaseClient.Delete(): %v", err)
	}

	return status, nil

}

// Cancel will cancel a request is in a cancellable state
func (c *AccessPackageAssignmentRequestClient) Cancel(ctx context.Context, id string) (int, error) {
	var status int
	entity := getEntity(c.BaseClient.ApiVersion)
	_, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("%s/%s/cancel", entity, id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("AccessPackageAssignmentRequestClient.BaseClient.Post(): %v", err)
	}

	return status, nil
}

// Reprocess re-processes an access package assignment request
func (c *AccessPackageAssignmentRequestClient) Reprocess(ctx context.Context, id string) (int, error) {
	var status int
	entity := getEntity(c.BaseClient.ApiVersion)
	_, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		ValidStatusCodes: []int{http.StatusAccepted},
		Uri: Uri{
			Entity:      fmt.Sprintf("/%s/%s/reprocess", entity, id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("AccessPackageAssignmentRequestClient.BaseClient.Post(): %v", err)
	}

	return status, nil
}

func getEntity(api ApiVersion) string {
	if api == VersionBeta {
		return "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests"
	}
	return "/identityGovernance/entitlementManagement/assignmentRequests"
}
