// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synchronization

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationsecret"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/synchronization/migrations"
)

func synchronizationSecretResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: synchronizationSecretResourceCreate,
		ReadContext:   synchronizationSecretResourceRead,
		UpdateContext: synchronizationSecretResourceUpdate,
		DeleteContext: synchronizationSecretResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, errs := stable.ValidateServicePrincipalID(id, "id"); len(errs) > 0 {
				out := ""
				for _, err := range errs {
					out += err.Error()
				}
				return errors.New(out)
			}
			return nil
		}),

		SchemaVersion: 1,
		StateUpgraders: []pluginsdk.StateUpgrader{
			{
				Type:    migrations.ResourceSynchronizationSecretInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceSynchronizationSecretInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*pluginsdk.Schema{
			"service_principal_id": {
				Description:  "The object ID of the service principal for which this synchronization secret should be created",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: stable.ValidateServicePrincipalID,
			},

			"credential": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"key": {
							Description: "Name for this key-value pair.",
							Type:        pluginsdk.TypeString,
							Required:    true,
						},
						"value": {
							Description: "Value for this key-value pair.",
							Type:        pluginsdk.TypeString,
							Required:    true,
							Sensitive:   true,
						},
					},
				},
			},
		},
	}
}

func synchronizationSecretResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationSecretClient

	id, err := stable.ParseServicePrincipalID(d.Get("service_principal_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "service_principal_id", "Parsing `service_principal_id`")
	}

	tf.LockByName(servicePrincipalResourceName, id.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ServicePrincipalId)

	synchronizationSecrets := synchronizationsecret.SetSynchronizationSecretRequest{
		Value: expandSynchronizationSecretKeyStringValuePair(d.Get("credential").([]interface{})),
	}

	if _, err := client.SetSynchronizationSecret(ctx, *id, synchronizationSecrets, synchronizationsecret.SetSynchronizationSecretOperationOptions{RetryFunc: synchronizationRetryFunc()}); err != nil {
		return tf.ErrorDiagF(err, "Creating synchronization secret for %s", id)
	}

	// Wait for the secret to appear
	if err := consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.ListSynchronizationSecrets(ctx, *id, synchronizationsecret.ListSynchronizationSecretsOperationOptions{RetryFunc: synchronizationRetryFunc()})
		if err != nil {
			return pointer.To(false), fmt.Errorf("retrieving synchronization secret")
		}
		newSynchronizationSecrets := resp.Model
		if newSynchronizationSecrets == nil {
			return pointer.To(false), nil
		}
		if len(pointer.From(synchronizationSecrets.Value)) == len(pointer.From(newSynchronizationSecrets)) {
			return pointer.To(true), nil
		}
		return pointer.To(false), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for synchronization secrets for %s", id)
	}

	d.SetId(id.ID())
	tf.Set(d, "credential", flattenSynchronizationSecretKeyStringValuePair(synchronizationSecrets.Value, nil))

	return synchronizationSecretResourceRead(ctx, d, meta)
}

func synchronizationSecretResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationSecretClient

	id, err := stable.ParseServicePrincipalID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization secret ID %q", d.Id())
	}

	tf.LockByName(servicePrincipalResourceName, id.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ServicePrincipalId)

	synchronizationSecrets := synchronizationsecret.SetSynchronizationSecretRequest{
		Value: expandSynchronizationSecretKeyStringValuePair(d.Get("credential").([]interface{})),
	}

	if _, err = client.SetSynchronizationSecret(ctx, *id, synchronizationSecrets, synchronizationsecret.SetSynchronizationSecretOperationOptions{RetryFunc: synchronizationRetryFunc()}); err != nil {
		return tf.ErrorDiagF(err, "Updating synchronization secret for %s", id)
	}

	// Wait for the secret to update
	if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.ListSynchronizationSecrets(ctx, *id, synchronizationsecret.ListSynchronizationSecretsOperationOptions{RetryFunc: synchronizationRetryFunc()})
		if err != nil {
			return pointer.To(false), fmt.Errorf("retrieving synchronization secret")
		}
		newSynchronizationSecrets := resp.Model
		if newSynchronizationSecrets == nil {
			return pointer.To(false), nil
		}
		if len(pointer.From(synchronizationSecrets.Value)) == len(pointer.From(newSynchronizationSecrets)) {
			return pointer.To(true), nil
		}
		return pointer.To(false), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for synchronization secrets for %s", id)
	}

	tf.Set(d, "credential", flattenSynchronizationSecretKeyStringValuePair(synchronizationSecrets.Value, nil))

	return synchronizationSecretResourceRead(ctx, d, meta)
}

func synchronizationSecretResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationSecretClient

	id, err := stable.ParseServicePrincipalID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization secret ID %q", d.Id())
	}

	resp, err := client.ListSynchronizationSecrets(ctx, *id, synchronizationsecret.ListSynchronizationSecretsOperationOptions{RetryFunc: synchronizationRetryFunc()})
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] Synchronization secrets for %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving synchronization secrets for %s", id)
	}

	synchronizationSecrets := resp.Model
	if synchronizationSecrets == nil {
		log.Printf("[DEBUG] Synchronization secrets for %s was nil - removing from state!", id)
		d.SetId("")
		return nil
	}

	tf.Set(d, "service_principal_id", id.ID())
	tf.Set(d, "credential", flattenSynchronizationSecretKeyStringValuePair(synchronizationSecrets, d.Get("credential").([]interface{})))

	return nil
}

func synchronizationSecretResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationSecretClient

	id, err := stable.ParseServicePrincipalID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization secret with ID %q", d.Id())
	}

	tf.LockByName(servicePrincipalResourceName, id.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ServicePrincipalId)

	// We delete secrets by setting values to empty strings
	credentials := emptySynchronizationSecretKeyStringValuePair(d.Get("credential").([]interface{}))

	synchronizationSecrets := synchronizationsecret.SetSynchronizationSecretRequest{
		Value: credentials,
	}
	if _, err = client.SetSynchronizationSecret(ctx, *id, synchronizationSecrets, synchronizationsecret.SetSynchronizationSecretOperationOptions{RetryFunc: synchronizationRetryFunc()}); err != nil {
		return tf.ErrorDiagF(err, "Removing synchronization secrets for %s", id)
	}

	// Wait for synchronization secret to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.ListSynchronizationSecrets(ctx, *id, synchronizationsecret.ListSynchronizationSecretsOperationOptions{RetryFunc: synchronizationRetryFunc()})
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}

		synchronizationSecrets := resp.Model
		if synchronizationSecrets == nil {
			return pointer.To(false), nil
		}

		// Test if credentials are removed
		if allCredentialsRemoved(*credentials, *synchronizationSecrets) {
			return pointer.To(false), nil
		}

		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of synchronization secrets %s", id)
	}

	return nil
}
