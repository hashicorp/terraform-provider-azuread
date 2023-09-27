// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func servicePrincipalClaimsMappingPolicyAssignmentResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: servicePrincipalClaimsMappingPolicyAssignmentResourceCreate,
		ReadContext:   servicePrincipalClaimsMappingPolicyAssignmentResourceRead,
		DeleteContext: servicePrincipalClaimsMappingPolicyAssignmentResourceDelete,

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.ObjectSubResourceID(id, "claimsMappingPolicy")
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"claims_mapping_policy_id": {
				Description: "ID of the claims mapping policy to assign",
				Type:        pluginsdk.TypeString,
				ForceNew:    true,
				Required:    true,
			},

			"service_principal_id": {
				Description: "Object ID of the service principal for which to assign the policy",
				Type:        pluginsdk.TypeString,
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	tenantId := meta.(*clients.Client).TenantID

	policyId := d.Get("claims_mapping_policy_id").(string)

	properties := msgraph.ServicePrincipal{
		DirectoryObject: msgraph.DirectoryObject{
			Id: utils.String(d.Get("service_principal_id").(string)),
		},
		ClaimsMappingPolicies: &[]msgraph.ClaimsMappingPolicy{
			{
				DirectoryObject: msgraph.DirectoryObject{
					ODataId: (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
						client.BaseClient.Endpoint, tenantId, policyId))),
					Id: &policyId,
				},
			},
		},
	}

	_, err := client.AssignClaimsMappingPolicy(ctx, &properties)
	if err != nil {
		return tf.ErrorDiagF(
			err,
			"Could not create ClaimsMappingPolicyAssignment, service_principal_id: %q, claims_mapping_policy_id: %q",
			*properties.DirectoryObject.ID(),
			*(*properties.ClaimsMappingPolicies)[0].DirectoryObject.ID(),
		)
	}

	id := parse.NewClaimsMappingPolicyAssignmentID(
		*properties.DirectoryObject.ID(),
		*(*properties.ClaimsMappingPolicies)[0].DirectoryObject.ID(),
	)

	d.SetId(id.String())

	return servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx, d, meta)
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	id, err := parse.ClaimsMappingPolicyAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Claims Mapping Policy Assignment ID %q", d.Id())
	}

	spID := id.ServicePrincipalId

	policyList, status, err := client.ListClaimsMappingPolicy(ctx, spID)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Service Principal with Object ID %q was not found - removing claims mapping policy assignment from state!", spID)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "listing Claims Mapping Policy Assignments for Service Principal with object ID: %q", d.Id())
	}

	policyID := id.ClaimsMappingPolicyId
	var foundPolicy *msgraph.ClaimsMappingPolicy

	// Check the assignment is found in the currently assigned policies
	for _, policy := range *policyList {
		if *policy.ID() == policyID {
			foundPolicy = &policy
			break
		}
	}
	if foundPolicy == nil {
		d.SetId("")
		log.Printf("[DEBUG] Claims Mapping Policy with Object ID %q was not found - removing assignment from state!", policyID)
		return nil
	}

	tf.Set(d, "service_principal_id", spID)
	tf.Set(d, "claims_mapping_policy_id", policyID)

	return nil
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	id, err := parse.ClaimsMappingPolicyAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Claims Mapping Policy Assignment ID %q", d.Id())
	}

	claimIDs := []string{id.ClaimsMappingPolicyId}

	spID := id.ServicePrincipalId

	sp := msgraph.ServicePrincipal{
		DirectoryObject: msgraph.DirectoryObject{
			Id: &spID,
		},
	}
	_, err = client.RemoveClaimsMappingPolicy(ctx, &sp, &claimIDs)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not Remove ClaimsMappingPolicyAssignment, service_principal_id: %q, claims_mapping_policy_ids: %q", spID, claimIDs)
	}

	return servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx, d, meta)
}
