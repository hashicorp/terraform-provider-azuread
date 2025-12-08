// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"fmt"

	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/glueckkanja/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ sdk.DataSource = GroupRoleManagementPolicyDataSource{}

type GroupRoleManagementPolicyDataSourceModel struct {
	Description string `tfschema:"description"`
	DisplayName string `tfschema:"display_name"`
	GroupId     string `tfschema:"group_id"`
	RoleId      string `tfschema:"role_id"`
}

type GroupRoleManagementPolicyDataSource struct{}

func (r GroupRoleManagementPolicyDataSource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_id": {
			Description:  "ID of the group to which this policy is assigned",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.IsUUID,
		},

		"role_id": {
			Description:  "The ID of the role of this policy to the group",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(possibleValuesForRoleDefinitionId, false),
		},
	}
}

func (r GroupRoleManagementPolicyDataSource) Attributes() map[string]*schema.Schema {
	return map[string]*pluginsdk.Schema{
		"display_name": {
			Description: "The display name of the policy",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"description": {
			Description: "Description of the policy",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (r GroupRoleManagementPolicyDataSource) ModelObject() interface{} {
	return &GroupRoleManagementPolicyDataSourceModel{}
}

func (r GroupRoleManagementPolicyDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			policyClient := metadata.Client.Policies.RoleManagementPolicyClient
			assignmentClient := metadata.Client.Policies.RoleManagementPolicyAssignmentClient

			var model GroupRoleManagementPolicyDataSourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			policyId, err := getPolicyId(ctx, metadata, model.GroupId, model.RoleId)
			if err != nil {
				return fmt.Errorf("determining Policy ID: %v", err)
			}

			id := stable.NewPolicyRoleManagementPolicyID(policyId.ID())

			policyOptions := rolemanagementpolicy.GetRoleManagementPolicyOperationOptions{
				Expand: &odata.Expand{
					Relationship: "*",
				},
			}

			policyResp, err := policyClient.GetRoleManagementPolicy(ctx, id, policyOptions)
			if err != nil {
				return fmt.Errorf("retrieving %s: %v", id, err)
			}

			policy := policyResp.Model
			if policy == nil {
				return fmt.Errorf("retrieving %s: API error, model was nil", id)
			}

			options := rolemanagementpolicyassignment.ListRoleManagementPolicyAssignmentsOperationOptions{
				Filter: pointer.To(fmt.Sprintf("scopeType eq 'Group' and scopeId eq '%s' and policyId eq '%s'", odata.EscapeSingleQuote(policyId.ScopeId), odata.EscapeSingleQuote(id.UnifiedRoleManagementPolicyId))),
			}
			resp, err := assignmentClient.ListRoleManagementPolicyAssignments(ctx, options)
			if err != nil {
				return fmt.Errorf("retrieving %s: %v", id, err)
			}

			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: expected 1 assignment, got nil result", id)
			}
			if len(*resp.Model) != 1 {
				return fmt.Errorf("retrieving %s: expected 1 assignment, got %d", id, len(*resp.Model))
			}

			assignment := (*resp.Model)[0]

			state := GroupRoleManagementPolicyDataSourceModel{
				Description: pointer.From(policy.Description),
				DisplayName: pointer.From(policy.DisplayName),
				GroupId:     policy.ScopeId,
				RoleId:      assignment.RoleDefinitionId.GetOrZero(),
			}

			metadata.ResourceData.SetId(id.ID())
			return metadata.Encode(&state)
		},
	}
}

func (r GroupRoleManagementPolicyDataSource) ResourceType() string {
	return "azuread_group_role_management_policy"
}
