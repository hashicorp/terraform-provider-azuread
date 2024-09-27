// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synchronization

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationjob"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
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
				Description:  "The object ID of the service principal for which this synchronization job should be provisioned",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: stable.ValidateServicePrincipalID,
			},

			"synchronization_job_id": {
				Description:  "The identifier for the synchronization jop.",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: stable.ValidateServicePrincipalIdSynchronizationJobID,
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
	servicePrincipalClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalClient

	servicePrincipalId, err := stable.ParseServicePrincipalID(d.Get("service_principal_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "service_principal_id", "Parsing `service_principal_id`")
	}

	tf.LockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)
	defer tf.UnlockByName(servicePrincipalResourceName, servicePrincipalId.ServicePrincipalId)

	servicePrincipalResp, err := servicePrincipalClient.GetServicePrincipal(ctx, *servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
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
	if servicePrincipal.Id == nil {
		return tf.ErrorDiagF(errors.New("model has nil ID"), "Retrieving %s", servicePrincipalId)
	}

	jobId, err := stable.ParseServicePrincipalIdSynchronizationJobID(d.Get("synchronization_job_id").(string))
	if err != nil {
		return tf.ErrorDiagPathF(err, "synchronization_job_id", "Parsing `synchronization_job_id`")
	}

	jobResp, err := client.GetSynchronizationJob(ctx, *jobId, synchronizationjob.GetSynchronizationJobOperationOptions{RetryFunc: synchronizationRetryFunc()})
	if err != nil {
		if response.WasNotFound(jobResp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "synchronization_job_id", "%s was not found", jobId)
		}
		return tf.ErrorDiagPathF(err, "job_id", "Retrieving %s", jobId)
	}

	job := jobResp.Model
	if job == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", jobId)
	}
	if job.Id == nil {
		return tf.ErrorDiagF(errors.New("model has nil ID"), "Retrieving %s", jobId)
	}

	properties := synchronizationjob.ProvisionSynchronizationJobOnDemandRequest{
		Parameters: expandSynchronizationJobApplicationParameters(d.Get("parameter").([]interface{})),
	}

	if _, err = client.ProvisionSynchronizationJobOnDemand(ctx, *jobId, properties, synchronizationjob.DefaultProvisionSynchronizationJobOnDemandOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Provisioning %s", jobId)
	}

	id, _ := uuid.GenerateUUID()
	d.SetId(id)

	return synchronizationProvisionOnDemandResourceRead(ctx, d, meta)
}

func synchronizationProvisionOnDemandResourceRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	// Nothing to read
	return nil
}

func synchronizationProvisionOnDemandResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	// Nothing to destroy
	return nil
}
