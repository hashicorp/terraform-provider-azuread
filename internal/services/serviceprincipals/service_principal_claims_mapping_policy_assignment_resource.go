// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/claimsmappingpolicy"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
)

func servicePrincipalClaimsMappingPolicyAssignmentResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: servicePrincipalClaimsMappingPolicyAssignmentResourceCreate,
		ReadContext:   servicePrincipalClaimsMappingPolicyAssignmentResourceRead,
		DeleteContext: servicePrincipalClaimsMappingPolicyAssignmentResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.ObjectSubResourceID(id, "claimsMappingPolicy")
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"claims_mapping_policy_id": {
				Description:  "ID of the claims mapping policy to assign",
				Type:         pluginsdk.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"service_principal_id": {
				Description:  "Object ID of the service principal for which to assign the policy",
				Type:         pluginsdk.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient

	servicePrincipalId := stable.NewServicePrincipalID(d.Get("service_principal_id").(string))
	policyId := stable.NewDirectoryObjectID(d.Get("claims_mapping_policy_id").(string))

	ref := stable.ReferenceCreate{
		ODataId: pointer.To(client.Client.BaseUri + policyId.ID()),
	}

	if _, err := client.AddClaimsMappingPolicyRef(ctx, servicePrincipalId, ref, claimsmappingpolicy.DefaultAddClaimsMappingPolicyRefOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Creating ClaimsMappingPolicyAssignment for %s", servicePrincipalId)
	}

	id := parse.NewClaimsMappingPolicyAssignmentID(servicePrincipalId.ServicePrincipalId, policyId.DirectoryObjectId)
	d.SetId(id.String())

	return servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx, d, meta)
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient

	id, err := parse.ClaimsMappingPolicyAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Claims Mapping Policy Assignment ID %q", d.Id())
	}

	servicePrincipalId := stable.NewServicePrincipalID(id.ServicePrincipalId)

	resp, err := client.ListClaimsMappingPolicies(ctx, servicePrincipalId, claimsmappingpolicy.DefaultListClaimsMappingPoliciesOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing claims mapping policy assignment from state!", servicePrincipalId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "listing Claims Mapping Policy Assignments for %s", servicePrincipalId)
	}

	policies := resp.Model
	if policies == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "listing Claims Mapping Policy Assignments for %s", servicePrincipalId)
	}

	var policy *stable.ClaimsMappingPolicy

	// Check the assignment is found in the currently assigned policies
	for _, p := range *policies {
		if pointer.From(p.Id) == id.ClaimsMappingPolicyId {
			policy = &p
			break
		}
	}
	if policy == nil {
		d.SetId("")
		log.Printf("[DEBUG] Claims Mapping Policy with Object ID %q was not found - removing assignment from state!", id.ClaimsMappingPolicyId)
		return nil
	}

	tf.Set(d, "service_principal_id", id.ServicePrincipalId)
	tf.Set(d, "claims_mapping_policy_id", id.ClaimsMappingPolicyId)

	return nil
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient

	id, err := parse.ClaimsMappingPolicyAssignmentID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Claims Mapping Policy Assignment ID %q", d.Id())
	}

	refId := stable.NewServicePrincipalIdClaimsMappingPolicyID(id.ServicePrincipalId, id.ClaimsMappingPolicyId)
	if _, err = client.RemoveClaimsMappingPolicyRef(ctx, refId, claimsmappingpolicy.DefaultRemoveClaimsMappingPolicyRefOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "removing %s", refId)
	}

	return servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx, d, meta)
}
