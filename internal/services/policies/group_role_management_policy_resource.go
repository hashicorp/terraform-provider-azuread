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

type GroupRoleManagementPolicyModel struct {
	Description             string                                             `tfschema:"description"`
	DisplayName             string                                             `tfschema:"display_name"`
	GroupId                 string                                             `tfschema:"group_id"`
	ScopeType               msgraph.UnifiedRoleManagementPolicyScope           `tfschema:"assignment_type"`
	ActiveAssignmentRules   []GroupRoleManagementPolicyActiveAssignmentRules   `tfschema:"active_assignment_rules"`
	EligibleAssignmentRules []GroupRoleManagementPolicyEligibleAssignmentRules `tfschema:"eligible_assignment_rules"`
	ActivationRules         []GroupRoleManagementPolicyActivationRules         `tfschema:"activation_rules"`
	NotificationRules       []GroupRoleManagementPolicyNotificationEvents      `tfschema:"notification_rules"`
}

type GroupRoleManagementPolicyActiveAssignmentRules struct {
	ExpirationRequired     bool   `tfschema:"expiration_required"`
	ExpireAfter            string `tfschema:"expire_after"`
	RequireMultiFactorAuth bool   `tfschema:"require_multifactor_authentication"`
	RequireJustification   bool   `tfschema:"require_justification"`
	RequireTicketInfo      bool   `tfschema:"require_ticket_info"`
}

type GroupRoleManagementPolicyEligibleAssignmentRules struct {
	ExpirationRequired bool   `tfschema:"expiration_required"`
	ExpireAfter        string `tfschema:"expire_after"`
}

type GroupRoleManagementPolicyActivationRules struct {
	MaximumDuration                 string                                   `tfschema:"maximum_duration"`
	RequireApproval                 bool                                     `tfschema:"require_approval"`
	ApprovalStages                  []GroupRoleManagementPolicyApprovalStage `tfschema:"approval_stage"`
	RequireConditionalAccessContext string                                   `tfschema:"required_conditional_access_authentication_context"`
	RequireMultiFactorAuth          bool                                     `tfschema:"require_multifactor_authentication"`
	RequireJustification            bool                                     `tfschema:"require_justification"`
	RequireTicketInfo               bool                                     `tfschema:"require_ticket_info"`
}

type GroupRoleManagementPolicyApprovalStage struct {
	PrimaryApprovers []GroupRoleManagementPolicyApprover `tfschema:"primary_approver"`
}

type GroupRoleManagementPolicyApprover struct {
	ID   string `tfschema:"object_id"`
	Type string `tfschema:"type"`
}

type GroupRoleManagementPolicyNotificationEvents struct {
	ActiveAssignments   []GroupRoleManagementPolicyNotificationRule `tfschema:"active_assignments"`
	EligibleActivations []GroupRoleManagementPolicyNotificationRule `tfschema:"eligible_activations"`
	EligibleAssignments []GroupRoleManagementPolicyNotificationRule `tfschema:"eligible_assignments"`
}

type GroupRoleManagementPolicyNotificationRule struct {
	AdminNotifications    []GroupRoleManagementPolicyNotificationSettings `tfschema:"admin_notifications"`
	ApproverNotifications []GroupRoleManagementPolicyNotificationSettings `tfschema:"approver_notifications"`
	AssigneeNotifications []GroupRoleManagementPolicyNotificationSettings `tfschema:"assignee_notifications"`
}

type GroupRoleManagementPolicyNotificationSettings struct {
	NotificationLevel    msgraph.UnifiedRoleManagementPolicyRuleNotificationLevel `tfschema:"notification_level"`
	DefaultRecipients    bool                                                     `tfschema:"default_recipients"`
	AdditionalRecipients []string                                                 `tfschema:"additional_recipients"`
}

type GroupRoleManagementPolicyResource struct{}

func (r GroupRoleManagementPolicyResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateRoleManagementPolicyAssignmentID
}

