// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/policies/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type RoleManagementPolicyModel struct {
	Description            string                                   `tfschema:"description"`
	DisplayName            string                                   `tfschema:"display_name"`
	GroupId                string                                   `tfschema:"object_id"`
	ScopeType              msgraph.UnifiedRoleManagementPolicyScope `tfschema:"assignment_type"`
	ActiveAssignmentRules  []ActiveAssignmentRules                  `tfschema:"active_assignment_rules"`
	EligbleAssignmentRules []EligibleAssignmentRules                `tfschema:"eligible_assignment_rules"`
	ActivationRules        []ActivationRules                        `tfschema:"activation_rules"`
	NotificationRules      []NotificationRules                      `tfschema:"notification_rules"`
}

type ActiveAssignmentRules struct {
	ExpirationRequired     bool   `tfschema:"expiration_required"`
	ExpireAfter            string `tfschema:"expire_after"`
	RequireMultiFactorAuth bool   `tfschema:"require_multifactor_authentication"`
	RequireJustification   bool   `tfschema:"require_justification"`
}

type EligibleAssignmentRules struct {
	ExpirationRequired bool   `tfschema:"expiration_required"`
	ExpireAfter        string `tfschema:"expire_after"`
}

type ActivationRules struct {
	MaximumDuration                 string          `tfschema:"maximum_duration"`
	RequireApproval                 bool            `tfschema:"require_approval"`
	ApprovalStages                  []ApprovalStage `tfschema:"approval_stages"`
	RequireConditionalAccessContext string          `tfschema:"require_conditional_access_authentication_context"`
	RequireMultiFactorAuth          bool            `tfschema:"require_multifactor_authentication"`
	RequireJustification            bool            `tfschema:"require_justification"`
	RequireTicketInfo               bool            `tfschema:"require_ticket_info"`
}

type ApprovalStage struct {
	PrimaryApprovers []Approver `tfschema:"primary_approvers"`
}

type Approver struct {
	Description string `tfschema:"description"`
	ObjectId    string `tfschema:"object_id"`
	ObjectType  string `tfschema:"object_type"`
}

type NotificationRules struct {
	AdminNotifications    []NotificationRule `tfschema:"admin_notifications"`
	ApproverNotifications []NotificationRule `tfschema:"approver_notifications"`
	AssigneeNotifications []NotificationRule `tfschema:"assignee_notifications"`
}

type NotificationRule struct {
	EligibleAssignments []NotificationSettings `tfschema:"eligible_assignments"`
	ActiveAssignments   []NotificationSettings `tfschema:"active_assignments"`
	Activations         []NotificationSettings `tfschema:"activations"`
}

type NotificationSettings struct {
	NotificationLevel    msgraph.UnifiedRoleManagementPolicyRuleNotificationLevel `tfschema:"notification_level"`
	DefaultRecipients    bool                                                     `tfschema:"default_recipients"`
	AdditionalRecipients []string                                                 `tfschema:"additional_recipients"`
}

type RoleManagementPolicyResource struct{}

func (r RoleManagementPolicyResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validation.IsUUID
}

var _ sdk.Resource = RoleManagementPolicyResource{}

func (r RoleManagementPolicyResource) ResourceType() string {
	return "azuread_group_role_management_policy"
}

func (r RoleManagementPolicyResource) ModelObject() interface{} {
	return &RoleManagementPolicyModel{}
}

