// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

var _ sdk.DataSource = GroupRoleManagementPolicyDataSource{}

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
	return &GroupRoleManagementPolicyModel{}
}

func (r GroupRoleManagementPolicyDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Groups.GroupsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			groupID := metadata.ResourceData.Get("group_id").(string)
			roleID := metadata.ResourceData.Get("role_id").(string)
			id, err := getPolicyId(ctx, metadata, groupID, roleID)
			if err != nil {
				return errors.New("Bad API response")
			}
			metadata.ResourceData.SetId(id.ID())
			return nil
		},
	}
}

func (r GroupRoleManagementPolicyDataSource) ResourceType() string {
	return "azuread_group_role_management_policy"
}
