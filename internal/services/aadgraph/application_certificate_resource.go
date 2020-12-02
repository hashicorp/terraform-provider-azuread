package aadgraph

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/go-cty/cty"
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
		di := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Generating certificate credentials for application with object ID %q", objectId),
			Detail:   err.Error(),
		}
		if kerr, ok := err.(graph.CredentialError); ok {
			di.AttributePath = cty.Path{cty.GetAttrStep{Name: kerr.Attr()}}
		}
		return diag.Diagnostics{di}
	}

	id := graph.CredentialIdFrom(objectId, "certificate", *cred.KeyID)

	tf.LockByName(resourceApplicationName, id.ObjectId)
	defer tf.UnlockByName(resourceApplicationName, id.ObjectId)

	existingCreds, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Listing certificate credentials for application with ID %q", objectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "application_object_id"}},
		}}
	}

	newCreds, err := graph.KeyCredentialResultAdd(existingCreds, cred)
	if err != nil {
		if _, ok := err.(*graph.AlreadyExistsError); ok {
			return tf.ImportAsExistsDiag("azuread_application_certificate", id.String())
		}
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Adding application certificate",
			Detail:   err.Error(),
		}}
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Creating certificate credentials %q for application with object ID %q", id.KeyId, id.ObjectId),
			Detail:   err.Error(),
		}}
	}

	_, err = graph.WaitForKeyCredentialReplication(ctx, id.KeyId, d.Timeout(schema.TimeoutCreate), func() (graphrbac.KeyCredentialListResult, error) {
		return client.ListKeyCredentials(ctx, id.ObjectId)
	})
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Waiting for certificate credential replication for application (AppID %q, KeyID %q)", id.ObjectId, id.KeyId),
			Detail:   err.Error(),
		}}
	}

	d.SetId(id.String())

	return applicationCertificateResourceRead(ctx, d, meta)
}

func applicationCertificateResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseCertificateId(d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Parsing certificate credential with ID %q", d.Id()),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
		}}
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
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Retrieving application with ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	credentials, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Listing certificate credentials for application with object ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	credential := graph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] Certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	if err := d.Set("application_object_id", id.ObjectId); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "application_object_id"}},
		}}
	}

	if err := d.Set("key_id", id.KeyId); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "key_id"}},
		}}
	}

	keyType := ""
	if v := credential.Type; v != nil {
		keyType = *v
	}
	if err := d.Set("type", keyType); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "type"}},
		}}
	}

	startDate := ""
	if v := credential.StartDate; v != nil {
		startDate = v.Format(time.RFC3339)
	}
	if err := d.Set("start_date", startDate); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "start_date"}},
		}}
	}

	endDate := ""
	if v := credential.EndDate; v != nil {
		endDate = v.Format(time.RFC3339)
	}
	if err := d.Set("end_date", endDate); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "end_date"}},
		}}
	}

	return nil
}

func applicationCertificateResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ApplicationsClient

	id, err := graph.ParseCertificateId(d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Parsing certificate credential with ID %q", d.Id()),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
		}}
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
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Retrieving application with ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	existing, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Listing certificate credential for application with object ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	newCreds, err := graph.KeyCredentialResultRemoveByKeyId(existing, id.KeyId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Removing certificate credential %q from application with object ID %q", id.KeyId, id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Removing certificate credential %q from application with object ID %q", id.KeyId, id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	return nil
}
