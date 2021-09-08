package conditionalaccess

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
)

func conditionalAccessPolicyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: conditionalAccessPolicyResourceCreate,
		ReadContext:   conditionalAccessPolicyResourceRead,
		UpdateContext: conditionalAccessPolicyResourceUpdate,
		DeleteContext: conditionalAccessPolicyResourceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"state": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.ConditionalAccessPolicyStateDisabled,
					msgraph.ConditionalAccessPolicyStateEnabled,
					msgraph.ConditionalAccessPolicyStateEnabledForReportingButNotEnforced,
				}, false),
			},

			"conditions": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"applications": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"included_applications": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
									"excluded_applications": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
									"included_user_actions": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
								},
							},
						},
						"users": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"included_users": {
										Type:         schema.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users"},
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
									"excluded_users": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
									"included_groups": {
										Type:         schema.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users"},
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
									"excluded_groups": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
									"included_roles": {
										Type:         schema.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users"},
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
									"excluded_roles": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
								},
							},
						},
						"client_app_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									"all",
									"browser",
									"mobileAppsAndDesktopClients",
									"exchangeActiveSync",
									"easSupported",
									"other",
								}, false),
							},
						},
						"locations": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"included_locations": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
									"excluded_locations": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
								},
							},
						},
						"platforms": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"included_platforms": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
											ValidateFunc: validation.StringInSlice([]string{
												"all",
												"android",
												"iOS",
												"macOS",
												"unknownFutureValue",
												"windows",
												"windowsPhone",
											}, false),
										},
									},
									"excluded_platforms": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
											ValidateFunc: validation.StringInSlice([]string{
												"all",
												"android",
												"iOS",
												"macOS",
												"unknownFutureValue",
												"windows",
												"windowsPhone",
											}, false),
										},
									},
								},
							},
						},
						"sign_in_risk_levels": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									"hidden",
									"high",
									"low",
									"medium",
									"none",
									"unknownFutureValue",
								}, false),
							},
						},
						"user_risk_levels": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									"hidden",
									"high",
									"low",
									"medium",
									"none",
									"unknownFutureValue",
								}, false),
							},
						},
					},
				},
			},
			"grant_controls": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operator": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"AND", "OR"}, false),
						},

						"built_in_controls": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									"approvedApplication",
									"block",
									"compliantApplication",
									"compliantDevice",
									"domainJoinedDevice",
									"mfa",
									"passwordChange",
									"unknownFutureValue",
								}, false),
							},
						},
						"custom_authentication_factors": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.NoEmptyStrings,
							},
						},
						"terms_of_use": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type:             schema.TypeString,
								ValidateDiagFunc: validate.NoEmptyStrings,
							},
						},
					},
				},
			},
			"session_controls": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_enforced_restrictions_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"cloud_app_security_policy": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"blockDownloads",
								"mcasConfigured",
								"monitorOnly",
								"unknownFutureValue",
							}, false),
						},
						"persistent_browser_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"always", "never"}, false),
						},
						"sign_in_frequency": {
							Type:         schema.TypeInt,
							Optional:     true,
							RequiredWith: []string{"session_controls.0.sign_in_frequency_period"},
							ValidateFunc: validation.IntAtLeast(0),
						},
						"sign_in_frequency_period": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: []string{"session_controls.0.sign_in_frequency"},
							ValidateFunc: validation.StringInSlice([]string{"days", "hours"}, false),
						},
					},
				},
			},
		},
	}
}

func conditionalAccessPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PoliciesClient

	properties := msgraph.ConditionalAccessPolicy{
		DisplayName:     utils.String(d.Get("display_name").(string)),
		State:           utils.String(d.Get("state").(string)),
		Conditions:      expandConditionalAccessConditionSet(d.Get("conditions").([]interface{})),
		GrantControls:   expandConditionalAccessGrantControls(d.Get("grant_controls").([]interface{})),
		SessionControls: expandConditionalAccessSessionControls(d.Get("session_controls").([]interface{})),
	}

	policy, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create conditional access policy")
	}

	if policy.ID == nil || *policy.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for conditional access policy is nil/empty")
	}

	d.SetId(*policy.ID)

	return conditionalAccessPolicyResourceRead(ctx, d, meta)
}

func conditionalAccessPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PoliciesClient

	properties := msgraph.ConditionalAccessPolicy{
		ID:              utils.String(d.Id()),
		DisplayName:     utils.String(d.Get("display_name").(string)),
		State:           utils.String(d.Get("state").(string)),
		Conditions:      expandConditionalAccessConditionSet(d.Get("conditions").([]interface{})),
		GrantControls:   expandConditionalAccessGrantControls(d.Get("grant_controls").([]interface{})),
		SessionControls: expandConditionalAccessSessionControls(d.Get("session_controls").([]interface{})),
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update conditional access policy with ID: %q", d.Id())
	}

	return nil
}

func conditionalAccessPolicyResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PoliciesClient

	policy, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Conditional Access Policy with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Conditional Access Policy with object ID %q", d.Id())
	}

	tf.Set(d, "display_name", policy.DisplayName)
	tf.Set(d, "state", policy.State)
	tf.Set(d, "conditions", flattenConditionalAccessConditionSet(policy.Conditions))
	tf.Set(d, "grant_controls", flattenConditionalAccessGrantControls(policy.GrantControls))
	tf.Set(d, "session_controls", flattenConditionalAccessSessionControls(policy.SessionControls))

	return nil
}

func conditionalAccessPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PoliciesClient

	_, status, err := client.Get(ctx, d.Id(), odata.Query{})
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

func flattenConditionalAccessConditionSet(in *msgraph.ConditionalAccessConditionSet) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"applications":        flattenConditionalAccessApplications(in.Applications),
			"users":               flattenConditionalAccessUsers(in.Users),
			"client_app_types":    tf.FlattenStringSlicePtr(in.ClientAppTypes),
			"locations":           flattenConditionalAccessLocations(in.Locations),
			"platforms":           flattenConditionalAccessPlatforms(in.Platforms),
			"sign_in_risk_levels": tf.FlattenStringSlicePtr(in.SignInRiskLevels),
			"user_risk_levels":    tf.FlattenStringSlicePtr(in.UserRiskLevels),
		},
	}
}

func flattenConditionalAccessApplications(in *msgraph.ConditionalAccessApplications) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_applications": tf.FlattenStringSlicePtr(in.IncludeApplications),
			"excluded_applications": tf.FlattenStringSlicePtr(in.ExcludeApplications),
			"included_user_actions": tf.FlattenStringSlicePtr(in.IncludeUserActions),
		},
	}
}

func flattenConditionalAccessUsers(in *msgraph.ConditionalAccessUsers) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_users":  tf.FlattenStringSlicePtr(in.IncludeUsers),
			"excluded_users":  tf.FlattenStringSlicePtr(in.ExcludeUsers),
			"included_groups": tf.FlattenStringSlicePtr(in.IncludeGroups),
			"excluded_groups": tf.FlattenStringSlicePtr(in.ExcludeGroups),
			"included_roles":  tf.FlattenStringSlicePtr(in.IncludeRoles),
			"excluded_roles":  tf.FlattenStringSlicePtr(in.ExcludeRoles),
		},
	}
}

func flattenConditionalAccessLocations(in *msgraph.ConditionalAccessLocations) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_locations": tf.FlattenStringSlicePtr(in.IncludeLocations),
			"excluded_locations": tf.FlattenStringSlicePtr(in.ExcludeLocations),
		},
	}
}

func flattenConditionalAccessPlatforms(in *msgraph.ConditionalAccessPlatforms) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_platforms": tf.FlattenStringSlicePtr(in.IncludePlatforms),
			"excluded_platforms": tf.FlattenStringSlicePtr(in.ExcludePlatforms),
		},
	}
}

func flattenConditionalAccessGrantControls(in *msgraph.ConditionalAccessGrantControls) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"operator":                      in.Operator,
			"built_in_controls":             tf.FlattenStringSlicePtr(in.BuiltInControls),
			"custom_authentication_factors": tf.FlattenStringSlicePtr(in.CustomAuthenticationFactors),
			"terms_of_use":                  tf.FlattenStringSlicePtr(in.TermsOfUse),
		},
	}
}

