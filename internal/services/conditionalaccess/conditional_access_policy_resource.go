// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesspolicy"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/conditionalaccess/migrations"
)

func conditionalAccessPolicyResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: conditionalAccessPolicyResourceCreate,
		ReadContext:   conditionalAccessPolicyResourceRead,
		UpdateContext: conditionalAccessPolicyResourceUpdate,
		DeleteContext: conditionalAccessPolicyResourceDelete,

		CustomizeDiff: conditionalAccessPolicyCustomizeDiff,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(15 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, errs := stable.ValidateIdentityConditionalAccessPolicyID(id, "id"); len(errs) > 0 {
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
				Type:    migrations.ResourceConditionalAccessPolicyInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceConditionalAccessPolicyInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"state": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessPolicyState(), false),
			},

			"conditions": {
				Type:     pluginsdk.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"applications": {
							Type:     pluginsdk.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"included_applications": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										ExactlyOneOf: []string{"conditions.0.applications.0.included_applications", "conditions.0.applications.0.included_user_actions"},
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"excluded_applications": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"included_user_actions": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										ExactlyOneOf: []string{"conditions.0.applications.0.included_applications", "conditions.0.applications.0.included_user_actions"},
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},
								},
							},
						},

						"client_applications": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"included_service_principals": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"excluded_service_principals": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},
								},
							},
						},

						"users": {
							Type:     pluginsdk.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"included_users": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users", "conditions.0.users.0.included_guests_or_external_users"},
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"excluded_users": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"included_groups": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users", "conditions.0.users.0.included_guests_or_external_users"},
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"excluded_groups": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"included_roles": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users", "conditions.0.users.0.included_guests_or_external_users"},
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"excluded_roles": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"included_guests_or_external_users": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users", "conditions.0.users.0.included_guests_or_external_users"},
										Elem: &pluginsdk.Resource{
											Schema: map[string]*pluginsdk.Schema{
												"guest_or_external_user_types": {
													Type:     pluginsdk.TypeList,
													Required: true,
													Elem: &pluginsdk.Schema{
														Type:         pluginsdk.TypeString,
														ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessGuestOrExternalUserTypes(), false),
													},
												},

												"external_tenants": {
													Type:     pluginsdk.TypeList,
													Optional: true,
													Elem: &pluginsdk.Resource{
														Schema: map[string]*pluginsdk.Schema{
															"membership_kind": {
																Type:         pluginsdk.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessExternalTenantsMembershipKind(), false),
															},

															"members": {
																Type:     pluginsdk.TypeList,
																Optional: true,
																Elem: &pluginsdk.Schema{
																	Type:         pluginsdk.TypeString,
																	ValidateFunc: validation.StringIsNotEmpty,
																},
															},
														},
													},
												},
											},
										},
									},

									"excluded_guests_or_external_users": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Resource{
											Schema: map[string]*pluginsdk.Schema{
												"guest_or_external_user_types": {
													Type:     pluginsdk.TypeList,
													Required: true,
													Elem: &pluginsdk.Schema{
														Type:         pluginsdk.TypeString,
														ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessGuestOrExternalUserTypes(), false),
													},
												},

												"external_tenants": {
													Type:     pluginsdk.TypeList,
													Optional: true,
													Elem: &pluginsdk.Resource{
														Schema: map[string]*pluginsdk.Schema{
															"membership_kind": {
																Type:         pluginsdk.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessExternalTenantsMembershipKind(), false),
															},

															"members": {
																Type:     pluginsdk.TypeList,
																Optional: true,
																Elem: &pluginsdk.Schema{
																	Type:         pluginsdk.TypeString,
																	ValidateFunc: validation.StringIsNotEmpty,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},

						"client_app_types": {
							Type:     pluginsdk.TypeList,
							Required: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessClientApp(), false),
							},
						},

						"devices": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"filter": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &pluginsdk.Resource{
											Schema: map[string]*pluginsdk.Schema{
												"mode": {
													Type:         pluginsdk.TypeString,
													Required:     true,
													ValidateFunc: validation.StringInSlice(stable.PossibleValuesForFilterMode(), false),
												},

												"rule": {
													Type:         pluginsdk.TypeString,
													Required:     true,
													ValidateFunc: validation.StringIsNotEmpty,
												},
											},
										},
									},
								},
							},
						},

						"locations": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"included_locations": {
										Type:     pluginsdk.TypeList,
										Required: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},

									"excluded_locations": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},
								},
							},
						},

						"platforms": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"included_platforms": {
										Type:     pluginsdk.TypeList,
										Required: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessDevicePlatform(), false),
										},
									},

									"excluded_platforms": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:         pluginsdk.TypeString,
											ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessDevicePlatform(), false),
										},
									},
								},
							},
						},

						"service_principal_risk_levels": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice(stable.PossibleValuesForRiskLevel(), false),
							},
						},

						"sign_in_risk_levels": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice(stable.PossibleValuesForRiskLevel(), false),
							},
						},

						"user_risk_levels": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice(stable.PossibleValuesForRiskLevel(), false),
							},
						},
					},
				},
			},

			"grant_controls": {
				Type:         pluginsdk.TypeList,
				Optional:     true,
				AtLeastOneOf: []string{"grant_controls", "session_controls"},
				MaxItems:     1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"operator": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"AND", "OR"}, false),
						},

						"built_in_controls": {
							Type:         pluginsdk.TypeList,
							Optional:     true,
							AtLeastOneOf: []string{"grant_controls.0.built_in_controls", "grant_controls.0.authentication_strength_policy_id", "grant_controls.0.terms_of_use"},
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice(stable.PossibleValuesForConditionalAccessGrantControl(), false),
							},
						},

						"authentication_strength_policy_id": {
							AtLeastOneOf: []string{"grant_controls.0.built_in_controls", "grant_controls.0.authentication_strength_policy_id", "grant_controls.0.terms_of_use"},
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: stable.ValidatePolicyAuthenticationStrengthPolicyID,
						},

						"custom_authentication_factors": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},

						"terms_of_use": {
							Type:         pluginsdk.TypeList,
							Optional:     true,
							AtLeastOneOf: []string{"grant_controls.0.built_in_controls", "grant_controls.0.authentication_strength_policy_id", "grant_controls.0.terms_of_use"},
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},
					},
				},
			},

			"session_controls": {
				Type:             pluginsdk.TypeList,
				Optional:         true,
				AtLeastOneOf:     []string{"grant_controls", "session_controls"},
				MaxItems:         1,
				DiffSuppressFunc: conditionalAccessPolicyDiffSuppress,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"application_enforced_restrictions_enabled": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"cloud_app_security_policy": {
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(stable.PossibleValuesForCloudAppSecuritySessionControlType(), false),
						},

						"disable_resilience_defaults": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"persistent_browser_mode": {
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(stable.PossibleValuesForPersistentBrowserSessionMode(), false),
						},

						"sign_in_frequency": {
							Type:         pluginsdk.TypeInt,
							Optional:     true,
							RequiredWith: []string{"session_controls.0.sign_in_frequency_period"},
							ValidateFunc: validation.IntAtLeast(0),
						},

						"sign_in_frequency_authentication_type": {
							Type:         pluginsdk.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringInSlice(stable.PossibleValuesForSignInFrequencyAuthenticationType(), false),
						},

						"sign_in_frequency_interval": {
							Type:         pluginsdk.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringInSlice(stable.PossibleValuesForSignInFrequencyInterval(), false),
						},

						"sign_in_frequency_period": {
							Type:         pluginsdk.TypeString,
							Optional:     true,
							RequiredWith: []string{"session_controls.0.sign_in_frequency"},
							ValidateFunc: validation.StringInSlice(stable.PossibleValuesForSigninFrequencyType(), false),
						},
					},
				},
			},

			"object_id": {
				Description: "The object ID of the policy",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func conditionalAccessPolicyCustomizeDiff(_ context.Context, diff *pluginsdk.ResourceDiff, _ interface{}) error {
	// The API does not like sessionControls being set with ineffectual properties, so this additional validation complements
	// AtLeastOneOf: []string{"grant_controls", "session_controls"} by helping to ensure that either `grant_controls` or a
	// _useful_ `session_controls` block has been set in the configuration.
	var sessionControlsSetButIneffective bool
	if diff.Get("session_controls.#").(int) == 1 && !diff.Get("session_controls.0.application_enforced_restrictions_enabled").(bool) &&
		diff.Get("session_controls.0.cloud_app_security_policy").(string) == "" && !diff.Get("session_controls.0.disable_resilience_defaults").(bool) &&
		diff.Get("session_controls.0.persistent_browser_mode").(string) == "" && diff.Get("session_controls.0.sign_in_frequency").(int) == 0 &&
		diff.Get("session_controls.0.sign_in_frequency_authentication_type").(string) == string(stable.SignInFrequencyAuthenticationType_PrimaryAndSecondaryAuthentication) &&
		diff.Get("session_controls.0.sign_in_frequency_interval").(string) == string(stable.SignInFrequencyInterval_TimeBased) {
		sessionControlsSetButIneffective = true
	}
	if diff.Get("grant_controls.#").(int) == 0 && sessionControlsSetButIneffective {
		return fmt.Errorf("when specifying `session_controls` but not `grant_controls`, one of the properties in the `session_controls` block must be set to an effective value in order for session controls to work")
	}

	return nil
}

func conditionalAccessPolicyDiffSuppress(k, old, new string, d *pluginsdk.ResourceData) bool {
	suppress := false

	// When ineffectual `session_controls` are specified, you must send `sessionControls: null`, and when policy has ineffectual
	// `sessionControls`, the API condenses it to `sessionControls: null` in the response.
	if k == "session_controls.#" && old == "0" && new == "1" {
		sessionControlsRaw := d.Get("session_controls").([]interface{})
		if len(sessionControlsRaw) == 1 && sessionControlsRaw[0] != nil {
			sessionControls := sessionControlsRaw[0].(map[string]interface{})

			// Suppress by default, but only if all the block properties have a non-default value
			suppress = true
			if v, ok := sessionControls["application_enforced_restrictions_enabled"]; ok && v.(bool) {
				suppress = false
			}
			if v, ok := sessionControls["cloud_app_security_policy"]; ok && v.(string) != "" {
				suppress = false
			}
			if v, ok := sessionControls["disable_resilience_defaults"]; ok && v.(bool) {
				suppress = false
			}
			if v, ok := sessionControls["persistent_browser_mode"]; ok && v.(string) != "" {
				suppress = false
			}
			if v, ok := sessionControls["sign_in_frequency"]; ok && v.(int) > 0 {
				suppress = false
			}
			if v, ok := sessionControls["sign_in_frequency_authentication_type"]; ok && v.(string) != "" {
				suppress = false
			}
			if v, ok := sessionControls["sign_in_frequency_interval"]; ok && v.(string) != "" {
				suppress = false
			}
			if v, ok := sessionControls["sign_in_frequency_period"]; ok && v.(string) != "" {
				suppress = false
			}
		}
	}

	return suppress
}

func conditionalAccessPolicyResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PolicyClient

	var err error

	var grantControls *stable.ConditionalAccessGrantControls
	if v, ok := d.GetOk("grant_controls"); ok {
		grantControls, err = expandConditionalAccessGrantControls(v.([]interface{}))
		if err != nil {
			return tf.ErrorDiagPathF(err, "grant_controls", "Parsing `grant_controls`")
		}
	}

	var sessionControls *stable.ConditionalAccessSessionControls
	if v, ok := d.GetOk("session_controls"); ok {
		sessionControls = expandConditionalAccessSessionControls(v.([]interface{}))
	}

	properties := stable.ConditionalAccessPolicy{
		DisplayName:     pointer.To(d.Get("display_name").(string)),
		State:           pointer.To(stable.ConditionalAccessPolicyState(d.Get("state").(string))),
		Conditions:      expandConditionalAccessConditionSet(d.Get("conditions").([]interface{})),
		GrantControls:   grantControls,
		SessionControls: sessionControls,
	}

	resp, err := client.CreateConditionalAccessPolicy(ctx, properties, conditionalaccesspolicy.DefaultCreateConditionalAccessPolicyOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create conditional access policy")
	}

	policy := resp.Model
	if policy == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Could not create conditional access policy")
	}

	if policy.Id == nil || *policy.Id == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for conditional access policy is nil/empty")
	}

	id := stable.NewIdentityConditionalAccessPolicyID(pointer.From(policy.Id))
	d.SetId(id.ID())

	return conditionalAccessPolicyResourceRead(ctx, d, meta)
}

func conditionalAccessPolicyResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PolicyClient

	id, err := stable.ParseIdentityConditionalAccessPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Conditional Access Policy ID")
	}

	var grantControls *stable.ConditionalAccessGrantControls
	if v, ok := d.GetOk("grant_controls"); ok {
		grantControls, err = expandConditionalAccessGrantControls(v.([]interface{}))
		if err != nil {
			return tf.ErrorDiagPathF(err, "grant_controls", "Parsing `grant_controls`")
		}
	}

	var sessionControls *stable.ConditionalAccessSessionControls
	if v, ok := d.GetOk("session_controls"); ok {
		sessionControls = expandConditionalAccessSessionControls(v.([]interface{}))
	}

	properties := stable.ConditionalAccessPolicy{
		DisplayName:     pointer.To(d.Get("display_name").(string)),
		State:           pointer.To(stable.ConditionalAccessPolicyState(d.Get("state").(string))),
		Conditions:      expandConditionalAccessConditionSet(d.Get("conditions").([]interface{})),
		GrantControls:   grantControls,
		SessionControls: sessionControls,
	}

	if _, err := client.UpdateConditionalAccessPolicy(ctx, *id, properties, conditionalaccesspolicy.DefaultUpdateConditionalAccessPolicyOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Could not update conditional access policy with ID: %q", d.Id())
	}

	// Poll for 5 retrievals of the updated policy. We don't check every property as this is prone to getting stuck
	// in a timeout loop, instead we're hoping that this allows enough time/activity for the update to be reflected.
	log.Printf("[DEBUG] Waiting for conditional access policy %q to be updated", d.Id())
	timeout, _ := ctx.Deadline()
	stateConf := &pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Pending"},
		Target:                    []string{"Done"},
		Timeout:                   time.Until(timeout),
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.GetConditionalAccessPolicy(ctx, *id, conditionalaccesspolicy.DefaultGetConditionalAccessPolicyOperationOptions())
			if err != nil {
				return nil, "Error", err
			}

			policy := resp.Model
			if policy == nil {
				return "stub", "Pending", nil
			}
			if policy.DisplayName == nil || *policy.DisplayName != d.Get("display_name").(string) {
				return "stub", "Pending", nil
			}
			if policy.State == nil || *policy.State != stable.ConditionalAccessPolicyState(d.Get("state").(string)) {
				return "stub", "Pending", nil
			}

			return "stub", "Done", nil
		},
	}
	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		return tf.ErrorDiagF(err, "waiting for update of conditional access policy with ID %q", d.Id())
	}

	return nil
}

func conditionalAccessPolicyResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PolicyClient

	id, err := stable.ParseIdentityConditionalAccessPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Conditional Access Policy ID")
	}

	resp, err := client.GetConditionalAccessPolicy(ctx, *id, conditionalaccesspolicy.DefaultGetConditionalAccessPolicyOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s not found - removing from state", id)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "retrieving %s", id)
	}

	policy := resp.Model
	if policy == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "retrieving %s", id)
	}

	tf.Set(d, "object_id", pointer.From(policy.Id))
	tf.Set(d, "display_name", pointer.From(policy.DisplayName))
	tf.Set(d, "state", pointer.From(policy.State))
	tf.Set(d, "conditions", flattenConditionalAccessConditionSet(policy.Conditions))
	tf.Set(d, "grant_controls", flattenConditionalAccessGrantControls(policy.GrantControls))
	tf.Set(d, "session_controls", flattenConditionalAccessSessionControls(policy.SessionControls))

	return nil
}

func conditionalAccessPolicyResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PolicyClient

	id, err := stable.ParseIdentityConditionalAccessPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Conditional Access Policy ID")
	}

	resp, err := client.GetConditionalAccessPolicy(ctx, *id, conditionalaccesspolicy.DefaultGetConditionalAccessPolicyOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s already deleted", id)
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "retrieving %s", id)
	}

	if _, err = client.DeleteConditionalAccessPolicy(ctx, *id, conditionalaccesspolicy.DefaultDeleteConditionalAccessPolicyOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetConditionalAccessPolicy(ctx, *id, conditionalaccesspolicy.DefaultGetConditionalAccessPolicyOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of %s", id)
	}

	return nil
}
