package directory_roles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers "github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func directoryRoleResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: directoryRoleResourceCreate,
		ReadContext:   directoryRoleResourceRead,
		DeleteContext: directoryRoleResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role_template_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"members": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.UUID,
				},
			},
		},
	}
}

func directoryRoleResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient
	dirRoleTemplatesClient := meta.(*clients.Client).DirectoryRoles.DirRoleTemplatesMsClient

	var displayName string
	if v, ok := d.GetOk("display_name"); ok && v.(string) != "" {
		displayName = v.(string)
	}

	dirRoles, err := directoryRoleFindByName(ctx, client, displayName)
	if err != nil {
		return tf.ErrorDiagPathF(err, "name", "Could not check for existing directory role template(s)")
	}
	if dirRoles == nil {
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
		dirRoles = append(dirRoles, *dirRole)
	}

	d.SetId(*dirRoles[0].ID)

	return directoryRoleResourceRead(ctx, d, meta)
}

func directoryRoleResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRolesClient

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

func directoryRoleResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