var _ sdk.Resource = GroupRoleManagementPolicyResource{}

func (r GroupRoleManagementPolicyResource) ResourceType() string {
	return "azuread_group_role_management_policy"
}

func (r GroupRoleManagementPolicyResource) ModelObject() interface{} {
	return &GroupRoleManagementPolicyModel{}
}

func (r GroupRoleManagementPolicyResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"group_id": {
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
						Computed:    true,
					},

					"expire_after": {
						Description:      "The duration after which assignments expire",
						Type:             pluginsdk.TypeString,
						Optional:         true,
						Computed:         true,
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
						Computed:    true,
					},

					"expire_after": {
						Description:      "The duration after which assignments expire",
						Type:             pluginsdk.TypeString,
						Optional:         true,
						Computed:         true,
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

					"require_ticket_info": {
						Description: "Whether ticket information is required to make an assignment",
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

					"approval_stage": {
						Description: "The approval stages for the activation",
						Type:        pluginsdk.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"primary_approver": {
									Description: "The IDs of the users or groups who can approve the activation",
									Type:        pluginsdk.TypeSet,
									Required:    true,
									MinItems:    1,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"object_id": {
												Description:      "The ID of the object to act as an approver",
												Type:             pluginsdk.TypeString,
												Required:         true,
												ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
											},

											"type": {
												Description:      "The type of object acting as an approver",
												Type:             pluginsdk.TypeString,
												Optional:         true,
												ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"singleUser", "groupMembers"}, false)),
											},
										},
									},
								},
							},
						},
					},

					"required_conditional_access_authentication_context": {
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
						ConflictsWith: []string{"activation_rules.0.required_conditional_access_authentication_context"},
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
					"active_assignments": {
						Description: "Notifications about active assignments",
						Type:        pluginsdk.TypeList,
						Optional:    true,
						Computed:    true,
						MaxItems:    1,
						Elem: &pluginsdk.Resource{
							Schema: notificationRuleSchema(),
						},
					},
					"eligible_activations": {
						Description: "Notifications about activations of eligible assignments",
						Type:        pluginsdk.TypeList,
						Optional:    true,
						Computed:    true,
						MaxItems:    1,
						Elem: &pluginsdk.Resource{
							Schema: notificationRuleSchema(),
						},
					},
					"eligible_assignments": {
						Description: "Notifications about eligible assignments",
						Type:        pluginsdk.TypeList,
						Optional:    true,
						Computed:    true,
						MaxItems:    1,
						Elem: &pluginsdk.Resource{
							Schema: notificationRuleSchema(),
						},
					},
				},
			},
		},
	}
}

