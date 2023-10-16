// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
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
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"state": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					msgraph.ConditionalAccessPolicyStateDisabled,
					msgraph.ConditionalAccessPolicyStateEnabled,
					msgraph.ConditionalAccessPolicyStateEnabledForReportingButNotEnforced,
				}, false),
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
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"excluded_applications": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"included_user_actions": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										ExactlyOneOf: []string{"conditions.0.applications.0.included_applications", "conditions.0.applications.0.included_user_actions"},
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
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
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"excluded_service_principals": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
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
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users"},
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"excluded_users": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"included_groups": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users"},
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"excluded_groups": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"included_roles": {
										Type:         pluginsdk.TypeList,
										Optional:     true,
										AtLeastOneOf: []string{"conditions.0.users.0.included_groups", "conditions.0.users.0.included_roles", "conditions.0.users.0.included_users"},
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"excluded_roles": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},
								},
							},
						},

						"client_app_types": {
							Type:     pluginsdk.TypeList,
							Required: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									msgraph.ConditionalAccessClientAppTypeAll,
									msgraph.ConditionalAccessClientAppTypeBrowser,
									msgraph.ConditionalAccessClientAppTypeEasSupported,
									msgraph.ConditionalAccessClientAppTypeExchangeActiveSync,
									msgraph.ConditionalAccessClientAppTypeMobileAppsAndDesktopClients,
									msgraph.ConditionalAccessClientAppTypeOther,
								}, false),
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
													Type:     pluginsdk.TypeString,
													Required: true,
													ValidateFunc: validation.StringInSlice([]string{
														msgraph.ConditionalAccessFilterModeExclude,
														msgraph.ConditionalAccessFilterModeInclude,
													}, false),
												},

												"rule": {
													Type:             pluginsdk.TypeString,
													Required:         true,
													ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
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
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
										},
									},

									"excluded_locations": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type:             pluginsdk.TypeString,
											ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
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
											Type: pluginsdk.TypeString,
											ValidateFunc: validation.StringInSlice([]string{
												msgraph.ConditionalAccessDevicePlatformAll,
												msgraph.ConditionalAccessDevicePlatformAndroid,
												msgraph.ConditionalAccessDevicePlatformIos,
												msgraph.ConditionalAccessDevicePlatformLinux,
												msgraph.ConditionalAccessDevicePlatformMacOs,
												msgraph.ConditionalAccessDevicePlatformUnknownFutureValue,
												msgraph.ConditionalAccessDevicePlatformWindows,
												msgraph.ConditionalAccessDevicePlatformWindowsPhone,
											}, false),
										},
									},

									"excluded_platforms": {
										Type:     pluginsdk.TypeList,
										Optional: true,
										Elem: &pluginsdk.Schema{
											Type: pluginsdk.TypeString,
											ValidateFunc: validation.StringInSlice([]string{
												msgraph.ConditionalAccessDevicePlatformAll,
												msgraph.ConditionalAccessDevicePlatformAndroid,
												msgraph.ConditionalAccessDevicePlatformIos,
												msgraph.ConditionalAccessDevicePlatformLinux,
												msgraph.ConditionalAccessDevicePlatformMacOs,
												msgraph.ConditionalAccessDevicePlatformUnknownFutureValue,
												msgraph.ConditionalAccessDevicePlatformWindows,
												msgraph.ConditionalAccessDevicePlatformWindowsPhone,
											}, false),
										},
									},
								},
							},
						},

						"service_principal_risk_levels": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									msgraph.ConditionalAccessRiskLevelHigh,
									msgraph.ConditionalAccessRiskLevelLow,
									msgraph.ConditionalAccessRiskLevelMedium,
									msgraph.ConditionalAccessRiskLevelNone,
									msgraph.ConditionalAccessRiskLevelUnknownFutureValue,
								}, false),
							},
						},

						"sign_in_risk_levels": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									msgraph.ConditionalAccessRiskLevelHidden,
									msgraph.ConditionalAccessRiskLevelHigh,
									msgraph.ConditionalAccessRiskLevelLow,
									msgraph.ConditionalAccessRiskLevelMedium,
									msgraph.ConditionalAccessRiskLevelNone,
									msgraph.ConditionalAccessRiskLevelUnknownFutureValue,
								}, false),
							},
						},

						"user_risk_levels": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									msgraph.ConditionalAccessRiskLevelHidden,
									msgraph.ConditionalAccessRiskLevelHigh,
									msgraph.ConditionalAccessRiskLevelLow,
									msgraph.ConditionalAccessRiskLevelMedium,
									msgraph.ConditionalAccessRiskLevelNone,
									msgraph.ConditionalAccessRiskLevelUnknownFutureValue,
								}, false),
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
							AtLeastOneOf: []string{"grant_controls.0.built_in_controls", "grant_controls.0.terms_of_use"},
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									msgraph.ConditionalAccessGrantControlApprovedApplication,
									msgraph.ConditionalAccessGrantControlBlock,
									msgraph.ConditionalAccessGrantControlCompliantApplication,
									msgraph.ConditionalAccessGrantControlCompliantDevice,
									msgraph.ConditionalAccessGrantControlDomainJoinedDevice,
									msgraph.ConditionalAccessGrantControlMfa,
									msgraph.ConditionalAccessGrantControlPasswordChange,
									msgraph.ConditionalAccessGrantControlUnknownFutureValue,
								}, false),
							},
						},

						"custom_authentication_factors": {
							Type:     pluginsdk.TypeList,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type:             pluginsdk.TypeString,
								ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
							},
						},

						"terms_of_use": {
							Type:         pluginsdk.TypeList,
							Optional:     true,
							AtLeastOneOf: []string{"grant_controls.0.built_in_controls", "grant_controls.0.terms_of_use"},
							Elem: &pluginsdk.Schema{
								Type:             pluginsdk.TypeString,
								ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
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
							Type:     pluginsdk.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.ConditionalAccessCloudAppSecuritySessionControlTypeBlockDownloads,
								msgraph.ConditionalAccessCloudAppSecuritySessionControlTypeMcasConfigured,
								msgraph.ConditionalAccessCloudAppSecuritySessionControlTypeMonitorOnly,
								msgraph.ConditionalAccessCloudAppSecuritySessionControlTypeUnknownFutureValue,
							}, false),
						},

						"disable_resilience_defaults": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},

						"persistent_browser_mode": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.PersistentBrowserSessionModeAlways,
								msgraph.PersistentBrowserSessionModeNever,
							}, false),
						},

						"sign_in_frequency": {
							Type:         pluginsdk.TypeInt,
							Optional:     true,
							RequiredWith: []string{"session_controls.0.sign_in_frequency_period"},
							ValidateFunc: validation.IntAtLeast(0),
						},

						"sign_in_frequency_period": {
							Type:         pluginsdk.TypeString,
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

func conditionalAccessPolicyCustomizeDiff(ctx context.Context, diff *pluginsdk.ResourceDiff, meta interface{}) error {
	// See https://github.com/microsoftgraph/msgraph-metadata/issues/93
	if old, new := diff.GetChange("session_controls.0.sign_in_frequency"); old.(int) > 0 && new.(int) == 0 {
		diff.ForceNew("session_controls.0.sign_in_frequency")
	}
	if old, new := diff.GetChange("session_controls.0.sign_in_frequency_period"); old.(string) != "" && new.(string) == "" {
		diff.ForceNew("session_controls.0.sign_in_frequency")
	}

	if old, new := diff.GetChange("conditions.0.devices.#"); old.(int) > 0 && new.(int) == 0 {
		diff.ForceNew("conditions.0.devices")
	}
	if old, new := diff.GetChange("conditions.0.devices.0.filter.#"); old.(int) > 0 && new.(int) == 0 {
		diff.ForceNew("conditions.0.devices.0.filter")
	}

	return nil
}

func conditionalAccessPolicyDiffSuppress(k, old, new string, d *pluginsdk.ResourceData) bool {
	suppress := false

	if k == "session_controls.#" && old == "0" && new == "1" {
		// When an ineffectual `session_controls` block is configured, the API just ignores it and returns
		// sessionControls: null
		sessionControlsRaw := d.Get("session_controls").([]interface{})
		if len(sessionControlsRaw) == 1 && sessionControlsRaw[0] != nil {
			sessionControls := sessionControlsRaw[0].(map[string]interface{})
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
			if v, ok := sessionControls["sign_in_frequency_period"]; ok && v.(string) != "" {
				suppress = false
			}
		}
	}

	return suppress
}

func conditionalAccessPolicyResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PoliciesClient

	properties := msgraph.ConditionalAccessPolicy{
		DisplayName: utils.String(d.Get("display_name").(string)),
		State:       utils.String(d.Get("state").(string)),
		Conditions:  expandConditionalAccessConditionSet(d.Get("conditions").([]interface{})),
	}

	if v, ok := d.GetOk("grant_controls"); ok {
		properties.GrantControls = expandConditionalAccessGrantControls(v.([]interface{}))
	}

	if v, ok := d.GetOk("session_controls"); ok {
		properties.SessionControls = expandConditionalAccessSessionControls(v.([]interface{}))
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

func conditionalAccessPolicyResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PoliciesClient

	properties := msgraph.ConditionalAccessPolicy{
		ID:          utils.String(d.Id()),
		DisplayName: utils.String(d.Get("display_name").(string)),
		State:       utils.String(d.Get("state").(string)),
		Conditions:  expandConditionalAccessConditionSet(d.Get("conditions").([]interface{})),
	}

	if v, ok := d.GetOk("grant_controls"); ok {
		properties.GrantControls = expandConditionalAccessGrantControls(v.([]interface{}))
	}

	if v, ok := d.GetOk("session_controls"); ok {
		properties.SessionControls = expandConditionalAccessSessionControls(v.([]interface{}))
	}

	if _, err := client.Update(ctx, properties); err != nil {
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
			client.BaseClient.DisableRetries = true
			policy, _, err := client.Get(ctx, d.Id(), odata.Query{})
			if err != nil {
				return nil, "Error", err
			}

			if policy == nil {
				return "stub", "Pending", nil
			}
			if policy.DisplayName == nil || *policy.DisplayName != d.Get("display_name").(string) {
				return "stub", "Pending", nil
			}
			if policy.State == nil || *policy.State != d.Get("state").(string) {
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

func conditionalAccessPolicyResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.PoliciesClient
	policyId := d.Id()

	_, status, err := client.Get(ctx, policyId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Conditional Access Policy with ID %q already deleted", policyId)
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving conditional access policy with ID %q", policyId)
	}

	status, err = client.Delete(ctx, policyId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting conditional access policy with ID %q, got status %d", policyId, status)
	}

	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, policyId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of conditional access policy with ID %q", policyId)
	}

	return nil
}