func (r RoleManagementPolicyResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"object_id": {
			Description:      "ID of the group to which this policy is assigned",
			Type:             pluginsdk.TypeString,
			Required:         true,
			ForceNew:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
		},

		"assignment_type": {
			Description: "The ID of the assignment to the group",
			Type:        pluginsdk.TypeString,
			Required:    true,
			ForceNew:    true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
				msgraph.PrivilegedAccessGroupRelationshipMember,
				msgraph.PrivilegedAccessGroupRelationshipOwner,
				msgraph.PrivilegedAccessGroupRelationshipUnknown,
			}, false)),
		},

		"eligible_assignment_rules": {
			Description: "The rules for eligible assignment of the policy",
			Type:        pluginsdk.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"expiration_required": {
						Description: "Must the assignment have an expiry date",
						Type:        pluginsdk.TypeBool,
						Optional:    true,
					},

					"expire_after": {
						Description:      "The duration after which assignments expire",
						Type:             pluginsdk.TypeString,
						Optional:         true,
						ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"P15D", "P30D", "P90D", "P180D", "P365D"}, false)),
					},
				},
			},
		},

		"active_assignment_rules": {
			Description: "The rules for active assignment of the policy",
			Type:        pluginsdk.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"expiration_required": {
						Description: "Must the assignment have an expiry date",
						Type:        pluginsdk.TypeBool,
						Optional:    true,
					},

					"expire_after": {
						Description:      "The duration after which assignments expire",
						Type:             pluginsdk.TypeString,
						Optional:         true,
						ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"P15D", "P30D", "P90D", "P180D", "P365D"}, false)),
					},

					"require_multifactor_authentication": {
						Description: "Whether multi-factor authentication is required to make an assignment",
						Type:        pluginsdk.TypeBool,
						Optional:    true,
						Computed:    true,
					},

					"require_justification": {
						Description: "Whether a justification is required to make an assignment",
						Type:        pluginsdk.TypeBool,
						Optional:    true,
						Computed:    true,
					},
				},
			},
		},

		"activation_rules": {
			Description: "The activation rules of the policy",
			Type:        pluginsdk.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"maximum_duration": {
						Description: "The time after which the an activation can be valid for",
						Type:        pluginsdk.TypeString,
						Optional:    true,
						Computed:    true,
						ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
							"PT30M", "PT1H", "PT1H30M", "PT2H", "PT2H30M", "PT3H", "PT3H30M", "PT4H", "PT4H30M", "PT5H", "PT5H30M", "PT6H",
							"PT6H30M", "PT7H", "PT7H30M", "PT8H", "PT8H30M", "PT9H", "PT9H30M", "PT10H", "PT10H30M", "PT11H", "PT11H30M", "PT12H",
							"PT12H30M", "PT13H", "PT13H30M", "PT14H", "PT14H30M", "PT15H", "PT15H30M", "PT16H", "PT16H30M", "PT17H", "PT17H30M", "PT18H",
							"PT18H30M", "PT19H", "PT19H30M", "PT20H", "PT20H30M", "PT21H", "PT21H30M", "PT22H", "PT22H30M", "PT23H", "PT23H30M", "P1D",
						}, false)),
					},

					"require_approval": {
						Description: "Whether an approval is required for activation",
						Type:        pluginsdk.TypeBool,
						Optional:    true,
						Computed:    true,
					},

					"approval_stages": {
						Description: "The approval stages for the activation",
						Type:        pluginsdk.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"primary_approvers": {
									Description: "The IDs of the users or groups who can approve the activation",
									Type:        pluginsdk.TypeList,
									Required:    true,
									MinItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"description": {
												Description:      "The description of the approver",
												Type:             pluginsdk.TypeString,
												Required:         true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
											},

											"object_id": {
												Description:      "The ID of the useror group to act as an approver",
												Type:             pluginsdk.TypeString,
												Required:         true,
												ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
											},

											"object_type": {
												Description:      "The type of the object to act as an approver",
												Type:             pluginsdk.TypeString,
												Required:         true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"user", "group"}, false)),
											},
										},
									},
								},
							},
						},
					},

					"require_conditional_access_authentication_context": {
						Description:      "Whether a conditional access context is required during activation",
						Type:             pluginsdk.TypeString,
						Optional:         true,
						Computed:         true,
						ConflictsWith:    []string{"activation_rules.0.require_multifactor_authentication"},
						ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
					},

					"require_multifactor_authentication": {
						Description:   "Whether multi-factor authentication is required during activation",
						Type:          pluginsdk.TypeBool,
						Optional:      true,
						Computed:      true,
						ConflictsWith: []string{"activation_rules.0.require_conditional_access_authentication_context"},
					},

					"require_justification": {
						Description: "Whether a justification is required during activation",
						Type:        pluginsdk.TypeBool,
						Optional:    true,
						Computed:    true,
					},

					"require_ticket_info": {
						Description: "Whether ticket information is required during activation",
						Type:        pluginsdk.TypeBool,
						Optional:    true,
						Computed:    true,
					},
				},
			},
		},

		"notification_rules": {
			Description: "The notification rules of the policy",
			Type:        pluginsdk.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"admin_notifications": {
						Description: "The admin notifications on assignment",
						Type:        pluginsdk.TypeList,
						Optional:    true,
						Computed:    true,
						MaxItems:    1,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"eligible_assignments": {
									Description: "The admin notifications for eligible assignments",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description: "What level of notifications are sent",
												Type:        pluginsdk.TypeString,
												Optional:    true,
												Computed:    true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelAll,
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelCritical,
												}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
												},
											},
										},
									},
								},

								"active_assignments": {
									Description: "The admin notifications for active assignments",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description: "What level of notifications are sent",
												Type:        pluginsdk.TypeString,
												Optional:    true,
												Computed:    true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelAll,
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelCritical,
												}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
												},
											},
										},
									},
								},

								"activations": {
									Description: "The admin notifications for role activation",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description: "What level of notifications are sent",
												Type:        pluginsdk.TypeString,
												Optional:    true,
												Computed:    true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelAll,
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelCritical,
												}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
												},
											},
										},
									},
								},
							},
						},
					},

					"approver_notifications": {
						Description: "The admin notifications on assignment",
						Type:        pluginsdk.TypeList,
						Optional:    true,
						Computed:    true,
						MaxItems:    1,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"eligible_assignments": {
									Description: "The admin notifications for eligible assignments",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description: "What level of notifications are sent",
												Type:        pluginsdk.TypeString,
												Optional:    true,
												Computed:    true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelAll,
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelCritical,
												}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
												},
											},
										},
									},
								},

								"active_assignments": {
									Description: "The admin notifications for active assignments",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description: "What level of notifications are sent",
												Type:        pluginsdk.TypeString,
												Optional:    true,
												Computed:    true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelAll,
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelCritical,
												}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
												},
											},
										},
									},
								},

								"activations": {
									Description: "The admin notifications for role activation",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description: "What level of notifications are sent",
												Type:        pluginsdk.TypeString,
												Optional:    true,
												Computed:    true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelAll,
													msgraph.UnifiedRoleManagementPolicyRuleNotificationLevelCritical,
												}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
												},
											},
										},
									},
								},
							},
						},
					},

					"assignee_notifications": {
						Description: "The admin notifications on assignment",
						Type:        pluginsdk.TypeList,
						Optional:    true,
						Computed:    true,
						MaxItems:    1,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"eligible_assignments": {
									Description: "The admin notifications for eligible assignments",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description:      "What level of notifications are sent",
												Type:             pluginsdk.TypeString,
												Optional:         true,
												Computed:         true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
												},
											},
										},
									},
								},

								"active_assignments": {
									Description: "The admin notifications for active assignments",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description:      "What level of notifications are sent",
												Type:             pluginsdk.TypeString,
												Optional:         true,
												Computed:         true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
												},
											},
										},
									},
								},

								"activations": {
									Description: "The admin notifications for role activation",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Computed:    true,
									MaxItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"notification_level": {
												Description:      "What level of notifications are sent",
												Type:             pluginsdk.TypeString,
												Optional:         true,
												Computed:         true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
											},
											"default_recipients": {
												Description: "Whether the default recipients are notified",
												Type:        pluginsdk.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"additional_recipients": {
												Description: "The additional recipients to notify",
												Type:        pluginsdk.TypeList,
												Optional:    true,
												Elem: &pluginsdk.Schema{
													Type: pluginsdk.TypeString,
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
	}
}

func (r RoleManagementPolicyResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"display_name": {
			Description: "The display name of the policy",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"description": {
			Description: "Description of the policy",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (r RoleManagementPolicyResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			policyClient := metadata.Client.Policies.RoleManagementPolicyClient
			assignmentClient := metadata.Client.Policies.RoleManagementPolicyAssignmentClient

			// Fetch the existing policy, as they already exist
			policies, _, err := assignmentClient.List(ctx, odata.Query{
				Filter: fmt.Sprintf("scopeId eq '%s' and scopeType eq 'Group' and roleDefinitionId eq '%s'", metadata.ResourceData.Get("object_id").(string), metadata.ResourceData.Get("assignment_type").(string)),
			})
			if err != nil {
				return fmt.Errorf("Could not list existing policy, %+v", err)
			}
			if len(*policies) != 1 {
				return fmt.Errorf("Got the wrong number of policies, expected 1, got %d", len(*policies))
			}

			assignmentId, err := parse.ParseRoleManagementPolicyAssignmentID(*(*policies)[0].ID)
			if err != nil {
				return fmt.Errorf("Could not parse policy assignment ID, %+v", err)
			}

			id := parse.NewRoleManagementPolicyID(assignmentId.ScopeType, assignmentId.ScopeId, assignmentId.PolicyId)
			metadata.SetID(id)

			policy, _, err := policyClient.Get(ctx, id.ID())
			if err != nil {
				return fmt.Errorf("Could not retrieve existing policy, %+v", err)
			}
			if policy == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			policyUpdate, err := buildPolicyForUpdate(pointer.To(metadata), policy)
			if err != nil {
				return fmt.Errorf("Could not build update request, %+v", err)
			}

			_, err = policyClient.Update(ctx, *policyUpdate)
			if err != nil {
				return fmt.Errorf("Could not create assignment schedule request, %+v", err)
			}

			return nil
		},
	}
}

func (r RoleManagementPolicyResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.RoleManagementPolicyClient

			id, err := parse.ParseRoleManagementPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("Could not parse policy ID, %+v", err)
			}

			var model RoleManagementPolicyModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			result, _, err := client.Get(ctx, id.ID())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			model.Description = *result.Description
			model.DisplayName = *result.DisplayName
			model.GroupId = *result.ScopeId

			if len(model.EligbleAssignmentRules) == 0 {
				model.EligbleAssignmentRules = make([]EligibleAssignmentRules, 1)
			}
			if len(model.ActiveAssignmentRules) == 0 {
				model.ActiveAssignmentRules = make([]ActiveAssignmentRules, 1)
			}
			if len(model.ActivationRules) == 0 {
				model.ActivationRules = make([]ActivationRules, 1)
			}
			if len(model.NotificationRules) == 0 {
				model.NotificationRules = make([]NotificationRules, 1)
			}
			if len(model.NotificationRules[0].AdminNotifications) == 0 {
				model.NotificationRules[0].AdminNotifications = make([]NotificationRule, 1)
			}
			if len(model.NotificationRules[0].ApproverNotifications) == 0 {
				model.NotificationRules[0].ApproverNotifications = make([]NotificationRule, 1)
			}
			if len(model.NotificationRules[0].AssigneeNotifications) == 0 {
				model.NotificationRules[0].AssigneeNotifications = make([]NotificationRule, 1)
			}

			for _, rule := range *result.Rules {
				switch *rule.ID {
				case "Expiration_Admin_Eligibility":
					model.EligbleAssignmentRules[0].ExpirationRequired = *rule.IsExpirationRequired
					model.EligbleAssignmentRules[0].ExpireAfter = *rule.MaximumDuration

				case "Enablement_Admin_Assignment":
					model.ActiveAssignmentRules[0].RequireMultiFactorAuth = false
					model.ActiveAssignmentRules[0].RequireJustification = false
					for _, enabledRule := range *rule.EnabledRules {
						switch enabledRule {
						case "MultiFactorAuthentication":
							model.ActiveAssignmentRules[0].RequireMultiFactorAuth = true
						case "Justification":
							model.ActiveAssignmentRules[0].RequireJustification = true
						}
					}

				case "Expiration_Admin_Assignment":
					model.ActiveAssignmentRules[0].ExpirationRequired = *rule.IsExpirationRequired
					model.ActiveAssignmentRules[0].ExpireAfter = *rule.MaximumDuration

				case "Expiration_EndUser_Assignment":
					model.ActivationRules[0].MaximumDuration = *rule.MaximumDuration

				case "Approval_EndUser_Assignment":
					model.ActivationRules[0].RequireApproval = *rule.Setting.IsApprovalRequired
					model.ActivationRules[0].ApprovalStages = make([]ApprovalStage, 0)
					for _, stage := range *rule.Setting.ApprovalStages {
						primaryApprovers := make([]Approver, 0)
						for _, approver := range *stage.PrimaryApprovers {
							switch {
							case *approver.ODataType == "#microsoft.graph.singleUser":
								primaryApprovers = append(primaryApprovers, Approver{
									Description: *approver.Description,
									ObjectId:    *approver.ID,
									ObjectType:  "user",
								})
							case *approver.ODataType == "#microsoft.graph.groupMembers":
								primaryApprovers = append(primaryApprovers, Approver{
									Description: *approver.Description,
									ObjectId:    *approver.ID,
									ObjectType:  "group",
								})
							default:
								return fmt.Errorf("unknown approver type: %s", *approver.ODataType)
							}
							model.ActivationRules[0].ApprovalStages = append(model.ActivationRules[0].ApprovalStages, ApprovalStage{
								PrimaryApprovers: primaryApprovers,
							})
						}
					}

				case "AuthenticationContext_EndUser_Assignment":
					if rule.ClaimValue != nil && *rule.ClaimValue != "" {
						model.ActivationRules[0].RequireConditionalAccessContext = *rule.ClaimValue
					}

				case "Enablement_EndUser_Assignment":
					model.ActivationRules[0].RequireMultiFactorAuth = false
					model.ActivationRules[0].RequireJustification = false
					model.ActivationRules[0].RequireTicketInfo = false
					for _, enabledRule := range *rule.EnabledRules {
						switch enabledRule {
						case "MultiFactorAuthentication":
							model.ActivationRules[0].RequireMultiFactorAuth = true
						case "Justification":
							model.ActivationRules[0].RequireJustification = true
						case "Ticketing":
							model.ActivationRules[0].RequireTicketInfo = true
						}
					}

				case "Notification_Admin_Admin_Eligibility":
					if len(model.NotificationRules[0].AdminNotifications[0].EligibleAssignments) == 0 {
						model.NotificationRules[0].AdminNotifications[0].EligibleAssignments = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].AdditionalRecipients = *rule.NotificationRecipients

				case "Notification_Admin_Admin_Assignment":
					if len(model.NotificationRules[0].AdminNotifications[0].ActiveAssignments) == 0 {
						model.NotificationRules[0].AdminNotifications[0].ActiveAssignments = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].AdminNotifications[0].ActiveAssignments[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].AdminNotifications[0].ActiveAssignments[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].AdminNotifications[0].ActiveAssignments[0].AdditionalRecipients = *rule.NotificationRecipients

				case "Notification_Admin_EndUser_Assignment":
					if len(model.NotificationRules[0].AdminNotifications[0].Activations) == 0 {
						model.NotificationRules[0].AdminNotifications[0].Activations = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].AdminNotifications[0].Activations[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].AdminNotifications[0].Activations[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].AdminNotifications[0].Activations[0].AdditionalRecipients = *rule.NotificationRecipients

				case "Notification_Approver_Admin_Eligibility":
					if len(model.NotificationRules[0].ApproverNotifications[0].EligibleAssignments) == 0 {
						model.NotificationRules[0].ApproverNotifications[0].EligibleAssignments = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].ApproverNotifications[0].EligibleAssignments[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].ApproverNotifications[0].EligibleAssignments[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].ApproverNotifications[0].EligibleAssignments[0].AdditionalRecipients = *rule.NotificationRecipients

				case "Notification_Approver_Admin_Assignment":
					if len(model.NotificationRules[0].ApproverNotifications[0].ActiveAssignments) == 0 {
						model.NotificationRules[0].ApproverNotifications[0].ActiveAssignments = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].ApproverNotifications[0].ActiveAssignments[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].ApproverNotifications[0].ActiveAssignments[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].ApproverNotifications[0].ActiveAssignments[0].AdditionalRecipients = *rule.NotificationRecipients

				case "Notification_Approver_EndUser_Assignment":
					if len(model.NotificationRules[0].ApproverNotifications[0].Activations) == 0 {
						model.NotificationRules[0].ApproverNotifications[0].Activations = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].ApproverNotifications[0].Activations[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].ApproverNotifications[0].Activations[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].ApproverNotifications[0].Activations[0].AdditionalRecipients = *rule.NotificationRecipients

				case "Notification_Requestor_Admin_Eligibility":
					if len(model.NotificationRules[0].AssigneeNotifications[0].EligibleAssignments) == 0 {
						model.NotificationRules[0].AssigneeNotifications[0].EligibleAssignments = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].AssigneeNotifications[0].EligibleAssignments[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].AssigneeNotifications[0].EligibleAssignments[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].AssigneeNotifications[0].EligibleAssignments[0].AdditionalRecipients = *rule.NotificationRecipients

				case "Notification_Requestor_Admin_Assignment":
					if len(model.NotificationRules[0].AssigneeNotifications[0].ActiveAssignments) == 0 {
						model.NotificationRules[0].AssigneeNotifications[0].ActiveAssignments = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].AssigneeNotifications[0].ActiveAssignments[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].AssigneeNotifications[0].ActiveAssignments[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].AssigneeNotifications[0].ActiveAssignments[0].AdditionalRecipients = *rule.NotificationRecipients

				case "Notification_Requestor_EndUser_Assignment":
					if len(model.NotificationRules[0].AssigneeNotifications[0].Activations) == 0 {
						model.NotificationRules[0].AssigneeNotifications[0].Activations = make([]NotificationSettings, 1)
					}
					model.NotificationRules[0].AssigneeNotifications[0].Activations[0].NotificationLevel = rule.NotificationLevel
					model.NotificationRules[0].AssigneeNotifications[0].Activations[0].DefaultRecipients = *rule.IsDefaultRecipientsEnabled
					model.NotificationRules[0].AssigneeNotifications[0].Activations[0].AdditionalRecipients = *rule.NotificationRecipients

				}
			}

			return metadata.Encode(&model)
		},
	}
}

