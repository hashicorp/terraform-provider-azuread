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
)

type ApplicationFallbackPublicClientModel struct {
	ApplicationId string `tfschema:"application_id"`
	Enabled       bool   `tfschema:"enabled"`
}

type ApplicationFallbackPublicClientResource struct{}

func (r ApplicationFallbackPublicClientResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateFallbackPublicClientID
}

var _ sdk.Resource = ApplicationFallbackPublicClientResource{}

func (r ApplicationFallbackPublicClientResource) ResourceType() string {
	return "azuread_application_fallback_public_client"
}

func (r ApplicationFallbackPublicClientResource) ModelObject() interface{} {
	return &ApplicationFallbackPublicClientModel{}
}

func (r ApplicationFallbackPublicClientResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which the fallback public client setting should be applied",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: parse.ValidateApplicationID,
		},

		"enabled": {
			Description: "Specifies explicitly whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI",
			Type:        pluginsdk.TypeBool,
			Optional:    true,
			Default:     false,
			ForceNew:    true,
		},
	}
}

func (r ApplicationFallbackPublicClientResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationFallbackPublicClientResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationFallbackPublicClientModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := parse.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewFallbackPublicClientID(applicationId.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			if _, err = client.SetFallbackPublicClient(ctx, id.ApplicationId, pointer.To(model.Enabled)); err != nil {
				return fmt.Errorf("setting %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationFallbackPublicClientResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseFallbackPublicClientID(metadata.ResourceData.Id())
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
			if result.IsFallbackPublicClient == nil {
				return metadata.MarkAsGone(id)
			}

			state := ApplicationFallbackPublicClientModel{
				ApplicationId: applicationId.ID(),
				Enabled:       pointer.From(result.IsFallbackPublicClient),
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationFallbackPublicClientResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseFallbackPublicClientID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationFallbackPublicClientModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			_, err = client.SetFallbackPublicClient(ctx, id.ApplicationId, nil)
			if err != nil {
				return fmt.Errorf("unsetting %s: %+v", id, err)
			}

			return nil
		},
	}
}
