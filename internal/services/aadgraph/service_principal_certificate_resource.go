package aadgraph

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

func servicePrincipalCertificateResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalCertificateResourceCreate,
		ReadContext:   servicePrincipalCertificateResourceRead,
		DeleteContext: servicePrincipalCertificateResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := graph.ParseCertificateId(id)
			return err
		}),

		Schema: graph.CertificateResourceSchema("service_principal_id"),
	}
}

func servicePrincipalCertificateResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	objectId := d.Get("service_principal_id").(string)

	cred, err := graph.KeyCredentialForResource(d)
	if err != nil {
		attr := ""
		if kerr, ok := err.(graph.CredentialError); ok {
			attr = kerr.Attr()
		}
		return tf.ErrorDiag(fmt.Sprintf("Generating certificate credentials for service principal with object ID %q", objectId), err.Error(), attr)
	}

	id := graph.CredentialIdFrom(objectId, "certificate", *cred.KeyID)

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	existingCreds, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Listing certificate credentials for service principal with ID %q", objectId), err.Error(), "application_object_id")
	}

	newCreds, err := graph.KeyCredentialResultAdd(existingCreds, cred)
	if err != nil {
		if _, ok := err.(*graph.AlreadyExistsError); ok {
			return tf.ImportAsExistsDiag("azuread_service_principal_certificate", id.String())
		}
		return tf.ErrorDiag("Adding service principal certificate", err.Error(), "")
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Creating certificate credentials %q for service principal with object ID %q", id.KeyId, id.ObjectId), err.Error(), "")
	}

	_, err = graph.WaitForKeyCredentialReplication(ctx, id.KeyId, d.Timeout(schema.TimeoutCreate), func() (graphrbac.KeyCredentialListResult, error) {
		return client.ListKeyCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Waiting for certificate credential replication for application (AppID %q, KeyID %q)", id.ObjectId, id.KeyId), err.Error(), "")
	}

	d.SetId(id.String())

	return servicePrincipalCertificateResourceRead(ctx, d, meta)
}

func servicePrincipalCertificateResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	id, err := graph.ParseCertificateId(d.Id())
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Parsing certificate credential with ID %q", d.Id()), err.Error(), "id")
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
		return tf.ErrorDiag(fmt.Sprintf("Retrieving service principal with ID %q", id.ObjectId), err.Error(), "name")
	}

	credentials, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Listing certificate credentials for service principal with object ID %q", id.ObjectId), err.Error(), "name")
	}

	credential := graph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	if err := d.Set("service_principal_id", id.ObjectId); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "service_principal_id")
	}

	if err := d.Set("key_id", id.KeyId); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "key_id")
	}

	keyType := ""
	if v := credential.Type; v != nil {
		keyType = *v
	}
	if err := d.Set("type", keyType); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "type")
	}

	startDate := ""
	if v := credential.StartDate; v != nil {
		startDate = v.Format(time.RFC3339)
	}
	if err := d.Set("start_date", startDate); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "start_date")
	}

	endDate := ""
	if v := credential.EndDate; v != nil {
		endDate = v.Format(time.RFC3339)
	}
	if err := d.Set("end_date", endDate); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "end_date")
	}

	return nil
}

func servicePrincipalCertificateResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	id, err := graph.ParseCertificateId(d.Id())
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Parsing certificate credential with ID %q", d.Id()), err.Error(), "id")
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
		return tf.ErrorDiag(fmt.Sprintf("Retrieving service principal with ID %q", id.ObjectId), err.Error(), "")
	}

	existing, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Listing certificate credentials for service principal with object ID %q", id.ObjectId), err.Error(), "")
	}

	newCreds, err := graph.KeyCredentialResultRemoveByKeyId(existing, id.KeyId)
	if err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Removing certificate credential %q from service principal with object ID %q", id.KeyId, id.ObjectId), err.Error(), "")
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiag(fmt.Sprintf("Removing certificate credential %q from service principal with object ID %q", id.KeyId, id.ObjectId), err.Error(), "")
	}

	return nil
}
