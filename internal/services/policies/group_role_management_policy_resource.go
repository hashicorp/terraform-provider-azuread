// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

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
	ActiveAssignmentRules  AssignmentRules                          `tfschema:"active_assignment_rules`
	EligbleAssignmentRules AssignmentRules                          `tfschema:"eligible_assignment_rules`
	ActivationRules        ActivationRules                          `tfschema:"activation_rules"`
	NotificationRules      NotificationRules                        `tfschema:"notification_rules"`
}

type AssignmentRules struct {
	AllowPermanent        bool   `tfschema:"allow_permanent"`
	ExpireAfter           string `tfschema:"expire_after"`
	ReqireMultiFactorAuth bool   `tfschema:"require_multifactor_authentication"`
	RequireJustification  bool   `tfschema:"require_justification"`
}

type ActivationRules struct {
	MaximumDuration string          `tfschema:"maximum_duration"`
	RequireApproval bool            `tfschema:"require_approval"`
	ApprovalStages  []ApprovalStage `tfschema:"approval_stages"`
}

type ApprovalStage struct {
	PrimaryApprovers []Approver `tfschema:"primary_approvers"`
}

type Approver struct {
	Description string `tfschema:"description"`
	UserId      string `tfschema:"user_id"`
	GroupId     string `tfschema:"group_id"`
}

type NotificationRules struct {
	AdminNotifications    NotificationRule `tfschema:"admin_notifications"`
	ApproverNotifications NotificationRule `tfschema:"approver_notifications"`
	AssigneeNotifications NotificationRule `tfschema:"assignee_notifications"`
}

type NotificationRule struct {
	EligibleAssignments NotificationSettings `tfschema:"eligible_assignments"`
	ActiveAssignments   NotificationSettings `tfschema:"active_assignments"`
	Activations         NotificationSettings `tfschema:"activations"`
}

