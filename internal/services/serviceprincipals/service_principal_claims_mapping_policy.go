package serviceprincipals

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func servicePrincipalClaimsMappingPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: servicePrincipalClaimsMappingPolicyResourceCreate,
		ReadContext:   servicePrincipalClaimsMappingPolicyResourceRead,
		UpdateContext: servicePrincipalClaimsMappingPolicyResourceUpdate,
		DeleteContext: servicePrincipalClaimsMappingPolicyResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier for this policy",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"definition": {
				Description: "A string collection containing a JSON string " +
					"that defines the rules and settings for this policy.",
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"display_name": {
				Description: "Display name for this policy",
				Type:        schema.TypeString,
				Required:    true,
			},

			"description": {
				Description: "Description for this policy",
				Optional:    true,
				Type:        schema.TypeString,
				Required:    false,
			},
		},
	}
}

func servicePrincipalClaimsMappingPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient
	var definitions []string
	for _, v := range d.Get("definition").([]interface{}) {
		definitions = append(definitions, v.(string))
	}

	displayName := d.Get("display_name").(string)

	claimsMappingPolicy := msgraph.ClaimsMappingPolicy{
		Definition:  &definitions,
		DisplayName: &displayName,
	}
	policy, _, err := client.Create(ctx, claimsMappingPolicy)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create ClaimsMappingPolicy %q", displayName)
	}

	if policy != nil {
		d.SetId(*policy.ID)
	}

	return servicePrincipalClaimsMappingPolicyResourceRead(ctx, d, meta)
}

func servicePrincipalClaimsMappingPolicyResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient
	objectId := d.Id()

	policy, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Claims Mapping Policy with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "retrieving Claims Mapping Policy with object ID: %q", d.Id())
	}

	tf.Set(d, "id", policy.ID)
	tf.Set(d, "definition", policy.Definition)
	tf.Set(d, "display_name", policy.DisplayName)

	return nil
}

func servicePrincipalClaimsMappingPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient
	objectId := d.Id()

	var definitions []string
	for _, v := range d.Get("definition").([]interface{}) {
		definitions = append(definitions, v.(string))
	}

	displayName := d.Get("display_name").(string)

	claimsMappingPolicy := msgraph.ClaimsMappingPolicy{
		DirectoryObject: msgraph.DirectoryObject{
			ID: &objectId,
		},
		Definition:  &definitions,
		DisplayName: &displayName,
	}
	_, err := client.Update(ctx, claimsMappingPolicy)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not update ClaimsMappingPolicy %q", displayName)
	}

	return servicePrincipalClaimsMappingPolicyResourceRead(ctx, d, meta)
}

func servicePrincipalClaimsMappingPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ServicePrincipals.ClaimsMappingPolicyClient
	objectId := d.Id()

	_, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(
				fmt.Errorf(
					"Claims Mapping Policy was not found"),
				"id", "Retrieving Claims Mapping Policy with object ID %q",
				objectId,
			)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Claims Mapping Policy with object ID %q", objectId)
	}

	status, err = client.Delete(ctx, objectId)
	if err != nil {
		return tf.ErrorDiagF(err, "Deleting Claims Mapping Policy with object ID %q, received status %d", objectId, status)
	}

	return nil
}
