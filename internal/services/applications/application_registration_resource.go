// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationRegistrationModel struct {
	ClientId                           string   `tfschema:"client_id"`
	Description                        string   `tfschema:"description"`
	DisabledByMicrosoft                string   `tfschema:"disabled_by_microsoft"`
	DisplayName                        string   `tfschema:"display_name"`
	GroupMembershipClaims              []string `tfschema:"group_membership_claims"`
	HomepageUrl                        string   `tfschema:"homepage_url"`
	ImplicitAccessTokenIssuanceEnabled bool     `tfschema:"implicit_access_token_issuance_enabled"`
	ImplicitIdTokenIssuanceEnabled     bool     `tfschema:"implicit_id_token_issuance_enabled"`
	LogoutUrl                          string   `tfschema:"logout_url"`
	MarketingUrl                       string   `tfschema:"marketing_url"`
	Notes                              string   `tfschema:"notes"`
	ObjectId                           string   `tfschema:"object_id"`
	PrivacyStatementUrl                string   `tfschema:"privacy_statement_url"`
	PublisherDomain                    string   `tfschema:"publisher_domain"`
	RequestedAccessTokenVersion        int      `tfschema:"requested_access_token_version"`
	ServiceManagementReference         string   `tfschema:"service_management_reference"`
	SignInAudience                     string   `tfschema:"sign_in_audience"`
	SupportUrl                         string   `tfschema:"support_url"`
	TermsOfServiceUrl                  string   `tfschema:"terms_of_service_url"`
}

var _ sdk.ResourceWithUpdate = ApplicationRegistrationResource{}

type ApplicationRegistrationResource struct{}

func (r ApplicationRegistrationResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateApplicationID
}

func (r ApplicationRegistrationResource) ResourceType() string {
	return "azuread_application_registration"
}

func (r ApplicationRegistrationResource) ModelObject() interface{} {
	return &ApplicationRegistrationModel{}
}

func (r ApplicationRegistrationResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"display_name": {
			Description:  "The display name for the application",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"description": {
			Description:  "Description of the application as shown to end users",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 1024),
		},

		"group_membership_claims": {
			Description: "Configures the `groups` claim that the app expects issued in a user or OAuth access token",
			Type:        pluginsdk.TypeSet,
			Optional:    true,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.StringInSlice(possibleValuesForGroupMembershipClaim, false),
			},
		},

		"homepage_url": {
			Description:  "URL of the home page for the application",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsHttpOrHttpsUrl,
		},

		"implicit_access_token_issuance_enabled": {
			Description: "Whether this application can request an access token using OAuth implicit flow",
			Type:        pluginsdk.TypeBool,
			Optional:    true,
		},

		"implicit_id_token_issuance_enabled": {
			Description: "Whether this application can request an ID token using OAuth implicit flow",
			Type:        pluginsdk.TypeBool,
			Optional:    true,
		},

		"logout_url": {
			Description:  "URL of the logout page for the application, where the session is cleared for single sign-out",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsLogoutUrl,
		},

		"marketing_url": {
			Description:  "URL of the marketing page for the application",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsHttpOrHttpsUrl,
		},

		"notes": {
			Description:  "User-specified notes relevant for the management of the application",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"privacy_statement_url": {
			Description:  "URL of the privacy statement for the application",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsHttpOrHttpsUrl,
		},

		"requested_access_token_version": {
			Description:  "The access token version expected by this resource",
			Type:         pluginsdk.TypeInt,
			Optional:     true,
			Default:      2,
			ValidateFunc: validation.IntBetween(1, 2),
		},

		"service_management_reference": {
			Description:  "References application or contact information from a service or asset management database",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"sign_in_audience": {
			Description:  "The Microsoft account types that are supported for the current application",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			Default:      SignInAudienceAzureADMyOrg,
			ValidateFunc: validation.StringInSlice(possibleValuesForSignInAudience, false),
		},

		"support_url": {
			Description:  "URL of the support page for the application",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsHttpOrHttpsUrl,
		},

		"terms_of_service_url": {
			Description:  "URL of the terms of service statement for the application",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsHttpOrHttpsUrl,
		},
	}
}

