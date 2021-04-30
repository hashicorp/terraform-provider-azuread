package directory_roles

import (
	"context"
	"errors"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers `github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph`
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func directoryRoleDataSourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.MsClient

	var dirRole msgraph.DirectoryRole
	var displayName string

	if v, ok := d.GetOk("display_name"); ok {
		displayName = v.(string)
	}

	if displayName != "" {
		dirRoles, err := helpers.DirectoryRoleFindByName(ctx, client, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "display_name", "Retrieving directory roles (%s)", displayName)
		}

		count := len(dirRoles)
		if count > 1 {
			return tf.ErrorDiagPathF(err, "display_name", "More than one directory found matching display_name (%s)", displayName)
		} else if count == 0 {
			return tf.ErrorDiagPathF(err, "display_name", "No directory role found matching specified display_name (%s)", displayName)
		}

		dirRole = *dirRoles[0]
	} else if objectId, ok := d.Get("object_id").(string); ok && objectId != "" {
		getDirRole, status, err := client.Get(ctx, objectId)
		if err != nil {
			if status == http.StatusNotFound {
				return tf.ErrorDiagPathF(nil, "object_id", "No directory role found with object ID: %q", objectId)
			}
			return tf.ErrorDiagF(err, "Retrieving directory role with object ID: %q", objectId)
		}
		if getDirRole == nil {
			return tf.ErrorDiagPathF(nil, "object_id", "Directory role not found with object ID: %q", objectId)
		}

		dirRole = *getDirRole
	}

	if dirRole.ID == nil {
		return tf.ErrorDiagF(errors.New("API returned directory role with nil object ID"), "Bad API Response")
	}

	d.SetId(*dirRole.ID)

	tf.Set(d, "description", dirRole.Description)
	tf.Set(d, "display_name", dirRole.DisplayName)
	tf.Set(d, "object_id", dirRole.ID)
	tf.Set(d, "role_template_id", dirRole.RoleTemplateId)

	members, _, err := client.ListMembers(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve directory role members for directory role with object ID: %q", d.Id())
	}
	tf.Set(d, "members", members)

	return nil
}