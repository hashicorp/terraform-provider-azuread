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

func ResourceServicePrincipalPassword() *schema.Resource {
	return &schema.Resource{
		Create: resourceServicePrincipalPasswordCreate,
		Read:   resourceServicePrincipalPasswordRead,
		Delete: resourceServicePrincipalPasswordDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: graph.PasswordResourceSchema("service_principal"),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceServicePrincipalPasswordInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceServicePrincipalPasswordInstanceStateUpgradeV0,
				Version: 0,
			},
		},
	}
}

func resourceServicePrincipalPasswordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).ServicePrincipalsClient
	ctx := meta.(*clients.AadClient).StopContext

	objectId := d.Get("service_principal_id").(string)

	cred, err := graph.PasswordCredentialForResource(d)
	if err != nil {
		return fmt.Errorf("Error generating Service Principal Credentials for Object ID %q: %+v", objectId, err)
	}
	id := graph.CredentialIdFrom(objectId, "password", *cred.KeyID)

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	existingCreds, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Service Principal %q: %+v", id.ObjectId, err)
	}

	newCreds, err := graph.PasswordCredentialResultAdd(existingCreds, cred, true)
	if err != nil {
		return tf.ImportAsExistsError("azuread_service_principal_password", id.String())
	}

	if _, err = client.UpdatePasswordCredentials(ctx, objectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return fmt.Errorf("Error creating Password Credential %q for Service Principal %q: %+v", id.KeyId, id.ObjectId, err)
	}

	d.SetId(id.String())

	_, err = graph.WaitForPasswordCredentialReplication(id.KeyId, func() (graphrbac.PasswordCredentialListResult, error) {
		return client.ListPasswordCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return fmt.Errorf("Error waiting for Service Principal Password replication (SP %q, KeyID %q: %+v", id.ObjectId, id.KeyId, err)
	}

	return resourceServicePrincipalPasswordRead(d, meta)
}

func resourceServicePrincipalPasswordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).ServicePrincipalsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseCredentialId(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Application Password ID: %v", err)
	}

	// ensure the parent Service Principal exists
	servicePrincipal, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Service Principal has been removed - skip it
		if ar.ResponseWasNotFound(servicePrincipal.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error retrieving Service Principal ID %q: %+v", id.ObjectId, err)
	}

	credentials, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Service Principal with Object ID %q: %+v", id.ObjectId, err)
	}

	credential := graph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Service Principal %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	// value is available in the SDK but isn't returned from the API
	d.Set("key_id", credential.KeyID)
	d.Set("service_principal_id", id.ObjectId)

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

func resourceServicePrincipalPasswordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).ServicePrincipalsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseCredentialId(d.Id())
	if err != nil {
		return fmt.Errorf("Error parsing Application Password ID: %v", err)
	}

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	// ensure the parent Service Principal exists
	servicePrincipal, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Service Principal was removed - skip it
		if ar.ResponseWasNotFound(servicePrincipal.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return fmt.Errorf("Error retrieving Service Principal ID %q: %+v", id.ObjectId, err)
	}

	existing, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return fmt.Errorf("Error Listing Password Credentials for Service Principal with Object ID %q: %+v", id.ObjectId, err)
	}

	newCreds := graph.PasswordCredentialResultRemoveByKeyId(existing, id.KeyId)
	if _, err = client.UpdatePasswordCredentials(ctx, id.ObjectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return fmt.Errorf("Error removing Password %q from Service Principal %q: %+v", id.KeyId, id.ObjectId, err)
	}

	return nil
}

func resourceServicePrincipalPasswordInstanceResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: graph.PasswordResourceSchema("service_principal"),
	}
}

func resourceServicePrincipalPasswordInstanceStateUpgradeV0(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	newId, err := graph.ParseOldCredentialId(rawState["id"].(string), "password")
	if err != nil {
		return rawState, fmt.Errorf("generating new ID: %s", err)
	}

	rawState["id"] = newId.String()
	return rawState, nil
}