func (r ApplicationRegistrationResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"client_id": {
			Description: "The Client ID (also called Application ID)",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"disabled_by_microsoft": {
			Description: "If the application has been disabled by Microsoft, this shows the status or reason",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"object_id": {
			Description: "The object ID of the application within the tenant",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"publisher_domain": {
			Description: "The verified publisher domain for the application",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (r ApplicationRegistrationResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			var model ApplicationRegistrationModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			properties := stable.Application{
				DisplayName:                nullable.Value(model.DisplayName),
				Description:                nullable.NoZero(model.Description),
				GroupMembershipClaims:      expandApplicationGroupMembershipClaims(tf.FlattenStringSlice(model.GroupMembershipClaims)),
				Notes:                      nullable.NoZero(model.Notes),
				ServiceManagementReference: nullable.NoZero(model.ServiceManagementReference),
				SignInAudience:             nullable.Value(model.SignInAudience),

				Api: &stable.ApiApplication{
					RequestedAccessTokenVersion: nullable.Value(int64(model.RequestedAccessTokenVersion)),
				},

				Info: &stable.InformationalUrl{
					MarketingUrl:        nullable.NoZero(model.MarketingUrl),
					PrivacyStatementUrl: nullable.NoZero(model.PrivacyStatementUrl),
					SupportUrl:          nullable.NoZero(model.SupportUrl),
					TermsOfServiceUrl:   nullable.NoZero(model.TermsOfServiceUrl),
				},

				Web: &stable.WebApplication{
					HomePageUrl: nullable.NoZero(model.HomepageUrl),
					LogoutUrl:   nullable.NoZero(model.LogoutUrl),

					ImplicitGrantSettings: &stable.ImplicitGrantSettings{
						EnableAccessTokenIssuance: nullable.Value(model.ImplicitAccessTokenIssuanceEnabled),
						EnableIdTokenIssuance:     nullable.Value(model.ImplicitIdTokenIssuanceEnabled),
					},
				},
			}

			resp, err := client.CreateApplication(ctx, properties)
			if err != nil {
				return fmt.Errorf("creating applicatoin: %+v", err)
			}

			app := resp.Model
			if app == nil || pointer.From(app.Id) == "" {
				return errors.New("creating applicatoin: object ID returned for application is nil/empty")
			}

			id := stable.NewApplicationID(*app.Id)
			metadata.SetID(id)

			return nil
		},
	}
}

func (r ApplicationRegistrationResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := stable.ParseApplicationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.GetApplication(ctx, *id, application.DefaultGetApplicationOperationOptions())
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

			state := ApplicationRegistrationModel{
				ClientId:                   app.AppId.GetOrZero(),
				Description:                app.Description.GetOrZero(),
				DisplayName:                app.DisplayName.GetOrZero(),
				GroupMembershipClaims:      tf.ExpandStringSlice(flattenApplicationGroupMembershipClaims(app.GroupMembershipClaims)),
				Notes:                      app.Notes.GetOrZero(),
				ObjectId:                   pointer.From(app.Id),
				PublisherDomain:            app.PublisherDomain.GetOrZero(),
				ServiceManagementReference: app.ServiceManagementReference.GetOrZero(),
				SignInAudience:             app.SignInAudience.GetOrZero(),
			}

			if api := app.Api; api != nil {
				state.RequestedAccessTokenVersion = int(api.RequestedAccessTokenVersion.GetOrZero())
			}

			if info := app.Info; info != nil {
				state.MarketingUrl = info.MarketingUrl.GetOrZero()
				state.PrivacyStatementUrl = info.PrivacyStatementUrl.GetOrZero()
				state.SupportUrl = info.SupportUrl.GetOrZero()
				state.TermsOfServiceUrl = info.TermsOfServiceUrl.GetOrZero()
			}

			if web := app.Web; web != nil {
				state.HomepageUrl = web.HomePageUrl.GetOrZero()
				state.LogoutUrl = web.LogoutUrl.GetOrZero()

				if implicitGrant := web.ImplicitGrantSettings; implicitGrant != nil {
					state.ImplicitAccessTokenIssuanceEnabled = implicitGrant.EnableAccessTokenIssuance.GetOrZero()
					state.ImplicitIdTokenIssuanceEnabled = implicitGrant.EnableIdTokenIssuance.GetOrZero()
				}
			}

			if app.DisabledByMicrosoftStatus != nil {
				state.DisabledByMicrosoft = fmt.Sprintf("%v", app.DisabledByMicrosoftStatus)
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationRegistrationResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient
			rd := metadata.ResourceData

			id, err := stable.ParseApplicationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationRegistrationModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			properties := stable.Application{}

			if rd.HasChange("display_name") {
				properties.DisplayName = nullable.Value(model.DisplayName)
			}

			if rd.HasChange("description") {
				properties.Description = nullable.NoZero(model.Description)
			}

			if rd.HasChange("group_membership_claims") {
				properties.GroupMembershipClaims = expandApplicationGroupMembershipClaims(tf.FlattenStringSlice(model.GroupMembershipClaims))
			}

			if rd.HasChange("notes") {
				properties.Notes = nullable.NoZero(model.Notes)
			}

			if rd.HasChange("requested_access_token_version") {
				properties.Api = &stable.ApiApplication{
					RequestedAccessTokenVersion: nullable.Value(int64(model.RequestedAccessTokenVersion)),
				}
			}

			if rd.HasChange("service_management_reference") {
				properties.ServiceManagementReference = nullable.NoZero(model.ServiceManagementReference)
			}

			if rd.HasChange("sign_in_audience") {
				properties.SignInAudience = nullable.Value(model.SignInAudience)
			}

			if rd.HasChange("marketing_url") || rd.HasChange("privacy_statement_url") || rd.HasChange("support_url") || rd.HasChange("terms_of_service_url") {
				properties.Info = &stable.InformationalUrl{}

				if rd.HasChange("marketing_url") {
					properties.Info.MarketingUrl = nullable.NoZero(model.MarketingUrl)
				}

				if rd.HasChange("privacy_statement_url") {
					properties.Info.PrivacyStatementUrl = nullable.NoZero(model.PrivacyStatementUrl)
				}

				if rd.HasChange("support_url") {
					properties.Info.SupportUrl = nullable.NoZero(model.SupportUrl)
				}

				if rd.HasChange("terms_of_service_url") {
					properties.Info.TermsOfServiceUrl = nullable.NoZero(model.TermsOfServiceUrl)
				}
			}

			if rd.HasChange("implicit_access_token_issuance_enabled") || rd.HasChange("homepage_url") || rd.HasChange("implicit_id_token_issuance_enabled") || rd.HasChange("logout_url") {
				properties.Web = &stable.WebApplication{}

				if rd.HasChange("homepage_url") {
					properties.Web.HomePageUrl = nullable.NoZero(model.HomepageUrl)
				}

				if rd.HasChange("logout_url") {
					properties.Web.LogoutUrl = nullable.NoZero(model.LogoutUrl)
				}

				if rd.HasChange("implicit_access_token_issuance_enabled") || rd.HasChange("implicit_id_token_issuance_enabled") {
					properties.Web.ImplicitGrantSettings = &stable.ImplicitGrantSettings{}

					if rd.HasChange("implicit_access_token_issuance_enabled") {
						properties.Web.ImplicitGrantSettings.EnableAccessTokenIssuance = nullable.Value(model.ImplicitAccessTokenIssuanceEnabled)
					}

					if rd.HasChange("implicit_id_token_issuance_enabled") {
						properties.Web.ImplicitGrantSettings.EnableIdTokenIssuance = nullable.Value(model.ImplicitIdTokenIssuanceEnabled)
					}
				}
			}

			if _, err = client.UpdateApplication(ctx, *id, properties); err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationRegistrationResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := stable.ParseApplicationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			if _, err = client.DeleteApplication(ctx, *id, application.DefaultDeleteApplicationOperationOptions()); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			// Wait for application object to be deleted
			if err = consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
				if resp, err := client.GetApplication(ctx, *id, application.DefaultGetApplicationOperationOptions()); err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						return pointer.To(false), nil
					}
					return nil, err
				}
				return pointer.To(true), nil
			}); err != nil {
				return fmt.Errorf("waiting for deletion of %s: %q", id, err)
			}

			return nil
		},
	}
}
