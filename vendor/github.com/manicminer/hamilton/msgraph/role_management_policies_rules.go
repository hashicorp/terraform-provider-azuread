package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type RoleManagementPolicyRuleClient struct {
	BaseClient Client
}

func NewRoleManagementPolicyRuleClient() *RoleManagementPolicyRuleClient {
	return &RoleManagementPolicyRuleClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List retrieves a list of Rules from a Role Management Policy
func (c *RoleManagementPolicyRuleClient) List(ctx context.Context, policyId string) (*[]UnifiedRoleManagementPolicyRule, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/roleManagementPolicies/%s/rules", policyId),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("RoleManagementPolicyRuleClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		UnifiedRoleManagementPolicyRule []UnifiedRoleManagementPolicyRule `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.UnifiedRoleManagementPolicyRule, status, nil
}

// Get retrieves a UnifiedRoleManagementPolicyRule
func (c *RoleManagementPolicyRuleClient) Get(ctx context.Context, policyId, ruleId string) (*UnifiedRoleManagementPolicyRule, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		OData: odata.Query{
			Expand: odata.Expand{Relationship: "*"},
		},
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/roleManagementPolicies/%s/rules/%s", policyId, ruleId),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("RoleManagementPolicyRuleClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var rule UnifiedRoleManagementPolicyRule
	if err := json.Unmarshal(respBody, &rule); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &rule, status, nil
}

// Update amends an existing UnifiedRoleManagementPolicyRule.
func (c *RoleManagementPolicyRuleClient) Update(ctx context.Context, policyId string, rule UnifiedRoleManagementPolicyRule) (int, error) {
	var status int

	body, err := json.Marshal(rule)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}

	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/policies/roleManagementPolicies/%s/rules/%s", policyId, *rule.ID),
		},
	})
	if err != nil {
		return status, fmt.Errorf("RoleManagementPolicyRuleClient.BaseClient.Patch(): %v", err)
	}

	return status, nil
}
