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

type ApplicationRedirectUrisModel struct {
	ApplicationId string   `tfschema:"application_id"`
	UriType       string   `tfschema:"type"`
	RedirectUris  []string `tfschema:"redirect_uris"`
}

var _ sdk.ResourceWithUpdate = ApplicationRedirectUrisResource{}

type ApplicationRedirectUrisResource struct{}

func (r ApplicationRedirectUrisResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateRedirectUrisID
}

func (r ApplicationRedirectUrisResource) ResourceType() string {
	return "azuread_application_redirect_uris"
}

func (r ApplicationRedirectUrisResource) ModelObject() interface{} {
	return &ApplicationRedirectUrisModel{}
}

func (r ApplicationRedirectUrisResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which these redirect URIs belong",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: parse.ValidateApplicationID,
		},

		"type": {
			Description:  "The type of redirect URIs to assign to the application",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice(possibleValuesForRedirectUriType, false),
		},

		"redirect_uris": {
			Description: "A set of redirect URIs",
			Type:        pluginsdk.TypeSet,
			Required:    true,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.IsRedirectUriFunc(true, true),
			},
		},
	}
}

func (r ApplicationRedirectUrisResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationRedirectUrisResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			var model ApplicationRedirectUrisModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := stable.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewRedirectUrisID(applicationId.ApplicationId, model.UriType)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			resp, err := client.GetApplication(ctx, *applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil {
				return fmt.Errorf("retrieving %s: app was nil", applicationId)
			}

			// Check for existing redirect URIs
			if existingUris := r.getRedirectUrisByType(*app, model.UriType); len(existingUris) > 0 {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			properties := stable.Application{}
			r.setRedirectUrisByType(&properties, model)

			if _, err = client.UpdateApplication(ctx, *applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationRedirectUrisResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseRedirectUrisID(metadata.ResourceData.Id())
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
				return fmt.Errorf("retrieving %s: app was nil", id)
			}

			redirectUris := r.getRedirectUrisByType(*app, id.UriType)

			if len(redirectUris) == 0 {
				return metadata.MarkAsGone(id)
			}

			state := ApplicationRedirectUrisModel{
				ApplicationId: applicationId.ID(),
				UriType:       id.UriType,
				RedirectUris:  redirectUris,
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationRedirectUrisResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseRedirectUrisID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			var model ApplicationRedirectUrisModel
			if err = metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			properties := stable.Application{}
			r.setRedirectUrisByType(&properties, model)

			if _, err = client.UpdateApplication(ctx, applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationRedirectUrisResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseRedirectUrisID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			properties := stable.Application{}
			r.deleteRedirectUrisByType(&properties, id.UriType)

			if _, err = client.UpdateApplication(ctx, applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationRedirectUrisResource) getRedirectUrisByType(application stable.Application, uriType string) []string {
	switch uriType {
	case RedirectUriTypePublicClient:
		if application.PublicClient != nil {
			return pointer.From(application.PublicClient.RedirectUris)
		}
	case RedirectUriTypeSPA:
		if application.Spa != nil {
			return pointer.From(application.Spa.RedirectUris)
		}
	case RedirectUriTypeWeb:
		if application.Web != nil {
			return pointer.From(application.Web.RedirectUris)
		}
	}

	return nil
}

func (r ApplicationRedirectUrisResource) setRedirectUrisByType(application *stable.Application, model ApplicationRedirectUrisModel) {
	switch model.UriType {
	case RedirectUriTypePublicClient:
		application.PublicClient = &stable.PublicClientApplication{
			RedirectUris: pointer.To(model.RedirectUris),
		}
	case RedirectUriTypeSPA:
		application.Spa = &stable.SpaApplication{
			RedirectUris: pointer.To(model.RedirectUris),
		}
	case RedirectUriTypeWeb:
		application.Web = &stable.WebApplication{
			RedirectUris: pointer.To(model.RedirectUris),
		}
	}
}
func (r ApplicationRedirectUrisResource) deleteRedirectUrisByType(application *stable.Application, uriType string) {
	switch uriType {
	case RedirectUriTypePublicClient:
		application.PublicClient = &stable.PublicClientApplication{
			RedirectUris: &[]string{},
		}
	case RedirectUriTypeSPA:
		application.Spa = &stable.SpaApplication{
			RedirectUris: &[]string{},
		}
	case RedirectUriTypeWeb:
		application.Web = &stable.WebApplication{
			RedirectUris: &[]string{},
		}
	}
}
