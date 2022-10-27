package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func synchronizationJobResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: synchronizationJobResourceCreate,
		ReadContext:   synchronizationJobResourceRead,
		UpdateContext: synchronizationJobResourceUpdate,
		DeleteContext: synchronizationJobResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.SynchronizationJobID(id)
			return err
		}),

		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"service_principal_id": {
				Description:      "The object ID of the service principal for which this synchronization job should be created",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validate.UUID,
			},
			"template_id": {
				Description: "Identifier of the synchronization template this job is based on.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"enabled": {
				Description: "Whether or not the synchronization job is enabled",
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
			},
			"schedule": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiration": {
							Description: "Date and time when this job will expire, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"interval": {
							Description: "The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"state": {
							Description: "State.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func synchronizationJobResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	if servicePrincipal == nil || servicePrincipal.ID == nil {
		return tf.ErrorDiagF(errors.New("nil service principal or service principal with nil ID was returned"), "API error retrieving service principal with object ID %q", objectId)
	}

	// Create a new synchronization job
	synchronizationJob := msgraph.SynchronizationJob{
		TemplateId: utils.String(d.Get("template_id").(string)),
	}

	newJob, _, err := client.Create(ctx, synchronizationJob, *servicePrincipal.ID)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating synchronization job for service principal ID %q", *servicePrincipal.ID)
	}
	if newJob == nil {
		return tf.ErrorDiagF(errors.New("nil received when creating synchronization Job"), "API error creating synchronization job for service principal ID %q", *servicePrincipal.ID)
	}
	if newJob.ID == nil {
		return tf.ErrorDiagF(errors.New("nil or empty id received"), "API error creating synchronization job for service principal ID %q", *servicePrincipal.ID)
	}

	id := parse.NewSynchronizationJobID(*servicePrincipal.ID, *newJob.ID)

	// Wait for the job to appear, this can take several moments
	timeout, _ := ctx.Deadline()
	_, err = (&resource.StateChangeConf{
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			_, status, err := client.Get(ctx, id.JobId, id.ServicePrincipalId)
			if err != nil {
				if status == http.StatusNotFound {
					return "stub", "Waiting", nil
				}
				return nil, "Error", fmt.Errorf("retrieving synchronization job")
			}
			return "stub", "Done", nil
		},
	}).WaitForStateContext(ctx)

	if err != nil {
		return tf.ErrorDiagF(err, "Waiting for synchronization job %q", id.JobId)
	}

	d.SetId(id.String())

	// Start job if desired
	if d.Get("enabled").(bool) {
		_, err := client.Start(ctx, id.JobId, id.ServicePrincipalId)
		if err != nil {
			return tf.ErrorDiagF(err, "Starting synchronization job %q", id.JobId)
		}
	}

	return synchronizationJobResourceRead(ctx, d, meta)
}

func synchronizationJobResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.SynchronizationJobClient

	id, err := parse.SynchronizationJobID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization job with ID %q", d.Id())
	}

	job, status, err := client.Get(ctx, id.JobId, id.ServicePrincipalId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Synchronization job with ID %q for service principal %q was not found - removing from state!", id.JobId, id.ServicePrincipalId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving synchronization job with object ID %q", id.JobId)
	}
	tf.Set(d, "service_principal_id", id.ServicePrincipalId)
	tf.Set(d, "schedule", flattenSynchronizationSchedule(job.Schedule))
	tf.Set(d, "template_id", job.TemplateId)
	tf.Set(d, "enabled", *job.Schedule.State == "Active")
	return nil
}

func synchronizationJobResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.SynchronizationJobClient
	id, err := parse.SynchronizationJobID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization job with ID %q", d.Id())
	}
	if d.HasChange("enabled") {
		if d.Get("enabled").(bool) {
			_, err := client.Start(ctx, id.JobId, id.ServicePrincipalId)
			if err != nil {
				return tf.ErrorDiagF(err, "Starting synchronization job %q", id.JobId)
			}
		} else {
			_, err := client.Pause(ctx, id.JobId, id.ServicePrincipalId)
			if err != nil {
				return tf.ErrorDiagF(err, "Pausing synchronization job %q", id.JobId)
			}
		}
	}
	return synchronizationJobResourceRead(ctx, d, meta)
}

func synchronizationJobResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.SynchronizationJobClient

	id, err := parse.SynchronizationJobID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing synchronization job with ID %q", d.Id())
	}

	tf.LockByName(servicePrincipalResourceName, id.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, id.ServicePrincipalId)

	if _, err := client.Delete(ctx, id.JobId, id.ServicePrincipalId); err != nil {
		return tf.ErrorDiagF(err, "Removing job %q from service principal with object ID %q", id.JobId, id.ServicePrincipalId)
	}

	// Wait for synchronization job to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true

		job, _, _ := client.Get(ctx, id.JobId, id.ServicePrincipalId)
		if job == nil {
			return utils.Bool(false), nil
		}

		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of synchronization job %q from service principal with object ID %q", id.JobId, id.ServicePrincipalId)
	}

	return nil
}