func flattenConditionalAccessSessionControls(in *msgraph.ConditionalAccessSessionControls) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"application_enforced_restrictions_enabled": in.ApplicationEnforcedRestrictions.IsEnabled,
			"cloud_app_security_policy":                 in.CloudAppSecurity.CloudAppSecurityType,
			"persistent_browser_mode":                   in.PersistentBrowser.Mode,
			"sign_in_frequency":                         in.SignInFrequency.Value,
			"sign_in_frequency_period":                  in.SignInFrequency.Type,
		},
	}
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
	platforms := config["platforms"].([]interface{})
	signInRiskLevels := config["sign_in_risk_levels"].([]interface{})
	userRiskLevels := config["user_risk_levels"].([]interface{})

	result.Applications = expandConditionalAccessApplications(applications)
	result.Users = expandConditionalAccessUsers(users)
	result.ClientAppTypes = tf.ExpandStringSlicePtr(clientAppTypes)
	result.Locations = expandConditionalAccessLocations(locations)
	result.Platforms = expandConditionalAccessPlatforms(platforms)
	result.SignInRiskLevels = tf.ExpandStringSlicePtr(signInRiskLevels)
	result.UserRiskLevels = tf.ExpandStringSlicePtr(userRiskLevels)

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

func expandConditionalAccessPlatforms(in []interface{}) *msgraph.ConditionalAccessPlatforms {
	if len(in) == 0 {
		return nil
	}

	result := msgraph.ConditionalAccessPlatforms{}
	config := in[0].(map[string]interface{})

	includePlatforms := config["included_platforms"].([]interface{})
	excludePlatforms := config["excluded_platforms"].([]interface{})

	result.IncludePlatforms = tf.ExpandStringSlicePtr(includePlatforms)
	result.ExcludePlatforms = tf.ExpandStringSlicePtr(excludePlatforms)

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
	customAuthenticationFactors := config["custom_authentication_factors"].([]interface{})
	termsOfUse := config["terms_of_use"].([]interface{})

	result.Operator = &operator
	result.BuiltInControls = tf.ExpandStringSlicePtr(builtInControls)
	result.CustomAuthenticationFactors = tf.ExpandStringSlicePtr(customAuthenticationFactors)
	result.TermsOfUse = tf.ExpandStringSlicePtr(termsOfUse)

	return &result
}

func expandConditionalAccessSessionControls(in []interface{}) *msgraph.ConditionalAccessSessionControls {
	if len(in) == 0 {
		return nil
	}

	result := msgraph.ConditionalAccessSessionControls{}
	config := in[0].(map[string]interface{})

	if applicationEnforcedRestrictions := config["application_enforced_restrictions_enabled"]; applicationEnforcedRestrictions != nil {
		result.ApplicationEnforcedRestrictions = &msgraph.ApplicationEnforcedRestrictionsSessionControl{
			IsEnabled: utils.Bool(applicationEnforcedRestrictions.(bool)),
		}
	}

	if cloudAppSecurity := config["cloud_app_security_policy"]; cloudAppSecurity != nil {
		result.CloudAppSecurity = &msgraph.CloudAppSecurityControl{
			IsEnabled:            utils.Bool(true),
			CloudAppSecurityType: utils.String(cloudAppSecurity.(string)),
		}
	}

	if persistentBrowser := config["persistent_browser_mode"]; persistentBrowser != nil {
		result.PersistentBrowser = &msgraph.PersistentBrowserSessionControl{
			IsEnabled: utils.Bool(true),
			Mode:      utils.String(persistentBrowser.(string)),
		}
	}

	if signInFrequency := config["sign_in_frequency"]; signInFrequency != nil {
		result.SignInFrequency = &msgraph.SignInFrequencySessionControl{
			IsEnabled: utils.Bool(true),
			Type:      utils.String(config["sign_in_frequency_period"].(string)),
			Value:     utils.Int32(int32(signInFrequency.(int))),
		}
	}

	return &result
}
