package serviceprincipals

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func servicePrincipalClaimsMappingPolicyAssignment() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalClaimsMappingPolicyAssignmentResourceCreate,
		ReadContext:   servicePrincipalClaimsMappingPolicyAssignmentResourceRead,
		DeleteContext: servicePrincipalClaimsMappingPolicyAssignmentResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.ObjectSubResourceID(id, "azuread_claims_mapping_policy")
			return err
		}),

		Schema: map[string]*schema.Schema{
			"claims_mapping_policy_id": {
				Description: "ID of the claims mapping policy to assign",
				ForceNew:    true,
				Type:        schema.TypeString,
				Required:    true,
			},

			"service_principal_id": {
				Description: "ID of the service principal to assign the policy to",
				ForceNew:    true,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	spClient := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient
	policyClient := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient

	policyID := d.Get("claims_mapping_policy_id").(string)
	policy, _, err := policyClient.Get(ctx, policyID, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(
			err,
			"Could not find ClaimsMappingPolicy, claims_mapping_policy_id: %q",
			policyID,
		)
	}

	properties := msgraph.ServicePrincipal{
		DirectoryObject: msgraph.DirectoryObject{
			ID: utils.String(d.Get("service_principal_id").(string)),
		},
		ClaimsMappingPolicies: &[]msgraph.ClaimsMappingPolicy{
			*policy,
		},
	}
	_, err = spClient.AssignClaimsMappingPolicy(ctx, &properties)
	if err != nil {
		return tf.ErrorDiagF(
			err,
			"Could not create ClaimsMappingPolicyAssignment, service_principal_id: %q, claims_mapping_policy_id: %q",
			*properties.DirectoryObject.ID,
			*(*properties.ClaimsMappingPolicies)[0].DirectoryObject.ID,
		)
	}

	resourceID := parse.NewClaimsMappingPolicyAssignmentID(
		*properties.DirectoryObject.ID,
		*(*properties.ClaimsMappingPolicies)[0].DirectoryObject.ID,
	)

	d.SetId(resourceID.String())

	return servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx, d, meta)
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	id, err := parse.ClaimsMappingPolicyAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Claims Mapping Policy Assignment ID %q", d.Id())
	}

	spID := id.ServicePolicyId

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
		if *policy.ID == policyID {
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

func servicePrincipalClaimsMappingPolicyAssignmentResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ServicePrincipalsClient

	id, err := parse.ClaimsMappingPolicyAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Claims Mapping Policy Assignment ID %q", d.Id())
	}

	claimIDs := []string{id.ClaimsMappingPolicyId}

	spID := id.ServicePolicyId

	sp := msgraph.ServicePrincipal{
		DirectoryObject: msgraph.DirectoryObject{
			ID: &spID,
		},
	}
	_, err = client.RemoveClaimsMappingPolicy(ctx, &sp, &claimIDs)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not Remove ClaimsMappingPolicyAssignment, service_principal_id: %q, claims_mapping_policy_ids: %q", spID, claimIDs)
	}

	return servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx, d, meta)
}
