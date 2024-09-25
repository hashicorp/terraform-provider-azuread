// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/member"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func directoryRoleGetMember(ctx context.Context, client *member.MemberClient, id stable.DirectoryRoleIdMemberId) (*stable.DirectoryObject, error) {
	options := member.ListMembersOperationOptions{
		Filter: pointer.To(fmt.Sprintf("id eq '%s'", id.DirectoryObjectId)),
	}

	resp, err := client.ListMembers(ctx, stable.NewDirectoryRoleID(id.DirectoryRoleId), options)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if resp.Model != nil {
		for _, member := range *resp.Model {
			if member.DirectoryObject().Id != nil && *member.DirectoryObject().Id == id.DirectoryObjectId {
				return &member, nil
			}
		}
	}

	return nil, nil
}

func expandCustomRolePermissions(in []interface{}) []stable.UnifiedRolePermission {
	result := make([]stable.UnifiedRolePermission, 0)
	for _, permRaw := range in {
		perm := permRaw.(map[string]interface{})

		var allowedResourceActions []string
		if v, ok := perm["allowed_resource_actions"]; ok {
			allowedResourceActions = tf.ExpandStringSlice(v.(*pluginsdk.Set).List())
		}

		result = append(result, stable.UnifiedRolePermission{
			AllowedResourceActions: allowedResourceActions,
		})
	}

	return result
}

func flattenCustomRolePermissions(in []stable.UnifiedRolePermission) []interface{} {
	result := make([]interface{}, 0)
	for _, perm := range in {
		result = append(result, map[string]interface{}{
			"allowed_resource_actions": tf.FlattenStringSlice(perm.AllowedResourceActions),
		})
	}

	return result
}
