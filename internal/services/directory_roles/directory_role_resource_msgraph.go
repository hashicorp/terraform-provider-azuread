package directory_roles

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers "github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func directoryRoleResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.MsClient
	dirRoleTemplatesClient := meta.(*clients.Client).DirectoryRoles.DirRoleTemplatesMsClient

	var displayName string
	if v, ok := d.GetOk("display_name"); ok && v.(string) != "" {
		displayName = v.(string)
	} else {
		displayName = d.Get("name").(string)
	}

	existingDirRole, err := helpers.DirectoryRoleFindByName(ctx, client, displayName)
	if err != nil {
		return tf.ErrorDiagPathF(err, "name", "Could not check for existing directory role template(s)")
	}
	if existingDirRole == nil {
		existingDirRoleTemplate, err := helpers.DirectoryRoleTemplateFindByName(ctx, dirRoleTemplatesClient, displayName)
		if err != nil {
			return tf.ErrorDiagPathF(err, "name", "Could not check for existing directory role template(s)")
		}
		if existingDirRoleTemplate != nil && existingDirRoleTemplate.ID == nil {
			return tf.ErrorDiagF(errors.New("API returned directory role template with nil object ID"), "Bad API response")
		}
		dirRole, _, err := client.Activate(ctx, *existingDirRoleTemplate.ID)
		if err != nil {
			return tf.ErrorDiagF(err, "Activation directory role %q", displayName)
		}

		if dirRole.ID == nil {
			return tf.ErrorDiagF(errors.New("API returned directory role with nil object ID"), "Bad API Response")
		}
		existingDirRole = dirRole
	}

	d.SetId(*existingDirRole.ID)

	_, err = helpers.WaitForCreationReplication(ctx, func() (interface{}, int, error) {
		return client.Get(ctx, *existingDirRole.ID)
	})

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for Directory Role with object ID: %q", *existingDirRole.ID)
	}

	return directoryRoleResourceReadMsGraph(ctx, d, meta)
}

func directoryRoleResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.MsClient

	dirRole, status, err := client.Get(ctx, d.Id())
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Directory Role with ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving Directory Role with object ID: %q", d.Id())
	}

	tf.Set(d, "description", dirRole.Description)
	tf.Set(d, "display_name", dirRole.DisplayName)
	tf.Set(d, "object_id", dirRole.ID)
	tf.Set(d, "role_template_id", dirRole.RoleTemplateId)

	members, _, err := client.ListMembers(ctx, *dirRole.ID)
	if err != nil {
		return tf.ErrorDiagPathF(err, "owners", "Could not retrieve members for directory roles with object ID %q", d.Id())
	}
	tf.Set(d, "members", members)

	return nil
}
