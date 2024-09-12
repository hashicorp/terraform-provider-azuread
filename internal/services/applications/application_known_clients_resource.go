// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationKnownClientsModel struct {
	ApplicationId  string   `tfschema:"application_id"`
	KnownClientIds []string `tfschema:"known_client_ids"`
}

var _ sdk.ResourceWithUpdate = ApplicationKnownClientsResource{}

type ApplicationKnownClientsResource struct{}

func (r ApplicationKnownClientsResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateKnownClientsID
}

func (r ApplicationKnownClientsResource) ResourceType() string {
	return "azuread_application_known_clients"
}

func (r ApplicationKnownClientsResource) ModelObject() interface{} {
	return &ApplicationKnownClientsModel{}
}

func (r ApplicationKnownClientsResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which this API access is granted",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: parse.ValidateApplicationID,
		},

		"known_client_ids": {
			Description: "A list of known client IDs, used for bundling consent if you have a solution that includes an API and a client application",
			Type:        pluginsdk.TypeSet,
			Required:    true,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func (r ApplicationKnownClientsResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationKnownClientsResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			var model ApplicationKnownClientsModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := stable.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewKnownClientsID(applicationId.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			resp, err := client.GetApplication(ctx, *applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil {
				return fmt.Errorf("retrieving %s: model was nil", applicationId)
			}

			// Check for existing known clients
			if app.Api != nil && app.Api.KnownClientApplications != nil && len(*app.Api.KnownClientApplications) > 0 {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			properties := stable.Application{
				Api: &stable.ApiApplication{
					KnownClientApplications: pointer.To(model.KnownClientIds),
				},
			}

			if _, err = client.UpdateApplication(ctx, *applicationId, properties); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationKnownClientsResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseKnownClientsID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			app := resp.Model
			if app == nil {
				return fmt.Errorf("retrieving %s: model was nil", id)
			}
			if app.Api == nil || app.Api.KnownClientApplications == nil {
				return metadata.MarkAsGone(id)
			}

			state := ApplicationKnownClientsModel{
				ApplicationId:  applicationId.ID(),
				KnownClientIds: pointer.From(app.Api.KnownClientApplications),
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationKnownClientsResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseKnownClientsID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			var model ApplicationKnownClientsModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			properties := stable.Application{
				Api: &stable.ApiApplication{
					KnownClientApplications: pointer.To(model.KnownClientIds),
				},
			}

			if _, err = client.UpdateApplication(ctx, applicationId, properties); err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationKnownClientsResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseKnownClientsID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			properties := stable.Application{
				Api: &stable.ApiApplication{
					KnownClientApplications: &[]string{},
				},
			}

			if _, err = client.UpdateApplication(ctx, applicationId, properties); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
