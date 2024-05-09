package synchronization

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func synchronizationJobProvisionOnDemandResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: synchronizationProvisionOnDemandResourceCreate,
		ReadContext:   synchronizationProvisionOnDemandResourceRead,
		DeleteContext: synchronizationProvisionOnDemandResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"service_principal_id": {
				Description:      "The object ID of the service principal for which this synchronization job should be provisioned",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"synchronization_job_id": {
				Description: "The identifier for the synchronization jop.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"parameter": {
				Description: "Represents the objects that will be provisioned and the synchronization rules executed. The resource is primarily used for on-demand provisioning.",
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Description: "The identifier of the synchronization rule to be applied. This rule ID is defined in the schema for a given synchronization job or template.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},

						"subject": {
							Description: "The identifiers of one or more objects to which a synchronizationJob is to be applied.",
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_id": {
										Description: "The identifier of an object to which a synchronization job is to be applied. Can be one of the following: (1) An onPremisesDistinguishedName for synchronization from Active Directory to Azure AD. (2) The user ID for synchronization from Azure AD to a third-party. (3) The Worker ID of the Workday worker for synchronization from Workday to either Active Directory or Azure AD.",
										Type:        schema.TypeString,
										Required:    true,
									},

									"object_type_name": {
										Description:  "The type of the object to which a synchronization job is to be applied. Can be one of the following: `user` for synchronizing between Active Directory and Azure AD, `User` for synchronizing a user between Azure AD and a third-party application, `Worker` for synchronization a user between Workday and either Active Directory or Azure AD, `Group` for synchronizing a group between Azure AD and a third-party application.",
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"Group", "user", "User", "Worker"}, false),
									},
								},
							},
						},
					},
				},
			},

			"triggers": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func synchronizationProvisionOnDemandResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.SynchronizationJobClient
	spClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	objectId := d.Get("service_principal_id").(string)
	jobId := d.Get("synchronization_job_id").(string)

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

	job, status, err := client.Get(ctx, jobId, objectId)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "job_id", "Job with object ID %q was not found for service principle %q", jobId, objectId)
		}
		return tf.ErrorDiagPathF(err, "job_id", "Retrieving job with object ID %q for service principle %q", jobId, objectId)
	}
	if job == nil || job.ID == nil {
		return tf.ErrorDiagF(errors.New("nil job or job with nil ID was returned"), "API error retrieving job with object ID %q/%s", objectId, jobId)
	}

	// Create a new synchronization job
	synchronizationProvisionOnDemand := &msgraph.SynchronizationJobProvisionOnDemand{
		Parameters: expandSynchronizationJobApplicationParameters(d.Get("parameter").([]interface{})),
	}

	_, err = client.ProvisionOnDemand(ctx, jobId, synchronizationProvisionOnDemand, *servicePrincipal.ID())
	if err != nil {
		return tf.ErrorDiagF(err, "Creating synchronization job for service principal ID %q", *servicePrincipal.ID())
	}

	id, _ := uuid.GenerateUUID()
	d.SetId(id)

	return synchronizationProvisionOnDemandResourceRead(ctx, d, meta)
}

func synchronizationProvisionOnDemandResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func synchronizationProvisionOnDemandResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
