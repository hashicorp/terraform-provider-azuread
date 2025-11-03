// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func applicationFlexibleFederatedIdentityCredentialResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: applicationFlexibleFederatedIdentityCredentialResourceCreate,
		UpdateContext: applicationFlexibleFederatedIdentityCredentialResourceUpdate,
		ReadContext:   applicationFlexibleFederatedIdentityCredentialResourceRead,
		DeleteContext: applicationFlexibleFederatedIdentityCredentialResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(15 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := beta.ValidateApplicationIdFederatedIdentityCredentialID(id, "")
			return errors.Join(err...)
		}),

		Schema: map[string]*pluginsdk.Schema{
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
				Description: "The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.",
				Type:        pluginsdk.TypeString,
				Required:    true,
			},

			"description": {
				Description:  "A description for the flexible federated identity credential",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 600),
			},

			"credential_id": {
				Description: "A UUID used to uniquely identify this flexible federated identity credential",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func applicationFlexibleFederatedIdentityCredentialResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics { // nolint
	client := meta.(*clients.Client).Applications.ApplicationClientBeta
	federatedIdentityCredentialClient := meta.(*clients.Client).Applications.ApplicationFlexibleFederatedIdentityCredential

	applicationId, err := beta.ParseApplicationID(d.Get("application_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_id", "Parsing `application_id`")
	}

	tf.LockByName(applicationResourceName, applicationId.ApplicationId)
	defer tf.UnlockByName(applicationResourceName, applicationId.ApplicationId)

	resp, err := client.GetApplication(ctx, *applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "retrieving %s: %+v", applicationId, err)
	}

	app := resp.Model
	if app == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "retrieving %s", applicationId)
	}

	credential := beta.FederatedIdentityCredential{
		Audiences: []string{d.Get("audience").(string)},
		ClaimsMatchingExpression: &beta.FederatedIdentityExpression{
			Value:           d.Get("claims_matching_expression").(string),
			LanguageVersion: 1, // Note - from docs: the language version to be used. Should always be set to 1, and is required to be set/sent.
		},
		Description: nullable.Value(d.Get("description").(string)),
		Issuer:      d.Get("issuer").(string),
		Name:        d.Get("display_name").(string),
	}

	federatedIdentityCredentialResp, err := federatedIdentityCredentialClient.CreateFederatedIdentityCredential(ctx, *applicationId, credential, federatedidentitycredential.DefaultCreateFederatedIdentityCredentialOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Adding flexible federated identity credential for %s", applicationId)
	}

	newCredential := federatedIdentityCredentialResp.Model
	if newCredential == nil {
		return tf.ErrorDiagF(errors.New("nil credential received when adding flexible federated identity credential"), "API error adding flexible federated identity credential for %s", applicationId)
	}
	if newCredential.Id == nil {
		return tf.ErrorDiagF(errors.New("nil or empty ID received"), "API error adding flexible federated identity credential for %s", applicationId)
	}

	id := beta.NewApplicationIdFederatedIdentityCredentialID(applicationId.ApplicationId, *newCredential.Id)

	// Wait for the credential to replicate
	timeout, _ := ctx.Deadline()
	polledForCredential, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			resp, err := federatedIdentityCredentialClient.GetFederatedIdentityCredential(ctx, id, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return nil, "Waiting", nil
				}
				return nil, "Error", err
			}
			credential := resp.Model
			if credential == nil {
				return nil, "Waiting", nil
			}

			return credential, "Done", nil
		},
	}).WaitForStateContext(ctx)

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for %s", id)
	} else if polledForCredential == nil {
		return tf.ErrorDiagF(errors.New("flexible federated identity credential not found in application manifest"), "Waiting for flexible federated identity credential%s", id)
	}

	d.SetId(id.ID())

	return applicationFlexibleFederatedIdentityCredentialResourceRead(ctx, d, meta)
}

func applicationFlexibleFederatedIdentityCredentialResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics { // nolint
	federatedIdentityCredentialClient := meta.(*clients.Client).Applications.ApplicationFlexibleFederatedIdentityCredential

	id, err := beta.ParseApplicationIdFederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing flexible federated identity credential with ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ApplicationId)
	defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

	credential := beta.FederatedIdentityCredential{
		Id:          pointer.To(id.FederatedIdentityCredentialId),
		Audiences:   []string{d.Get("audience").(string)},
		Description: nullable.Value(d.Get("description").(string)),
		Issuer:      d.Get("issuer").(string),
		ClaimsMatchingExpression: &beta.FederatedIdentityExpression{
			Value:           d.Get("claims_matching_expression").(string),
			LanguageVersion: 1, // Note - from docs: the language version to be used. Should always be set to 1, and is required to be set/sent.
		},
		// Name is immutable but must be specified as it is a required field
		Name: d.Get("display_name").(string),
	}

	if _, err = federatedIdentityCredentialClient.UpdateFederatedIdentityCredential(ctx, *id, credential, federatedidentitycredential.DefaultUpdateFederatedIdentityCredentialOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Updating flexible federated identity credential with ID %q for application with object ID %q", id.FederatedIdentityCredentialId, id.ApplicationId)
	}

	return applicationFlexibleFederatedIdentityCredentialResourceRead(ctx, d, meta)
}

func applicationFlexibleFederatedIdentityCredentialResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics { // nolint
	federatedIdentityCredentialClient := meta.(*clients.Client).Applications.ApplicationFlexibleFederatedIdentityCredential

	id, err := beta.ParseApplicationIdFederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing flexible federated identity credential with ID %q", d.Id())
	}

	applicationId := beta.NewApplicationID(id.ApplicationId)

	resp, err := federatedIdentityCredentialClient.GetFederatedIdentityCredential(ctx, *id, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] Flexible Federated Identity Credential with ID %q for Application %s was not found - removing from state!", id.ApplicationId, id.FederatedIdentityCredentialId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving flexible federated identity credential with ID %q for application with object ID %q", id.FederatedIdentityCredentialId, id.ApplicationId)
	}

	credential := resp.Model
	if credential == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "retrieving %s", *id)
	}

	tf.Set(d, "application_id", applicationId.ID())
	tf.Set(d, "credential_id", id.FederatedIdentityCredentialId)
	if credential.ClaimsMatchingExpression != nil {
		tf.Set(d, "claims_matching_expression", credential.ClaimsMatchingExpression.Value)
	}
	if len(credential.Audiences) > 0 {
		tf.Set(d, "audience", credential.Audiences[0])
	}
	tf.Set(d, "description", credential.Description.GetOrZero())
	tf.Set(d, "display_name", credential.Name)
	tf.Set(d, "issuer", credential.Issuer)

	return nil
}

func applicationFlexibleFederatedIdentityCredentialResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics { // nolint
	federatedIdentityCredentialClient := meta.(*clients.Client).Applications.ApplicationFlexibleFederatedIdentityCredential

	id, err := beta.ParseApplicationIdFederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing flexible federated identity credential with ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ApplicationId)
	defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

	if _, err := federatedIdentityCredentialClient.DeleteFederatedIdentityCredential(ctx, *id, federatedidentitycredential.DefaultDeleteFederatedIdentityCredentialOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing %s", *id)
	}

	// Wait for credential to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
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
		return tf.ErrorDiagF(err, "Waiting for deletion of %s", *id)
	}

	return nil
}