func (r GroupRoleManagementPolicyResource) Attributes() map[string]*pluginsdk.Schema {
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

func (r GroupRoleManagementPolicyResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			policyClient := metadata.Client.Policies.RoleManagementPolicyClient
			assignmentClient := metadata.Client.Policies.RoleManagementPolicyAssignmentClient

			// Fetch the existing policy, as they already exist
			policies, _, err := assignmentClient.List(ctx, odata.Query{
				Filter: fmt.Sprintf("scopeId eq '%s' and scopeType eq 'Group' and roleDefinitionId eq '%s'", metadata.ResourceData.Get("group_id").(string), metadata.ResourceData.Get("assignment_type").(string)),
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

func (r GroupRoleManagementPolicyResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.RoleManagementPolicyClient

			id, err := parse.ParseRoleManagementPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("Could not parse policy ID, %+v", err)
			}

			var model GroupRoleManagementPolicyModel
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

			if len(model.EligibleAssignmentRules) == 0 {
				model.EligibleAssignmentRules = make([]GroupRoleManagementPolicyEligibleAssignmentRules, 1)
			}
			if len(model.ActiveAssignmentRules) == 0 {
				model.ActiveAssignmentRules = make([]GroupRoleManagementPolicyActiveAssignmentRules, 1)
			}
			if len(model.ActivationRules) == 0 {
				model.ActivationRules = make([]GroupRoleManagementPolicyActivationRules, 1)
			}
			if len(model.NotificationRules) == 0 {
				model.NotificationRules = make([]GroupRoleManagementPolicyNotificationEvents, 1)
			}
			if len(model.NotificationRules[0].EligibleActivations) == 0 {
				model.NotificationRules[0].EligibleActivations = make([]GroupRoleManagementPolicyNotificationRule, 1)
			}
			if len(model.NotificationRules[0].ActiveAssignments) == 0 {
				model.NotificationRules[0].ActiveAssignments = make([]GroupRoleManagementPolicyNotificationRule, 1)
			}
			if len(model.NotificationRules[0].EligibleAssignments) == 0 {
				model.NotificationRules[0].EligibleAssignments = make([]GroupRoleManagementPolicyNotificationRule, 1)
			}

			for _, rule := range *result.Rules {
				switch *rule.ID {
				case "Approval_EndUser_Assignment":
					model.ActivationRules[0].RequireApproval = *rule.Setting.IsApprovalRequired

					primaryApprovers := make([]GroupRoleManagementPolicyApprover, 0)
					for _, approver := range *(*rule.Setting.ApprovalStages)[0].PrimaryApprovers {
						switch {
						case *approver.ODataType == "#microsoft.graph.singleUser":
							primaryApprovers = append(primaryApprovers, GroupRoleManagementPolicyApprover{
								ID:   pointer.ToString(approver.UserID),
								Type: "singleUser",
							})
						case *approver.ODataType == "#microsoft.graph.groupMembers":
							primaryApprovers = append(primaryApprovers, GroupRoleManagementPolicyApprover{
								ID:   pointer.ToString(approver.GroupID),
								Type: "groupMembers",
							})
						default:
							return fmt.Errorf("unknown approver type: %s", *approver.ODataType)
						}
					}
					model.ActivationRules[0].ApprovalStages = []GroupRoleManagementPolicyApprovalStage{{PrimaryApprovers: primaryApprovers}}

				case "AuthenticationContext_EndUser_Assignment":
					if rule.ClaimValue != nil && *rule.ClaimValue != "" {
						model.ActivationRules[0].RequireConditionalAccessContext = *rule.ClaimValue
					}

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

				case "Expiration_Admin_Eligibility":
					model.EligibleAssignmentRules[0].ExpirationRequired = *rule.IsExpirationRequired
					model.EligibleAssignmentRules[0].ExpireAfter = *rule.MaximumDuration

				case "Expiration_Admin_Assignment":
					model.ActiveAssignmentRules[0].ExpirationRequired = *rule.IsExpirationRequired
					model.ActiveAssignmentRules[0].ExpireAfter = *rule.MaximumDuration

				case "Expiration_EndUser_Assignment":
					model.ActivationRules[0].MaximumDuration = *rule.MaximumDuration

				case "Notification_Admin_Admin_Assignment":
					model.NotificationRules[0].ActiveAssignments[0].AdminNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}

				case "Notification_Admin_Admin_Eligibility":
					model.NotificationRules[0].EligibleAssignments[0].AdminNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}

				case "Notification_Admin_EndUser_Assignment":
					model.NotificationRules[0].EligibleActivations[0].AdminNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}

				case "Notification_Approver_Admin_Assignment":
					model.NotificationRules[0].ActiveAssignments[0].ApproverNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}

				case "Notification_Approver_Admin_Eligibility":
					model.NotificationRules[0].EligibleAssignments[0].ApproverNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}

				case "Notification_Approver_EndUser_Assignment":
					model.NotificationRules[0].EligibleActivations[0].ApproverNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}

				case "Notification_Requestor_Admin_Assignment":
					model.NotificationRules[0].ActiveAssignments[0].AssigneeNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}

				case "Notification_Requestor_Admin_Eligibility":
					model.NotificationRules[0].EligibleAssignments[0].AssigneeNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}

				case "Notification_Requestor_EndUser_Assignment":
					model.NotificationRules[0].EligibleActivations[0].AssigneeNotifications = []GroupRoleManagementPolicyNotificationSettings{
						*flattenNotificationSettings(pointer.To(rule)),
					}
				}
			}

			return metadata.Encode(&model)
		},
	}
}

