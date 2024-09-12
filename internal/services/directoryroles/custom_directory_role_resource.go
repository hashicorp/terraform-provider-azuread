// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroledefinition"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
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
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleDefinitionClient

	displayName := d.Get("display_name").(string)

	properties := stable.UnifiedRoleDefinition{
		Description:     nullable.NoZero(d.Get("description").(string)),
		DisplayName:     nullable.Value(displayName),
		IsEnabled:       nullable.Value(d.Get("enabled").(bool)),
		RolePermissions: expandCustomRolePermissions(d.Get("permissions").(*pluginsdk.Set).List()),
		TemplateId:      nullable.Value(d.Get("template_id").(string)),
		Version:         nullable.Value(d.Get("version").(string)),
	}

	resp, err := client.CreateDirectoryRoleDefinition(ctx, properties, directoryroledefinition.DefaultCreateDirectoryRoleDefinitionOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Creating custom directory role %q", displayName)
	}

	role := resp.Model
	if role == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Creating custom directory role %q", displayName)
	}

	if role.Id == nil || *role.Id == "" {
		return tf.ErrorDiagF(errors.New("API returned custom directory role with nil ID"), "Bad API Response")
	}

	d.SetId(*role.Id)

	return customDirectoryRoleResourceRead(ctx, d, meta)
}

func customDirectoryRoleResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleDefinitionClient
	id := stable.NewRoleManagementDirectoryRoleDefinitionID(d.Id())

	displayName := d.Get("display_name").(string)

	properties := stable.UnifiedRoleDefinition{
		Description:     nullable.NoZero(d.Get("description").(string)),
		DisplayName:     nullable.Value(displayName),
		IsEnabled:       nullable.Value(d.Get("enabled").(bool)),
		RolePermissions: expandCustomRolePermissions(d.Get("permissions").(*pluginsdk.Set).List()),
		TemplateId:      nullable.Value(d.Get("template_id").(string)),
		Version:         nullable.Value(d.Get("version").(string)),
	}

	_, err := client.UpdateDirectoryRoleDefinition(ctx, id, properties, directoryroledefinition.DefaultUpdateDirectoryRoleDefinitionOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Updating custom directory role %q", displayName)
	}

	return customDirectoryRoleResourceRead(ctx, d, meta)
}

func customDirectoryRoleResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleDefinitionClient
	id := stable.NewRoleManagementDirectoryRoleDefinitionID(d.Id())

	resp, err := client.GetDirectoryRoleDefinition(ctx, id, directoryroledefinition.DefaultGetDirectoryRoleDefinitionOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "template_id", "Retrieving %s: %+v", id, err)
	}

	role := resp.Model
	if role == nil {
		return tf.ErrorDiagF(errors.New("API error: nil unifiedDirectoryRole was returned"), "Retrieving %s", id)
	}

	tf.Set(d, "description", role.Description.GetOrZero())
	tf.Set(d, "display_name", role.DisplayName.GetOrZero())
	tf.Set(d, "enabled", role.IsEnabled.GetOrZero())
	tf.Set(d, "object_id", id.UnifiedRoleDefinitionId)
	tf.Set(d, "permissions", flattenCustomRolePermissions(role.RolePermissions))
	tf.Set(d, "template_id", role.TemplateId.GetOrZero())
	tf.Set(d, "version", role.Version.GetOrZero())

	return nil
}

func customDirectoryRoleResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).DirectoryRoles.DirectoryRoleDefinitionClient
	id := stable.NewRoleManagementDirectoryRoleDefinitionID(d.Id())

	resp, err := client.GetDirectoryRoleDefinition(ctx, id, directoryroledefinition.DefaultGetDirectoryRoleDefinitionOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(fmt.Errorf("Custom Directory Role was not found"), "id", "Retrieving %s", id)
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving %s", id)
	}

	if _, err = client.DeleteDirectoryRoleDefinition(ctx, id, directoryroledefinition.DefaultDeleteDirectoryRoleDefinitionOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	return nil
}
