package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type TokenIssuancePolicyClient struct {
	BaseClient Client
}

// NewTokenIssuancePolicyClient returns a new TokenIssuancePolicyClient
func NewTokenIssuancePolicyClient() *TokenIssuancePolicyClient {
	return &TokenIssuancePolicyClient{
		BaseClient: NewClient(Version10),
	}
}

// Create creates a new TokenIssuancePolicy.
func (c *TokenIssuancePolicyClient) Create(ctx context.Context, policy TokenIssuancePolicy) (*TokenIssuancePolicy, int, error) {
	var status int

	body, err := json.Marshal(policy)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		OData:            odata.Query{Metadata: odata.MetadataFull},
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity: "/policies/tokenIssuancePolicies",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("TokenIssuancePolicyClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var newPolicy TokenIssuancePolicy
	if err := json.Unmarshal(respBody, &newPolicy); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &newPolicy, status, nil
}

// List returns a list of TokenIssuancePolicy, optionally queried using OData.
func (c *TokenIssuancePolicyClient) List(ctx context.Context, query odata.Query) (*[]TokenIssuancePolicy, int, error) {
	query.Metadata = odata.MetadataFull

	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		DisablePaging:    query.Top > 0,
		OData:            query,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: "/policies/tokenIssuancePolicies",
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("TokenIssuancePolicyClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		TokenIssuancePolicies []TokenIssuancePolicy `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.TokenIssuancePolicies, status, nil
}

// Get retrieves a TokenIssuancePolicy.
func (c *TokenIssuancePolicyClient) Get(ctx context.Context, id string, query odata.Query) (*TokenIssuancePolicy, int, error) {
	query.Metadata = odata.MetadataFull

	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		OData:                  query,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/tokenIssuancePolicies/%s", id),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("TokenIssuancePolicyClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var tokenIssuancePolicy TokenIssuancePolicy
	if err := json.Unmarshal(respBody, &tokenIssuancePolicy); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &tokenIssuancePolicy, status, nil
}

// Update amends an existing TokenIssuancePolicy.
func (c *TokenIssuancePolicyClient) Update(ctx context.Context, tokenIssuancePolicy TokenIssuancePolicy) (int, error) {
	var status int

	if tokenIssuancePolicy.ID() == nil {
		return status, fmt.Errorf("cannot update TokenIssuancePolicy with nil ID")
	}

	tokenIssuancePolicyId := *tokenIssuancePolicy.ID()
	tokenIssuancePolicy.Id = nil
	tokenIssuancePolicy.ObjectId = nil

	body, err := json.Marshal(tokenIssuancePolicy)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}

	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:                   body,
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes: []int{
			http.StatusOK,
			http.StatusNoContent,
		},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/tokenIssuancePolicies/%s", tokenIssuancePolicyId),
		},
	})
	if err != nil {
		return status, fmt.Errorf("TokenIssuancePolicyClient.BaseClient.Patch(): %v", err)
	}

	return status, nil
}

// Delete removes a TokenIssuancePolicy.
func (c *TokenIssuancePolicyClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusNoContent},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/tokenIssuancePolicies/%s", id),
		},
	})
	if err != nil {
		return status, fmt.Errorf("TokenIssuancePolicyClient.BaseClient.Delete(): %v", err)
	}

	return status, nil
}