func (r RoleManagementPolicyResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.RoleManagementPolicyClient

			id, err := parse.ParseRoleManagementPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("Could not parse policy ID, %+v", err)
			}

			metadata.SetID(id)

			policy, _, err := client.Get(ctx, id.ID())
			if err != nil {
				return fmt.Errorf("Could not retrieve existing policy, %+v", err)
			}
			if policy == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			policyUpdate, err := buildPolicyForUpdate(pointer.To(metadata), policy)
			if err != nil {
				return fmt.Errorf("Could not build update request, %+v", err)
			}

			_, err = client.Update(ctx, *policyUpdate)
			if err != nil {
				return fmt.Errorf("Could not create assignment schedule request, %+v", err)
			}

			return nil
		},
	}
}

func (r RoleManagementPolicyResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			id, err := parse.ParseRoleManagementPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			return metadata.MarkAsGone(id)
		},
	}
}

func buildPolicyForUpdate(metadata *sdk.ResourceMetaData, policy *msgraph.UnifiedRoleManagementPolicy) (*msgraph.UnifiedRoleManagementPolicy, error) {
	var model RoleManagementPolicyModel
	if err := metadata.Decode(&model); err != nil {
		return nil, fmt.Errorf("decoding: %+v", err)
	}

	// Take the slice of rules and convert it to a map with the ID as the key
	policyRules := make(map[string]msgraph.UnifiedRoleManagementPolicyRule)
	for _, rule := range *policy.Rules {
		policyRules[*rule.ID] = rule
	}
	updatedRules := make([]msgraph.UnifiedRoleManagementPolicyRule, 0)

	if metadata.ResourceData.HasChange("eligible_assignment_rules") {
		expirationRequired := policyRules["Expiration_Admin_Eligibility"].IsExpirationRequired
		maximumDuration := policyRules["Expiration_Admin_Eligibility"].MaximumDuration

		if metadata.ResourceData.HasChange("eligible_assignment_rules.0.expiration_required") {
			expirationRequired = pointer.To(model.EligbleAssignmentRules[0].ExpirationRequired)
		}
		if metadata.ResourceData.HasChange("eligible_assignment_rules.0.expire_after") {
			maximumDuration = pointer.To(model.EligbleAssignmentRules[0].ExpireAfter)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                   policyRules["Expiration_Admin_Eligibility"].ID,
			ODataType:            policyRules["Expiration_Admin_Eligibility"].ODataType,
			Target:               policyRules["Expiration_Admin_Eligibility"].Target,
			IsExpirationRequired: expirationRequired,
			MaximumDuration:      maximumDuration,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("active_assignment_rules.0.require_multifactor_authentication") ||
		metadata.ResourceData.HasChange("active_assignment_rules.0.require_justification") {
		enabledRules := make([]string, 0)
		if model.ActiveAssignmentRules[0].RequireMultiFactorAuth {
			enabledRules = append(enabledRules, "MultiFactorAuthentication")
		}
		if model.ActiveAssignmentRules[0].RequireJustification {
			enabledRules = append(enabledRules, "Justification")
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:           policyRules["Enablement_Admin_Assignment"].ID,
			ODataType:    policyRules["Enablement_Admin_Assignment"].ODataType,
			Target:       policyRules["Enablement_Admin_Assignment"].Target,
			EnabledRules: pointer.To(enabledRules),
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("active_assignment_rules.0.expiration_required") ||
		metadata.ResourceData.HasChange("active_assignment_rules.0.expire_after") {
		expirationRequired := policyRules["Expiration_Admin_Assignment"].IsExpirationRequired
		maximumDuration := policyRules["Expiration_Admin_Assignment"].MaximumDuration

		if metadata.ResourceData.HasChange("active_assignment_rules.0.expiration_required") {
			expirationRequired = pointer.To(model.EligbleAssignmentRules[0].ExpirationRequired)
		}
		if metadata.ResourceData.HasChange("active_assignment_rules.0.expire_after") {
			maximumDuration = pointer.To(model.EligbleAssignmentRules[0].ExpireAfter)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                   policyRules["Expiration_Admin_Assignment"].ID,
			ODataType:            policyRules["Expiration_Admin_Assignment"].ODataType,
			Target:               policyRules["Expiration_Admin_Assignment"].Target,
			IsExpirationRequired: expirationRequired,
			MaximumDuration:      maximumDuration,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("activation_rules.0.maximum_duration") {
		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:              policyRules["Expiration_EndUser_Assignment"].ID,
			ODataType:       policyRules["Expiration_EndUser_Assignment"].ODataType,
			Target:          policyRules["Expiration_EndUser_Assignment"].Target,
			MaximumDuration: pointer.To(model.EligbleAssignmentRules[0].ExpireAfter),
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("activation_rules.0.require_approval") ||
		metadata.ResourceData.HasChange("activation_rules.0.approval_stages") {
		isApprovalRequired := policyRules["Approval_EndUser_Assignment"].Setting.IsApprovalRequired
		var approvalStages []msgraph.ApprovalStage
		if metadata.ResourceData.HasChange("activation_rules.0.require_approval") {
			isApprovalRequired = pointer.To(model.ActivationRules[0].RequireApproval)
		}
		if metadata.ResourceData.HasChange("activation_rules.0.approval_stages") {
			approvalStages = make([]msgraph.ApprovalStage, 0)
			for _, stage := range model.ActivationRules[0].ApprovalStages {
				primaryApprovers := make([]msgraph.UserSet, 0)
				for _, approver := range stage.PrimaryApprovers {
					if approver.ObjectType == "user" {
						primaryApprovers = append(primaryApprovers, msgraph.UserSet{
							ODataType:   pointer.To("#microsoft.graph.singleUser"),
							ID:          &approver.ObjectId,
							Description: &approver.Description,
						})
					} else if approver.ObjectType == "group" {
						primaryApprovers = append(primaryApprovers, msgraph.UserSet{
							ODataType:   pointer.To("#microsoft.graph.groupMembers"),
							ID:          &approver.ObjectId,
							Description: &approver.Description,
						})
					} else {
						return nil, fmt.Errorf("either user_id or group_id must be set")
					}
				}

				approvalStages = append(approvalStages, msgraph.ApprovalStage{
					PrimaryApprovers: &primaryApprovers,
				})
			}
		} else {
			approvalStages = *policyRules["Approval_EndUser_Assignment"].Setting.ApprovalStages
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:        policyRules["Approval_EndUser_Assignment"].ID,
			ODataType: policyRules["Approval_EndUser_Assignment"].ODataType,
			Target:    policyRules["Approval_EndUser_Assignment"].Target,
			Setting: pointer.To(msgraph.ApprovalSettings{
				IsApprovalRequired: isApprovalRequired,
				ApprovalStages:     &approvalStages,
			}),
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("activation_rules.0.require_conditional_access_authentication_context") {
		isEnabled := policyRules["AuthenticationContext_EndUser_Assignment"].IsEnabled
		claimValue := policyRules["AuthenticationContext_EndUser_Assignment"].ClaimValue

		if _, set := metadata.ResourceData.GetOk("activation_rules.0.require_conditional_access_authentication_context"); set {
			isEnabled = pointer.To(true)
			claimValue = pointer.To(model.ActivationRules[0].RequireConditionalAccessContext)
		} else {
			isEnabled = pointer.To(false)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:         policyRules["AuthenticationContext_EndUser_Assignment"].ID,
			ODataType:  policyRules["AuthenticationContext_EndUser_Assignment"].ODataType,
			Target:     policyRules["AuthenticationContext_EndUser_Assignment"].Target,
			IsEnabled:  isEnabled,
			ClaimValue: claimValue,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("activation_rules.0.require_multifactor_authentication") ||
		metadata.ResourceData.HasChange("activation_rules.0.require_justification") ||
		metadata.ResourceData.HasChange("activation_rules.0.require_ticket_info") {
		enabledRules := make([]string, 0)
		if model.ActivationRules[0].RequireMultiFactorAuth {
			enabledRules = append(enabledRules, "MultiFactorAuthentication")
		}
		if model.ActivationRules[0].RequireJustification {
			enabledRules = append(enabledRules, "Justification")
		}
		if model.ActivationRules[0].RequireTicketInfo {
			enabledRules = append(enabledRules, "Ticketing")
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:           policyRules["Enablement_EndUser_Assignment"].ID,
			ODataType:    policyRules["Enablement_EndUser_Assignment"].ODataType,
			Target:       policyRules["Enablement_EndUser_Assignment"].Target,
			EnabledRules: &enabledRules,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.eligible_assignments") {
		level := policyRules["Notification_Admin_Admin_Eligibility"].NotificationLevel
		defaultRecipients := policyRules["Notification_Admin_Admin_Eligibility"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Admin_Admin_Eligibility"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.eligible_assignments.0.notification_level") {
			level = model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.eligible_assignments.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.eligible_assignments.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Admin_Admin_Eligibility"].ID,
			ODataType:                  policyRules["Notification_Admin_Admin_Eligibility"].ODataType,
			Target:                     policyRules["Notification_Admin_Admin_Eligibility"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.active_assignments") {
		level := policyRules["Notification_Admin_Admin_Assignment"].NotificationLevel
		defaultRecipients := policyRules["Notification_Admin_Admin_Assignment"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Admin_Admin_Assignment"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.active_assignments.0.notification_level") {
			level = model.NotificationRules[0].AdminNotifications[0].ActiveAssignments[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.active_assignments.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].AdminNotifications[0].ActiveAssignments[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.active_assignments.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].AdminNotifications[0].ActiveAssignments[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Admin_Admin_Assignment"].ID,
			ODataType:                  policyRules["Notification_Admin_Admin_Assignment"].ODataType,
			Target:                     policyRules["Notification_Admin_Admin_Assignment"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.activations") {
		level := policyRules["Notification_Admin_EndUser_Assignment"].NotificationLevel
		defaultRecipients := policyRules["Notification_Admin_EndUser_Assignment"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Admin_EndUser_Assignment"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.activations.0.notification_level") {
			level = model.NotificationRules[0].AdminNotifications[0].Activations[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.activations.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].AdminNotifications[0].Activations[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.admin_notifications.0.activations.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].AdminNotifications[0].Activations[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Admin_EndUser_Assignment"].ID,
			ODataType:                  policyRules["Notification_Admin_EndUser_Assignment"].ODataType,
			Target:                     policyRules["Notification_Admin_EndUser_Assignment"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.eligible_assignments") {
		level := policyRules["Notification_Approver_Admin_Eligibility"].NotificationLevel
		defaultRecipients := policyRules["Notification_Approver_Admin_Eligibility"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Approver_Admin_Eligibility"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.eligible_assignments.0.notification_level") {
			level = model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.eligible_assignments.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.eligible_assignments.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].AdminNotifications[0].EligibleAssignments[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Approver_Admin_Eligibility"].ID,
			ODataType:                  policyRules["Notification_Approver_Admin_Eligibility"].ODataType,
			Target:                     policyRules["Notification_Approver_Admin_Eligibility"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.active_assignments") {
		level := policyRules["Notification_Approver_Admin_Assignment"].NotificationLevel
		defaultRecipients := policyRules["Notification_Approver_Admin_Assignment"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Approver_Admin_Assignment"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.active_assignments.0.notification_level") {
			level = model.NotificationRules[0].ApproverNotifications[0].ActiveAssignments[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.active_assignments.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].ApproverNotifications[0].ActiveAssignments[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.active_assignments.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].ApproverNotifications[0].ActiveAssignments[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Approver_Admin_Assignment"].ID,
			ODataType:                  policyRules["Notification_Approver_Admin_Assignment"].ODataType,
			Target:                     policyRules["Notification_Approver_Admin_Assignment"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.activations") {
		level := policyRules["Notification_Approver_EndUser_Assignment"].NotificationLevel
		defaultRecipients := policyRules["Notification_Approver_EndUser_Assignment"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Approver_EndUser_Assignment"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.activations.0.notification_level") {
			level = model.NotificationRules[0].ApproverNotifications[0].Activations[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.activations.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].ApproverNotifications[0].Activations[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.approver_notifications.0.activations.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].ApproverNotifications[0].Activations[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Approver_EndUser_Assignment"].ID,
			ODataType:                  policyRules["Notification_Approver_EndUser_Assignment"].ODataType,
			Target:                     policyRules["Notification_Approver_EndUser_Assignment"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.eligible_assignments") {
		level := policyRules["Notification_Requestor_Admin_Eligibility"].NotificationLevel
		defaultRecipients := policyRules["Notification_Requestor_Admin_Eligibility"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Requestor_Admin_Eligibility"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.eligible_assignments.0.notification_level") {
			level = model.NotificationRules[0].AssigneeNotifications[0].EligibleAssignments[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.eligible_assignments.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].AssigneeNotifications[0].EligibleAssignments[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.eligible_assignments.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].AssigneeNotifications[0].EligibleAssignments[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Requestor_Admin_Eligibility"].ID,
			ODataType:                  policyRules["Notification_Requestor_Admin_Eligibility"].ODataType,
			Target:                     policyRules["Notification_Requestor_Admin_Eligibility"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.active_assignments") {
		level := policyRules["Notification_Requestor_Admin_Assignment"].NotificationLevel
		defaultRecipients := policyRules["Notification_Requestor_Admin_Assignment"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Requestor_Admin_Assignment"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.active_assignments.0.notification_level") {
			level = model.NotificationRules[0].AssigneeNotifications[0].ActiveAssignments[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.active_assignments.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].AssigneeNotifications[0].ActiveAssignments[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.active_assignments.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].AssigneeNotifications[0].ActiveAssignments[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Requestor_Admin_Assignment"].ID,
			ODataType:                  policyRules["Notification_Requestor_Admin_Assignment"].ODataType,
			Target:                     policyRules["Notification_Requestor_Admin_Assignment"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.activations") {
		level := policyRules["Notification_Requestor_EndUser_Assignment"].NotificationLevel
		defaultRecipients := policyRules["Notification_Requestor_EndUser_Assignment"].IsDefaultRecipientsEnabled
		additionalRecipients := policyRules["Notification_Requestor_EndUser_Assignment"].NotificationRecipients

		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.activations.0.notification_level") {
			level = model.NotificationRules[0].AssigneeNotifications[0].Activations[0].NotificationLevel
		}
		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.activations.0.default_recipients") {
			defaultRecipients = pointer.To(model.NotificationRules[0].AssigneeNotifications[0].Activations[0].DefaultRecipients)
		}
		if metadata.ResourceData.HasChange("notification_rules.0.assignee_notifications.0.activations.0.additional_recipients") {
			additionalRecipients = pointer.To(model.NotificationRules[0].AssigneeNotifications[0].Activations[0].AdditionalRecipients)
		}

		rule := msgraph.UnifiedRoleManagementPolicyRule{
			ID:                         policyRules["Notification_Requestor_EndUser_Assignment"].ID,
			ODataType:                  policyRules["Notification_Requestor_EndUser_Assignment"].ODataType,
			Target:                     policyRules["Notification_Requestor_EndUser_Assignment"].Target,
			NotificationLevel:          level,
			IsDefaultRecipientsEnabled: defaultRecipients,
			NotificationRecipients:     additionalRecipients,
		}
		updatedRules = append(updatedRules, rule)
	}

	return &msgraph.UnifiedRoleManagementPolicy{
		ID:    policy.ID,
		Rules: pointer.To(updatedRules),
	}, nil
}
