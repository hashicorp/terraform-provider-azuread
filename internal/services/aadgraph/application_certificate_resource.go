package aadgraph

import (
	"context"
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

func applicationCertificateResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationCertificateResourceCreate,
		ReadContext:   applicationCertificateResourceRead,
		DeleteContext: applicationCertificateResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := graph.ParseCertificateId(id)
			return err
		}),

		Schema: graph.CertificateResourceSchema("application_object_id"),
	}
}

func applicationCertificateResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	objectId := d.Get("application_object_id").(string)

	cred, err := graph.KeyCredentialForResource(d)
	if err != nil {
		attr := ""
		if kerr, ok := err.(graph.CredentialError); ok {
			attr = kerr.Attr()
		}
		return tf.ErrorDiagPathF(err, attr, "Generating certificate credentials for application with object ID %q", objectId)
	}

	id := graph.CredentialIdFrom(objectId, "certificate", *cred.KeyID)

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	existingCreds, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_object_id", "Listing certificate credentials for application with ID %q", objectId)
	}

	newCreds, err := graph.KeyCredentialResultAdd(existingCreds, cred)
	if err != nil {
		if _, ok := err.(*graph.AlreadyExistsError); ok {
			return tf.ImportAsExistsDiag("azuread_application_certificate", id.String())
		}
		return tf.ErrorDiagF(err, "Adding application certificate")
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiagF(err, "Creating certificate credentials %q for application with object ID %q", id.KeyId, id.ObjectId)
	}

	_, err = graph.WaitForKeyCredentialReplication(ctx, id.KeyId, d.Timeout(schema.TimeoutCreate), func() (graphrbac.KeyCredentialListResult, error) {
		return client.ListKeyCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for certificate credential replication for application (AppID %q, KeyID %q)", id.ObjectId, id.KeyId)
	}

	d.SetId(id.String())

	return applicationCertificateResourceRead(ctx, d, meta)
}

func applicationCertificateResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseCertificateId(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
	}

	// ensure the Application Object exists
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

	credentials, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "application_object_id", "Listing certificate credentials for application with object ID %q", id.ObjectId)
	}

	credential := graph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	if dg := tf.Set(d, "application_object_id", id.ObjectId); dg != nil {
		return dg
	}

	if dg := tf.Set(d, "key_id", id.KeyId); dg != nil {
		return dg
	}

	keyType := ""
	if v := credential.Type; v != nil {
		keyType = *v
	}
	if dg := tf.Set(d, "type", keyType); dg != nil {
		return dg
	}

	startDate := ""
	if v := credential.StartDate; v != nil {
		startDate = v.Format(time.RFC3339)
	}
	if dg := tf.Set(d, "start_date", startDate); dg != nil {
		return dg
	}

	endDate := ""
	if v := credential.EndDate; v != nil {
		endDate = v.Format(time.RFC3339)
	}
	if dg := tf.Set(d, "end_date", endDate); dg != nil {
		return dg
	}

	return nil
}

func applicationCertificateResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseCertificateId(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing certificate credential with ID %q", d.Id())
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
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving application with object ID %q", id.ObjectId)
	}

	existing, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return tf.ErrorDiagF(err, "Listing certificate credential for application with object ID %q", id.ObjectId)
	}

	newCreds, err := graph.KeyCredentialResultRemoveByKeyId(existing, id.KeyId)
	if err != nil {
		return tf.ErrorDiagF(err, "Removing certificate credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return tf.ErrorDiagF(err, "Removing certificate credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}
