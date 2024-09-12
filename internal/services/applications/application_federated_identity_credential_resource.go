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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

func applicationFederatedIdentityCredentialResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: applicationFederatedIdentityCredentialResourceCreate,
		UpdateContext: applicationFederatedIdentityCredentialResourceUpdate,
		ReadContext:   applicationFederatedIdentityCredentialResourceRead,
		DeleteContext: applicationFederatedIdentityCredentialResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(15 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.FederatedIdentityCredentialID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"application_id": {
				Description:  "The resource ID of the application for which this federated identity credential should be created",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true, // TODO remove Computed in v3.0
				ForceNew:     true,
				ExactlyOneOf: []string{"application_id", "application_object_id"},
				ValidateFunc: parse.ValidateApplicationID,
			},

			"application_object_id": {
				Description:  "The object ID of the application for which this federated identity credential should be created",
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"application_id", "application_object_id"},
				Deprecated:   "The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider",
				ValidateFunc: validation.Any(validation.IsUUID, parse.ValidateApplicationID),
				DiffSuppressFunc: func(_, oldValue, newValue string, _ *pluginsdk.ResourceData) bool {
					// Where oldValue is a UUID (i.e. the bare object ID), and newValue is a properly formed application
					// resource ID, we'll ignore a diff where these point to the same application resource.
					// This maintains compatibility with configurations mixing the ID attributes, e.g.
					//     application_object_id = azuread_application.example.id
					if _, err := uuid.ParseUUID(oldValue); err == nil {
						if applicationId, err := parse.ParseApplicationID(newValue); err == nil {
							if applicationId.ApplicationId == oldValue {
								return true
							}
						}
					}
					return false
				},
			},

			"audiences": {
				Description: "List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.",
				Type:        pluginsdk.TypeList,
				Required:    true,
				MaxItems:    1,
				// TODO: consider making this a scalar value instead of a list in v3.0 (the API now only accepts a single value)
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"display_name": {
				Description:  "A unique display name for the federated identity credential",
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

			"subject": {
				Description: "The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.",
				Type:        pluginsdk.TypeString,
				Required:    true,
			},

			"description": {
				Description: "A description for the federated identity credential",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"credential_id": {
				Description: "A UUID used to uniquely identify this federated identity credential",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func applicationFederatedIdentityCredentialResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics { //nolint
	client := meta.(*clients.Client).Applications.ApplicationClient
	federatedIdentityCredentialClient := meta.(*clients.Client).Applications.ApplicationFederatedIdentityCredential

	var applicationId *stable.ApplicationId
	var err error
	if v := d.Get("application_id").(string); v != "" {
		if applicationId, err = stable.ParseApplicationID(v); err != nil {
			return tf.ErrorDiagPathF(err, "application_id", "Parsing `application_id`: %q", v)
		}
	} else {
		// TODO: this permits parsing the application_object_id as either a structured ID or a bare UUID, to avoid
		// breaking users who might have `application_object_id = azuread_application.foo.id` in their config, and
		// should be removed in version 3.0 along with the application_object_id property
		v = d.Get("application_object_id").(string)
		if _, err = uuid.ParseUUID(v); err == nil {
			applicationId = pointer.To(stable.NewApplicationID(v))
		} else {
			if applicationId, err = stable.ParseApplicationID(v); err != nil {
				return tf.ErrorDiagPathF(err, "application_id", "Parsing `application_object_id`: %q", v)
			}
		}
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

	credential := stable.FederatedIdentityCredential{
		Audiences:   tf.ExpandStringSlice(d.Get("audiences").([]interface{})),
		Description: nullable.Value(d.Get("description").(string)),
		Issuer:      d.Get("issuer").(string),
		Name:        d.Get("display_name").(string),
		Subject:     d.Get("subject").(string),
	}

	federatedIdentityCredentialResp, err := federatedIdentityCredentialClient.CreateFederatedIdentityCredential(ctx, *applicationId, credential)
	if err != nil {
		return tf.ErrorDiagF(err, "Adding federated identity credential for %s", applicationId)
	}

	newCredential := federatedIdentityCredentialResp.Model
	if newCredential == nil {
		return tf.ErrorDiagF(errors.New("nil credential received when adding federated identity credential"), "API error adding federated identity credential for %s", applicationId)
	}
	if newCredential.Id == nil {
		return tf.ErrorDiagF(errors.New("nil or empty ID received"), "API error adding federated identity credential for %s", applicationId)
	}

	id, err := stable.ParseApplicationIdFederatedIdentityCredentialID(*newCredential.Id)
	if err != nil {
		return tf.ErrorDiagF(err, "parsing federated identity credential ID for %s: %q", applicationId, *newCredential.Id)
	}

	// Wait for the credential to replicate
	timeout, _ := ctx.Deadline()
	polledForCredential, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			resp, err := federatedIdentityCredentialClient.GetFederatedIdentityCredential(ctx, *id, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
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
		return tf.ErrorDiagF(errors.New("federated identity credential not found in application manifest"), "Waiting for federated identity credential%s", id)
	}

	// TODO: migrate this to a stable.ApplicationIdFederatedIdentityCredentialId
	resourceId := parse.NewCredentialID(applicationId.ApplicationId, "federatedIdentityCredential", *newCredential.Id)
	d.SetId(resourceId.String())

	return applicationFederatedIdentityCredentialResourceRead(ctx, d, meta)
}

func applicationFederatedIdentityCredentialResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics { //nolint
	federatedIdentityCredentialClient := meta.(*clients.Client).Applications.ApplicationFederatedIdentityCredential

	id, err := parse.FederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing federated identity credential with ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	credential := stable.FederatedIdentityCredential{
		Id:          pointer.To(id.KeyId),
		Audiences:   tf.ExpandStringSlice(d.Get("audiences").([]interface{})),
		Description: nullable.Value(d.Get("description").(string)),
		Issuer:      d.Get("issuer").(string),
		Subject:     d.Get("subject").(string),
	}

	credentialId := stable.NewApplicationIdFederatedIdentityCredentialID(id.ObjectId, id.KeyId)

	if _, err = federatedIdentityCredentialClient.UpdateFederatedIdentityCredential(ctx, credentialId, credential); err != nil {
		return tf.ErrorDiagF(err, "Updating federated identity credential with ID %q for application with object ID %q", id.KeyId, id.ObjectId)
	}

	return applicationFederatedIdentityCredentialResourceRead(ctx, d, meta)
}

func applicationFederatedIdentityCredentialResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics { //nolint
	federatedIdentityCredentialClient := meta.(*clients.Client).Applications.ApplicationFederatedIdentityCredential

	id, err := parse.FederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing federated identity credential with ID %q", d.Id())
	}

	applicationId := parse.NewApplicationID(id.ObjectId)
	credentialId := stable.NewApplicationIdFederatedIdentityCredentialID(id.ObjectId, id.KeyId)

	resp, err := federatedIdentityCredentialClient.GetFederatedIdentityCredential(ctx, credentialId, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] Federated Identity Credential with ID %q for Application %s was not found - removing from state!", id.KeyId, id.ObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving federated identity credential with ID %q for application with object ID %q", id.KeyId, id.ObjectId)
	}

	credential := resp.Model
	if credential == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "retrieving %s", credentialId)
	}

	tf.Set(d, "application_id", applicationId.ID())
	tf.Set(d, "credential_id", id.KeyId)

	tf.Set(d, "audiences", tf.FlattenStringSlice(credential.Audiences))
	tf.Set(d, "description", credential.Description)
	tf.Set(d, "display_name", credential.Name)
	tf.Set(d, "issuer", credential.Issuer)
	tf.Set(d, "subject", credential.Subject)

	if v := d.Get("application_object_id").(string); v != "" {
		tf.Set(d, "application_object_id", v)
	} else {
		tf.Set(d, "application_object_id", id.ObjectId)
	}

	return nil
}

func applicationFederatedIdentityCredentialResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics { //nolint
	federatedIdentityCredentialClient := meta.(*clients.Client).Applications.ApplicationFederatedIdentityCredential

	id, err := parse.FederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing federated identity credential with ID %q", d.Id())
	}

	credentialId := stable.NewApplicationIdFederatedIdentityCredentialID(id.ObjectId, id.KeyId)

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	if _, err := federatedIdentityCredentialClient.DeleteFederatedIdentityCredential(ctx, credentialId, federatedidentitycredential.DefaultDeleteFederatedIdentityCredentialOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing %s", credentialId)
	}

	// Wait for credential to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := federatedIdentityCredentialClient.GetFederatedIdentityCredential(ctx, credentialId, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
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
		return tf.ErrorDiagF(err, "Waiting for deletion of %s", credentialId)
	}

	return nil
}
