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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
)

func conditionalAccessPolicyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: conditionalAccessPolicyResourceCreate,
		ReadContext:   conditionalAccessPolicyResourceRead,
		UpdateContext: conditionalAccessPolicyResourceUpdate,
		DeleteContext: conditionalAccessPolicyResourceDelete,

		CustomizeDiff: conditionalAccessPolicyCustomizeDiff,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(15 * time.Minute),
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
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"included_applications": {
										Type:         schema.TypeList,
										Optional:     true,
										ExactlyOneOf: []string{"conditions.0.applications.0.included_applications", "conditions.0.applications.0.included_user_actions"},
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
										Type:         schema.TypeList,
										Optional:     true,
										ExactlyOneOf: []string{"conditions.0.applications.0.included_applications", "conditions.0.applications.0.included_user_actions"},
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},
								},
							},
						},

						"client_applications": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"included_service_principals": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validate.NoEmptyStrings,
										},
									},

									"excluded_service_principals": {
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
							Required: true,
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
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filter": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type:     schema.TypeString,
													Required: true,
													ValidateFunc: validation.StringInSlice([]string{
														msgraph.ConditionalAccessFilterModeExclude,
														msgraph.ConditionalAccessFilterModeInclude,
													}, false),
												},

												"rule": {
													Type:             schema.TypeString,
													Required:         true,
													ValidateDiagFunc: validate.NoEmptyStrings,
												},
											},
										},
									},
								},
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
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
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
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
				Type:         schema.TypeList,
				Optional:     true,
				AtLeastOneOf: []string{"grant_controls", "session_controls"},
				MaxItems:     1,
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
				Type:         schema.TypeList,
				Optional:     true,
				AtLeastOneOf: []string{"grant_controls", "session_controls"},
				MaxItems:     1,
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
								msgraph.ConditionalAccessCloudAppSecuritySessionControlTypeBlockDownloads,
								msgraph.ConditionalAccessCloudAppSecuritySessionControlTypeMcasConfigured,
								msgraph.ConditionalAccessCloudAppSecuritySessionControlTypeMonitorOnly,
								msgraph.ConditionalAccessCloudAppSecuritySessionControlTypeUnknownFutureValue,
							}, false),
						},

						"disable_resilience_defaults": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"persistent_browser_mode": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.PersistentBrowserSessionModeAlways,
								msgraph.PersistentBrowserSessionModeNever,
							}, false),
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

func conditionalAccessPolicyCustomizeDiff(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
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

func conditionalAccessPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func conditionalAccessPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	stateConf := &resource.StateChangeConf{ //nolint:staticcheck
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
