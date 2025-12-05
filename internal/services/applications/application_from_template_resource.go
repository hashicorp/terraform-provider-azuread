// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/stable/applicationtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationFromTemplateModel struct {
	TemplateId  string `tfschema:"template_id"`
	DisplayName string `tfschema:"display_name"`

	ApplicationId            string `tfschema:"application_id"`
	ApplicationObjectId      string `tfschema:"application_object_id"`
	ServicePrincipalId       string `tfschema:"service_principal_id"`
	ServicePrincipalObjectId string `tfschema:"service_principal_object_id"`
}

var _ sdk.ResourceWithUpdate = ApplicationFromTemplateResource{}

type ApplicationFromTemplateResource struct{}

func (r ApplicationFromTemplateResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateFromTemplateID
}

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
			client := metadata.Client.Applications.ApplicationTemplateClient

			var model ApplicationFromTemplateModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			templateId := stable.NewApplicationTemplateID(model.TemplateId)

			request := applicationtemplate.InstantiateRequest{
				DisplayName: nullable.Value(model.DisplayName),
			}

			resp, err := client.Instantiate(ctx, templateId, request, applicationtemplate.DefaultInstantiateOperationOptions())
			if err != nil {
				return fmt.Errorf("creating %s: %+v", templateId, err)
			}
			if resp.Model == nil {
				return fmt.Errorf("creating %s: model was nil", templateId)
			}
			if resp.Model.Application == nil {
				return fmt.Errorf("creating %s: application was nil", templateId)
			}
			if resp.Model.ServicePrincipal == nil {
				return fmt.Errorf("creating %s: servicePrincipal was nil", templateId)
			}

			id := parse.NewFromTemplateID(model.TemplateId, *resp.Model.Application.Id, *resp.Model.ServicePrincipal.Id)
			metadata.SetID(id)

			if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
				client := metadata.Client.Applications.ApplicationClient

				resp, err := client.GetApplication(ctx, stable.NewApplicationID(id.ApplicationId), application.DefaultGetApplicationOperationOptions())
				if err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						return pointer.To(false), nil
					}
					return nil, err
				}
				return pointer.To(resp.Model != nil), nil
			}); err != nil {
				return fmt.Errorf("creating %s: timed out waiting for replication of new application", templateId)
			}

			return nil
		},
	}
}

func (r ApplicationFromTemplateResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseFromTemplateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)
			servicePrincipalId := stable.NewServicePrincipalID(id.ServicePrincipalId)

			// Check the application exists
			resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: model was nil", applicationId)
			}

			state := ApplicationFromTemplateModel{
				DisplayName:              resp.Model.DisplayName.GetOrZero(),
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
			client := metadata.Client.Applications.ApplicationClient
			rd := metadata.ResourceData

			id, err := parse.ParseFromTemplateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationFromTemplateModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			if rd.HasChange("display_name") {
				properties := stable.Application{
					DisplayName: nullable.Value(model.DisplayName),
				}

				if _, err = client.UpdateApplication(ctx, applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
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
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseFromTemplateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			applicationId := stable.NewApplicationID(id.ApplicationId)

			if _, err = client.DeleteApplication(ctx, applicationId, application.DefaultDeleteApplicationOperationOptions()); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
