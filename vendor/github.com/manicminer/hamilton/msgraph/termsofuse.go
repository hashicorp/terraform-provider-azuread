package msgraph

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// TermsOfUseAgreementClient performs operations on TermsOfUseAgreement.
type TermsOfUseAgreementClient struct {
	BaseClient Client
}

// NewTermsOfUseAgreementClient returns a new TermsOfUseAgreementClient
func NewTermsOfUseAgreementClient(tenantId string) *TermsOfUseAgreementClient {
	return &TermsOfUseAgreementClient{
		BaseClient: NewClient(VersionBeta, tenantId),
	}
}

// List returns a list of TermsOfUseAgreement agreements, optionally filtered using OData.
func (c *TermsOfUseAgreementClient) List(ctx context.Context, filter string) (*[]TermsOfUseAgreement, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/identityGovernance/termsOfUse/agreements",
			Params:      params,
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("TermsOfUseAgreementClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}
	var data struct {
		TermsOfUseAgreements []TermsOfUseAgreement `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &data.TermsOfUseAgreements, status, nil
}

// Create creates a new TermsOfUse agreement.
func (c *TermsOfUseAgreementClient) Create(ctx context.Context, termsOfUseAgreement TermsOfUseAgreement) (*TermsOfUseAgreement, int, error) {
	var status int
	body, err := json.Marshal(termsOfUseAgreement)

	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}
	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity:      "/identityGovernance/termsOfUse/agreements",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("TermsOfUseAgreementClient.BaseClient.Post(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}
	var newTermsOfUseAgreement TermsOfUseAgreement
	if err := json.Unmarshal(respBody, &newTermsOfUseAgreement); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &newTermsOfUseAgreement, status, nil
}

// Get retrieves an TermsOfUseAgreement agreement.
func (c *TermsOfUseAgreementClient) Get(ctx context.Context, id string) (*TermsOfUseAgreement, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identityGovernance/termsOfUse/agreements/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("TermsOfUseAgreementClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}
	var termsOfUseAgreement TermsOfUseAgreement
	if err := json.Unmarshal(respBody, &termsOfUseAgreement); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &termsOfUseAgreement, status, nil
}

// Update amends an existing TermsOfUseAgreement agreement.
func (c *TermsOfUseAgreementClient) Update(ctx context.Context, termsOfUseAgreement TermsOfUseAgreement) (int, error) {
	var status int
	if termsOfUseAgreement.ID == nil {
		return status, errors.New("cannot update TermsOfUseAgreement agreement with nil ID")
	}

	body, err := json.Marshal(termsOfUseAgreement)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}
	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identityGovernance/termsOfUse/agreements/%s", *termsOfUseAgreement.ID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("TermsOfUseAgreementClient.BaseClient.Patch(): %v", err)
	}
	return status, nil
}

// Delete removes a TermsOfUseAgreement agreement.
func (c *TermsOfUseAgreementClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identityGovernance/termsOfUse/agreements/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("TermsOfUseAgreementClient.BaseClient.Delete(): %v", err)
	}
	return status, nil
}
