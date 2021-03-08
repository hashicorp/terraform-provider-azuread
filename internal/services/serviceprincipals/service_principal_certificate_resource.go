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

func servicePrincipalCertificateResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalCertificateResourceCreate,
		ReadContext:   servicePrincipalCertificateResourceRead,
		DeleteContext: servicePrincipalCertificateResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.CertificateID(id)
			return err
		}),

		Schema: aadgraph.CertificateResourceSchema("service_principal_id"),
	}
}

func servicePrincipalCertificateResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.AadClient

	objectId := d.Get("service_principal_id").(string)

	cred, err := aadgraph.KeyCredentialForResource(d)
	if err != nil {
		attr := ""
		if kerr, ok := err.(aadgraph.CredentialError); ok {
			attr = kerr.Attr()
		}
		return tf.ErrorDiagPathF(err, attr, "Generating certificate credentials for service principal with object ID %q", objectId)
	}

	id := parse.NewCredentialID(objectId, "certificate", *cred.KeyID)

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	existingCreds, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "service_principal_id", "Listing certificate credentials for service principal with ID %q", objectId)
	}

	newCreds, err := aadgraph.KeyCredentialResultAdd(existingCreds, cred)
	if err != nil {
		if _, ok := err.(*aadgraph.AlreadyExistsError); ok {
			return tf.ImportAsExistsDiag("azuread_service_principal_certificate", id.String())
		}
		return tf.ErrorDiagF(err, "Adding service principal certificate")
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiagF(err, "Creating certificate credentials %q for service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	_, err = aadgraph.WaitForKeyCredentialReplication(ctx, id.KeyId, d.Timeout(schema.TimeoutCreate), func() (graphrbac.KeyCredentialListResult, error) {
		return client.ListKeyCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for certificate credential replication for service principal (ObjectID %q, KeyID %q)", id.ObjectId, id.KeyId)
	}

	d.SetId(id.String())

	return servicePrincipalCertificateResourceRead(ctx, d, meta)
}

func servicePrincipalCertificateResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.AadClient

	id, err := parse.CertificateID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	// ensure the Service Principal Object exists
	sp, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Service Principal has been removed - skip it
		if utils.ResponseWasNotFound(sp.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", id.ObjectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", id.ObjectId)
	}

	credentials, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "service_principal_id", "Listing certificate credentials for service principal with object ID %q", id.ObjectId)
	}

	credential := aadgraph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "service_principal_id", id.ObjectId)
	tf.Set(d, "key_id", id.KeyId)

	keyType := ""
	if v := credential.Type; v != nil {
		keyType = *v
	}
	tf.Set(d, "type", keyType)

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

func servicePrincipalCertificateResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.AadClient

	id, err := parse.CertificateID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	// ensure the parent Service Principal exists
	sp, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		// the parent Service Principal has been removed - skip it
		if utils.ResponseWasNotFound(sp.Response) {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing from state!", id.ObjectId)
			return nil
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", id.ObjectId)
	}

	existing, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagF(err, "Listing certificate credentials for service principal with object ID %q", id.ObjectId)
	}

	newCreds, err := aadgraph.KeyCredentialResultRemoveByKeyId(existing, id.KeyId)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing certificate credential %q from service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiagF(err, "Removing certificate credential %q from service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}