type NotificationSettings struct {
	NotificationLevel    string   `tfschema:"notification_level"`
	DefaultRecipients    bool     `tfschema:"default_recipients"`
	AdditionalRecipients []string `tfschema:"additional_recipients"`
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
			ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
		},

		"assignment_type": {
			Description: "The ID of the assignment to the group",
			Type:        pluginsdk.TypeString,
			Required:    true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{
				msgraph.PrivilegedAccessGroupRelationshipMember,
				msgraph.PrivilegedAccessGroupRelationshipOwner,
				msgraph.PrivilegedAccessGroupRelationshipUnknown,
			}, false)),
		},

		"eligible_assignment_rules": {
			Description: "The rules for eligible assignment of the policy",
			Type:        pluginsdk.TypeMap,
			Optional:    true,
			Computed:    true,
			Elem: map[string]*pluginsdk.Schema{
				// Expiration_Admin_Eligibility #microsoft.graph.unifiedRoleManagementPolicyExpirationRule
				"allow_permanent": { // isExpirationRequired
					Description:   "Whether assignments can be permanent",
					Type:          pluginsdk.TypeBool,
					Optional:      true,
					Computed:      true,
					ConflictsWith: []string{"eligible_assignment.0.allow_permanent"},
				},

				"expire_after": { // maximumDuration
					Description:      "The duration after which assignments expire",
					Type:             pluginsdk.TypeString,
					Optional:         true,
					Computed:         true,
					ConflictsWith:    []string{"eligible_assignment.0.allow_permanent"},
					ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"P15D", "P1M", "P3M", "P6M", "P1Y"}, false)),
				},
			},
		},

		"active_assignment_rules": {
			Description: "The rules for active assignment of the policy",
			Type:        pluginsdk.TypeMap,
			Optional:    true,
			Computed:    true,
			Elem: map[string]*pluginsdk.Schema{
				// Expiration_Admin_Assignment #microsoft.graph.unifiedRoleManagementPolicyExpirationRule
				"allow_permanent": { // isExpirationRequired
					Description:   "Whether assignments can be permanent",
					Type:          pluginsdk.TypeBool,
					Optional:      true,
					Computed:      true,
					ConflictsWith: []string{"eligible_assignment.0.allow_permanent"},
				},

				"expire_after": { // maximumDuration
					Description:      "The duration after which assignments expire",
					Type:             pluginsdk.TypeString,
					Optional:         true,
					Computed:         true,
					ConflictsWith:    []string{"eligible_assignment.0.allow_permanent"},
					ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"P15D", "P1M", "P3M", "P6M", "P1Y"}, false)),
				},

				// Enablement_Admin_Assignment #microsoft.graph.unifiedRoleManagementPolicyEnablementRule
				"require_multifactor_authentication": { // enabledRule "MultiFactorAuthentication"
					Description: "Whether multi-factor authentication is required to make an assignment",
					Type:        pluginsdk.TypeBool,
					Optional:    true,
					Computed:    true,
				},
				"require_justification": { // enabledRule "Justification"
					Description: "Whether a justification is required to make an assignment",
					Type:        pluginsdk.TypeBool,
					Optional:    true,
					Computed:    true,
				},
			},
		},

		"activation_rules": {
			Description: "The activation rules of the policy",
			Type:        pluginsdk.TypeMap,
			Optional:    true,
			Computed:    true,
			Elem: map[string]*pluginsdk.Schema{
				// Expiration_EndUser_Assignment #microsoft.graph.unifiedRoleManagementPolicyExpirationRule
				"maximum_duration": { // maximumDuration
					Description: "The maximum duration an activation can be valid for",
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

				// Approval_EndUser_Assignment #microsoft.graph.unifiedRoleManagementPolicyApprovalRule
				"require_approval": { // setting.isApprovalRequired
					Description: "Whether an approval is required for activation",
					Type:        pluginsdk.TypeBool,
					Optional:    true,
					Computed:    true,
				},

				"approval_stages": { // setting.approvalStages
					Description: "The approval stages for the activation",
					Type:        pluginsdk.TypeMap,
					Optional:    true,
					MinItems:    1,
					MaxItems:    1,
					Elem: map[string]*pluginsdk.Schema{
						"primary_approvers": { // primaryApprovers
							Description: "The IDs of the users or groups who can approve the activation",
							Type:        pluginsdk.TypeList,
							Optional:    true,
							Computed:    true,
							MinItems:    1,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"description": { // description
										Description:      "The description of the approver",
										Type:             pluginsdk.TypeString,
										Required:         true,
										ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
									},

									"user_id": { // "@odata.type" : "#microsoft.graph.singleUser", user
										Description:      "The ID of the user to act as an approver",
										Type:             pluginsdk.TypeString,
										Required:         true,
										ConflictsWith:    []string{"approvers.0.group_id"},
										ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
									},

									"group_id": { // "@odata.type" : "#microsoft.graph.groupMembers"
										Description:      "The ID of the group to act as an approver",
										Type:             pluginsdk.TypeString,
										Required:         true,
										ConflictsWith:    []string{"approvers.0.user_id"},
										ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
									},
								},
							},
						},
					},
				},

				// AuthenticationContext_EndUser_Assignment #microsoft.graph.unifiedRoleManagementPolicyAuthenticationContextRule
				"require_conditional_access_authentication_context": { // isEnabled, claimValue
					Description:      "Whether a conditional access context is required during activation",
					Type:             pluginsdk.TypeString,
					Optional:         true,
					Computed:         true,
					ConflictsWith:    []string{"activation.0.require_multifactor_authentication"},
					ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
				},

				// Enablement_EndUser_Assignment #microsoft.graph.unifiedRoleManagementPolicyEnablementRule
				"require_multifactor_authentication": { // enabledRule "MultiFactorAuthentication"
					Description:   "Whether multi-factor authentication is required during activation",
					Type:          pluginsdk.TypeBool,
					Optional:      true,
					Computed:      true,
					ConflictsWith: []string{"activation.0.require_conditional_access"},
				},

				"require_justification": { // enabledRules "Justification"
					Description: "Whether a justification is required during activation",
					Type:        pluginsdk.TypeBool,
					Optional:    true,
					Computed:    true,
				},

				"require_ticket_info": { // enabledRules "Ticketing"
					Description: "Whether ticket information is required during activation",
					Type:        pluginsdk.TypeBool,
					Optional:    true,
					Computed:    true,
				},
			},
		},

		"notification_rules": {
			Description: "The notification rules of the policy",
			Type:        pluginsdk.TypeMap,
			Optional:    true,
			Computed:    true,
			Elem: map[string]*pluginsdk.Schema{
				"admin_notifications": {
					Description: "The admin notifications on assignment",
					Type:        pluginsdk.TypeMap,
					Optional:    true,
					Computed:    true,
					Elem: map[string]*pluginsdk.Schema{
						// Notification_Admin_Admin_Eligibility #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"eligible_assignments": {
							Description: "The admin notifications for eligible assignments",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
									Description: "The additional recipients to notify",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},

						// Notification_Admin_Admin_Assignment #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"active_assignments": {
							Description: "The admin notifications for active assignments",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
									Description: "The additional recipients to notify",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},

						// Notification_Admin_EndUser_Assignment #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"activations": {
							Description: "The admin notifications for role activation",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
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

				"approver_notifications": {
					Description: "The admin notifications on assignment",
					Type:        pluginsdk.TypeMap,
					Optional:    true,
					Computed:    true,
					Elem: map[string]*pluginsdk.Schema{
						// Notification_Approver_Admin_Eligibility #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"eligible_assignments": {
							Description: "The admin notifications for eligible assignments",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
									Description: "The additional recipients to notify",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},

						// Notification_Approver_Admin_Assignment #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"active_assignments": {
							Description: "The admin notifications for active assignments",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
									Description: "The additional recipients to notify",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},

						// Notification_Approver_EndUser_Assignment #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"activations": {
							Description: "The admin notifications for role activation",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
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

				"assignee_notifications": {
					Description: "The admin notifications on assignment",
					Type:        pluginsdk.TypeMap,
					Optional:    true,
					Computed:    true,
					Elem: map[string]*pluginsdk.Schema{
						// Notification_Requestor_Admin_Eligibility #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"eligible_assignments": {
							Description: "The admin notifications for eligible assignments",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
									Description: "The additional recipients to notify",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},

						// Notification_Requestor_Admin_Assignment #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"active_assignments": {
							Description: "The admin notifications for active assignments",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
									Description: "The additional recipients to notify",
									Type:        pluginsdk.TypeList,
									Optional:    true,
									Elem: &pluginsdk.Schema{
										Type: pluginsdk.TypeString,
									},
								},
							},
						},

						// Notification_Requestor_EndUser_Assignment #microsoft.graph.unifiedRoleManagementPolicyNotificationRule
						"activations": {
							Description: "The admin notifications for role activation",
							Type:        pluginsdk.TypeMap,
							Optional:    true,
							Computed:    true,
							Elem: map[string]*pluginsdk.Schema{
								"notification_level": { // notificationLevel All/Critical
									Description:      "What level of notifications are sent",
									Type:             pluginsdk.TypeString,
									Optional:         true,
									Computed:         true,
									ValidateDiagFunc: validation.ValidateDiag(validation.StringInSlice([]string{"All", "Critical"}, false)),
								},
								"default_recipients": { // isDefaultRecipientsEnabled
									Description: "Whether the default recipients are notified",
									Type:        pluginsdk.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"additional_recipients": { // notificationRecipients
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
	}
}

func (r RoleManagementPolicyResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"display_name": {
			Description:      "The display name of the policy",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},

		"description": {
			Description:      "Description of the policy",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},
	}
}

func (r RoleManagementPolicyResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
		},
	}
}

func (r RoleManagementPolicyResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
		},
	}
}

func (r RoleManagementPolicyResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
		},
	}
}

func (r RoleManagementPolicyResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
		},
	}
}
