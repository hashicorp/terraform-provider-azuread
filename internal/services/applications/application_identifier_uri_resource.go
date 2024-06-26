// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type ApplicationIdentifierUriModel struct {
	ApplicationId string `tfschema:"application_id"`
	IdentifierUri string `tfschema:"identifier_uri"`
}

var _ sdk.Resource = ApplicationIdentifierUriResource{}

type ApplicationIdentifierUriResource struct{}

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
			ValidateFunc: parse.ValidateApplicationID,
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
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationIdentifierUriModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := parse.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			identifierUriSegment := base64.StdEncoding.EncodeToString([]byte(model.IdentifierUri))
			id := parse.NewIdentifierUriID(applicationId.ApplicationId, identifierUriSegment)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", applicationId)
			}

			newIdentifierUris := make([]string, 0)

			// Don't forget any existing identifier URIs, since they must be updated together
			if result.IdentifierUris != nil {
				newIdentifierUris = *result.IdentifierUris
			}

			// Check for existing identifier URI
			for _, uri := range newIdentifierUris {
				if uri == model.IdentifierUri {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			newIdentifierUris = append(newIdentifierUris, model.IdentifierUri)

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &id.ApplicationId,
				},
				IdentifierUris: &newIdentifierUris,
			}

			if _, err = client.Update(ctx, properties); err != nil {
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
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseIdentifierUriID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			uriFromIdSegment, err := base64.StdEncoding.DecodeString(id.IdentifierUri)
			if err != nil {
				return fmt.Errorf("failed to decode identifierUri from resource ID: %+v", err)
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
			if result.IdentifierUris == nil {
				return metadata.MarkAsGone(id)
			}

			// Match the identifier URI
			var identifierUri *string
			for _, existingUri := range *result.IdentifierUris {
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
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

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

			applicationId := parse.NewApplicationID(id.ApplicationId)
			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil || result.IdentifierUris == nil {
				return fmt.Errorf("retrieving %s: identifierUris was nil", applicationId)
			}

			// Look for the identifier URI to remove
			newIdentifierUris := make([]string, 0)
			found := false
			for _, existingUri := range *result.IdentifierUris {
				if existingUri == model.IdentifierUri {
					found = true
				} else {
					newIdentifierUris = append(newIdentifierUris, existingUri)
				}
			}
			if !found {
				return fmt.Errorf("deleting %s: could not identify existing identifier URI", id)
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
				IdentifierUris: &newIdentifierUris,
			}

			// Patch the application with the new set of identifier URIs
			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
