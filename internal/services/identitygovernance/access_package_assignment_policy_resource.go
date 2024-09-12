// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

const accessPackageAssignmentPolicyResourceName = "azuread_access_package_assignment_policy"

func accessPackageAssignmentPolicyResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: accessPackageAssignmentPolicyResourceCreate,
		ReadContext:   accessPackageAssignmentPolicyResourceRead,
		UpdateContext: accessPackageAssignmentPolicyResourceUpdate,
		DeleteContext: accessPackageAssignmentPolicyResourceDelete,

		CustomizeDiff: assignmentPolicyCustomDiff,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"access_package_id": {
				Description:      "The ID of the access package that will contain the policy",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"display_name": {
				Description:      "The display name of the policy",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"description": {
				Description:      "The description of the policy",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"duration_in_days": {
				Description:   "How many days this assignment is valid for",
				Type:          pluginsdk.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"expiration_date"},
				ValidateFunc:  validation.IntBetween(0, 3660),
			},

			"expiration_date": {
				Description:   "The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z)",
				Type:          pluginsdk.TypeString,
				Optional:      true,
				ConflictsWith: []string{"duration_in_days"},
				ValidateFunc:  validation.IsRFC3339Time,
				//DiffSuppressFunc: assignmentPolicyDiffSuppress,
			},

			"extension_enabled": {
				Description: "When enabled, users will be able to request extension of their access to this package before their access expires",
				Type:        pluginsdk.TypeBool,
				Optional:    true,
			},

			"requestor_settings": {
				Description:      "This block configures the users who can request access",
				Type:             pluginsdk.TypeList,
				Optional:         true,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
				MaxItems:         1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"requests_accepted": {
							Description: "Whether to accept requests now, when disabled, no new requests can be made using this policy",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"scope_type": {
							Description:  "Specify the scopes of the requestors",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(possibleValuesForRequestorScopeType, false),
						},

						"requestor": {
							Description: "The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers",
							Type:        pluginsdk.TypeList,
							Optional:    true,
							Elem:        schemaUserSet(),
						},
					},
				},
			},

			"approval_settings": {
				Description:      "Settings of whether approvals are required and how they are obtained",
				Type:             pluginsdk.TypeList,
				Optional:         true,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
				MaxItems:         1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"approval_required": {
							Description: "Whether an approval is required",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"approval_required_for_extension": {
							Description: "Whether an approval is required to grant extension. Same approval settings used to approve initial access will apply",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"requestor_justification_required": {
							Description: "Whether requestor are required to provide a justification to request an access package. Justification is visible to other approvers and the requestor",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"approval_stage": {
							Description: "The process to obtain an approval",
							Type:        pluginsdk.TypeList,
							Optional:    true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"approval_timeout_in_days": {
										Description: "Decision must be made in how many days? If a request is not approved within this time period after it is made, it will be automatically rejected",
										Type:        pluginsdk.TypeInt,
										Required:    true,
									},

									"approver_justification_required": {
										Description: "Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor",
										Type:        pluginsdk.TypeBool,
										Optional:    true,
									},

									"alternative_approval_enabled": {
										Description: "If no action taken, forward to alternate approvers?",
										Type:        pluginsdk.TypeBool,
										Optional:    true,
									},

									"enable_alternative_approval_in_days": {
										Description: "Forward to alternate approver(s) after how many days?",
										Type:        pluginsdk.TypeInt,
										Optional:    true,
									},

									"primary_approver": {
										Description: "The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection",
										Type:        pluginsdk.TypeList,
										Optional:    true,
										Elem:        schemaUserSet(),
									},

									"alternative_approver": {
										Description: "If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection",
										Type:        pluginsdk.TypeList,
										Optional:    true,
										Elem:        schemaUserSet(),
									},
								},
							},
						},
					},
				},
			},

			"assignment_review_settings": {
				Description:      "The settings of whether assignment review is needed and how it's conducted",
				Type:             pluginsdk.TypeList,
				Optional:         true,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
				MaxItems:         1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"enabled": {
							Description: "Whether to enable assignment review",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"review_frequency": {
							Description:  "This will determine how often the access review campaign runs",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(possibleValuesForAccessReviewRecurrenceType, false),
						},

						"review_type": {
							Description:  "Self review or specific reviewers",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(possibleValuesForAccessReviewReviewerType, false),
						},

						"starting_on": {
							Description:  "This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsRFC3339Time,
						},

						"duration_in_days": {
							Description: "How many days each occurrence of the access review series will run",
							Type:        pluginsdk.TypeInt,
							Optional:    true,
						},

						"reviewer": {
							Description: "If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers",
							Type:        pluginsdk.TypeList,
							Optional:    true,
							Elem:        schemaUserSet(),
						},

						"access_recommendation_enabled": {
							Description: "Whether to show Show reviewer decision helpers. If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"approver_justification_required": {
							Description: "Whether a reviewer need provide a justification for their decision. Justification is visible to other reviewers and the requestor",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"access_review_timeout_behavior": {
							Description:  "What actions the system takes if reviewers don't respond in time",
							Type:         pluginsdk.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(beta.PossibleValuesForAccessReviewTimeoutBehavior(), false),
						},
					},
				},
			},

			"question": {
				Description:      "One or more questions to the requestor",
				Type:             pluginsdk.TypeList,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
				Optional:         true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"required": {
							Description: "Whether this question is required",
							Type:        pluginsdk.TypeBool,
							Optional:    true,
						},

						"sequence": {
							Description: "The sequence number of this question",
							Type:        pluginsdk.TypeInt,
							Optional:    true,
						},

						"choice": {
							Description: "Configuration of a choice to the question",
							Type:        pluginsdk.TypeList,
							Optional:    true,
							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{
									"actual_value": {
										Description: "The actual value of this choice",
										Type:        pluginsdk.TypeString,
										Required:    true,
									},

									"display_value": {
										Description: "The display text of this choice",
										Type:        pluginsdk.TypeList,
										Required:    true,
										MaxItems:    1,
										Elem:        schemaLocalizedContent(),
									},
								},
							},
						},

						"text": {
							Description: "The content of this question",
							Type:        pluginsdk.TypeList,
							Required:    true,
							MaxItems:    1,
							Elem:        schemaLocalizedContent(),
						},
					},
				},
			},
		},
	}
}

func accessPackageAssignmentPolicyResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	properties, err := buildAssignmentPolicyResourceData(ctx, d, meta)
	if err != nil {
		return tf.ErrorDiagF(err, "Building resource data from supplied parameters")
	}

	resp, err := client.CreateEntitlementManagementAccessPackageAssignmentPolicy(ctx, *properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating access package assignment policy %q", d.Get("display_name").(string))
	}

	if resp.Model == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Creating access package assignment policy")
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(*resp.Model.Id)
	d.SetId(id.AccessPackageAssignmentPolicyId)

	return accessPackageAssignmentPolicyResourceRead(ctx, d, meta)
}

func accessPackageAssignmentPolicyResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(d.Id())

	properties, err := buildAssignmentPolicyResourceData(ctx, d, meta)
	if err != nil {
		return tf.ErrorDiagF(err, "Building resource data from supplied parameters")
	}

	tf.LockByName(accessPackageAssignmentPolicyResourceName, id.AccessPackageAssignmentPolicyId)
	defer tf.UnlockByName(accessPackageAssignmentPolicyResourceName, id.AccessPackageAssignmentPolicyId)

	if _, err = client.SetEntitlementManagementAccessPackageAssignmentPolicy(ctx, id, *properties); err != nil {
		return tf.ErrorDiagF(err, "Updating %s", id)
	}

	return accessPackageAssignmentPolicyResourceRead(ctx, d, meta)
}

func accessPackageAssignmentPolicyResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(d.Id())

	resp, err := client.GetEntitlementManagementAccessPackageAssignmentPolicy(ctx, id, entitlementmanagementaccesspackageassignmentpolicy.DefaultGetEntitlementManagementAccessPackageAssignmentPolicyOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "Retrieving %s", id)
	}

	accessPackageAssignmentPolicy := resp.Model
	if accessPackageAssignmentPolicy == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "access_package_id", accessPackageAssignmentPolicy.AccessPackageId.GetOrZero())
	tf.Set(d, "approval_settings", flattenApprovalSettings(accessPackageAssignmentPolicy.RequestApprovalSettings))
	tf.Set(d, "assignment_review_settings", flattenAssignmentReviewSettings(accessPackageAssignmentPolicy.AccessReviewSettings))
	tf.Set(d, "description", accessPackageAssignmentPolicy.Description.GetOrZero())
	tf.Set(d, "display_name", accessPackageAssignmentPolicy.DisplayName.GetOrZero())
	tf.Set(d, "duration_in_days", int(accessPackageAssignmentPolicy.DurationInDays.GetOrZero()))
	tf.Set(d, "expiration_date", accessPackageAssignmentPolicy.ExpirationDateTime.GetOrZero())
	tf.Set(d, "extension_enabled", accessPackageAssignmentPolicy.CanExtend.GetOrZero())
	tf.Set(d, "question", flattenAccessPackageQuestions(accessPackageAssignmentPolicy.Questions))
	tf.Set(d, "requestor_settings", flattenRequestorSettings(accessPackageAssignmentPolicy.RequestorSettings))

	return nil
}

func accessPackageAssignmentPolicyResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(d.Id())

	if _, err := client.DeleteEntitlementManagementAccessPackageAssignmentPolicy(ctx, id, entitlementmanagementaccesspackageassignmentpolicy.DefaultDeleteEntitlementManagementAccessPackageAssignmentPolicyOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	// Wait for user object to be deleted
	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetEntitlementManagementAccessPackageAssignmentPolicy(ctx, id, entitlementmanagementaccesspackageassignmentpolicy.DefaultGetEntitlementManagementAccessPackageAssignmentPolicyOperationOptions()); err != nil {
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

func buildAssignmentPolicyResourceData(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) (*beta.AccessPackageAssignmentPolicy, error) {
	accessPackageClient := meta.(*clients.Client).IdentityGovernance.AccessPackageClient
	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(d.Get("access_package_id").(string))

	resp, err := accessPackageClient.GetEntitlementManagementAccessPackage(ctx, id, entitlementmanagementaccesspackage.DefaultGetEntitlementManagementAccessPackageOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state!", id)
		}

		return nil, fmt.Errorf("retrieving %s: %v", id, err)
	}

	properties := beta.AccessPackageAssignmentPolicy{
		AccessPackageId:    nullable.NoZero(d.Get("access_package_id").(string)),
		CanExtend:          nullable.NoZero(d.Get("extension_enabled").(bool)),
		Description:        nullable.NoZero(d.Get("description").(string)),
		DisplayName:        nullable.NoZero(d.Get("display_name").(string)),
		DurationInDays:     nullable.NoZero(int64(d.Get("duration_in_days").(int))),
		ExpirationDateTime: nullable.NoZero(d.Get("expiration_date").(string)),
		Questions:          expandAccessPackageQuestions(d.Get("question").([]interface{})),
	}

	requestApprovalSettings, err := expandApprovalSettings(d.Get("approval_settings").([]interface{}))
	if err != nil {
		return nil, fmt.Errorf("expanding `approval_settings`: %v", err)
	}
	properties.RequestApprovalSettings = requestApprovalSettings

	requestorSettings, err := expandRequestorSettings(d.Get("requestor_settings").([]interface{}))
	if err != nil {
		return nil, fmt.Errorf("building `requestor_settings`: %v", err)
	}
	properties.RequestorSettings = requestorSettings

	reviewSettings, err := expandAssignmentReviewSettings(d.Get("assignment_review_settings").([]interface{}))
	if err != nil {
		return nil, fmt.Errorf("building `assignment_review_settings`: %v", err)
	}
	properties.AccessReviewSettings = reviewSettings

	return &properties, nil
}

func assignmentPolicyDiffSuppress(k, old, new string, d *pluginsdk.ResourceData) bool {
	if k == "approval_settings.#" && old == "1" && new == "0" {
		return true
	}

	if k == "requestor_settings.#" && old == "1" && new == "0" {
		return true
	}

	if k == "requestor_settings.0.scope_type" && old == RequestorScopeTypeNoSubjects && len(new) == 0 {
		return true
	}

	if k == "assignment_review_settings.0.starting_on" && len(new) == 0 {
		return true
	}

	if k == "assignment_review_settings.#" && old == "1" && new == "0" {
		return true
	}

	if k == "question.#" && old == "1" && new == "0" {
		return true
	}

	return false
}

func assignmentPolicyCustomDiff(ctx context.Context, diff *pluginsdk.ResourceDiff, meta interface{}) error {
	if reviewSettings := diff.Get("assignment_review_settings").([]interface{}); len(reviewSettings) > 0 {
		reviewSetting := reviewSettings[0].(map[string]interface{})
		if reviewSetting["enabled"].(bool) &&
			(reviewSetting["duration_in_days"] == 0 ||
				len(reviewSetting["review_frequency"].(string)) == 0 ||
				len(reviewSetting["access_review_timeout_behavior"].(string)) == 0) {
			return fmt.Errorf("`duration_in_days`, `review_frequency`, `access_review_timeout_behavior` must be set when review is enabled")
		}
	}

	return nil
}
