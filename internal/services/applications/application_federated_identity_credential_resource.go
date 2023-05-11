package applications

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
)

func applicationFederatedIdentityCredentialResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationFederatedIdentityCredentialResourceCreate,
		UpdateContext: applicationFederatedIdentityCredentialResourceUpdate,
		ReadContext:   applicationFederatedIdentityCredentialResourceRead,
		DeleteContext: applicationFederatedIdentityCredentialResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.FederatedIdentityCredentialID(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"application_object_id": {
				Description:      "The object ID of the application for which this federated identity credential should be created",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"audiences": {
				Description: "List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.",
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				// TODO: consider making this a scalar value instead of a list in v3.0 (the API now only accepts a single value)
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validate.ValidateDiag(validation.StringIsNotEmpty),
				},
			},

			"display_name": {
				Description:      "A unique display name for the federated identity credential",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.ValidateDiag(validation.StringLenBetween(1, 120)),
			},

			"issuer": {
				Description: "The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"subject": {
				Description: "The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"description": {
				Description: "A description for the federated identity credential",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"credential_id": {
				Description: "A UUID used to uniquely identify this federated identity credential",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func applicationFederatedIdentityCredentialResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).Applications.ApplicationsClient
	objectId := d.Get("application_object_id").(string)

	tf.LockByName(applicationResourceName, objectId)
	defer tf.UnlockByName(applicationResourceName, objectId)

	app, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "application_object_id", "Application with object ID %q was not found", objectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving application with object ID %q", objectId)
	}
	if app == nil || app.ID() == nil {
		return tf.ErrorDiagF(errors.New("nil application or application with nil ID was returned"), "API error retrieving application with object ID %q", objectId)
	}

	credential := msgraph.FederatedIdentityCredential{
		Audiences:   tf.ExpandStringSlicePtr(d.Get("audiences").([]interface{})),
		Description: utils.NullableString(d.Get("description").(string)),
		Issuer:      utils.String(d.Get("issuer").(string)),
		Name:        utils.String(d.Get("display_name").(string)),
		Subject:     utils.String(d.Get("subject").(string)),
	}

	newCredential, _, err := client.CreateFederatedIdentityCredential(ctx, *app.ID(), credential)
	if err != nil {
		return tf.ErrorDiagF(err, "Adding federated identity credential for application with object ID %q", *app.ID())
	}
	if newCredential == nil {
		return tf.ErrorDiagF(errors.New("nil credential received when adding federated identity credential"), "API error adding federated identity credential for application with object ID %q", *app.ID())
	}
	if newCredential.ID == nil {
		return tf.ErrorDiagF(errors.New("nil or empty ID received"), "API error adding federated identity credential for application with object ID %q", *app.ID())
	}

	id := parse.NewCredentialID(*app.ID(), "federatedIdentityCredential", *newCredential.ID)

	// Wait for the credential to replicate
	timeout, _ := ctx.Deadline()
	polledForCredential, err := (&resource.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			credentials, _, err := client.ListFederatedIdentityCredentials(ctx, id.ObjectId, odata.Query{})
			if err != nil {
				return nil, "Error", err
			}

			if credentials != nil {
				for _, cred := range *credentials {
					if cred.ID != nil && strings.EqualFold(*cred.ID, id.KeyId) {
						return &cred, "Done", nil
					}
				}
			}

			return nil, "Waiting", nil
		},
	}).WaitForStateContext(ctx)

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for federated identity credential for application with object ID %q", id.ObjectId)
	} else if polledForCredential == nil {
		return tf.ErrorDiagF(errors.New("federated identity credential not found in application manifest"), "Waiting for federated identity credential for application with object ID %q", id.ObjectId)
	}

	d.SetId(id.String())

	return applicationFederatedIdentityCredentialResourceRead(ctx, d, meta)
}

func applicationFederatedIdentityCredentialResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).Applications.ApplicationsClient

	id, err := parse.FederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing federated identity credential with ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	credential := msgraph.FederatedIdentityCredential{
		ID:          utils.String(id.KeyId),
		Audiences:   tf.ExpandStringSlicePtr(d.Get("audiences").([]interface{})),
		Description: utils.NullableString(d.Get("description").(string)),
		Issuer:      utils.String(d.Get("issuer").(string)),
		Subject:     utils.String(d.Get("subject").(string)),
	}

	_, err = client.UpdateFederatedIdentityCredential(ctx, id.ObjectId, credential)
	if err != nil {
		return tf.ErrorDiagF(err, "Updating federated identity credential with ID %q for application with object ID %q", id.KeyId, id.ObjectId)
	}

	return applicationFederatedIdentityCredentialResourceRead(ctx, d, meta)
}

func applicationFederatedIdentityCredentialResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).Applications.ApplicationsClient

	id, err := parse.FederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing federated identity credential with ID %q", d.Id())
	}

	credential, status, err := client.GetFederatedIdentityCredential(ctx, id.ObjectId, id.KeyId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Federated Identity Credential with ID %q for Application %s was not found - removing from state!", id.KeyId, id.ObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "id", "Retrieving federated identity credential with ID %q for application with object ID %q", id.KeyId, id.ObjectId)
	}

	tf.Set(d, "application_object_id", id.ObjectId)
	tf.Set(d, "credential_id", id.KeyId)

	tf.Set(d, "audiences", tf.FlattenStringSlicePtr(credential.Audiences))
	tf.Set(d, "description", credential.Description)
	tf.Set(d, "display_name", credential.Name)
	tf.Set(d, "issuer", credential.Issuer)
	tf.Set(d, "subject", credential.Subject)

	return nil
}

func applicationFederatedIdentityCredentialResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).Applications.ApplicationsClient

	id, err := parse.FederatedIdentityCredentialID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing federated identity credential with ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	if _, err := client.DeleteFederatedIdentityCredential(ctx, id.ObjectId, id.KeyId); err != nil {
		return tf.ErrorDiagF(err, "Removing federated identity credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	// Wait for credential to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true

		credentials, _, err := client.ListFederatedIdentityCredentials(ctx, id.ObjectId, odata.Query{})
		if err != nil {
			return nil, err
		}

		if credentials != nil {
			for _, cred := range *credentials {
				if cred.ID != nil && strings.EqualFold(*cred.ID, id.KeyId) {
					return utils.Bool(true), nil
				}
			}
		}

		return utils.Bool(false), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of federated identity credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}
