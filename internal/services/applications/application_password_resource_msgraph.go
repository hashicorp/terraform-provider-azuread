package applications

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
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func applicationPasswordResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).Applications.MsClient
	objectId := d.Get("application_object_id").(string)

	if val, ok := d.GetOk("key_id"); ok && val.(string) != "" {
		return tf.ErrorDiagPathF(fmt.Errorf("`key_id` is a read-only field when using Microsoft Graph. Please remove the `key_id` field from your configuration"), "key_id", "Creating application password")
	}

	if val, ok := d.GetOk("value"); ok && val.(string) != "" {
		return tf.ErrorDiagPathF(fmt.Errorf("`value` is a read-only field when using Microsoft Graph. Please remove the `value` field from your configuration"), "value", "Creating application password")
	}

	credential, err := helpers.PasswordCredentialForResource(d)
	if err != nil {
		attr := ""
		if kerr, ok := err.(helpers.CredentialError); ok {
			attr = kerr.Attr()
		}
		return tf.ErrorDiagPathF(err, attr, "Generating password credentials for application with object ID %q", objectId)
	}
	if credential == nil {
		return tf.ErrorDiagF(errors.New("nil credential was returned"), "Generating password credentials for application with object ID %q", objectId)
	}

	tf.LockByName(applicationResourceName, objectId)
	defer tf.UnlockByName(applicationResourceName, objectId)

	app, status, err := client.Get(ctx, objectId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "application_object_id", "Application with object ID %q was not found", objectId)
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving application with object ID %q", objectId)
	}
	if app == nil || app.ID == nil {
		return tf.ErrorDiagF(errors.New("nil application or application with nil ID was returned"), "API error retrieving application with object ID %q", objectId)
	}

	newCredential, _, err := client.AddPassword(ctx, *app.ID, *credential)
	if err != nil {
		return tf.ErrorDiagF(err, "Adding password for application with object ID %q", *app.ID)
	}
	if newCredential == nil {
		return tf.ErrorDiagF(errors.New("nil credential received when adding password"), "API error adding password for application with object ID %q", *app.ID)
	}
	if newCredential.KeyId == nil {
		return tf.ErrorDiagF(errors.New("nil or empty keyId received"), "API error adding password for application with object ID %q", *app.ID)
	}
	if newCredential.SecretText == nil || len(*newCredential.SecretText) == 0 {
		return tf.ErrorDiagF(errors.New("nil or empty password received"), "API error adding password for application with object ID %q", *app.ID)
	}

	id := parse.NewCredentialID(*app.ID, "password", *newCredential.KeyId)
	d.SetId(id.String())
	d.Set("value", newCredential.SecretText)

	return applicationPasswordResourceReadMsGraph(ctx, d, meta)
}

func applicationPasswordResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).Applications.MsClient

	id, err := parse.PasswordID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing password credential with ID %q", d.Id())
	}

	app, status, err := client.Get(ctx, id.ObjectId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Application with ID %q for %s credential %q was not found - removing from state!", id.ObjectId, id.KeyType, id.KeyId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagPathF(err, "application_object_id", "Retrieving application with object ID %q", id.ObjectId)
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

	tf.Set(d, "application_object_id", id.ObjectId)
	tf.Set(d, "description", credential.DisplayName)
	tf.Set(d, "display_name", credential.DisplayName)
	tf.Set(d, "key_id", id.KeyId)

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

func applicationPasswordResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { //nolint
	client := meta.(*clients.Client).Applications.MsClient

	id, err := parse.PasswordID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing password credential with ID %q", d.Id())
	}

	tf.LockByName(applicationResourceName, id.ObjectId)
	defer tf.UnlockByName(applicationResourceName, id.ObjectId)

	if _, err := client.RemovePassword(ctx, id.ObjectId, id.KeyId); err != nil {
		return tf.ErrorDiagF(err, "Removing password credential %q from application with object ID %q", id.KeyId, id.ObjectId)
	}

	return nil
}
