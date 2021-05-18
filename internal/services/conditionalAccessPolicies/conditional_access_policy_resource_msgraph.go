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

	conditionsRaw := d.Get("conditions").([]interface{})
	conditions := expandConditionalAccessConditionSet(conditionsRaw)

	grantControlsRaw := d.Get("grant_controls").([]interface{})
	grantControls := expandConditionalAccessGrantControls(grantControlsRaw)

	properties := msgraph.ConditionalAccessPolicy{
		DisplayName:   utils.String(displayName),
		State:         utils.String(state),
		Conditions:    conditions,
		GrantControls: grantControls,
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

	properties := msgraph.ConditionalAccessPolicy{
		ID: utils.String(d.Id()),
	}

	if d.HasChange("display_name") {
		displayName := d.Get("display_name").(string)
		properties.DisplayName = &displayName
	}

	if d.HasChange("state") {
		state := d.Get("state").(string)
		properties.State = &state
	}

	if d.HasChange("conditions") {
		conditionsRaw := d.Get("conditions").([]interface{})
		conditions := expandConditionalAccessConditionSet(conditionsRaw)
		properties.Conditions = conditions
	}

	if d.HasChange("grant_controls") {
		grantControlsRaw := d.Get("grant_controls").([]interface{})
		grantControls := expandConditionalAccessGrantControls(grantControlsRaw)
		properties.GrantControls = grantControls
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

func expandConditionalAccessConditionSet(in []interface{}) *msgraph.ConditionalAccessConditionSet {
	if len(in) == 0 {
		return nil
	}

	result := msgraph.ConditionalAccessConditionSet{}
	config := in[0].(map[string]interface{})

	applications := config["applications"].([]interface{})
	users := config["users"].([]interface{})
	clientAppTypes := config["client_app_types"].([]interface{})
	locations := config["locations"].([]interface{})

	result.Applications = expandConditionalAccessApplications(applications)
	result.Users = expandConditionalAccessUsers(users)
	result.ClientAppTypes = tf.ExpandStringSlicePtr(clientAppTypes)
	result.Locations = expandConditionalAccessLocations(locations)

	return &result
}

func expandConditionalAccessApplications(in []interface{}) *msgraph.ConditionalAccessApplications {
	if len(in) == 0 {
		return nil
	}

	result := msgraph.ConditionalAccessApplications{}
	config := in[0].(map[string]interface{})

	includeApplications := config["included_applications"].([]interface{})
	excludeApplications := config["excluded_applications"].([]interface{})
	includeUserActions := config["included_user_actions"].([]interface{})

	result.IncludeApplications = tf.ExpandStringSlicePtr(includeApplications)
	result.ExcludeApplications = tf.ExpandStringSlicePtr(excludeApplications)
	result.IncludeUserActions = tf.ExpandStringSlicePtr(includeUserActions)

	return &result
}

func expandConditionalAccessUsers(in []interface{}) *msgraph.ConditionalAccessUsers {
	if len(in) == 0 {
		return nil
	}

	result := msgraph.ConditionalAccessUsers{}
	config := in[0].(map[string]interface{})

	includeUsers := config["included_users"].([]interface{})
	excludeUsers := config["excluded_users"].([]interface{})
	includeGroups := config["included_groups"].([]interface{})
	excludeGroups := config["excluded_groups"].([]interface{})
	includeRoles := config["included_roles"].([]interface{})
	excludeRoles := config["excluded_roles"].([]interface{})

	result.IncludeUsers = tf.ExpandStringSlicePtr(includeUsers)
	result.ExcludeUsers = tf.ExpandStringSlicePtr(excludeUsers)
	result.IncludeGroups = tf.ExpandStringSlicePtr(includeGroups)
	result.ExcludeGroups = tf.ExpandStringSlicePtr(excludeGroups)
	result.IncludeRoles = tf.ExpandStringSlicePtr(includeRoles)
	result.ExcludeRoles = tf.ExpandStringSlicePtr(excludeRoles)

	return &result
}

func expandConditionalAccessLocations(in []interface{}) *msgraph.ConditionalAccessLocations {
	if len(in) == 0 {
		return nil
	}

	result := msgraph.ConditionalAccessLocations{}
	config := in[0].(map[string]interface{})

	includeLocations := config["included_locations"].([]interface{})
	excludeLocations := config["excluded_locations"].([]interface{})

	result.IncludeLocations = tf.ExpandStringSlicePtr(includeLocations)
	result.ExcludeLocations = tf.ExpandStringSlicePtr(excludeLocations)

	return &result
}

func expandConditionalAccessGrantControls(in []interface{}) *msgraph.ConditionalAccessGrantControls {
	if len(in) == 0 {
		return nil
	}

	result := msgraph.ConditionalAccessGrantControls{}
	config := in[0].(map[string]interface{})

	operator := config["operator"].(string)
	builtInControls := config["built_in_controls"].([]interface{})

	result.Operator = &operator
	result.BuiltInControls = tf.ExpandStringSlicePtr(builtInControls)

	return &result
}
