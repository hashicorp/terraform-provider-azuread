package conditionalAccessPolicies

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func conditionalAccessPolicyResourceCreateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccessPolicies.MsClient

	displayName := d.Get("display_name").(string)
	state := d.Get("state").(string)

	properties := msgraph.ConditionalAccessPolicy{
		DisplayName: utils.String(displayName),
		State:       utils.String(state),
		Conditions: &msgraph.ConditionalAccessConditionSet{
			ClientAppTypes: &[]string{"mobileAppsAndDesktopClients", "browser"},
			Applications: &msgraph.ConditionalAccessApplications{
				IncludeApplications: &[]string{"All"},
			},
			Users: &msgraph.ConditionalAccessUsers{
				IncludeUsers: &[]string{"All"},
				ExcludeUsers: &[]string{"GuestsOrExternalUsers"},
			},
			Locations: &msgraph.ConditionalAccessLocations{
				IncludeLocations: &[]string{"All"},
				ExcludeLocations: &[]string{"AllTrusted"},
			},
		},
		GrantControls: &msgraph.ConditionalAccessGrantControls{
			Operator:        utils.String("OR"),
			BuiltInControls: &[]string{"block"},
		},
	}

	policy, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create conditional access policy")
	}

	if policy.ID == nil || *policy.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for conditional access policy is nil/empty")
	}

	d.SetId(*policy.ID)

	return conditionalAccessPolicyResourceReadMsGraph(ctx, d, meta)
}

func conditionalAccessPolicyResourceUpdateMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccessPolicies.MsClient

	var displayName string
	if d.HasChange("display_name") {
		displayName = d.Get("display_name").(string)
	}

	var state string
	if d.HasChange("state") {
		state = d.Get("state").(string)
	}

	properties := msgraph.ConditionalAccessPolicy{
		ID: utils.String(d.Id()),
	}

	if displayName != "" {
		properties.DisplayName = &displayName
	}

	if state != "" {
		properties.State = &state
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update conditional access policy with ID: %q", d.Id())
	}

	return nil
}

func conditionalAccessPolicyResourceReadMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccessPolicies.MsClient

	policy, status, err := client.Get(ctx, d.Id())
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Conditional Access Policy with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Conditional Access Policy with object ID %q", d.Id())
	}

	tf.Set(d, "display_name", policy.DisplayName)
	tf.Set(d, "id", policy.ID)
	tf.Set(d, "state", policy.State)

	return nil
}

func conditionalAccessPolicyResourceDeleteMsGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccessPolicies.MsClient

	_, status, err := client.Get(ctx, d.Id())
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Conditional Access Policy with ID %q already deleted", d.Id())
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving conditional access policy with ID %q", d.Id())
	}

	status, err = client.Delete(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting conditional access policy with ID %q, got status %d", d.Id(), status)
	}

	return nil
}
