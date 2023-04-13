package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// ReportsClient Client performs operations on reports.
type ReportsClient struct {
	BaseClient Client
}

// NewReportsClient returns a new ReportsClient.
func NewReportsClient() *ReportsClient {
	return &ReportsClient{
		BaseClient: NewClient(VersionBeta),
	}
}

func (c *ReportsClient) GetCredentialUserRegistrationCount(ctx context.Context, query odata.Query) (*[]CredentialUserRegistrationCount, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/reports/getCredentialUserRegistrationCount",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ReportsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		CredentialUserRegistrationCount []CredentialUserRegistrationCount `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.CredentialUserRegistrationCount, status, nil
}

func (c *ReportsClient) GetCredentialUserRegistrationDetails(ctx context.Context, query odata.Query) (*[]CredentialUserRegistrationDetails, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/reports/credentialUserRegistrationDetails",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ReportsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		CredentialUserRegistrationDetails []CredentialUserRegistrationDetails `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.CredentialUserRegistrationDetails, status, nil
}

func (c *ReportsClient) GetUserCredentialUsageDetails(ctx context.Context, query odata.Query) (*[]UserCredentialUsageDetails, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/reports/userCredentialUsageDetails",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ReportsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		UserCredentialUsageDetails []UserCredentialUsageDetails `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.UserCredentialUsageDetails, status, nil
}

func (c *ReportsClient) GetCredentialUsageSummary(ctx context.Context, period CredentialUsageSummaryPeriod, query odata.Query) (*[]CredentialUsageSummary, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/reports/getCredentialUsageSummary(period='%s')", period),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ReportsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		CredentialUsageSummary []CredentialUsageSummary `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.CredentialUsageSummary, status, nil
}

func (c *ReportsClient) GetAuthenticationMethodsUsersRegisteredByFeature(ctx context.Context, query odata.Query) (*UserRegistrationFeatureSummary, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/reports/authenticationMethods/usersRegisteredByFeature",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ReportsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var userRegistrationFeatureSummary UserRegistrationFeatureSummary
	if err := json.Unmarshal(respBody, &userRegistrationFeatureSummary); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &userRegistrationFeatureSummary, status, nil
}

func (c *ReportsClient) GetAuthenticationMethodsUsersRegisteredByMethod(ctx context.Context, query odata.Query) (*UserRegistrationMethodSummary, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/reports/authenticationMethods/usersRegisteredByMethod",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ReportsClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var userRegistrationMethodSummary UserRegistrationMethodSummary
	if err := json.Unmarshal(respBody, &userRegistrationMethodSummary); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &userRegistrationMethodSummary, status, nil
}
