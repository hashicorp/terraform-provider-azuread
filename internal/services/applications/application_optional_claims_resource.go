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
	"github.com/manicminer/hamilton/msgraph"
)

type ApplicationOptionalClaimsModel struct {
	ApplicationId string          `tfschema:"application_id"`
	AccessTokens  []OptionalClaim `tfschema:"access_token"`
	IdTokens      []OptionalClaim `tfschema:"id_token"`
	Saml2Tokens   []OptionalClaim `tfschema:"saml2_token"`
}

type OptionalClaim struct {
	Name                 string   `tfschema:"name"`
	Source               string   `tfschema:"source"`
	Essential            bool     `tfschema:"essential"`
	AdditionalProperties []string `tfschema:"additional_properties"`
}

type ApplicationOptionalClaimsResource struct{}

func (r ApplicationOptionalClaimsResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateOptionalClaimsID
}

var _ sdk.ResourceWithUpdate = ApplicationOptionalClaimsResource{}

func (r ApplicationOptionalClaimsResource) ResourceType() string {
	return "azuread_application_optional_claims"
}

func (r ApplicationOptionalClaimsResource) ModelObject() interface{} {
	return &ApplicationOptionalClaimsModel{}
}

func (r ApplicationOptionalClaimsResource) Arguments() (ret map[string]*pluginsdk.Schema) {
	ret = map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which these optional claims belong",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: parse.ValidateApplicationID,
		},

		"access_token": schemaOptionalClaims(),
		"id_token":     schemaOptionalClaims(),
		"saml2_token":  schemaOptionalClaims(),
	}

	atLeastOneOf := []string{"access_token", "id_token", "saml2_token"}
	ret["access_token"].AtLeastOneOf = atLeastOneOf
	ret["id_token"].AtLeastOneOf = atLeastOneOf
	ret["saml2_token"].AtLeastOneOf = atLeastOneOf

	return
}

func (r ApplicationOptionalClaimsResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationOptionalClaimsResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationOptionalClaimsModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := parse.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewOptionalClaimsID(applicationId.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", applicationId)
			}

			// Check for existing optional claims
			if claims := result.OptionalClaims; claims != nil {
				if claims.AccessToken != nil && len(*claims.AccessToken) > 0 {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
				if claims.IdToken != nil && len(*claims.IdToken) > 0 {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
				if claims.Saml2Token != nil && len(*claims.Saml2Token) > 0 {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			// Assemble the optional claims
			optionalClaims := msgraph.OptionalClaims{}

			if len(model.AccessTokens) > 0 {
				accessTokenClaims := make([]msgraph.OptionalClaim, 0)
				for _, claim := range model.AccessTokens {
					accessTokenClaims = append(accessTokenClaims, msgraph.OptionalClaim{
						Name:                 pointer.To(claim.Name),
						Source:               pointer.To(claim.Source),
						Essential:            pointer.To(claim.Essential),
						AdditionalProperties: pointer.To(claim.AdditionalProperties),
					})
				}
				optionalClaims.AccessToken = &accessTokenClaims
			}

			if len(model.IdTokens) > 0 {
				idTokenClaims := make([]msgraph.OptionalClaim, 0)
				for _, claim := range model.IdTokens {
					idTokenClaims = append(idTokenClaims, msgraph.OptionalClaim{
						Name:                 pointer.To(claim.Name),
						Source:               pointer.To(claim.Source),
						Essential:            pointer.To(claim.Essential),
						AdditionalProperties: pointer.To(claim.AdditionalProperties),
					})
				}
				optionalClaims.IdToken = &idTokenClaims
			}

			if len(model.Saml2Tokens) > 0 {
				saml2TokenClaims := make([]msgraph.OptionalClaim, 0)
				for _, claim := range model.Saml2Tokens {
					saml2TokenClaims = append(saml2TokenClaims, msgraph.OptionalClaim{
						Name:                 pointer.To(claim.Name),
						Source:               pointer.To(claim.Source),
						Essential:            pointer.To(claim.Essential),
						AdditionalProperties: pointer.To(claim.AdditionalProperties),
					})
				}
				optionalClaims.Saml2Token = &saml2TokenClaims
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &id.ApplicationId,
				},
				OptionalClaims: &optionalClaims,
			}

			if _, err = client.Update(ctx, properties); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationOptionalClaimsResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseOptionalClaimsID(metadata.ResourceData.Id())
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

			if claims := result.OptionalClaims; claims == nil {
				return metadata.MarkAsGone(id)
			} else if (claims.AccessToken == nil || len(*claims.AccessToken) == 0) &&
				(claims.IdToken == nil || len(*claims.IdToken) == 0) &&
				(claims.Saml2Token == nil || len(*claims.Saml2Token) == 0) {
				return metadata.MarkAsGone(id)
			}

			state := ApplicationOptionalClaimsModel{
				ApplicationId: applicationId.ID(),
			}

			if accessTokenClaims := result.OptionalClaims.AccessToken; accessTokenClaims != nil {
				for _, claim := range *accessTokenClaims {
					state.AccessTokens = append(state.AccessTokens, OptionalClaim{
						Name:                 pointer.From(claim.Name),
						Source:               pointer.From(claim.Source),
						Essential:            pointer.From(claim.Essential),
						AdditionalProperties: pointer.From(claim.AdditionalProperties),
					})
				}
			}

			if idTokenClaims := result.OptionalClaims.IdToken; idTokenClaims != nil {
				for _, claim := range *idTokenClaims {
					state.IdTokens = append(state.IdTokens, OptionalClaim{
						Name:                 pointer.From(claim.Name),
						Source:               pointer.From(claim.Source),
						Essential:            pointer.From(claim.Essential),
						AdditionalProperties: pointer.From(claim.AdditionalProperties),
					})
				}
			}

			if idTokenClaims := result.OptionalClaims.Saml2Token; idTokenClaims != nil {
				for _, claim := range *idTokenClaims {
					state.Saml2Tokens = append(state.Saml2Tokens, OptionalClaim{
						Name:                 pointer.From(claim.Name),
						Source:               pointer.From(claim.Source),
						Essential:            pointer.From(claim.Essential),
						AdditionalProperties: pointer.From(claim.AdditionalProperties),
					})
				}
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationOptionalClaimsResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			rd := metadata.ResourceData

			id, err := parse.ParseOptionalClaimsID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationOptionalClaimsModel
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
			if result == nil || result.OptionalClaims == nil {
				return fmt.Errorf("retrieving %s: optionalClaims was nil", applicationId)
			}

			// Start with the existing claims, as they must be updated together, then update each type in turn as needed
			newOptionalClaims := *result.OptionalClaims

			if rd.HasChange("access_token") {
				newAccessTokenClaims := make([]msgraph.OptionalClaim, 0)
				for _, claim := range model.AccessTokens {
					newAccessTokenClaims = append(newAccessTokenClaims, msgraph.OptionalClaim{
						Name:                 pointer.To(claim.Name),
						Source:               pointer.To(claim.Source),
						Essential:            pointer.To(claim.Essential),
						AdditionalProperties: pointer.To(claim.AdditionalProperties),
					})
				}
				newOptionalClaims.AccessToken = &newAccessTokenClaims
			}

			if rd.HasChange("id_token") {
				newIdTokenClaims := make([]msgraph.OptionalClaim, 0)
				for _, claim := range model.IdTokens {
					newIdTokenClaims = append(newIdTokenClaims, msgraph.OptionalClaim{
						Name:                 pointer.To(claim.Name),
						Source:               pointer.To(claim.Source),
						Essential:            pointer.To(claim.Essential),
						AdditionalProperties: pointer.To(claim.AdditionalProperties),
					})
				}
				newOptionalClaims.IdToken = &newIdTokenClaims
			}

			if rd.HasChange("saml2_token") {
				newSaml2TokenClaims := make([]msgraph.OptionalClaim, 0)
				for _, claim := range model.Saml2Tokens {
					newSaml2TokenClaims = append(newSaml2TokenClaims, msgraph.OptionalClaim{
						Name:                 pointer.To(claim.Name),
						Source:               pointer.To(claim.Source),
						Essential:            pointer.To(claim.Essential),
						AdditionalProperties: pointer.To(claim.AdditionalProperties),
					})
				}
				newOptionalClaims.Saml2Token = &newSaml2TokenClaims
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
				OptionalClaims: &newOptionalClaims,
			}

			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationOptionalClaimsResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseOptionalClaimsID(metadata.ResourceData.Id())
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
				OptionalClaims: &msgraph.OptionalClaims{},
			}

			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
