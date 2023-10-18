// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

type ApplicationFromTemplateModel struct {
	TemplateId  string `tfschema:"template_id"`
	DisplayName string `tfschema:"display_name"`

	ApplicationId            string `tfschema:"application_id"`
	ApplicationObjectId      string `tfschema:"application_object_id"`
	ServicePrincipalId       string `tfschema:"service_principal_id"`
	ServicePrincipalObjectId string `tfschema:"service_principal_object_id"`
}

type ApplicationFromTemplateResource struct{}

func (r ApplicationFromTemplateResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateFromTemplateID
}

var _ sdk.ResourceWithUpdate = ApplicationFromTemplateResource{}

func (r ApplicationFromTemplateResource) ResourceType() string {
	return "azuread_application_from_template"
}

func (r ApplicationFromTemplateResource) ModelObject() interface{} {
	return &ApplicationFromTemplateModel{}
}

func (r ApplicationFromTemplateResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"display_name": {
			Description:  "The display name for the application",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"template_id": {
			Description:  "The UUID of the template to instantiate for this application",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},
	}
}

func (r ApplicationFromTemplateResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description: "The resource ID for this application",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"application_object_id": {
			Description: "The object ID for this application",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"service_principal_id": {
			Description: "The resource ID for this service principal",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"service_principal_object_id": {
			Description: "The object ID for this service principal",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (r ApplicationFromTemplateResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationTemplatesClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationFromTemplateModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			properties := msgraph.ApplicationTemplate{
				DisplayName: pointer.To(model.DisplayName),
				ID:          pointer.To(model.TemplateId),
			}

			result, _, err := client.Instantiate(ctx, properties)
			if err != nil {
				return fmt.Errorf("creating %s: %+v", parse.FromTemplateId{}, err)
			}
			if result == nil {
				return fmt.Errorf("creating %s: result was nil", parse.FromTemplateId{})
			}
			if result.Application == nil {
				return fmt.Errorf("creating %s: application was nil", parse.FromTemplateId{})
			}
			if result.ServicePrincipal == nil {
				return fmt.Errorf("creating %s: servicePrincipal was nil", parse.FromTemplateId{})
			}

			id := parse.NewFromTemplateID(model.TemplateId, *result.Application.ID(), *result.ServicePrincipal.ID())
			metadata.SetID(id)

			if err = helpers.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				client := metadata.Client.Applications.ApplicationsClient
				client.BaseClient.DisableRetries = true
				defer func() { client.BaseClient.DisableRetries = false }()

				result, status, err := client.Get(ctx, *result.Application.ID(), odata.Query{})
				if err != nil {
					if status == http.StatusNotFound {
						return utils.Bool(false), nil
					}
					return nil, err
				}
				return pointer.To(result != nil), nil
			}); err != nil {
				return fmt.Errorf("creating %s: timed out waiting for replication of new application", parse.FromTemplateId{})
			}

			return nil
		},
	}
}

func (r ApplicationFromTemplateResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseFromTemplateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			// Check the application exists
			result, status, err := client.Get(ctx, id.ApplicationId, odata.Query{})
			if err != nil {
				if status == http.StatusNotFound {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", id)
			}

			applicationId := parse.NewApplicationID(id.ApplicationId)
			servicePrincipalId := parse.NewServicePrincipalID(id.ServicePrincipalId)

			state := ApplicationFromTemplateModel{
				DisplayName:              pointer.From(result.DisplayName),
				TemplateId:               id.TemplateId,
				ApplicationId:            applicationId.ID(),
				ApplicationObjectId:      applicationId.ApplicationId,
				ServicePrincipalId:       servicePrincipalId.ID(),
				ServicePrincipalObjectId: servicePrincipalId.ServicePrincipalId,
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationFromTemplateResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			rd := metadata.ResourceData

			id, err := parse.ParseFromTemplateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationFromTemplateModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			if rd.HasChange("display_name") {
				properties := msgraph.Application{
					DirectoryObject: msgraph.DirectoryObject{
						Id: &id.ApplicationId,
					},
					DisplayName: pointer.To(model.DisplayName),
				}

				if _, err = client.Update(ctx, properties); err != nil {
					return fmt.Errorf("updating %s: %+v", id, err)
				}
			}

			return nil
		},
	}
}

func (r ApplicationFromTemplateResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseFromTemplateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			_, err = client.Delete(ctx, id.ApplicationId)
			if err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
