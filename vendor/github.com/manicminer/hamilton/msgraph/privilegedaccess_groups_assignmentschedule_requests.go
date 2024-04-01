package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type PrivilegedAccessGroupAssignmentScheduleRequestsClient struct {
	BaseClient Client
}

func NewPrivilegedAccessGroupAssignmentScheduleRequestsClient() *PrivilegedAccessGroupAssignmentScheduleRequestsClient {
	return &PrivilegedAccessGroupAssignmentScheduleRequestsClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List retrieves a list of PrivilegedAccessGroupAssignmentScheduleRequests
func (c *PrivilegedAccessGroupAssignmentScheduleRequestsClient) List(ctx context.Context, query odata.Query) (*[]PrivilegedAccessGroupAssignmentScheduleRequest, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		DisablePaging:          query.Top > 0,
		OData:                  query,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity: "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupAssignmentScheduleRequestsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		Requests []PrivilegedAccessGroupAssignmentScheduleRequest `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.Requests, status, nil
}

// Create creates a new PrivilegedAccessGroupAssignmentScheduleRequest.
func (c *PrivilegedAccessGroupAssignmentScheduleRequestsClient) Create(ctx context.Context, request PrivilegedAccessGroupAssignmentScheduleRequest) (*PrivilegedAccessGroupAssignmentScheduleRequest, int, error) {
	var status int

	body, err := json.Marshal(request)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		Body:                   body,
		ValidStatusCodes:       []int{http.StatusCreated},
		Uri: Uri{
			Entity: "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests",
		},
	})
	if err != nil && status != http.StatusNotFound {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupAssignmentScheduleRequestsClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var newRequest PrivilegedAccessGroupAssignmentScheduleRequest
	if err := json.Unmarshal(respBody, &newRequest); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &newRequest, status, nil
}

// Get retrieves a PrivilegedAccessGroupAssignmentScheduleRequest
func (c *PrivilegedAccessGroupAssignmentScheduleRequestsClient) Get(ctx context.Context, requestId string) (*PrivilegedAccessGroupAssignmentScheduleRequest, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/%s", requestId),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupAssignmentScheduleRequestsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var request PrivilegedAccessGroupAssignmentScheduleRequest
	if err := json.Unmarshal(respBody, &request); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &request, status, nil
}

// Cancel cancels a PrivilegedAccessGroupAssignmentScheduleRequest
func (c *PrivilegedAccessGroupAssignmentScheduleRequestsClient) Cancel(ctx context.Context, requestId string) (int, error) {
	_, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity: fmt.Sprintf("/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/%s/cancel", requestId),
		},
	})
	if err != nil {
		return status, fmt.Errorf("PrivilegedAccessGroupAssignmentScheduleRequestsClient.BaseClient.Post(): %v", err)
	}

	return status, nil
}
