// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/directoryrole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroletemplates/stable/directoryroletemplate"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/suppress"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles/migrations"
)

type DirectoryRoleModel struct {
	Description string `tfschema:"description"`
	DisplayName string `tfschema:"display_name"`
	ObjectId    string `tfschema:"object_id"`
	TemplateId  string `tfschema:"template_id"`
}

var _ sdk.Resource = DirectoryRoleResource{}
var _ sdk.ResourceWithStateMigration = DirectoryRoleResource{}

type DirectoryRoleResource struct{}

func (r DirectoryRoleResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return stable.ValidateDirectoryRoleID
}

func (r DirectoryRoleResource) ResourceType() string {
	return "azuread_directory_role"
}

func (r DirectoryRoleResource) ModelObject() interface{} {
	return &DirectoryRoleModel{}
}

func (r DirectoryRoleResource) StateUpgraders() sdk.StateUpgradeData {
	return sdk.StateUpgradeData{
		SchemaVersion: 1,
		Upgraders: map[int]pluginsdk.StateUpgrade{
			0: migrations.ResourceDirectoryRoleStateUpgradeV0{},
		},
	}
}

func (r DirectoryRoleResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"display_name": {
			Description:      "The display name of the directory role",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			Computed:         true,
			ForceNew:         true,
			ExactlyOneOf:     []string{"display_name", "template_id"},
			DiffSuppressFunc: suppress.CaseDifference,
			ValidateFunc:     validation.StringIsNotEmpty,
		},

		"template_id": {
			Description:  "The object ID of the template associated with the directory role",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			Computed:     true,
			ForceNew:     true,
			ExactlyOneOf: []string{"display_name", "template_id"},
			ValidateFunc: validation.IsUUID,
		},
	}
}

func (r DirectoryRoleResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"description": {
			Description: "The description of the directory role",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"object_id": {
			Description: "The object ID of the directory role",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (r DirectoryRoleResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DirectoryRoles.DirectoryRoleClient
			templateClient := metadata.Client.DirectoryRoles.DirectoryRoleTemplateClient

			var model DirectoryRoleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// First we find the directory role template
			var template *stable.DirectoryRoleTemplate
			if model.DisplayName != "" {
				resp, err := templateClient.ListDirectoryRoleTemplates(ctx, directoryroletemplate.DefaultListDirectoryRoleTemplatesOperationOptions())
				if err != nil {
					return fmt.Errorf("listing directory role templates: %+v", err)
				}

				templates := resp.Model
				if templates == nil {
					return fmt.Errorf("listing directory role templates: API error, result was nil")
				}

				for _, t := range *templates {
					if strings.EqualFold(model.DisplayName, t.DisplayName.GetOrZero()) {
						template = &t
						break
					}
				}
			} else if model.TemplateId != "" {
				resp, err := templateClient.GetDirectoryRoleTemplate(ctx, stable.NewDirectoryRoleTemplateID(model.TemplateId), directoryroletemplate.DefaultGetDirectoryRoleTemplateOperationOptions())
				if err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						return fmt.Errorf("no directory role template with object ID %q was found", model.TemplateId)
					}
					return fmt.Errorf("retrieving directory role template with object ID %q: %+v", model.TemplateId, err)
				}

				template = resp.Model
			}

			if template == nil {
				return fmt.Errorf("no directory role template found")
			}
			if template.Id == nil {
				return fmt.Errorf("received directory role template with nil ID (API error)")
			}

			templateId := *template.Id
			var directoryRole *stable.DirectoryRole

			// Now look for the directory role created from that template
			options := directoryrole.ListDirectoryRolesOperationOptions{
				Filter: pointer.To(fmt.Sprintf("roleTemplateId eq '%s'", odata.EscapeSingleQuote(templateId))),
			}

			if resp, err := client.ListDirectoryRoles(ctx, options); err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					// Directory role was not found, so activate it
					properties := stable.DirectoryRole{
						RoleTemplateId: nullable.Value(templateId),
					}
					if resp, err := client.CreateDirectoryRole(ctx, properties, directoryrole.DefaultCreateDirectoryRoleOperationOptions()); err != nil {
						return fmt.Errorf("activating directory role for template ID %q: %v", templateId, err)
					} else {
						directoryRole = resp.Model
					}
				} else {
					return fmt.Errorf("retrieving directory role with template ID %q: %v", templateId, err)
				}
			} else if resp.Model != nil {
				for _, role := range *resp.Model {
					directoryRole = &role
					break
				}
			}

			if directoryRole == nil {
				return fmt.Errorf("retrieving directory role for template ID %q: result was nil", templateId)
			}
			if directoryRole.Id == nil {
				return fmt.Errorf("retrieving directory role for template ID %q: ID was nil (API error)", templateId)
			}

			id := stable.NewDirectoryRoleID(*directoryRole.Id)
			metadata.SetID(id)

			return nil
		},
	}
}

func (r DirectoryRoleResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DirectoryRoles.DirectoryRoleClient

			id, err := stable.ParseDirectoryRoleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var state DirectoryRoleModel
			if err := metadata.Decode(&state); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			resp, err := client.GetDirectoryRole(ctx, *id, directoryrole.DefaultGetDirectoryRoleOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			directoryRole := resp.Model
			if directoryRole == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			state.Description = directoryRole.Description.GetOrZero()
			state.DisplayName = directoryRole.DisplayName.GetOrZero()
			state.ObjectId = pointer.From(directoryRole.Id)
			state.TemplateId = directoryRole.RoleTemplateId.GetOrZero()

			return metadata.Encode(&state)
		},
	}
}

func (r DirectoryRoleResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			// Directory roles cannot be deactivated or deleted, so this is a no-op
			return nil
		},
	}
}
