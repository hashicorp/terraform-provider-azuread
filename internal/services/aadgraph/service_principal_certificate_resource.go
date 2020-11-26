package aadgraph

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
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
		di := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Generating certificate credentials for service principal with object ID %q", objectId),
			Detail:   err.Error(),
		}
		if kerr, ok := err.(graph.CredentialError); ok {
			di.AttributePath = cty.Path{cty.GetAttrStep{Name: kerr.Attr()}}
		}
		return diag.Diagnostics{di}
	}

	id := graph.CredentialIdFrom(objectId, "certificate", *cred.KeyID)

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	existingCreds, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Listing certificate credentials for service principal with ID %q", objectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "application_object_id"}},
		}}
	}

	newCreds, err := graph.KeyCredentialResultAdd(existingCreds, cred)
	if err != nil {
		if _, ok := err.(*graph.AlreadyExistsError); ok {
			return tf.ImportAsExistsDiag("azuread_service_principal_certificate", id.String())
		}
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Adding service principal certificate",
			Detail:   err.Error(),
		}}
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Creating certificate credentials %q for service principal with object ID %q", id.KeyId, id.ObjectId),
			Detail:   err.Error(),
		}}
	}

	_, err = graph.WaitForKeyCredentialReplication(id.KeyId, d.Timeout(schema.TimeoutCreate), func() (graphrbac.KeyCredentialListResult, error) {
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

	return servicePrincipalCertificateResourceRead(ctx, d, meta)
}

func servicePrincipalCertificateResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	id, err := graph.ParseCertificateId(d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Parsing certificate credential with ID %q", d.Id()),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
		}}
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
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Retrieving service principal with ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	credentials, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Listing certificate credentials for service principal with object ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	credential := graph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
	if credential == nil {
		log.Printf("[DEBUG] certificate credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	d.Set("service_principal_id", id.ObjectId)
	d.Set("key_id", id.KeyId)

	if keyType := credential.Type; keyType != nil {
		d.Set("type", keyType)
	}

	if endDate := credential.EndDate; endDate != nil {
		d.Set("end_date", endDate.Format(time.RFC3339))
	}

	if startDate := credential.StartDate; startDate != nil {
		d.Set("start_date", startDate.Format(time.RFC3339))
	}

	return nil
}

func servicePrincipalCertificateResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient).AadGraph.ServicePrincipalsClient

	id, err := graph.ParseCertificateId(d.Id())
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Parsing certificate credential with ID %q", d.Id()),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "id"}},
		}}
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
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Retrieving service principal with ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	existing, err := client.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Listing certificate credentials for service principal with object ID %q", id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	newCreds, err := graph.KeyCredentialResultRemoveByKeyId(existing, id.KeyId)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Removing certificate credential %q from service principal with object ID %q", id.KeyId, id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	if _, err = client.UpdateKeyCredentials(ctx, id.ObjectId, graphrbac.KeyCredentialsUpdateParameters{Value: newCreds}); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       fmt.Sprintf("Removing certificate credential %q from service principal with object ID %q", id.KeyId, id.ObjectId),
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "name"}},
		}}
	}

	return nil
}
