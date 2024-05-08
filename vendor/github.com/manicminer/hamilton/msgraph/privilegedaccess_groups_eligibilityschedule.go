package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type PrivilegedAccessGroupEligibilityScheduleClient struct {
	BaseClient Client
}

func NewPrivilegedAccessGroupEligibilityScheduleClient() *PrivilegedAccessGroupEligibilityScheduleClient {
	return &PrivilegedAccessGroupEligibilityScheduleClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List retrieves a list of PrivilegedAccessGroupEligibilities
func (c *PrivilegedAccessGroupEligibilityScheduleClient) List(ctx context.Context, query odata.Query) (*[]PrivilegedAccessGroupEligibilitySchedule, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/identityGovernance/privilegedAccess/group/eligibilitySchedules",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupEligibilityScheduleClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		Schedules []PrivilegedAccessGroupEligibilitySchedule `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.Schedules, status, nil
}

// Get retrieves a PrivilegedAccessGroupEligibility
func (c *PrivilegedAccessGroupEligibilityScheduleClient) Get(ctx context.Context, scheduleId string) (*PrivilegedAccessGroupEligibilitySchedule, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/identityGovernance/privilegedAccess/group/eligibilitySchedules/%s", scheduleId),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupEligibilityScheduleClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var schedule PrivilegedAccessGroupEligibilitySchedule
	if err := json.Unmarshal(respBody, &schedule); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &schedule, status, nil
}