func (r GroupRoleManagementPolicyResource) Update() sdk.ResourceFunc {
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

func (r GroupRoleManagementPolicyResource) Delete() sdk.ResourceFunc {
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
	var model GroupRoleManagementPolicyModel
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

		if *expirationRequired != model.EligibleAssignmentRules[0].ExpirationRequired {
			expirationRequired = pointer.To(model.EligibleAssignmentRules[0].ExpirationRequired)
		}
		if *maximumDuration != model.EligibleAssignmentRules[0].ExpireAfter &&
			model.EligibleAssignmentRules[0].ExpireAfter != "" {
			maximumDuration = pointer.To(model.EligibleAssignmentRules[0].ExpireAfter)
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
		if model.ActiveAssignmentRules[0].RequireTicketInfo {
			enabledRules = append(enabledRules, "Ticketing")
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

		if *expirationRequired != model.ActiveAssignmentRules[0].ExpirationRequired {
			expirationRequired = pointer.To(model.ActiveAssignmentRules[0].ExpirationRequired)
		}
		if *maximumDuration != model.ActiveAssignmentRules[0].ExpireAfter &&
			model.ActiveAssignmentRules[0].ExpireAfter != "" {
			maximumDuration = pointer.To(model.ActiveAssignmentRules[0].ExpireAfter)
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
			MaximumDuration: pointer.To(model.ActivationRules[0].MaximumDuration),
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("activation_rules.0.require_approval") ||
		metadata.ResourceData.HasChange("activation_rules.0.approval_stage") {
		if model.ActivationRules[0].RequireApproval && len(model.ActivationRules[0].ApprovalStages) != 1 {
			return nil, fmt.Errorf("require_approval is true, but no approval_stages are provided")
		}

		isApprovalRequired := policyRules["Approval_EndUser_Assignment"].Setting.IsApprovalRequired
		var approvalStages []msgraph.ApprovalStage
		if *isApprovalRequired != model.ActivationRules[0].RequireApproval {
			isApprovalRequired = pointer.To(model.ActivationRules[0].RequireApproval)
		}
		if metadata.ResourceData.HasChange("activation_rules.0.approval_stage") {
			approvalStages = make([]msgraph.ApprovalStage, 0)
			for _, stage := range model.ActivationRules[0].ApprovalStages {
				primaryApprovers := make([]msgraph.UserSet, 0)
				for _, approver := range stage.PrimaryApprovers {
					switch approver.Type {
					case "singleUser":
						primaryApprovers = append(primaryApprovers, msgraph.UserSet{
							ODataType: pointer.To("#microsoft.graph.singleUser"),
							UserID:    pointer.To(approver.ID),
						})
					case "groupMembers":
						primaryApprovers = append(primaryApprovers, msgraph.UserSet{
							ODataType: pointer.To("#microsoft.graph.groupMembers"),
							GroupID:   pointer.To(approver.ID),
						})
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

	if metadata.ResourceData.HasChange("activation_rules.0.required_conditional_access_authentication_context") {
		isEnabled := policyRules["AuthenticationContext_EndUser_Assignment"].IsEnabled
		claimValue := policyRules["AuthenticationContext_EndUser_Assignment"].ClaimValue

		if _, set := metadata.ResourceData.GetOk("activation_rules.0.required_conditional_access_authentication_context"); set {
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

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.admin_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Admin_Admin_Eligibility"],
				model.NotificationRules[0].EligibleAssignments[0].AdminNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.admin_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.admin_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Admin_Admin_Assignment"],
				model.NotificationRules[0].ActiveAssignments[0].AdminNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.admin_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.admin_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Admin_EndUser_Assignment"],
				model.NotificationRules[0].EligibleActivations[0].AdminNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.admin_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.approver_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Approver_Admin_Eligibility"],
				model.NotificationRules[0].EligibleAssignments[0].ApproverNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.approver_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.approver_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Approver_Admin_Assignment"],
				model.NotificationRules[0].ActiveAssignments[0].ApproverNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.approver_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.approver_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Approver_EndUser_Assignment"],
				model.NotificationRules[0].EligibleActivations[0].ApproverNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.approver_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.assignee_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Requestor_Admin_Eligibility"],
				model.NotificationRules[0].EligibleAssignments[0].AssigneeNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.assignee_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.assignee_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Requestor_Admin_Assignment"],
				model.NotificationRules[0].ActiveAssignments[0].AssigneeNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.assignee_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.assignee_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Requestor_EndUser_Assignment"],
				model.NotificationRules[0].EligibleActivations[0].AssigneeNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.assignee_notifications.0.additional_recipients"),
			),
		)
	}

	return &msgraph.UnifiedRoleManagementPolicy{
		ID:    policy.ID,
		Rules: pointer.To(updatedRules),
	}, nil
}

func expandNotificationSettings(rule msgraph.UnifiedRoleManagementPolicyRule, data GroupRoleManagementPolicyNotificationSettings, recipientChange bool) msgraph.UnifiedRoleManagementPolicyRule {
	level := rule.NotificationLevel
	defaultRecipients := rule.IsDefaultRecipientsEnabled
	additionalRecipients := rule.NotificationRecipients

	if level != data.NotificationLevel {
		level = data.NotificationLevel
	}
	if *defaultRecipients != data.DefaultRecipients {
		defaultRecipients = pointer.To(data.DefaultRecipients)
	}
	if recipientChange {
		additionalRecipients = pointer.To(data.AdditionalRecipients)
	}

	return msgraph.UnifiedRoleManagementPolicyRule{
		ID:                         rule.ID,
		ODataType:                  rule.ODataType,
		Target:                     rule.Target,
		RecipientType:              rule.RecipientType,
		NotificationType:           rule.NotificationType,
		NotificationLevel:          level,
		IsDefaultRecipientsEnabled: defaultRecipients,
		NotificationRecipients:     additionalRecipients,
	}
}

func flattenNotificationSettings(rule *msgraph.UnifiedRoleManagementPolicyRule) *GroupRoleManagementPolicyNotificationSettings {
	return &GroupRoleManagementPolicyNotificationSettings{
		NotificationLevel:    rule.NotificationLevel,
		DefaultRecipients:    *rule.IsDefaultRecipientsEnabled,
		AdditionalRecipients: *rule.NotificationRecipients,
	}
}

func notificationRuleSchema() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"admin_notifications": {
			Description: "Admin notification settings",
			Type:        pluginsdk.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem: &pluginsdk.Resource{
				Schema: notificationSettingsSchema(),
			},
		},
		"approver_notifications": {
			Description: "Approver notification settings",
			Type:        pluginsdk.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem: &pluginsdk.Resource{
				Schema: notificationSettingsSchema(),
			},
		},
		"assignee_notifications": {
			Description: "Assignee notification settings",
			Type:        pluginsdk.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem: &pluginsdk.Resource{
				Schema: notificationSettingsSchema(),
			},
		},
	}
}

func notificationSettingsSchema() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"notification_level": {
			Description:      "What level of notifications are sent",
			Type:             pluginsdk.TypeString,
			Required:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
		},
		"default_recipients": {
			Description: "Whether the default recipients are notified",
			Type:        pluginsdk.TypeBool,
			Required:    true,
		},
		"additional_recipients": {
			Description: "The additional recipients to notify",
			Type:        pluginsdk.TypeSet,
			Optional:    true,
			Computed:    true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},
	}
}
