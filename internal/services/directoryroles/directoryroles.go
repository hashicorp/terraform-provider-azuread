package directoryroles

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func expandCustomRolePermissions(in []interface{}) *[]msgraph.UnifiedRolePermission {
	if in == nil {
		return nil
	}

	result := make([]msgraph.UnifiedRolePermission, 0)
	for _, permRaw := range in {
		perm := permRaw.(map[string]interface{})

		var allowedResourceActions *[]string
		if v, ok := perm["allowed_resource_actions"]; ok {
			allowedResourceActions = tf.ExpandStringSlicePtr(v.(*schema.Set).List())
		}

		result = append(result, msgraph.UnifiedRolePermission{
			AllowedResourceActions: allowedResourceActions,
		})
	}

	return &result
}

func flattenCustomRolePermissions(in *[]msgraph.UnifiedRolePermission) []interface{} {
	result := make([]interface{}, 0)

	if in == nil {
		return result
	}

	for _, perm := range *in {
		result = append(result, map[string]interface{}{
			"allowed_resource_actions": tf.FlattenStringSlicePtr(perm.AllowedResourceActions),
		})
	}

	return result
}
