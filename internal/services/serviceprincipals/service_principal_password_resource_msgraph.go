package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	helpers "github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func servicePrincipalPasswordResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).ServicePrincipals.MsClient
	objectId := d.Get("service_principal_id").(string)

	if val, ok := d.GetOk("display_name"); ok && val.(string) != "" {
		return tf.ErrorDiagPathF(fmt.Errorf("`display_name` is a read-only field when using Microsoft Graph. Please remove the `display_name` field from your configuration"), "display_name", "Creating service principal password")
	}

	if val, ok := d.GetOk("end_date"); ok && val.(string) != "" {
		return tf.ErrorDiagPathF(fmt.Errorf("`end_date` is a read-only field when using Microsoft Graph. Please remove the `end_date` field from your configuration"), "end_date", "Creating service principal password")
	}

	if val, ok := d.GetOk("end_date_relative"); ok && val.(string) != "" {
		return tf.ErrorDiagPathF(fmt.Errorf("`end_date_relative` is a read-only field when using Microsoft Graph. Please remove the `end_date_relative` field from your configuration"), "end_date_relative", "Creating service principal password")
	}

	if val, ok := d.GetOk("key_id"); ok && val.(string) != "" {
		return tf.ErrorDiagPathF(fmt.Errorf("`key_id` is a read-only field when using Microsoft Graph. Please remove the `key_id` field from your configuration"), "key_id", "Creating service principal password")
	}

	if val, ok := d.GetOk("start_date"); ok && val.(string) != "" {
		return tf.ErrorDiagPathF(fmt.Errorf("`start_date` is a read-only field when using Microsoft Graph. Please remove the `start_date` field from your configuration"), "start_date", "Creating service principal password")
	}

	if val, ok := d.GetOk("value"); ok && val.(string) != "" {
		return tf.ErrorDiagPathF(fmt.Errorf("`value` is a read-only field when using Microsoft Graph. Please remove the `value` field from your configuration"), "value", "Creating service principal password")
	}

	credential, err := helpers.PasswordCredentialForResource(d)
	if err != nil {
		attr := ""
		if kerr, ok := err.(helpers.CredentialError); ok {
			attr = kerr.Attr()
		}
		return tf.ErrorDiagPathF(err, attr, "Generating password credentials for service principal with object ID %q", objectId)
	}
	if credential == nil {
		return tf.ErrorDiagF(errors.New("nil credential was returned"), "Generating password credentials for service principal with object ID %q", objectId)
	}

	tf.LockByName(servicePrincipalResourceName, objectId)
	defer tf.UnlockByName(servicePrincipalResourceName, objectId)

	sp, status, err := client.Get(ctx, objectId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "service_principal_id", "Service principal with object ID %q was not found", objectId)
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", objectId)
	}
	if sp == nil || sp.ID == nil {
		return tf.ErrorDiagF(errors.New("nil service principal or service principal with nil ID was returned"), "API error retrieving service principal with object ID %q", objectId)
	}

	newCredential, _, err := client.AddPassword(ctx, *sp.ID, *credential)
	if err != nil {
		return tf.ErrorDiagF(err, "Adding password for service principal with object ID %q", *sp.ID)
	}
	if newCredential == nil {
		return tf.ErrorDiagF(errors.New("nil credential received when adding password"), "API error adding password for service principal with object ID %q", *sp.ID)
	}
	if newCredential.KeyId == nil {
		return tf.ErrorDiagF(errors.New("nil or empty keyId received"), "API error adding password for service principal with object ID %q", *sp.ID)
	}
	if newCredential.SecretText == nil || len(*newCredential.SecretText) == 0 {
		return tf.ErrorDiagF(errors.New("nil or empty password received"), "API error adding password for service principal with object ID %q", *sp.ID)
	}

	id := parse.NewCredentialID(*sp.ID, "password", *newCredential.KeyId)
	d.SetId(id.String())
	d.Set("value", newCredential.SecretText)

	return servicePrincipalPasswordResourceReadMsGraph(ctx, d, meta)
}

func servicePrincipalPasswordResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).ServicePrincipals.MsClient

	id, err := parse.PasswordID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing password credential with ID %q", d.Id())
	}

	app, status, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Service Principal with ID %q for %s credential %q was not found - removing from state!", id.ObjectId, id.KeyType, id.KeyId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", id.ObjectId)
	}

	var credential *msgraph.PasswordCredential
	if app.PasswordCredentials != nil {
		for _, cred := range *app.PasswordCredentials {
			if cred.KeyId != nil && *cred.KeyId == id.KeyId {
				credential = &cred
				break
			}
		}
	}

	if credential == nil {
		log.Printf("[DEBUG] Password credential %q (ID %q) was not found - removing from state!", id.KeyId, id.ObjectId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "description", credential.DisplayName)
	tf.Set(d, "display_name", credential.DisplayName)
	tf.Set(d, "key_id", id.KeyId)
	tf.Set(d, "service_principal_id", id.ObjectId)

	startDate := ""
	if v := credential.StartDateTime; v != nil {
		startDate = v.Format(time.RFC3339)
	}
	tf.Set(d, "start_date", startDate)

	endDate := ""
	if v := credential.EndDateTime; v != nil {
		endDate = v.Format(time.RFC3339)
	}
	tf.Set(d, "end_date", endDate)

	return nil
}

func servicePrincipalPasswordResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).ServicePrincipals.MsClient

	id, err := parse.PasswordID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing password credential with ID %q", d.Id())
	}

	tf.LockByName(servicePrincipalResourceName, id.ObjectId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ObjectId)

	if _, err := client.RemovePassword(ctx, id.ObjectId, id.KeyId); err != nil {
		return tf.ErrorDiagF(err, "Removing password credential %q from service principal with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}
