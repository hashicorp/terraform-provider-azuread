package msgraph

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type WindowsAutopilotDeploymentProfilesClient struct {
	BaseClient Client
}

// NewWindowsAutopilotDeploymentProfilesClient returns a new WindowsAutopilotDeploymentProfilesClient.
func NewWindowsAutopilotDeploymentProfilesClient() *WindowsAutopilotDeploymentProfilesClient {
	return &WindowsAutopilotDeploymentProfilesClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List returns a list of Windows Autopilot Deployment Profiles, optionally queried using OData.
func (c *WindowsAutopilotDeploymentProfilesClient) List(ctx context.Context, query odata.Query) (*[]WindowsAutopilotDeploymentProfile, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/deviceManagement/windowsAutopilotDeploymentProfiles",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("WindowsAutopilotDeploymentProfilesClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		WindowsAutopilotDeploymentProfiles []WindowsAutopilotDeploymentProfile `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.WindowsAutopilotDeploymentProfiles, status, nil
}

// Create creates a new WindowsAutopilotDeploymentProfile.
func (c *WindowsAutopilotDeploymentProfilesClient) Create(ctx context.Context, profile WindowsAutopilotDeploymentProfile) (*WindowsAutopilotDeploymentProfile, int, error) {
	var status int
	body, err := json.Marshal(profile)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity: "/deviceManagement/windowsAutopilotDeploymentProfiles",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("WindowsAutopilotDeploymentProfilesClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var newWindowsAutopilotDeploymentProfile WindowsAutopilotDeploymentProfile
	if err := json.Unmarshal(respBody, &newWindowsAutopilotDeploymentProfile); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &newWindowsAutopilotDeploymentProfile, status, nil
}

// Get retrieves a WindowsAutopilotDeploymentProfile.
func (c *WindowsAutopilotDeploymentProfilesClient) Get(ctx context.Context, id string, query odata.Query) (*WindowsAutopilotDeploymentProfile, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		OData:                  query,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/deviceManagement/windowsAutopilotDeploymentProfiles/%s", id),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("WindowsAutopilotDeploymentProfilesClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var windowsAutopilotDeploymentProfile WindowsAutopilotDeploymentProfile
	if err := json.Unmarshal(respBody, &windowsAutopilotDeploymentProfile); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &windowsAutopilotDeploymentProfile, status, nil
}

// Update amends an existing WindowsAutopilotDeploymentProfile.
func (c *WindowsAutopilotDeploymentProfilesClient) Update(ctx context.Context, profile WindowsAutopilotDeploymentProfile) (int, error) {
	var status int

	if profile.ID == nil {
		return status, errors.New("cannot update windowsAutopilotDeploymentProfile with nil ID")
	}

	body, err := json.Marshal(profile)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}

	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:                   body,
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusOK, http.StatusNoContent},
		Uri: Uri{
			Entity: fmt.Sprintf("/deviceManagement/windowsAutopilotDeploymentProfiles/%s", *profile.ID),
		},
	})
	if err != nil {
		return status, fmt.Errorf("WindowsAutopilotDeploymentProfilesClient.BaseClient.Patch(): %v", err)
	}

	return status, nil
}

// Delete removes a WindowsAutopilotDeploymentProfile.
func (c *WindowsAutopilotDeploymentProfilesClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusOK, http.StatusNoContent},
		Uri: Uri{
			Entity: fmt.Sprintf("/deviceManagement/windowsAutopilotDeploymentProfiles/%s", id),
		},
	})
	if err != nil {
		return status, fmt.Errorf("WindowsAutopilotDeploymentProfilesClient.BaseClient.Delete(): %v", err)
	}

	return status, nil
}
