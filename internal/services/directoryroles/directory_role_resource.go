// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/suppress"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type DirectoryRoleModel struct {
	Description string `tfschema:"description"`
	DisplayName string `tfschema:"display_name"`
	ObjectId    string `tfschema:"object_id"`
	TemplateId  string `tfschema:"template_id"`
}

var _ sdk.Resource = DirectoryRoleResource{}

type DirectoryRoleResource struct{}

func (r DirectoryRoleResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validation.IsUUID
}

func (r DirectoryRoleResource) ResourceType() string {
	return "azuread_directory_role"
}

func (r DirectoryRoleResource) ModelObject() interface{} {
	return &DirectoryRoleModel{}
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
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},

		"template_id": {
			Description:      "The object ID of the template associated with the directory role",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			Computed:         true,
			ForceNew:         true,
			ExactlyOneOf:     []string{"display_name", "template_id"},
			ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
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
			client := metadata.Client.DirectoryRoles.DirectoryRolesClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			directoryRoleTemplatesClient := metadata.Client.DirectoryRoles.DirectoryRoleTemplatesClient

			var model DirectoryRoleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// First we find the directory role template
			var template *msgraph.DirectoryRoleTemplate
			if model.DisplayName != "" {
				templates, _, err := directoryRoleTemplatesClient.List(ctx)
				if err != nil {
					return fmt.Errorf("listing directory role templates: %+v", err)
				}
				if templates == nil {
					return fmt.Errorf("listing directory role templates: API error, result was nil")
				}

				for _, t := range *templates {
					if strings.EqualFold(model.DisplayName, pointer.From(t.DisplayName)) {
						template = &t
						break
					}
				}

				if template == nil {
					return fmt.Errorf("no directory role template found with display name: %q", model.DisplayName)
				}
			} else if model.TemplateId != "" {
				var status int
				var err error
				template, status, err = directoryRoleTemplatesClient.Get(ctx, model.TemplateId)
				if err != nil {
					if status == http.StatusNotFound {
						return fmt.Errorf("no directory role template with object ID %q was found", model.TemplateId)
					}
					return fmt.Errorf("retrieving directory role template with object ID %q: %+v", model.TemplateId, err)
				}

				if template == nil {
					return fmt.Errorf("retrieving directory role template with object ID %q: API error, result was nil", model.TemplateId)
				}
			}

			if template == nil {
				return fmt.Errorf("no directory role template found")
			}

			if template.ID == nil {
				return fmt.Errorf("received directory role template with nil ID (API error)")
			}

			templateId := *template.ID

			// Now look for the directory role created from that template
			directoryRole, status, err := client.GetByTemplateId(ctx, templateId)
			if err != nil {
				if status == http.StatusNotFound {
					// Directory role was not found, so activate it
					directoryRole, _, err = client.Activate(ctx, templateId)
					if err != nil {
						return fmt.Errorf("activating directory role for template ID %q: %+v", templateId, err)
					}
				} else {
					return fmt.Errorf("retrieving directory role with template ID %q: %+v", templateId, err)
				}
			}

			if directoryRole == nil {
				return fmt.Errorf("retrieving directory role for template ID %q: result was nil", templateId)
			}
			if directoryRole.ID() == nil {
				return fmt.Errorf("retrieving directory role for template ID %q: ID was nil (API error)", templateId)
			}

			id := parse.NewDirectoryRoleID(*directoryRole.ID())
			metadata.SetID(id)

			return nil
		},
	}
}

func (r DirectoryRoleResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DirectoryRoles.DirectoryRolesClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id := parse.NewDirectoryRoleID(metadata.ResourceData.Id())

			var state DirectoryRoleModel
			if err := metadata.Decode(&state); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			directoryRole, status, err := client.Get(ctx, id.ID())
			if err != nil {
				if status == http.StatusNotFound {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if directoryRole == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			state.Description = pointer.From(directoryRole.Description)
			state.DisplayName = pointer.From(directoryRole.DisplayName)
			state.ObjectId = pointer.From(directoryRole.ID())
			state.TemplateId = pointer.From(directoryRole.RoleTemplateId)

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
