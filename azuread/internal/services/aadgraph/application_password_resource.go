package aadgraph

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/clients"
)

func ApplicationPasswordResource() *schema.Resource {
	return &schema.Resource{
		Create: applicationPasswordResourceCreate,
		Read:   applicationPasswordResourceRead,
		Delete: applicationPasswordResourceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: graph.PasswordResourceSchema("application"),

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

func applicationPasswordResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Get("application_object_id").(string)

	cred, err := graph.PasswordCredentialForResource(d)
	if err != nil {
		return fmt.Errorf("Error generating Application Credentials for Object ID %q: %+v", objectId, err)
	}
	id := graph.CredentialIdFrom(objectId, "password", *cred.KeyID)

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	existingCreds, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("Error Listing Application Credentials for Object ID %q: %+v", id.ObjectId, err)
	}

	newCreds, err := graph.PasswordCredentialResultAdd(existingCreds, cred, true)
	if err != nil {
		return tf.ImportAsExistsError("azuread_application_password", id.String())
	}

	if _, err = client.UpdatePasswordCredentials(ctx, id.ObjectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return fmt.Errorf("Error creating Application Credentials %q for Object ID %q: %+v", id.KeyId, id.ObjectId, err)
	}

	_, err = graph.WaitForPasswordCredentialReplication(id.KeyId, func() (graphrbac.PasswordCredentialListResult, error) {
		return client.ListPasswordCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return fmt.Errorf("Error waiting for Application Password replication (AppID %q, KeyID %q: %+v", id.ObjectId, id.KeyId, err)
	}

	d.SetId(id.String())

	return applicationPasswordResourceRead(d, meta)
}

func applicationPasswordResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseCredentialId(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Application Password ID: %v", err)
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
	d.Set("key_id", id.KeyId)

	if description := credential.CustomKeyIdentifier; description != nil {
		d.Set("description", string(*description))
	}

	if endDate := credential.EndDate; endDate != nil {
		d.Set("end_date", endDate.Format(time.RFC3339))
	}

	if startDate := credential.StartDate; startDate != nil {
		d.Set("start_date", startDate.Format(time.RFC3339))
	}

	return nil
}

func applicationPasswordResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseCredentialId(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Application Password ID: %v", err)
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

func resourceApplicationPasswordInstanceResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: graph.PasswordResourceSchema("application"),
	}
}

func resourceApplicationPasswordInstanceStateUpgradeV0(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	newId, err := graph.ParseOldCredentialId(rawState["id"].(string), "password")
	if err != nil {
		return rawState, fmt.Errorf("generating new ID: %s", err)
	}

	rawState["id"] = newId.String()
	return rawState, nil
}
