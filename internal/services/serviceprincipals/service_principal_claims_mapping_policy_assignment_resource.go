// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/claimsmappingpolicy"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/migrations"
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
			if _, errs := stable.ValidateServicePrincipalIdClaimsMappingPolicyID(id, "id"); len(errs) > 0 {
				out := ""
				for _, err := range errs {
					out += err.Error()
				}
				return fmt.Errorf(out)
			}
			return nil
		}),

		SchemaVersion: 1,
		StateUpgraders: []pluginsdk.StateUpgrader{
			{
				Type:    migrations.ResourceServicePrincipalClaimsMappingPolicyAssignmentInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceServicePrincipalClaimsMappingPolicyAssignmentInstanceStateUpgradeV0,
				Version: 0,
			},
		},

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

	id := stable.NewServicePrincipalIdClaimsMappingPolicyID(servicePrincipalId.ServicePrincipalId, policyId.DirectoryObjectId)
	d.SetId(id.ID())

	return servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx, d, meta)
}

func servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient

	id, err := stable.ParseServicePrincipalIdClaimsMappingPolicyID(d.Id())
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

	id, err := stable.ParseServicePrincipalIdClaimsMappingPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Claims Mapping Policy Assignment ID %q", d.Id())
	}

	if _, err = client.RemoveClaimsMappingPolicyRef(ctx, *id, claimsmappingpolicy.DefaultRemoveClaimsMappingPolicyRefOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "removing %s", id)
	}

	return servicePrincipalClaimsMappingPolicyAssignmentResourceRead(ctx, d, meta)
}
