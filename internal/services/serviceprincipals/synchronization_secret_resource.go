// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
)

func synchronizationSecretResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: synchronizationSecretResourceCreate,
		ReadContext:   synchronizationSecretResourceRead,
		UpdateContext: synchronizationSecretResourceUpdate,
		DeleteContext: synchronizationSecretResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(3 * time.Minute),
			Delete: schema.DefaultTimeout(3 * time.Minute),
		},

		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"service_principal_id": {
				Description:      "The object ID of the service principal for which this synchronization secret should be created",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},
			"credential": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "Name for this key-value pair.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "Value for this key-value pair.",
							Type:        schema.TypeString,
							Required:    true,
							Sensitive:   true,
						},
					},
				},
			},
		},
	}
}

func synchronizationSecretResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.SynchronizationJobClient
	spClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	objectId := d.Get("service_principal_id").(string)

	tf.LockByName(servicePrincipalResourceName, objectId)
	defer tf.UnlockByName(servicePrincipalResourceName, objectId)

	servicePrincipal, status, err := spClient.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "service_principal_id", "Service principal with object ID %q was not found", objectId)
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", objectId)
	}
	if servicePrincipal == nil || servicePrincipal.ID() == nil {
		return tf.ErrorDiagF(errors.New("nil service principal or service principal with nil ID was returned"), "API error retrieving service principal with object ID %q", objectId)
	}

	synchronizationSecret := msgraph.SynchronizationSecret{
		Credentials: expandSynchronizationSecretKeyStringValuePair(d.Get("credential").([]interface{})),
	}

	_, err = client.SetSecrets(ctx, synchronizationSecret, *servicePrincipal.ID())
	if err != nil {
		return tf.ErrorDiagF(err, "Creating synchronization secret for service principal ID %q", *servicePrincipal.ID())
	}
	id := parse.NewSynchronizationSecretID(*servicePrincipal.ID())

	// Wait for the secret to appear
	timeout, _ := ctx.Deadline()
	_, err = (&resource.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			newSynchronizationSecret, _, err := client.GetSecrets(ctx, id.ServicePrincipalId)
			if err != nil {
				return nil, "Error", fmt.Errorf("retrieving synchronization secret")
			}
			if newSynchronizationSecret != nil {
				if len(*synchronizationSecret.Credentials) == len(*newSynchronizationSecret.Credentials) {
					return "stub", "Done", nil
				}
				return "stub", "Waiting", nil
			} else {
				return "stub", "Waiting", nil
			}
		},
	}).WaitForStateContext(ctx)

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for synchronization secret %q", id.ServicePrincipalId)
	}

	if d.IsNewResource() {
		d.SetId(id.String())
	}
	tf.Set(d, "credential", flattenSynchronizationSecretKeyStringValuePair(synchronizationSecret.Credentials, nil))

	return synchronizationSecretResourceRead(ctx, d, meta)
}

func synchronizationSecretResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Update is same as create
	return synchronizationSecretResourceCreate(ctx, d, meta)
}

func synchronizationSecretResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.SynchronizationJobClient

	id, err := parse.SynchronizationSecretID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization secret with ID %q", d.Id())
	}

	secrets, status, err := client.GetSecrets(ctx, id.ServicePrincipalId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Synchronization secrets for service principal %q was not found - removing from state!", id.ServicePrincipalId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving synchronization secrets for service principal %q", id.ServicePrincipalId)
	}
	tf.Set(d, "credential", flattenSynchronizationSecretKeyStringValuePair(secrets.Credentials, d.Get("credential").([]interface{})))

	return nil
}

func synchronizationSecretResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.SynchronizationJobClient
	spClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	id, err := parse.SynchronizationSecretID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization secret with ID %q", d.Id())
	}

	tf.LockByName(servicePrincipalResourceName, id.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ServicePrincipalId)

	servicePrincipal, status, err := spClient.Get(ctx, id.ServicePrincipalId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "service_principal_id", "Service principal with object ID %q was not found", id.ServicePrincipalId)
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving service principal with object ID %q", id.ServicePrincipalId)
	}
	if servicePrincipal == nil || servicePrincipal.ID() == nil {
		return tf.ErrorDiagF(errors.New("nil service principal or service principal with nil ID was returned"), "API error retrieving service principal with object ID %q", id.ServicePrincipalId)
	}

	// We delete secrets by setting values to empty strings
	credentials := emptySynchronizationSecretKeyStringValuePair(d.Get("credential").([]interface{}))

	synchronizationSecret := msgraph.SynchronizationSecret{
		Credentials: credentials,
	}
	if _, err := client.SetSecrets(ctx, synchronizationSecret, id.ServicePrincipalId); err != nil {
		return tf.ErrorDiagF(err, "Removing synchronization secrets for service principal with object ID %q", id.ServicePrincipalId)
	}

	// Wait for synchronization secret to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true

		synchronizationSecrets, _, _ := client.GetSecrets(ctx, id.ServicePrincipalId)

		// Test if credentials are removed
		if allCredentialsRemoved(*credentials, *synchronizationSecrets.Credentials) {
			return utils.Bool(false), nil
		}

		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of synchronization secrets from service principal with object ID %q", id.ServicePrincipalId)
	}

	return nil
}
