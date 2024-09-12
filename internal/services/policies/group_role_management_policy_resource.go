// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/policies/parse"
)

type GroupRoleManagementPolicyModel struct {
	Description             string                                             `tfschema:"description"`
	DisplayName             string                                             `tfschema:"display_name"`
	GroupId                 string                                             `tfschema:"group_id"`
	RoleId                  string                                             `tfschema:"role_id"`
	ActiveAssignmentRules   []GroupRoleManagementPolicyActiveAssignmentRules   `tfschema:"active_assignment_rules"`
	EligibleAssignmentRules []GroupRoleManagementPolicyEligibleAssignmentRules `tfschema:"eligible_assignment_rules"`
	ActivationRules         []GroupRoleManagementPolicyActivationRules         `tfschema:"activation_rules"`
	NotificationRules       []GroupRoleManagementPolicyNotificationEvents      `tfschema:"notification_rules"`
}

type GroupRoleManagementPolicyActiveAssignmentRules struct {
	ExpirationRequired     bool   `tfschema:"expiration_required"`
	ExpireAfter            string `tfschema:"expire_after"`
	RequireJustification   bool   `tfschema:"require_justification"`
	RequireMultiFactorAuth bool   `tfschema:"require_multifactor_authentication"`
	RequireTicketInfo      bool   `tfschema:"require_ticket_info"`
}

type GroupRoleManagementPolicyEligibleAssignmentRules struct {
	ExpirationRequired bool   `tfschema:"expiration_required"`
	ExpireAfter        string `tfschema:"expire_after"`
}

type GroupRoleManagementPolicyActivationRules struct {
	ApprovalStages                  []GroupRoleManagementPolicyApprovalStage `tfschema:"approval_stage"`
	MaximumDuration                 string                                   `tfschema:"maximum_duration"`
	RequireApproval                 bool                                     `tfschema:"require_approval"`
	RequireConditionalAccessContext string                                   `tfschema:"required_conditional_access_authentication_context"`
	RequireJustification            bool                                     `tfschema:"require_justification"`
	RequireMultiFactorAuth          bool                                     `tfschema:"require_multifactor_authentication"`
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
	AdditionalRecipients []string `tfschema:"additional_recipients"`
	DefaultRecipients    bool     `tfschema:"default_recipients"`
	NotificationLevel    string   `tfschema:"notification_level"`
}

var _ sdk.ResourceWithUpdate = GroupRoleManagementPolicyResource{}

type GroupRoleManagementPolicyResource struct{}

func (r GroupRoleManagementPolicyResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateRoleManagementPolicyID
}

func (r GroupRoleManagementPolicyResource) ResourceType() string {
	return "azuread_group_role_management_policy"
}

func (r GroupRoleManagementPolicyResource) ModelObject() interface{} {
	return &GroupRoleManagementPolicyModel{}
}

func (r GroupRoleManagementPolicyResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"group_id": {
			Description:  "ID of the group to which this policy is assigned",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"role_id": {
			Description:  "The ID of the role of this policy to the group",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice(possibleValuesForRoleDefinitionId, false),
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
						Description:  "The duration after which assignments expire",
						Type:         pluginsdk.TypeString,
						Optional:     true,
						Computed:     true,
						ValidateFunc: validation.StringInSlice([]string{"P15D", "P30D", "P90D", "P180D", "P365D"}, false),
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
						Description:  "The duration after which assignments expire",
						Type:         pluginsdk.TypeString,
						Optional:     true,
						Computed:     true,
						ValidateFunc: validation.StringInSlice([]string{"P15D", "P30D", "P90D", "P180D", "P365D"}, false),
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
						ValidateFunc: validation.StringInSlice([]string{
							"PT30M", "PT1H", "PT1H30M", "PT2H", "PT2H30M", "PT3H", "PT3H30M", "PT4H", "PT4H30M", "PT5H", "PT5H30M", "PT6H",
							"PT6H30M", "PT7H", "PT7H30M", "PT8H", "PT8H30M", "PT9H", "PT9H30M", "PT10H", "PT10H30M", "PT11H", "PT11H30M", "PT12H",
							"PT12H30M", "PT13H", "PT13H30M", "PT14H", "PT14H30M", "PT15H", "PT15H30M", "PT16H", "PT16H30M", "PT17H", "PT17H30M", "PT18H",
							"PT18H30M", "PT19H", "PT19H30M", "PT20H", "PT20H30M", "PT21H", "PT21H30M", "PT22H", "PT22H30M", "PT23H", "PT23H30M", "P1D",
						}, false),
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
												Description:  "The ID of the object to act as an approver",
												Type:         pluginsdk.TypeString,
												Required:     true,
												ValidateFunc: validation.IsUUID,
											},

											"type": {
												Description:  "The type of object acting as an approver",
												Type:         pluginsdk.TypeString,
												Optional:     true,
												ValidateFunc: validation.StringInSlice([]string{"singleUser", "groupMembers"}, false),
											},
										},
									},
								},
							},
						},
					},

					"required_conditional_access_authentication_context": {
						Description:   "Whether a conditional access context is required during activation",
						Type:          pluginsdk.TypeString,
						Optional:      true,
						Computed:      true,
						ConflictsWith: []string{"activation_rules.0.require_multifactor_authentication"},
						ValidateFunc:  validation.StringIsNotEmpty,
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
			client := metadata.Client.Policies.RoleManagementPolicyClient

			var model GroupRoleManagementPolicyModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// Fetch the existing policy, as they already exist
			policyId, err := getPolicyId(ctx, metadata, model.GroupId, model.RoleId)
			if err != nil {
				return fmt.Errorf("parsing policy ID: %v", err)
			}

			id := stable.NewPolicyRoleManagementPolicyID(policyId.ID())

			resp, err := client.GetRoleManagementPolicy(ctx, id, rolemanagementpolicy.DefaultGetRoleManagementPolicyOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving existing %s: %v", id, err)
			}

			roleManagementPolicy := resp.Model
			if roleManagementPolicy == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			policyUpdate, err := buildPolicyForUpdate(pointer.To(metadata), roleManagementPolicy)
			if err != nil {
				return fmt.Errorf("building update request: %v", err)
			}

			if _, err = client.UpdateRoleManagementPolicy(ctx, id, *policyUpdate); err != nil {
				return fmt.Errorf("creating %s: %v", id, err)
			}

			// Update the ID as it changes on modification
			policyId, err = getPolicyId(ctx, metadata, model.GroupId, model.RoleId)
			if err != nil {
				return fmt.Errorf("parsing policy ID: %v", err)
			}

			metadata.SetID(policyId)

			return nil
		},
	}
}

func (r GroupRoleManagementPolicyResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.RoleManagementPolicyClient
			assignmentClient := metadata.Client.Policies.RoleManagementPolicyAssignmentClient

			policyId, err := parse.ParseRoleManagementPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("parsing policy ID: %v", err)
			}

			var model GroupRoleManagementPolicyModel
			if err = metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id := stable.NewPolicyRoleManagementPolicyID(policyId.ID())

			policyResp, err := client.GetRoleManagementPolicy(ctx, id, rolemanagementpolicy.DefaultGetRoleManagementPolicyOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %v", id, err)
			}

			policy := policyResp.Model
			if policy == nil {
				return fmt.Errorf("retrieving %s: API error, model was nil", id)
			}

			options := rolemanagementpolicyassignment.ListRoleManagementPolicyAssignmentsOperationOptions{
				Filter: pointer.To(fmt.Sprintf("scopeType eq 'Group' and scopeId eq '%s' and policyId eq '%s'", policyId.ScopeId, id.UnifiedRoleManagementPolicyId)),
			}
			resp, err := assignmentClient.ListRoleManagementPolicyAssignments(ctx, options)
			if err != nil {
				return fmt.Errorf("retrieving %s: %v", id, err)
			}

			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: expected 1 assignment, got nil result", id)
			}
			if len(*resp.Model) != 1 {
				return fmt.Errorf("retrieving %s: expected 1 assignment, got %d", id, len(*resp.Model))
			}

			assignment := (*resp.Model)[0]

			model.Description = pointer.From(policy.Description)
			model.DisplayName = pointer.From(policy.DisplayName)
			model.GroupId = policy.ScopeId
			model.RoleId = assignment.RoleDefinitionId.GetOrZero()

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

			if policy.Rules != nil {
				for _, rule := range *policy.Rules {
					switch pointer.From(rule.UnifiedRoleManagementPolicyRule().Id) {
					case "Approval_EndUser_Assignment":
						model.ActivationRules[0].RequireApproval = rule.(stable.UnifiedRoleManagementPolicyApprovalRule).Setting.IsApprovalRequired.GetOrZero()
						primaryApprovers := make([]GroupRoleManagementPolicyApprover, 0)

						if rule.(stable.UnifiedRoleManagementPolicyApprovalRule).Setting != nil && rule.(stable.UnifiedRoleManagementPolicyApprovalRule).Setting.ApprovalStages != nil {
							if approvers := (*rule.(stable.UnifiedRoleManagementPolicyApprovalRule).Setting.ApprovalStages)[0].PrimaryApprovers; approvers != nil {
								for _, approver := range *approvers {
									switch {
									case pointer.From(approver.SubjectSet().ODataType) == "#microsoft.graph.singleUser":
										primaryApprovers = append(primaryApprovers, GroupRoleManagementPolicyApprover{
											ID:   approver.(stable.SingleUser).UserId.GetOrZero(),
											Type: "singleUser",
										})
									case pointer.From(approver.SubjectSet().ODataType) == "#microsoft.graph.groupMembers":
										primaryApprovers = append(primaryApprovers, GroupRoleManagementPolicyApprover{
											ID:   approver.(stable.GroupMembers).GroupId.GetOrZero(),
											Type: "groupMembers",
										})
									default:
										return fmt.Errorf("unknown approver type: %s", *approver.SubjectSet().ODataType)
									}
								}
							}
						}

						model.ActivationRules[0].ApprovalStages = []GroupRoleManagementPolicyApprovalStage{{PrimaryApprovers: primaryApprovers}}

					case "AuthenticationContext_EndUser_Assignment":
						if rule.(stable.UnifiedRoleManagementPolicyAuthenticationContextRule).ClaimValue.GetOrZero() != "" {
							model.ActivationRules[0].RequireConditionalAccessContext = rule.(stable.UnifiedRoleManagementPolicyAuthenticationContextRule).ClaimValue.GetOrZero()
						}

					case "Enablement_Admin_Assignment":
						model.ActiveAssignmentRules[0].RequireMultiFactorAuth = false
						model.ActiveAssignmentRules[0].RequireJustification = false

						if enabledRules := rule.(stable.UnifiedRoleManagementPolicyEnablementRule).EnabledRules; enabledRules != nil {
							for _, enabledRule := range *enabledRules {
								switch enabledRule {
								case "MultiFactorAuthentication":
									model.ActiveAssignmentRules[0].RequireMultiFactorAuth = true
								case "Justification":
									model.ActiveAssignmentRules[0].RequireJustification = true
								}
							}
						}

					case "Enablement_EndUser_Assignment":
						model.ActivationRules[0].RequireMultiFactorAuth = false
						model.ActivationRules[0].RequireJustification = false
						model.ActivationRules[0].RequireTicketInfo = false

						if enabledRules := rule.(stable.UnifiedRoleManagementPolicyEnablementRule).EnabledRules; enabledRules != nil {
							for _, enabledRule := range *enabledRules {
								switch enabledRule {
								case "MultiFactorAuthentication":
									model.ActivationRules[0].RequireMultiFactorAuth = true
								case "Justification":
									model.ActivationRules[0].RequireJustification = true
								case "Ticketing":
									model.ActivationRules[0].RequireTicketInfo = true
								}
							}
						}

					case "Expiration_Admin_Eligibility":
						model.EligibleAssignmentRules[0].ExpirationRequired = rule.(stable.UnifiedRoleManagementPolicyExpirationRule).IsExpirationRequired.GetOrZero()
						model.EligibleAssignmentRules[0].ExpireAfter = rule.(stable.UnifiedRoleManagementPolicyExpirationRule).MaximumDuration.GetOrZero()

					case "Expiration_Admin_Assignment":
						model.ActiveAssignmentRules[0].ExpirationRequired = rule.(stable.UnifiedRoleManagementPolicyExpirationRule).IsExpirationRequired.GetOrZero()
						model.ActiveAssignmentRules[0].ExpireAfter = rule.(stable.UnifiedRoleManagementPolicyExpirationRule).MaximumDuration.GetOrZero()

					case "Expiration_EndUser_Assignment":
						model.ActivationRules[0].MaximumDuration = rule.(stable.UnifiedRoleManagementPolicyExpirationRule).MaximumDuration.GetOrZero()

					case "Notification_Admin_Admin_Assignment":
						model.NotificationRules[0].ActiveAssignments[0].AdminNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}

					case "Notification_Admin_Admin_Eligibility":
						model.NotificationRules[0].EligibleAssignments[0].AdminNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}

					case "Notification_Admin_EndUser_Assignment":
						model.NotificationRules[0].EligibleActivations[0].AdminNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}

					case "Notification_Approver_Admin_Assignment":
						model.NotificationRules[0].ActiveAssignments[0].ApproverNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}

					case "Notification_Approver_Admin_Eligibility":
						model.NotificationRules[0].EligibleAssignments[0].ApproverNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}

					case "Notification_Approver_EndUser_Assignment":
						model.NotificationRules[0].EligibleActivations[0].ApproverNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}

					case "Notification_Requestor_Admin_Assignment":
						model.NotificationRules[0].ActiveAssignments[0].AssigneeNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}

					case "Notification_Requestor_Admin_Eligibility":
						model.NotificationRules[0].EligibleAssignments[0].AssigneeNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}

					case "Notification_Requestor_EndUser_Assignment":
						model.NotificationRules[0].EligibleActivations[0].AssigneeNotifications = []GroupRoleManagementPolicyNotificationSettings{
							flattenNotificationSettings(rule.(stable.UnifiedRoleManagementPolicyNotificationRule)),
						}
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

			policyId, err := parse.ParseRoleManagementPolicyID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("parsing policy ID: %v", err)
			}

			var model GroupRoleManagementPolicyModel
			if err = metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id := stable.NewPolicyRoleManagementPolicyID(policyId.ID())

			policyResp, err := client.GetRoleManagementPolicy(ctx, id, rolemanagementpolicy.DefaultGetRoleManagementPolicyOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %v", id, err)
			}

			policy := policyResp.Model
			if policy == nil {
				return fmt.Errorf("retrieving %s: API error, model was nil", id)
			}

			policyUpdate, err := buildPolicyForUpdate(pointer.To(metadata), policy)
			if err != nil {
				return fmt.Errorf("building update request: %v", err)
			}

			if _, err = client.UpdateRoleManagementPolicy(ctx, id, *policyUpdate); err != nil {
				return fmt.Errorf("updating %s: %v", id, err)
			}

			// Update the ID as it changes on modification
			policyId, err = getPolicyId(ctx, metadata, model.GroupId, model.RoleId)
			if err != nil {
				return fmt.Errorf("parsing policy ID: %v", err)
			}

			metadata.SetID(policyId)

			return nil
		},
	}
}

func (r GroupRoleManagementPolicyResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			// Policy cannot be destroyed, so this is a noop
			return nil
		},
	}
}

func buildPolicyForUpdate(metadata *sdk.ResourceMetaData, policy *stable.UnifiedRoleManagementPolicy) (*stable.UnifiedRoleManagementPolicy, error) {
	var model GroupRoleManagementPolicyModel
	if err := metadata.Decode(&model); err != nil {
		return nil, fmt.Errorf("decoding: %+v", err)
	}

	// Take the slice of rules and convert it to a map with the ID as the key
	policyRules := make(map[string]stable.UnifiedRoleManagementPolicyRule)
	for _, rule := range *policy.Rules {
		id := rule.UnifiedRoleManagementPolicyRule().Id
		if id == nil {
			continue
		}
		policyRules[*id] = rule
	}
	updatedRules := make([]stable.UnifiedRoleManagementPolicyRule, 0)

	if metadata.ResourceData.HasChange("eligible_assignment_rules") {
		rule, ok := policyRules["Expiration_Admin_Eligibility"].(stable.UnifiedRoleManagementPolicyExpirationRule)
		if !ok {
			return nil, fmt.Errorf("policy rule was not a UnifiedRoleManagementPolicyExpirationRule")
		}

		expirationRequired := rule.IsExpirationRequired
		if expirationRequired.GetOrZero() != model.EligibleAssignmentRules[0].ExpirationRequired {
			expirationRequired = nullable.Value(model.EligibleAssignmentRules[0].ExpirationRequired)
		}

		maximumDuration := rule.MaximumDuration
		if maximumDuration.GetOrZero() != model.EligibleAssignmentRules[0].ExpireAfter && model.EligibleAssignmentRules[0].ExpireAfter != "" {
			maximumDuration = nullable.Value(model.EligibleAssignmentRules[0].ExpireAfter)
		}

		rule.IsExpirationRequired = expirationRequired
		rule.MaximumDuration = maximumDuration

		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("active_assignment_rules.0.require_multifactor_authentication") ||
		metadata.ResourceData.HasChange("active_assignment_rules.0.require_justification") {
		rule, ok := policyRules["Enablement_Admin_Assignment"].(stable.UnifiedRoleManagementPolicyEnablementRule)
		if !ok {
			return nil, fmt.Errorf("policy rule was not a UnifiedRoleManagementPolicyEnablementRule")
		}

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

		rule.EnabledRules = pointer.To(enabledRules)

		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("active_assignment_rules.0.expiration_required") ||
		metadata.ResourceData.HasChange("active_assignment_rules.0.expire_after") {
		rule, ok := policyRules["Expiration_Admin_Assignment"].(stable.UnifiedRoleManagementPolicyExpirationRule)
		if !ok {
			return nil, fmt.Errorf("policy rule was not a UnifiedRoleManagementPolicyExpirationRule")
		}

		expirationRequired := rule.IsExpirationRequired
		if expirationRequired.GetOrZero() != model.ActiveAssignmentRules[0].ExpirationRequired {
			expirationRequired = nullable.Value(model.ActiveAssignmentRules[0].ExpirationRequired)
		}

		maximumDuration := rule.MaximumDuration
		if maximumDuration.GetOrZero() != model.ActiveAssignmentRules[0].ExpireAfter && model.ActiveAssignmentRules[0].ExpireAfter != "" {
			maximumDuration = nullable.Value(model.EligibleAssignmentRules[0].ExpireAfter)
		}

		rule.IsExpirationRequired = expirationRequired
		rule.MaximumDuration = maximumDuration

		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("activation_rules.0.maximum_duration") {
		rule, ok := policyRules["Expiration_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyExpirationRule)
		if !ok {
			return nil, fmt.Errorf("policy rule was not a UnifiedRoleManagementPolicyExpirationRule")
		}

		rule.MaximumDuration = nullable.Value(model.ActivationRules[0].MaximumDuration)

		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("activation_rules.0.require_approval") ||
		metadata.ResourceData.HasChange("activation_rules.0.approval_stage") {
		rule, ok := policyRules["Approval_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyApprovalRule)
		if !ok {
			return nil, fmt.Errorf("policy rule was not a UnifiedRoleManagementPolicyApprovalRule")
		}

		if model.ActivationRules[0].RequireApproval && len(model.ActivationRules[0].ApprovalStages) != 1 {
			return nil, fmt.Errorf("require_approval is true, but no approval_stages are provided")
		}

		isApprovalRequired := rule.Setting.IsApprovalRequired.GetOrZero()
		var approvalStages []stable.UnifiedApprovalStage
		if isApprovalRequired != model.ActivationRules[0].RequireApproval {
			isApprovalRequired = model.ActivationRules[0].RequireApproval
		}

		if metadata.ResourceData.HasChange("activation_rules.0.approval_stage") {
			approvalStages = make([]stable.UnifiedApprovalStage, 0)
			for _, stage := range model.ActivationRules[0].ApprovalStages {
				primaryApprovers := make([]stable.SubjectSet, 0)
				for _, approver := range stage.PrimaryApprovers {
					switch approver.Type {
					case "singleUser":
						primaryApprovers = append(primaryApprovers, stable.SingleUser{
							ODataType: pointer.To("#microsoft.graph.singleUser"),
							UserId:    nullable.Value(approver.ID),
						})
					case "groupMembers":
						primaryApprovers = append(primaryApprovers, stable.GroupMembers{
							ODataType: pointer.To("#microsoft.graph.groupMembers"),
							GroupId:   nullable.Value(approver.ID),
						})
					}
				}

				approvalStages = append(approvalStages, stable.UnifiedApprovalStage{
					PrimaryApprovers: &primaryApprovers,
				})
			}
		} else {
			approvalStages = *policyRules["Approval_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyApprovalRule).Setting.ApprovalStages
		}

		rule = stable.UnifiedRoleManagementPolicyApprovalRule{
			Id:        policyRules["Approval_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyApprovalRule).Id,
			ODataType: policyRules["Approval_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyApprovalRule).ODataType,
			Target:    policyRules["Approval_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyApprovalRule).Target,
			Setting: &stable.ApprovalSettings{
				IsApprovalRequired: nullable.Value(isApprovalRequired),
				ApprovalStages:     &approvalStages,
			},
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("activation_rules.0.required_conditional_access_authentication_context") {
		claimValue := policyRules["AuthenticationContext_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyAuthenticationContextRule).ClaimValue

		var isEnabled bool
		if _, set := metadata.ResourceData.GetOk("activation_rules.0.required_conditional_access_authentication_context"); set {
			isEnabled = true
			claimValue = nullable.Value(model.ActivationRules[0].RequireConditionalAccessContext)
		}

		rule := stable.UnifiedRoleManagementPolicyAuthenticationContextRule{
			Id:         policyRules["AuthenticationContext_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyAuthenticationContextRule).Id,
			ODataType:  policyRules["AuthenticationContext_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyAuthenticationContextRule).ODataType,
			Target:     policyRules["AuthenticationContext_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyAuthenticationContextRule).Target,
			IsEnabled:  nullable.Value(isEnabled),
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

		rule := stable.UnifiedRoleManagementPolicyEnablementRule{
			Id:           policyRules["Enablement_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyEnablementRule).Id,
			ODataType:    policyRules["Enablement_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyEnablementRule).ODataType,
			Target:       policyRules["Enablement_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyEnablementRule).Target,
			EnabledRules: &enabledRules,
		}
		updatedRules = append(updatedRules, rule)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.admin_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Admin_Admin_Eligibility"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].EligibleAssignments[0].AdminNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.admin_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.admin_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Admin_Admin_Assignment"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].ActiveAssignments[0].AdminNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.admin_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.admin_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Admin_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].EligibleActivations[0].AdminNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.admin_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.approver_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Approver_Admin_Eligibility"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].EligibleAssignments[0].ApproverNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.approver_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.approver_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Approver_Admin_Assignment"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].ActiveAssignments[0].ApproverNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.approver_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.approver_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Approver_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].EligibleActivations[0].ApproverNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.approver_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.assignee_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Requestor_Admin_Eligibility"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].EligibleAssignments[0].AssigneeNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_assignments.0.assignee_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.assignee_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Requestor_Admin_Assignment"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].ActiveAssignments[0].AssigneeNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.active_assignments.0.assignee_notifications.0.additional_recipients"),
			),
		)
	}

	if metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.assignee_notifications") {
		updatedRules = append(updatedRules,
			expandNotificationSettings(
				policyRules["Notification_Requestor_EndUser_Assignment"].(stable.UnifiedRoleManagementPolicyNotificationRule),
				model.NotificationRules[0].EligibleActivations[0].AssigneeNotifications[0],
				metadata.ResourceData.HasChange("notification_rules.0.eligible_activations.0.assignee_notifications.0.additional_recipients"),
			),
		)
	}

	return &stable.UnifiedRoleManagementPolicy{
		Id:    policy.Id,
		Rules: pointer.To(updatedRules),
	}, nil
}

func expandNotificationSettings(rule stable.UnifiedRoleManagementPolicyNotificationRule, data GroupRoleManagementPolicyNotificationSettings, recipientChange bool) stable.UnifiedRoleManagementPolicyNotificationRule {
	level := rule.NotificationLevel
	defaultRecipients := rule.IsDefaultRecipientsEnabled
	additionalRecipients := rule.NotificationRecipients

	if level.GetOrZero() != data.NotificationLevel {
		level = nullable.Value(data.NotificationLevel)
	}
	if defaultRecipients.GetOrZero() != data.DefaultRecipients {
		defaultRecipients = nullable.Value(data.DefaultRecipients)
	}
	if recipientChange {
		additionalRecipients = pointer.To(data.AdditionalRecipients)
	}

	return stable.UnifiedRoleManagementPolicyNotificationRule{
		Id:                         rule.Id,
		ODataType:                  rule.ODataType,
		Target:                     rule.Target,
		RecipientType:              rule.RecipientType,
		NotificationType:           rule.NotificationType,
		NotificationLevel:          level,
		IsDefaultRecipientsEnabled: defaultRecipients,
		NotificationRecipients:     additionalRecipients,
	}
}

func flattenNotificationSettings(rule stable.UnifiedRoleManagementPolicyNotificationRule) GroupRoleManagementPolicyNotificationSettings {
	return GroupRoleManagementPolicyNotificationSettings{
		NotificationLevel:    rule.NotificationLevel.GetOrZero(),
		DefaultRecipients:    rule.IsDefaultRecipientsEnabled.GetOrZero(),
		AdditionalRecipients: pointer.From(rule.NotificationRecipients),
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
			Description:  "What level of notifications are sent",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"All", "Critical"}, false),
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
