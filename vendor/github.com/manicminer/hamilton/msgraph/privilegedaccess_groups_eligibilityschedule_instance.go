package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type PrivilegedAccessGroupEligibilityScheduleInstancesClient struct {
	BaseClient Client
}

func NewPrivilegedAccessGroupEligibilityScheduleInstancesClient() *PrivilegedAccessGroupEligibilityScheduleInstancesClient {
	return &PrivilegedAccessGroupEligibilityScheduleInstancesClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List retrieves a list of PrivilegedAccessGroupEligibilityScheduleInstances
func (c *PrivilegedAccessGroupEligibilityScheduleInstancesClient) List(ctx context.Context, query odata.Query) (*[]PrivilegedAccessGroupEligibilityScheduleInstance, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		Instances []PrivilegedAccessGroupEligibilityScheduleInstance `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.Instances, status, nil
}

// Get retrieves a PrivilegedAccessGroupEligibilityScheduleInstance
func (c *PrivilegedAccessGroupEligibilityScheduleInstancesClient) Get(ctx context.Context, instanceId string) (*PrivilegedAccessGroupEligibilityScheduleInstance, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/%s", instanceId),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("PrivilegedAccessGroupClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var instance PrivilegedAccessGroupEligibilityScheduleInstance
	if err := json.Unmarshal(respBody, &instance); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &instance, status, nil
}
