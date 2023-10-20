// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func customDirectoryRoleResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: customDirectoryRoleResourceCreate,
		UpdateContext: customDirectoryRoleResourceUpdate,
		ReadContext:   customDirectoryRoleResourceRead,
		DeleteContext: customDirectoryRoleResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description:      "The display name of the custom directory role",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"enabled": {
				Description: "Indicates whether the role is enabled for assignment",
				Type:        pluginsdk.TypeBool,
				Required:    true,
			},

			"permissions": {
				Description: "List of permissions that are included in the custom directory role",
				Type:        pluginsdk.TypeSet,
				Required:    true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"allowed_resource_actions": {
							Description: "Set of tasks that can be performed on a resource",
							Type:        pluginsdk.TypeSet,
							Required:    true,
							Elem: &pluginsdk.Schema{
								Type:             pluginsdk.TypeString,
								ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
							},
						},
					},
				},
			},

			"version": {
				Description:      "The version of the role definition.",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringLenBetween(1, 128)),
			},

			"description": {
				Description: "The description of the custom directory role",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"template_id": {
				Description:      "Custom template identifier that is typically used if one needs an identifier to be the same across different directories.",
				Type:             pluginsdk.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),

				// The template ID _can_ technically be changed but doing so mutates the role ID - essentially
				// causing the equivalent of a ForceNew by the API :/
				ForceNew: true,
			},

			"object_id": {
				Description: "The object ID of the directory role",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func customDirectoryRoleResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleDefinitionsClient

	displayName := d.Get("display_name").(string)

	properties := msgraph.UnifiedRoleDefinition{
		Description:     tf.NullableString(d.Get("description").(string)),
		DisplayName:     pointer.To(displayName),
		IsEnabled:       pointer.To(d.Get("enabled").(bool)),
		RolePermissions: expandCustomRolePermissions(d.Get("permissions").(*pluginsdk.Set).List()),
		TemplateId:      pointer.To(d.Get("template_id").(string)),
		Version:         pointer.To(d.Get("version").(string)),
	}

	role, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating custom directory role %q", displayName)
	}

	if role.ID() == nil || *role.ID() == "" {
		return tf.ErrorDiagF(errors.New("API returned custom directory role with nil ID"), "Bad API Response")
	}

	d.SetId(*role.ID())

	return customDirectoryRoleResourceRead(ctx, d, meta)
}

func customDirectoryRoleResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleDefinitionsClient
	roleId := d.Id()

	displayName := d.Get("display_name").(string)

	properties := msgraph.UnifiedRoleDefinition{
		DirectoryObject: msgraph.DirectoryObject{
			Id: &roleId,
		},
		Description:     tf.NullableString(d.Get("description").(string)),
		DisplayName:     pointer.To(displayName),
		IsEnabled:       pointer.To(d.Get("enabled").(bool)),
		RolePermissions: expandCustomRolePermissions(d.Get("permissions").(*pluginsdk.Set).List()),
		TemplateId:      pointer.To(d.Get("template_id").(string)),
		Version:         pointer.To(d.Get("version").(string)),
	}

	_, err := client.Update(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Updating custom directory role %q", displayName)
	}

	return customDirectoryRoleResourceRead(ctx, d, meta)
}

func customDirectoryRoleResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleDefinitionsClient
	roleId := d.Id()

	role, status, err := client.Get(ctx, roleId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Custom Directory Role with ID %q was not found - removing from state", roleId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "template_id", "Retrieving custom directory role with ID %q: %+v", roleId, err)
	}
	if role == nil {
		return tf.ErrorDiagF(errors.New("API error: nil unifiedDirectoryRole was returned"), "Retrieving custom directory role with ID %q", roleId)
	}

	tf.Set(d, "description", role.Description)
	tf.Set(d, "display_name", role.DisplayName)
	tf.Set(d, "enabled", role.IsEnabled)
	tf.Set(d, "object_id", role.ID())
	tf.Set(d, "permissions", flattenCustomRolePermissions(role.RolePermissions))
	tf.Set(d, "template_id", role.TemplateId)
	tf.Set(d, "version", role.Version)

	return nil
}

func customDirectoryRoleResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.RoleDefinitionsClient
	roleId := d.Id()

	_, status, err := client.Get(ctx, roleId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Custom Directory Role was not found"), "id", "Retrieving custom directory role with ID %q", roleId)
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving custom directory role with ID %q", roleId)
	}

	if status, err := client.Delete(ctx, roleId); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting custom directory role with ID %q, got status %d", roleId, status)
	}

	return nil
}
