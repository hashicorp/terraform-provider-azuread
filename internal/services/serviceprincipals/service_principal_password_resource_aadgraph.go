package serviceprincipals

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func servicePrincipalPasswordResourceCreateAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.AadClient
	objectId := d.Get("service_principal_id").(string)

	cred, err := aadgraph.PasswordCredentialForResource(d)
	if err != nil {
		attr := ""
		if kerr, ok := err.(aadgraph.CredentialError); ok {
			attr = kerr.Attr()
		}
		return tf.ErrorDiagPathF(err, attr, "Generating password credentials for service principal with object ID %q", objectId)
	}
	id := parse.NewCredentialID(objectId, "password", *cred.KeyID)

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	existingCreds, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "service_principal_id", "Listing password credentials for service principal with ID %q", objectId)
	}

	newCreds, err := aadgraph.PasswordCredentialResultAdd(existingCreds, cred)
	if err != nil {
		if _, ok := err.(*aadgraph.AlreadyExistsError); ok {
			return tf.ImportAsExistsDiag("azuread_service_principal_password", id.String())
		}
		return tf.ErrorDiagF(err, "Adding service principal password")
	}

	if _, err = client.UpdatePasswordCredentials(ctx, objectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiagF(err, "Creating password credentials %q for service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	d.SetId(id.String())

	_, err = aadgraph.WaitForPasswordCredentialReplication(ctx, id.KeyId, d.Timeout(schema.TimeoutCreate), func() (graphrbac.PasswordCredentialListResult, error) {
		return client.ListPasswordCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for password credential replication for service principal (ObjectID %q, KeyID %q)", id.ObjectId, id.KeyId)
	}

	return servicePrincipalPasswordResourceReadAadGraph(ctx, d, meta)
}

func servicePrincipalPasswordResourceReadAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.AadClient

	id, err := parse.PasswordID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing password credential with ID %q", d.Id())
	}

	// ensure the parent Service Principal exists
	servicePrincipal, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Service Principal has been removed - skip it
		if utils.ResponseWasNotFound(servicePrincipal.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", id.ObjectId)
	}

	credentials, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "service_principal_id", "Listing password credentials for service principal with object ID %q", id.ObjectId)
	}

	credential := aadgraph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Service Principal %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "service_principal_id", id.ObjectId)
	tf.Set(d, "key_id", id.KeyId)

	description := ""
	if v := credential.CustomKeyIdentifier; v != nil {
		description = string(*v)
	}
	tf.Set(d, "description", description)
	tf.Set(d, "display_name", description)

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

func servicePrincipalPasswordResourceDeleteAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.AadClient

	id, err := parse.PasswordID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing password credential with ID %q", d.Id())
	}

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	// HACK: We can't yet move this resource to MS Graph (see comments in application_password_resource.go
	// Since AAD Graph lags behind reality, this hack waits for the AAD Graph API to see
	// and return the application before attempting to manage its passwords.
	_, err = aadgraph.WaitForCreationReplication(ctx, 5*time.Minute, func() (interface{}, error) {
		return client.Get(ctx, id.ObjectId)
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Service Principal was not found with object ID: %q", id.ObjectId)
	}

	// ensure the parent Service Principal exists
	servicePrincipal, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Service Principal was removed - skip it
		if utils.ResponseWasNotFound(servicePrincipal.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", id.ObjectId)
	}

	existing, err := client.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "service_principal_id", "Listing password credentials for service principal with object ID %q", id.ObjectId)
	}

	newCreds, err := aadgraph.PasswordCredentialResultRemoveByKeyId(existing, id.KeyId)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing password credential %q from service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	if _, err = client.UpdatePasswordCredentials(ctx, id.ObjectId, graphrbac.PasswordCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiagF(err, "Removing password credential %q from service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}
