package azuread

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func resourceApplicationCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCertificateCreate,
		Read:   resourceApplicationCertificateRead,
		Delete: resourceApplicationCertificateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		// Schema: graph.PasswordResourceSchema("application_object"), //todo switch back to this in 1.0
		Schema: map[string]*schema.Schema{
			"application_object_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"AsymmetricX509Cert",
					"Symmetric",
				}, false),
			},

			"value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				//ValidateFunc: validation.StringLenBetween(1, 863), // Encrypted secret cannot be empty and can be at most 1024 bytes.
			},

			"start_date": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},

			"end_date": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"end_date_relative"},
				ValidateFunc:  validation.IsRFC3339Time,
			},

			"end_date_relative": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"end_date"},
				ValidateFunc:  validate.NoEmptyStrings,
			},
		},
	}
}

func resourceApplicationCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	objectId := d.Get("application_object_id").(string)

	cred, err := graph.KeyCredentialForResource(d)
	if err != nil {
		return fmt.Errorf("generating certificate credentials for object ID %q: %+v", objectId, err)
	}
	id := graph.CredentialIdFrom(objectId, *cred.KeyID)

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	existingCreds, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("listing certificate credentials for application with object ID %q: %+v", id.ObjectId, err)
	}

	newCreds, err := graph.KeyCredentialResultAdd(existingCreds, cred, true)
	if err != nil {
		return tf.ImportAsExistsError("azuread_application_certificate", id.String())
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return fmt.Errorf("creating certificate credentials %q for application with object ID %q: %+v", id.KeyId, id.ObjectId, err)
	}

	_, err = graph.WaitForKeyCredentialReplication(id.KeyId, func() (graphrbac.KeyCredentialListResult, error) {
		return client.ListKeyCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return fmt.Errorf("waiting for certificate credential replication for application (AppID %q, KeyID %q: %+v", id.ObjectId, id.KeyId, err)
	}

	d.SetId(id.String())

	return resourceApplicationCertificateRead(d, meta)
}

func resourceApplicationCertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	id, err := graph.ParseCredentialId(d.Id())
	if err != nil {
		return fmt.Errorf("parsing certificate credential with ID: %v", err)
	}
	// ensure the Application Object exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Application has been removed - skip it
		if ar.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving application with ID %q: %+v", id.ObjectId, err)
	}

	credentials, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("listing certificate credentials for application with object ID %q: %+v", id.ObjectId, err)
	}

	credential := graph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	// todo, move this into a graph helper function?
	d.Set("application_object_id", id.ObjectId)
	d.Set("key_id", id.KeyId)

	if keyType := credential.Type; keyType != nil {
		d.Set("type", string(*keyType))
	}

	if endDate := credential.EndDate; endDate != nil {
		d.Set("end_date", endDate.Format(time.RFC3339))
	}

	if startDate := credential.StartDate; startDate != nil {
		d.Set("start_date", startDate.Format(time.RFC3339))
	}

	return nil
}

func resourceApplicationCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	id, err := graph.ParseCredentialId(d.Id())
	if err != nil {
		return fmt.Errorf("parsing certificate credential with ID: %v", err)
	}

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	// ensure the parent Application exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Application has been removed - skip it
		if ar.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return fmt.Errorf("retrieving application with ID %q: %+v", id.ObjectId, err)
	}

	existing, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("listing certificate credentials for application %q: %+v", id.ObjectId, err)
	}

	newCreds := graph.KeyCredentialResultRemoveByKeyId(existing, id.KeyId)
	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return fmt.Errorf("removing certificate credentials %q from application with object ID %q: %+v", id.KeyId, id.ObjectId, err)
	}

	return nil
}
