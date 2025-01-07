// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/migrations"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationIdentifierUriModel struct {
	ApplicationId string `tfschema:"application_id"`
	IdentifierUri string `tfschema:"identifier_uri"`
}

var _ sdk.ResourceWithStateMigration = ApplicationIdentifierUriResource{}

type ApplicationIdentifierUriResource struct{}

func (r ApplicationIdentifierUriResource) StateUpgraders() sdk.StateUpgradeData {
	return sdk.StateUpgradeData{
		SchemaVersion: 1,
		Upgraders: map[int]pluginsdk.StateUpgrade{
			0: migrations.ResourceApplicationIdentifierUriStateUpgradeV0{},
		},
	}
}

func (r ApplicationIdentifierUriResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateIdentifierUriID
}

func (r ApplicationIdentifierUriResource) ResourceType() string {
	return "azuread_application_identifier_uri"
}

func (r ApplicationIdentifierUriResource) ModelObject() interface{} {
	return &ApplicationIdentifierUriModel{}
}

func (r ApplicationIdentifierUriResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which the identifier URI should be added",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: stable.ValidateApplicationID,
		},

		"identifier_uri": {
			Description: "The user-defined URI or URI-like string that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant",
			Type:        pluginsdk.TypeString,
			Required:    true,
			ForceNew:    true,
			// Extensive validation is intentionally avoided here, as the accepted values are undocumented, vary wildly and are
			// different for each user depending on the tenant domain configuration, whether the application is used for SSO etc
			ValidateFunc: validation.StringIsNotEmpty,
		},
	}
}

func (r ApplicationIdentifierUriResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationIdentifierUriResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			var model ApplicationIdentifierUriModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := stable.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			identifierUriSegment := base64.URLEncoding.EncodeToString([]byte(model.IdentifierUri))
			id := parse.NewIdentifierUriID(applicationId.ApplicationId, identifierUriSegment)

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

			newIdentifierUris := make([]string, 0)

			// Don't forget any existing identifier URIs, since they must be updated together
			if app.IdentifierUris != nil {
				newIdentifierUris = *app.IdentifierUris
			}

			// Check for existing identifier URI
			for _, uri := range newIdentifierUris {
				if uri == model.IdentifierUri {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			newIdentifierUris = append(newIdentifierUris, model.IdentifierUri)

			properties := stable.Application{
				IdentifierUris: &newIdentifierUris,
			}

			if _, err = client.UpdateApplication(ctx, *applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationIdentifierUriResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseIdentifierUriID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			uriFromIdSegment, err := base64.URLEncoding.DecodeString(id.IdentifierUri)
			if err != nil {
				return fmt.Errorf("failed to decode identifierUri from resource ID: %+v", err)
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
			if app.IdentifierUris == nil {
				return metadata.MarkAsGone(id)
			}

			// Match the identifier URI
			var identifierUri *string
			for _, existingUri := range *app.IdentifierUris {
				if existingUri == string(uriFromIdSegment) {
					identifierUri = &existingUri
					break
				}
			}

			if identifierUri == nil {
				return metadata.MarkAsGone(id)
			}

			state := ApplicationIdentifierUriModel{
				ApplicationId: applicationId.ID(),
				IdentifierUri: *identifierUri,
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationIdentifierUriResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseIdentifierUriID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationIdentifierUriModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			applicationId := stable.NewApplicationID(id.ApplicationId)
			resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil || app.IdentifierUris == nil {
				return fmt.Errorf("retrieving %s: identifierUris was nil", applicationId)
			}

			// Look for the identifier URI to remove
			newIdentifierUris := make([]string, 0)
			found := false
			for _, existingUri := range *app.IdentifierUris {
				if existingUri == model.IdentifierUri {
					found = true
				} else {
					newIdentifierUris = append(newIdentifierUris, existingUri)
				}
			}
			if !found {
				return fmt.Errorf("deleting %s: could not identify existing identifier URI", id)
			}

			properties := stable.Application{
				IdentifierUris: &newIdentifierUris,
			}

			// Patch the application with the new set of identifier URIs
			if _, err = client.UpdateApplication(ctx, applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
