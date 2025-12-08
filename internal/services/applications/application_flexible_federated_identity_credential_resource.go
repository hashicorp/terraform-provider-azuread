// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"time"

	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/glueckkanja/terraform-provider-azuread/internal/sdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/applications/custompollers"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

type flexibleFederatedIdentityCredentialResource struct{}

var _ sdk.ResourceWithUpdate = &flexibleFederatedIdentityCredentialResource{}

type flexibleFederatedIdentityCredentialModel struct {
	ApplicationId            string `tfschema:"application_id"`
	Audience                 string `tfschema:"audience"`
	ClaimsMatchingExpression string `tfschema:"claims_matching_expression"`
	Description              string `tfschema:"description"`
	DisplayName              string `tfschema:"display_name"`
	Issuer                   string `tfschema:"issuer"`
	CredentialId             string `tfschema:"credential_id"`
}

func (f flexibleFederatedIdentityCredentialResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application for which this flexible federated identity credential should be created",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: stable.ValidateApplicationID,
		},

		"audience": {
			Description:  "The audience that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringLenBetween(1, 600),
		},

		"claims_matching_expression": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotWhiteSpace,
			Description:  "The expression to match for claims.",
		},

		"display_name": {
			Description:  "A unique display name for the flexible federated identity credential",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringLenBetween(1, 120),
		},

		"issuer": {
			Description:  "The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotWhiteSpace,
		},

		"description": {
			Description:  "A description for the flexible federated identity credential",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 600),
		},
	}
}

func (f flexibleFederatedIdentityCredentialResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"credential_id": {
			Description: "A UUID used to uniquely identify this flexible federated identity credential",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (f flexibleFederatedIdentityCredentialResource) ModelObject() interface{} {
	return &flexibleFederatedIdentityCredentialModel{}
}

func (f flexibleFederatedIdentityCredentialResource) ResourceType() string {
	return "azuread_application_flexible_federated_identity_credential"
}

func (f flexibleFederatedIdentityCredentialResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 15 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClientBeta
			flexibleFederatedIdentityCredentialClient := metadata.Client.Applications.ApplicationFlexibleFederatedIdentityCredential

			data := &flexibleFederatedIdentityCredentialModel{}

			if err := metadata.Decode(data); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := beta.ParseApplicationID(data.ApplicationId)
			if err != nil {
				return err
			}

			tf.LockByName(applicationResourceName, applicationId.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, applicationId.ApplicationId)

			resp, err := client.GetApplication(ctx, *applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil {
				return fmt.Errorf("retrieving %s: model was nil", applicationId)
			}

			credential := beta.FederatedIdentityCredential{
				Audiences: []string{data.Audience},
				ClaimsMatchingExpression: &beta.FederatedIdentityExpression{
					Value:           data.ClaimsMatchingExpression,
					LanguageVersion: 1, // Note - from docs: the language version to be used. Should always be set to 1, and is required to be set/sent.
				},
				Description: nullable.Value(data.Description),
				Issuer:      data.Issuer,
				Name:        data.DisplayName,
			}

			federatedIdentityCredentialResp, err := flexibleFederatedIdentityCredentialClient.CreateFederatedIdentityCredential(ctx, *applicationId, credential, federatedidentitycredential.DefaultCreateFederatedIdentityCredentialOperationOptions())
			if err != nil {
				return fmt.Errorf("adding flexible federated identity credential for %s", applicationId)
			}

			newCredential := federatedIdentityCredentialResp.Model
			if newCredential == nil {
				return fmt.Errorf("api error adding flexible federated identity credential for %s. nil credential received when adding flexible federated identity credential", applicationId)
			}
			if newCredential.Id == nil {
				return fmt.Errorf("api error adding flexible federated identity credential for %s. nil or empty ID received", applicationId)
			}

			id := beta.NewApplicationIdFederatedIdentityCredentialID(applicationId.ApplicationId, *newCredential.Id)

			pollerType := custompollers.NewApplicationFlexibleFederatedCredentialCreationPoller(flexibleFederatedIdentityCredentialClient, id)
			poller := pollers.NewPoller(pollerType, 10*time.Second, pollers.DefaultNumberOfDroppedConnectionsToAllow)
			// Wait for the credential to replicate - TODO This may need converting to a consistency.WaitForUpdate
			if err := poller.PollUntilDone(ctx); err != nil {
				return err
			}

			metadata.SetID(id)

			return nil
		},
	}
}

func (f flexibleFederatedIdentityCredentialResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			federatedIdentityCredentialClient := metadata.Client.Applications.ApplicationFlexibleFederatedIdentityCredential

			state := flexibleFederatedIdentityCredentialModel{}

			id, err := beta.ParseApplicationIdFederatedIdentityCredentialID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := federatedIdentityCredentialClient.GetFederatedIdentityCredential(ctx, *id, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}

				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			state.ApplicationId = beta.NewApplicationID(id.ApplicationId).ID()
			state.CredentialId = id.FederatedIdentityCredentialId
			if model := resp.Model; model != nil {
				if model.ClaimsMatchingExpression != nil {
					state.ClaimsMatchingExpression = model.ClaimsMatchingExpression.Value
				}
				if len(model.Audiences) > 0 {
					state.Audience = model.Audiences[0]
				}
				state.Description = model.Description.GetOrZero()
				state.DisplayName = model.Name
				state.Issuer = model.Issuer
			}

			return metadata.Encode(&state)
		},
	}
}

func (f flexibleFederatedIdentityCredentialResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			federatedIdentityCredentialClient := metadata.Client.Applications.ApplicationFlexibleFederatedIdentityCredential

			id, err := beta.ParseApplicationIdFederatedIdentityCredentialID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			data := &flexibleFederatedIdentityCredentialModel{}

			if err := metadata.Decode(data); err != nil {
				return fmt.Errorf("decoding %s: %+v", id, err)
			}

			credential := beta.FederatedIdentityCredential{
				Id:          pointer.To(id.FederatedIdentityCredentialId),
				Audiences:   []string{data.Audience},
				Description: nullable.Value(data.Description),
				Issuer:      data.Issuer,
				ClaimsMatchingExpression: &beta.FederatedIdentityExpression{
					Value:           data.ClaimsMatchingExpression,
					LanguageVersion: 1, // Note - from docs: the language version to be used. Should always be set to 1, and is required to be set/sent.
				},
				// Name is immutable but must be specified as it is a required field
				Name: data.DisplayName,
			}

			if _, err = federatedIdentityCredentialClient.UpdateFederatedIdentityCredential(ctx, *id, credential, federatedidentitycredential.DefaultUpdateFederatedIdentityCredentialOperationOptions()); err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (f flexibleFederatedIdentityCredentialResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			federatedIdentityCredentialClient := metadata.Client.Applications.ApplicationFlexibleFederatedIdentityCredential

			id, err := beta.ParseApplicationIdFederatedIdentityCredentialID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			if _, err = federatedIdentityCredentialClient.DeleteFederatedIdentityCredential(ctx, *id, federatedidentitycredential.DefaultDeleteFederatedIdentityCredentialOperationOptions()); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			// Wait for credential to be deleted
			if err = consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
				resp, err := federatedIdentityCredentialClient.GetFederatedIdentityCredential(ctx, *id, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
				if err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						return pointer.To(false), nil
					}
					return nil, err
				}
				credential := resp.Model
				if credential == nil {
					return pointer.To(false), nil
				}

				return pointer.To(true), nil
			}); err != nil {
				return fmt.Errorf("waiting for deletion of %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (f flexibleFederatedIdentityCredentialResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return beta.ValidateApplicationIdFederatedIdentityCredentialID
}
