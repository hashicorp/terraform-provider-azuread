package applications

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func applicationPasswordResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationPasswordResourceCreate,
		ReadContext:   applicationPasswordResourceRead,
		DeleteContext: applicationPasswordResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.PasswordID(id)
			return err
		}),

		Schema: aadgraph.PasswordResourceSchema("application_object_id"),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceApplicationPasswordInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceApplicationPasswordInstanceStateUpgradeV0,
				Version: 0,
			},
		},
	}
}

func applicationPasswordResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	objectId := d.Get("application_object_id").(string)

	cred, err := aadgraph.PasswordCredentialForResource(d)
	if err != nil {
		attr := ""
		if kerr, ok := err.(aadgraph.CredentialError); ok {
			attr = kerr.Attr()
		}
		return tf.ErrorDiagPathF(err, attr, "Generating password credentials for application with object ID %q", objectId)
	}
	id := parse.NewCredentialID(objectId, "password", *cred.KeyID)

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	existingCreds, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_object_id", "Listing password credentials for application with ID %q", objectId)
	}

	newCreds, err := aadgraph.PasswordCredentialResultAdd(existingCreds, cred)
	if err != nil {
		if _, ok := err.(*aadgraph.AlreadyExistsError); ok {
			return tf.ImportAsExistsDiag("azuread_application_password", id.String())
		}
		return tf.ErrorDiagF(err, "Adding application password")
	}

	if _, err = client.UpdatePasswordCredentials(ctx, id.ObjectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiagF(err, "Creating password credentials %q for application with object ID %q", id.KeyId, id.ObjectId)
	}

	_, err = aadgraph.WaitForPasswordCredentialReplication(ctx, id.KeyId, d.Timeout(schema.TimeoutCreate), func() (graphrbac.PasswordCredentialListResult, error) {
		return client.ListPasswordCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for certificate credential replication for application (AppID %q, KeyID %q)", id.ObjectId, id.KeyId)
	}

	d.SetId(id.String())

	return applicationPasswordResourceRead(ctx, d, meta)
}

func applicationPasswordResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	id, err := parse.PasswordID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing password credential with ID %q", d.Id())
	}

	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Application has been removed - skip it
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving application with object ID %q", id.ObjectId)
	}

	credentials, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_object_id", "Listing password credentials for application with object ID %q", id.ObjectId)
	}

	credential := aadgraph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Password credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "application_object_id", id.ObjectId)
	tf.Set(d, "key_id", id.KeyId)

	description := ""
	if v := credential.CustomKeyIdentifier; v != nil {
		description = string(*v)
	}
	tf.Set(d, "description", description)

	startDate := ""
	if v := credential.StartDate; v != nil {
		startDate = v.Format(time.RFC3339)
	}
	tf.Set(d, "start_date", startDate)

	endDate := ""
	if v := credential.EndDate; v != nil {
		endDate = v.Format(time.RFC3339)
	}
	tf.Set(d, "end_date", endDate)

	return nil
}

func applicationPasswordResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Applications.AadClient

	id, err := parse.PasswordID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing password credential with ID %q", d.Id())
	}

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	// ensure the parent Application exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Application has been removed - skip it
		if utils.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving application with ID %q", id.ObjectId)
	}

	existing, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagF(err, "Listing password credentials for application with object ID %q", id.ObjectId)
	}

	newCreds, err := aadgraph.PasswordCredentialResultRemoveByKeyId(existing, id.KeyId)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing password credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	if _, err = client.UpdatePasswordCredentials(ctx, id.ObjectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiagF(err, "Removing password credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}

func resourceApplicationPasswordInstanceResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_object_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"key_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"value": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringLenBetween(1, 863),
			},

			"start_date": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},

			"end_date": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"end_date_relative"},
				ValidateFunc: validation.IsRFC3339Time,
			},

			"end_date_relative": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ExactlyOneOf:     []string{"end_date"},
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func resourceApplicationPasswordInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	newId, err := parse.OldPasswordID(rawState["id"].(string))
	if err != nil {
		return rawState, fmt.Errorf("generating new ID: %s", err)
	}

	rawState["id"] = newId.String()
	return rawState, nil
}
