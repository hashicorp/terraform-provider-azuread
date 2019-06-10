package azuread

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func resourceApplicationPassword() *schema.Resource {

	// temporary terrible hack/code to allow deprecation of `application_id`
	// todo remove in 1.0
	s := graph.PasswordResourceSchema("application_object")
	s["application_id"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ForceNew:      true,
		Computed:      true,
		ValidateFunc:  validate.UUID,
		Deprecated:    "Deprecated in favour of `application_object_id` to prevent confusion",
		ConflictsWith: []string{"application_id"},
	}
	// this is bad, i am aware of it, and I feel awful about it
	s["application_object_id"].Required = false
	s["application_object_id"].Optional = true
	s["application_object_id"].Computed = true
	s["application_object_id"].ConflictsWith = []string{"application_object_id"}

	return &schema.Resource{
		Create: resourceApplicationPasswordCreate,
		Read:   resourceApplicationPasswordRead,
		Delete: resourceApplicationPasswordDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		//Schema: graph.PasswordResourceSchema("application_object"),
		Schema: s,
	}
}

func resourceApplicationPasswordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	objectId := d.Get("application_object_id").(string)
	if objectId == "" { // todo remove in 1.0
		objectId = d.Get("application_id").(string)
	}
	if objectId == "" {
		return fmt.Errorf("one of `application_object_id` or `application_id` must be specifed")
	}

	cred, err := graph.PasswordCredentialForResource(d)
	if err != nil {
		return fmt.Errorf("Error generating Application Credentials for Object ID %q: %+v", objectId, err)
	}
	id := graph.PasswordCredentialIdFrom(objectId, *cred.KeyID)

	azureADLockByName(resourceApplicationName, id.ObjectId)
	defer azureADUnlockByName(resourceApplicationName, id.ObjectId)

	existingCreds, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("Error Listing Application Credentials for Object ID %q: %+v", id.ObjectId, err)
	}

	newCreds, err := graph.PasswordCredentialResultAdd(existingCreds, cred, requireResourcesToBeImported)
	if err != nil {
		return tf.ImportAsExistsError("azuread_application_password", id.String())
	}

	if _, err = client.UpdatePasswordCredentials(ctx, id.ObjectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return fmt.Errorf("Error creating Application Credentials %q for Object ID %q: %+v", *cred.KeyID, id.ObjectId, err)
	}

	d.SetId(id.String())

	return resourceApplicationPasswordRead(d, meta)
}

func resourceApplicationPasswordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	id, err := graph.ParsePasswordCredentialId(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Application Password ID: %v", err)
	}

	// ensure the Application Object exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Service Principal has been removed - skip it
		if ar.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error retrieving Application ID %q: %+v", id.ObjectId, err)
	}

	credentials, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("Error Listing Application Credentials for Application with Object ID %q: %+v", id.ObjectId, err)
	}

	credential := graph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Application Credentials %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	// todo, move this into a graph helper function?
	d.Set("application_object_id", id.ObjectId)
	d.Set("application_id", id.ObjectId) //todo remove in 2.0
	d.Set("key_id", id.KeyId)

	if endDate := credential.EndDate; endDate != nil {
		d.Set("end_date", endDate.Format(time.RFC3339))
	}

	if startDate := credential.StartDate; startDate != nil {
		d.Set("start_date", startDate.Format(time.RFC3339))
	}

	return nil
}

func resourceApplicationPasswordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).applicationsClient
	ctx := meta.(*ArmClient).StopContext

	id, err := graph.ParsePasswordCredentialId(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Application Password ID: %v", err)
	}

	azureADLockByName(resourceApplicationName, id.ObjectId)
	defer azureADUnlockByName(resourceApplicationName, id.ObjectId)

	// ensure the parent Application exists
	app, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Service Principal has been removed - skip it
		if ar.ResponseWasNotFound(app.Response) {
			log.Printf("[DEBUG] Application with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return fmt.Errorf("Error retrieving Application ID %q: %+v", id.ObjectId, err)
	}

	existing, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("Error Listing Application Credentials for %q: %+v", id.ObjectId, err)
	}

	newCreds := graph.PasswordCredentialResultRemoveByKeyId(existing, id.KeyId)
	if _, err = client.UpdatePasswordCredentials(ctx, id.ObjectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return fmt.Errorf("Error removing Application Credentials %q from Application Object ID %q: %+v", id.KeyId, id.ObjectId, err)
	}

	return nil
}
