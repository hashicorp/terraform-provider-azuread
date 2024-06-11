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
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

const (
	RedirectUriTypePublicClient = "PublicClient"
	RedirectUriTypeSPA          = "SPA"
	RedirectUriTypeWeb          = "Web"
)

type ApplicationRedirectUrisModel struct {
	ApplicationId string   `tfschema:"application_id"`
	UriType       string   `tfschema:"type"`
	RedirectUris  []string `tfschema:"redirect_uris"`
}

type ApplicationRedirectUrisResource struct{}

func (r ApplicationRedirectUrisResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateRedirectUrisID
}

var _ sdk.ResourceWithUpdate = ApplicationRedirectUrisResource{}

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
			Description: "The type of redirect URIs to assign to the application",
			Type:        pluginsdk.TypeString,
			Required:    true,
			ForceNew:    true,
			ValidateFunc: validation.StringInSlice([]string{
				RedirectUriTypePublicClient,
				RedirectUriTypeSPA,
				RedirectUriTypeWeb,
			}, false),
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
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationRedirectUrisModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := parse.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewRedirectUrisID(applicationId.ApplicationId, model.UriType)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", applicationId)
			}

			// Check for existing redirect URIs
			if existingUris := r.getRedirectUrisByType(*result, model.UriType); len(existingUris) > 0 {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &id.ApplicationId,
				},
			}

			r.setRedirectUrisByType(&properties, model)

			if _, err = client.Update(ctx, properties); err != nil {
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
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseRedirectUrisID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := parse.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

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

			redirectUris := r.getRedirectUrisByType(*result, id.UriType)

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
			client := metadata.Client.Applications.ApplicationsClient

			id, err := parse.ParseRedirectUrisID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := parse.NewApplicationID(id.ApplicationId)

			var model ApplicationRedirectUrisModel
			if err = metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
			}

			r.setRedirectUrisByType(&properties, model)

			_, err = client.Update(ctx, properties)
			if err != nil {
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
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseRedirectUrisID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := parse.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
			}

			r.deleteRedirectUrisByType(&properties, id.UriType)

			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationRedirectUrisResource) getRedirectUrisByType(application msgraph.Application, uriType string) []string {
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

func (r ApplicationRedirectUrisResource) setRedirectUrisByType(application *msgraph.Application, model ApplicationRedirectUrisModel) {
	switch model.UriType {
	case RedirectUriTypePublicClient:
		application.PublicClient = &msgraph.PublicClient{
			RedirectUris: pointer.To(model.RedirectUris),
		}
	case RedirectUriTypeSPA:
		application.Spa = &msgraph.ApplicationSpa{
			RedirectUris: pointer.To(model.RedirectUris),
		}
	case RedirectUriTypeWeb:
		application.Web = &msgraph.ApplicationWeb{
			RedirectUris: pointer.To(model.RedirectUris),
		}
	}
}
func (r ApplicationRedirectUrisResource) deleteRedirectUrisByType(application *msgraph.Application, uriType string) {
	switch uriType {
	case RedirectUriTypePublicClient:
		application.PublicClient = &msgraph.PublicClient{
			RedirectUris: &[]string{},
		}
	case RedirectUriTypeSPA:
		application.Spa = &msgraph.ApplicationSpa{
			RedirectUris: &[]string{},
		}
	case RedirectUriTypeWeb:
		application.Web = &msgraph.ApplicationWeb{
			RedirectUris: &[]string{},
		}
	}
}
