package policies

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
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func claimsMappingPolicyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: claimsMappingPolicyResourceCreate,
		ReadContext:   claimsMappingPolicyResourceRead,
		UpdateContext: claimsMappingPolicyResourceUpdate,
		DeleteContext: claimsMappingPolicyResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"definition": {
				Description: "A string collection containing a JSON string that defines the rules and settings for this policy",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"display_name": {
				Description: "Display name for this policy",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func claimsMappingPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient

	claimsMappingPolicy := msgraph.ClaimsMappingPolicy{
		Definition:  tf.ExpandStringSlicePtr(d.Get("definition").([]interface{})),
		DisplayName: utils.String(d.Get("display_name").(string)),
	}
	policy, _, err := client.Create(ctx, claimsMappingPolicy)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create Claims Mapping Policy")
	}

	if policy.ID == nil || *policy.ID == "" {
		return tf.ErrorDiagF(fmt.Errorf("Object ID returned for Claims Mapping Policy is nil"), "Bad API response")
	}

	d.SetId(*policy.ID)

	return claimsMappingPolicyResourceRead(ctx, d, meta)
}

func claimsMappingPolicyResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient
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

	tf.Set(d, "definition", policy.Definition)
	tf.Set(d, "display_name", policy.DisplayName)

	return nil
}

func claimsMappingPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient
	objectId := d.Id()

	claimsMappingPolicy := msgraph.ClaimsMappingPolicy{
		DirectoryObject: msgraph.DirectoryObject{
			ID: &objectId,
		},
		Definition:  tf.ExpandStringSlicePtr(d.Get("definition").([]interface{})),
		DisplayName: utils.String(d.Get("display_name").(string)),
	}
	_, err := client.Update(ctx, claimsMappingPolicy)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not update Claims Mapping Policy with object ID %q", objectId)
	}

	return claimsMappingPolicyResourceRead(ctx, d, meta)
}

func claimsMappingPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient
	objectId := d.Id()

	_, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Claims Mapping Policy was not found"), "id", "Retrieving Claims Mapping Policy with object ID %q", objectId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Claims Mapping Policy with object ID %q", objectId)
	}

	status, err = client.Delete(ctx, objectId)
	if err != nil {
		return tf.ErrorDiagF(err, "Deleting Claims Mapping Policy with object ID %q, received status %d", objectId, status)
	}

	return nil
}
