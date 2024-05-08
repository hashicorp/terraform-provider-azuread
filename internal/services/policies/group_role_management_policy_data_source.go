// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

var _ sdk.DataSource = GroupRoleManagementPolicyDataSource{}

type GroupRoleManagementPolicyDataSourceModel struct {
	Description string                                   `tfschema:"description"`
	DisplayName string                                   `tfschema:"display_name"`
	GroupId     string                                   `tfschema:"group_id"`
	RoleId      msgraph.UnifiedRoleManagementPolicyScope `tfschema:"role_id"`
}

type GroupRoleManagementPolicyDataSource struct{}

func (r GroupRoleManagementPolicyDataSource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_id": {
			Description:      "ID of the group to which this policy is assigned",
			Type:             pluginsdk.TypeString,
			Required:         true,
			ForceNew:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
		},

		"role_id": {
			Description: "The ID of the role of this policy to the group",
			Type:        pluginsdk.TypeString,
			Required:    true,
			ForceNew:    true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
				msgraph.PrivilegedAccessGroupRelationshipMember,
				msgraph.PrivilegedAccessGroupRelationshipOwner,
				msgraph.PrivilegedAccessGroupRelationshipUnknown,
			}, false)),
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
			clientPolicy := metadata.Client.Policies.RoleManagementPolicyClient
			clientAssignment := metadata.Client.Policies.RoleManagementPolicyAssignmentClient

			clientPolicy.BaseClient.DisableRetries = true
			clientAssignment.BaseClient.DisableRetries = true

			defer func() {
				clientPolicy.BaseClient.DisableRetries = false
				clientAssignment.BaseClient.DisableRetries = false
			}()

			groupID := metadata.ResourceData.Get("group_id").(string)
			roleID := metadata.ResourceData.Get("role_id").(string)
			id, err := getPolicyId(ctx, metadata, groupID, roleID)
			if err != nil {
				return errors.New("Bad API response")
			}

			result, _, err := clientPolicy.Get(ctx, id.ID())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			assignments, _, err := clientAssignment.List(ctx, odata.Query{
				Filter: fmt.Sprintf("scopeType eq 'Group' and scopeId eq '%s' and policyId eq '%s'", id.ScopeId, id.ID()),
			})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if len(*assignments) != 1 {
				return fmt.Errorf("retrieving %s: expected 1 assignment, got %d", id, len(*assignments))
			}

			state := GroupRoleManagementPolicyDataSourceModel{
				Description: *result.Description,
				DisplayName: *result.DisplayName,
				GroupId:     *result.ScopeId,
				RoleId:      *(*assignments)[0].RoleDefinitionId,
			}

			metadata.ResourceData.SetId(id.ID())
			return metadata.Encode(&state)
		},
	}
}

func (r GroupRoleManagementPolicyDataSource) ResourceType() string {
	return "azuread_group_role_management_policy"
}
