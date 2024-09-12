// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synchronization

import (
	"context"
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
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/synchronization/parse"
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

		Schema: map[string]*pluginsdk.Schema{
			"service_principal_id": {
				Description:      "The object ID of the service principal for which this synchronization secret should be created",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
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

	servicePrincipalId := stable.NewServicePrincipalID(d.Get("service_principal_id").(string))

	tf.LockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)

	synchronizationSecrets := synchronizationsecret.SetSynchronizationSecretRequest{
		Value: expandSynchronizationSecretKeyStringValuePair(d.Get("credential").([]interface{})),
	}

	if _, err := client.SetSynchronizationSecret(ctx, servicePrincipalId, synchronizationSecrets); err != nil {
		return tf.ErrorDiagF(err, "Creating synchronization secret for %s", servicePrincipalId)
	}

	// Wait for the secret to appear
	timeout, _ := ctx.Deadline()
	if _, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.ListSynchronizationSecrets(ctx, servicePrincipalId, synchronizationsecret.DefaultListSynchronizationSecretsOperationOptions())
			if err != nil {
				return nil, "Error", fmt.Errorf("retrieving synchronization secret")
			}
			newSynchronizationSecrets := resp.Model
			if newSynchronizationSecrets == nil {
				return "stub", "Waiting", nil
			}
			if len(pointer.From(synchronizationSecrets.Value)) == len(pointer.From(newSynchronizationSecrets)) {
				return "stub", "Done", nil
			}
			return "stub", "Waiting", nil
		},
	}).WaitForStateContext(ctx); err != nil {
		return tf.ErrorDiagF(err, "Waiting for synchronization secrets for %s", servicePrincipalId)
	}

	d.SetId(parse.NewSynchronizationSecretID(servicePrincipalId.ServicePrincipalId).String())
	tf.Set(d, "credential", flattenSynchronizationSecretKeyStringValuePair(synchronizationSecrets.Value, nil))

	return synchronizationSecretResourceRead(ctx, d, meta)
}

func synchronizationSecretResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationSecretClient

	id, err := parse.SynchronizationSecretID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization secret ID %q", d.Id())
	}

	servicePrincipalId := stable.NewServicePrincipalID(id.ServicePrincipalId)

	tf.LockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)

	synchronizationSecrets := synchronizationsecret.SetSynchronizationSecretRequest{
		Value: expandSynchronizationSecretKeyStringValuePair(d.Get("credential").([]interface{})),
	}

	if _, err := client.SetSynchronizationSecret(ctx, servicePrincipalId, synchronizationSecrets); err != nil {
		return tf.ErrorDiagF(err, "Updating synchronization secret for %s", servicePrincipalId)
	}

	// Wait for the secret to update
	timeout, _ := ctx.Deadline()
	if _, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.ListSynchronizationSecrets(ctx, servicePrincipalId, synchronizationsecret.DefaultListSynchronizationSecretsOperationOptions())
			if err != nil {
				return nil, "Error", fmt.Errorf("retrieving synchronization secret")
			}
			newSynchronizationSecrets := resp.Model
			if newSynchronizationSecrets == nil {
				return "stub", "Waiting", nil
			}
			if len(pointer.From(synchronizationSecrets.Value)) == len(pointer.From(newSynchronizationSecrets)) {
				return "stub", "Done", nil
			}
			return "stub", "Waiting", nil
		},
	}).WaitForStateContext(ctx); err != nil {
		return tf.ErrorDiagF(err, "Waiting for synchronization secrets for %s", servicePrincipalId)
	}

	tf.Set(d, "credential", flattenSynchronizationSecretKeyStringValuePair(synchronizationSecrets.Value, nil))

	return synchronizationSecretResourceRead(ctx, d, meta)
}

func synchronizationSecretResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationSecretClient

	id, err := parse.SynchronizationSecretID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization secret ID %q", d.Id())
	}

	servicePrincipalId := stable.NewServicePrincipalID(id.ServicePrincipalId)

	resp, err := client.ListSynchronizationSecrets(ctx, servicePrincipalId, synchronizationsecret.DefaultListSynchronizationSecretsOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] Synchronization secrets for %s was not found - removing from state!", servicePrincipalId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving synchronization secrets for %s", servicePrincipalId)
	}

	synchronizationSecrets := resp.Model
	if synchronizationSecrets == nil {
		log.Printf("[DEBUG] Synchronization secrets for %s was nil - removing from state!", servicePrincipalId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "credential", flattenSynchronizationSecretKeyStringValuePair(synchronizationSecrets, d.Get("credential").([]interface{})))

	return nil
}

func synchronizationSecretResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationSecretClient

	id, err := parse.SynchronizationSecretID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization secret with ID %q", d.Id())
	}

	servicePrincipalId := stable.NewServicePrincipalID(id.ServicePrincipalId)

	tf.LockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)

	// We delete secrets by setting values to empty strings
	credentials := emptySynchronizationSecretKeyStringValuePair(d.Get("credential").([]interface{}))

	synchronizationSecrets := synchronizationsecret.SetSynchronizationSecretRequest{
		Value: credentials,
	}
	if _, err := client.SetSynchronizationSecret(ctx, servicePrincipalId, synchronizationSecrets); err != nil {
		return tf.ErrorDiagF(err, "Removing synchronization secrets for %s", servicePrincipalId)
	}

	// Wait for synchronization secret to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.ListSynchronizationSecrets(ctx, servicePrincipalId, synchronizationsecret.DefaultListSynchronizationSecretsOperationOptions())
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
		return tf.ErrorDiagF(err, "Waiting for deletion of synchronization secrets %s", servicePrincipalId)
	}

	return nil
}
