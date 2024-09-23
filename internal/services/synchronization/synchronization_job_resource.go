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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationjob"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/synchronization/parse"
)

func synchronizationJobResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: synchronizationJobResourceCreate,
		ReadContext:   synchronizationJobResourceRead,
		UpdateContext: synchronizationJobResourceUpdate,
		DeleteContext: synchronizationJobResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(15 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.SynchronizationJobID(id)
			return err
		}),

		SchemaVersion: 0,

		Schema: map[string]*pluginsdk.Schema{
			"service_principal_id": {
				Description:  "The object ID of the service principal for which this synchronization job should be created",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"template_id": {
				Description: "Identifier of the synchronization template this job is based on.",
				Type:        pluginsdk.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"enabled": {
				Description: "Whether or not the synchronization job is enabled",
				Type:        pluginsdk.TypeBool,
				Default:     true,
				Optional:    true,
			},

			"schedule": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"expiration": {
							Description: "Date and time when this job will expire, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"interval": {
							Description: "The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},

						"state": {
							Description: "State.",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func synchronizationJobResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationJobClient
	servicePrincipalClient := meta.(*clients.Client).Synchronization.ServicePrincipalClient

	servicePrincipalId := stable.NewServicePrincipalID(d.Get("service_principal_id").(string))

	tf.LockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)

	servicePrincipalResp, err := servicePrincipalClient.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
	if err != nil {
		if response.WasNotFound(servicePrincipalResp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "service_principal_id", "%s was not found", servicePrincipalId)
		}
		return tf.ErrorDiagPathF(err, "service_principal_id", "Retrieving %s", servicePrincipalId)
	}

	servicePrincipal := servicePrincipalResp.Model
	if servicePrincipal == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", servicePrincipalId)
	}

	synchronizationJob := stable.SynchronizationJob{
		TemplateId: nullable.Value(d.Get("template_id").(string)),
	}

	resp, err := client.CreateSynchronizationJob(ctx, servicePrincipalId, synchronizationJob, synchronizationjob.CreateSynchronizationJobOperationOptions{RetryFunc: synchronizationRetryFunc()})
	if err != nil {
		return tf.ErrorDiagF(err, "Creating synchronization job for %s", servicePrincipalId)
	}
	if resp.Model == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "API error creating synchronization job for %s", servicePrincipalId)
	}
	if resp.Model.Id == nil {
		return tf.ErrorDiagF(errors.New("nil or empty id received"), "API error creating synchronization job for %s", servicePrincipalId)
	}

	id := stable.NewServicePrincipalIdSynchronizationJobID(servicePrincipalId.ServicePrincipalId, *resp.Model.Id)

	// Wait for the job to appear, this can take several moments
	if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetSynchronizationJob(ctx, id, synchronizationjob.DefaultGetSynchronizationJobOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return pointer.To(false), fmt.Errorf("retrieving synchronization job")
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for creation of %s", id)
	}

	resourceId := parse.NewSynchronizationJobID(id.ServicePrincipalId, id.SynchronizationJobId)
	d.SetId(resourceId.String())

	// Start job if desired
	if d.Get("enabled").(bool) {
		if _, err = client.StartSynchronizationJob(ctx, id, synchronizationjob.StartSynchronizationJobOperationOptions{RetryFunc: synchronizationRetryFunc()}); err != nil {
			return tf.ErrorDiagF(err, "Starting %s", id)
		}
	}

	return synchronizationJobResourceRead(ctx, d, meta)
}

func synchronizationJobResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationJobClient

	resourceId, err := parse.SynchronizationJobID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization job ID %q", d.Id())
	}

	id := stable.NewServicePrincipalIdSynchronizationJobID(resourceId.ServicePrincipalId, resourceId.JobId)

	resp, err := client.GetSynchronizationJob(ctx, id, synchronizationjob.GetSynchronizationJobOperationOptions{RetryFunc: synchronizationRetryFunc()})
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	synchronizationJob := resp.Model
	if synchronizationJob == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "service_principal_id", id.ServicePrincipalId)
	tf.Set(d, "schedule", flattenSynchronizationSchedule(synchronizationJob.Schedule))
	tf.Set(d, "template_id", synchronizationJob.TemplateId.GetOrZero())
	tf.Set(d, "enabled", pointer.From(synchronizationJob.Schedule.State) == stable.SynchronizationScheduleState_Active)
	return nil
}

func synchronizationJobResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationJobClient

	resourceId, err := parse.SynchronizationJobID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization job ID %q", d.Id())
	}

	id := stable.NewServicePrincipalIdSynchronizationJobID(resourceId.ServicePrincipalId, resourceId.JobId)

	if d.HasChange("enabled") {
		if d.Get("enabled").(bool) {
			if _, err = client.StartSynchronizationJob(ctx, id, synchronizationjob.StartSynchronizationJobOperationOptions{RetryFunc: synchronizationRetryFunc()}); err != nil {
				return tf.ErrorDiagF(err, "Starting %s", id)
			}
		} else {
			if _, err = client.PauseSynchronizationJob(ctx, id, synchronizationjob.PauseSynchronizationJobOperationOptions{RetryFunc: synchronizationRetryFunc()}); err != nil {
				return tf.ErrorDiagF(err, "Pausing %s", id)
			}
		}
	}

	return synchronizationJobResourceRead(ctx, d, meta)
}

func synchronizationJobResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Synchronization.SynchronizationJobClient

	resourceId, err := parse.SynchronizationJobID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization job with ID %q", d.Id())
	}

	id := stable.NewServicePrincipalIdSynchronizationJobID(resourceId.ServicePrincipalId, resourceId.JobId)

	tf.LockByName(servicePrincipalResourceName, id.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ServicePrincipalId)

	if _, err = client.DeleteSynchronizationJob(ctx, id, synchronizationjob.DeleteSynchronizationJobOperationOptions{RetryFunc: synchronizationRetryFunc()}); err != nil {
		return tf.ErrorDiagF(err, "Removing %s", id)
	}

	// Wait for synchronization job to be deleted
	if err = consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetSynchronizationJob(ctx, id, synchronizationjob.DefaultGetSynchronizationJobOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of %s", id)
	}

	return nil
}
