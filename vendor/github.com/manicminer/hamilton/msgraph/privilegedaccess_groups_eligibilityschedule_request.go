package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type PrivilegedAccessGroupEligibilityScheduleRequestsClient struct {
	BaseClient Client
}

func NewPrivilegedAccessGroupEligibilityScheduleRequestsClient() *PrivilegedAccessGroupEligibilityScheduleRequestsClient {
	return &PrivilegedAccessGroupEligibilityScheduleRequestsClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List retrieves a list of PrivilegedAccessGroupEligibilityScheduleRequests
func (c *PrivilegedAccessGroupEligibilityScheduleRequestsClient) List(ctx context.Context, query odata.Query) (*[]PrivilegedAccessGroupEligibilityScheduleRequest, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupEligibilityScheduleRequestsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		Requests []PrivilegedAccessGroupEligibilityScheduleRequest `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.Requests, status, nil
}

// Create creates a new PrivilegedAccessGroupEligibilityScheduleRequest.
func (c *PrivilegedAccessGroupEligibilityScheduleRequestsClient) Create(ctx context.Context, request PrivilegedAccessGroupEligibilityScheduleRequest) (*PrivilegedAccessGroupEligibilityScheduleRequest, int, error) {
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
			Entity: "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests",
		},
	})
	if err != nil && status != http.StatusNotFound {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupEligibilityScheduleRequestsClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var newRequest PrivilegedAccessGroupEligibilityScheduleRequest
	if err := json.Unmarshal(respBody, &newRequest); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &newRequest, status, nil
}

// Get retrieves a PrivilegedAccessGroupEligibilityScheduleRequest
func (c *PrivilegedAccessGroupEligibilityScheduleRequestsClient) Get(ctx context.Context, requestId string) (*PrivilegedAccessGroupEligibilityScheduleRequest, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/%s", requestId),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupEligibilityScheduleRequestsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var request PrivilegedAccessGroupEligibilityScheduleRequest
	if err := json.Unmarshal(respBody, &request); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &request, status, nil
}

// Cancel cancels a PrivilegedAccessGroupEligibilityScheduleRequest
func (c *PrivilegedAccessGroupEligibilityScheduleRequestsClient) Cancel(ctx context.Context, requestId string) (int, error) {
	_, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity: fmt.Sprintf("/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/%s/cancel", requestId),
		},
	})
	if err != nil {
		return status, fmt.Errorf("PrivilegedAccessGroupEligibilityScheduleRequestsClient.BaseClient.Post(): %v", err)
	}

	return status, nil
}
