package identitygovernance

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

const accessPackageAssignmentPolicyResourceName = "azuread_access_package_assignment_policy"

func accessPackageAssignmentPolicyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: accessPackageAssignmentPolicyResourceCreate,
		ReadContext:   accessPackageAssignmentPolicyResourceRead,
		UpdateContext: accessPackageAssignmentPolicyResourceUpdate,
		DeleteContext: accessPackageAssignmentPolicyResourceDelete,

		CustomizeDiff: assignmentPolicyCustomDiff,

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
				Description:      "The display name of the policy.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"description": {
				Description:      "The description of the policy.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},
			"can_extend": {
				Description: "When enabled, users will be able to request extension of their access to this package before their access expires.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"duration_in_days": {
				Description:   "How many days this assignment is valid for.",
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"expiration_date"},
				ValidateFunc:  validation.IntBetween(0, 3660),
			},
			"expiration_date": {
				Description:      "The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).",
				Type:             schema.TypeString,
				Optional:         true,
				ConflictsWith:    []string{"duration_in_days"},
				ValidateFunc:     validation.IsRFC3339Time,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
			},
			"requestor_settings": {
				Description:      "This block configures the users who can request access.",
				Type:             schema.TypeList,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
				MaxItems:         1,
				Optional:         true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_requests": {
							Description: "Whether to accept requets now, when disabled, no new requests can be made using this policy.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"scope_type": {
							Description: "Specify the scopes of the requestors. Valid values are `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`,`SpecificDirectorySubjects`.",
							Type:        schema.TypeString,
							Optional:    true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.RequestorSettingsScopeTypeAllConfiguredConnectedOrganizationSubjects,
								msgraph.RequestorSettingsScopeTypeAllExistingConnectedOrganizationSubjects,
								msgraph.RequestorSettingsScopeTypeAllExistingDirectoryMemberUsers,
								msgraph.RequestorSettingsScopeTypeAllExistingDirectorySubjects,
								msgraph.RequestorSettingsScopeTypeAllExternalSubjects,
								msgraph.RequestorSettingsScopeTypeNoSubjects,
								msgraph.RequestorSettingsScopeTypeSpecificConnectedOrganizationSubjects,
								msgraph.RequestorSettingsScopeTypeSpecificDirectorySubjects,
							}, false),
						},
						"requestor": {
							Description: "The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        schemaUserSet(),
						},
					},
				},
			},
			"approval_settings": {
				Description:      "Settings of whether apporvals are required and how they are obtained.",
				Type:             schema.TypeList,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
				MaxItems:         1,
				Optional:         true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_approval_required": {
							Description: "Whether an approval is required.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"is_approval_required_for_extension": {
							Description: "Whether an approval is required to grant extension. Same approval settings used to approve initial access will apply.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"is_requestor_justification_required": {
							Description: "Whether reuqirestor are required to provide a justification to request an access package. Justification is visible to other approvers and the requestor.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"approval_stage": {
							Description: "The process to obtain an approval",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"approval_timeout_in_days": {
										Description: "Decision must be made in how many days? If a request is not approved within this time period after it is made, it will be automatically rejected.",
										Type:        schema.TypeInt,
										Required:    true,
									},
									"is_approver_justification_required": {
										Description: "Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"is_alternative_approval_enabled": {
										Description: "If no action taken, forward to alternate approvers?",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"enable_alternative_approval_in_days": {
										Description: "Forward to alternate approver(s) after how many days?",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"primary_approver": {
										Description: "The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem:        schemaUserSet(),
									},
									"alternative_approver": {
										Description: "If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection.",
										Type:        schema.TypeList,
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
				Description:      "The settings of whether assignment review is needed and how it's conducted.",
				Type:             schema.TypeList,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
				MaxItems:         1,
				Optional:         true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_enabled": {
							Description: "Whether to enable assignment reivew.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"review_frequency": {
							Description: "This will determine how often the access review campaign runs, valid values are `weekly`,`monthly`,`quarterly`,`halfyearly`,`annual`.",
							Type:        schema.TypeString,
							Optional:    true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.AccessReviewRecurranceTypeAnnual,
								msgraph.AccessReviewRecurranceTypeHalfYearly,
								msgraph.AccessReviewRecurranceTypeQuarterly,
								msgraph.AccessReviewRecurranceTypeMonthly,
								msgraph.AccessReviewRecurranceTypeWeekly,
							}, false),
						},
						"review_type": {
							Description: "Self reivew or specific reviewers, valid values are `Self`, `Reviewers`.",
							Type:        schema.TypeString,
							Optional:    true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.AccessReviewReviewerTypeSelf,
								msgraph.AccessReviewReviewerTypeReviewers,
							}, false),
						},
						"starting_on": {
							Description:  "This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date",
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsRFC3339Time,
						},
						"duration_in_days": {
							Description: "How many days each occurrence of the access review series will run.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"reviewer": {
							Description: "If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as members of a group, using a collection of singleUser and groupMembers.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        schemaUserSet(),
						},
						"is_access_recommendation_enabled": {
							Description: "Whether to show Show reviewer decision helpers. If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"is_approver_justification_required": {
							Description: "Whether a reviewer need provide a justification for their decision. Justification is visible to other reviewers and the requestor.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"access_review_timeout_behavior": {
							Description: "What actions the system takes if reviewers don't respond in time, valid values are `keepAccess`, `removeAcces`, `acceptAccessRecommendation`.",
							Type:        schema.TypeString,
							Optional:    true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.AccessReviewTimeoutBehaviorTypeKeepAccess,
								msgraph.AccessReviewTimeoutBehaviorTypeRemoveAccess,
								msgraph.AccessReviewTimeoutBehaviorTypeAcceptAccessRecommendation,
							}, false),
						},
					},
				},
			},
			"question": {
				Description:      "One ore more questions to the requestor.",
				Type:             schema.TypeList,
				DiffSuppressFunc: assignmentPolicyDiffSuppress,
				Optional:         true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_required": {
							Description: "Whether this question is required.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"sequence": {
							Description: "The sequence number of this question.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"choice": {
							Description: "Configuration of a choice to the question.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"actual_value": {
										Description: "The actual value of this choice",
										Type:        schema.TypeString,
										Required:    true,
									},
									"display_value": {
										Description: "The display text of this choice",
										Type:        schema.TypeList,
										MaxItems:    1,
										Required:    true,
										Elem:        schemaLocalizedContent(),
									},
								},
							},
						},
						"text": {
							Description: "The content of this question.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Elem:        schemaLocalizedContent(),
						},
					},
				},
			},
			"access_package_id": {
				Description:      "The ID of the access package that will contain the policy.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.UUID,
			},
		},
	}
}

func accessPackageAssignmentPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	var properties msgraph.AccessPackageAssignmentPolicy
	var err error
	if properties, err = buildAssignmentPolicyResourceData(ctx, d, meta); err != nil {
		return tf.ErrorDiagF(err, "Error building resource data from supplied parameters!")
	}

	accessPackageAssignmentPolicy, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Creating access package assignment policy %q", d.Get("display_name").(string))
	}

	d.SetId(*accessPackageAssignmentPolicy.ID)
	return accessPackageAssignmentPolicyResourceRead(ctx, d, meta)
}

func accessPackageAssignmentPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	var properties msgraph.AccessPackageAssignmentPolicy
	var err error
	if properties, err = buildAssignmentPolicyResourceData(ctx, d, meta); err != nil {
		return tf.ErrorDiagF(err, "Error building resource data from supplied parameters!")
	}

	objectId := d.Id()
	tf.LockByName(accessPackageAssignmentPolicyResourceName, objectId)
	defer tf.UnlockByName(accessPackageAssignmentPolicyResourceName, objectId)
	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update access package assignment policy with ID: %q", objectId)
	}

	return accessPackageAssignmentPolicyResourceRead(ctx, d, meta)
}

func accessPackageAssignmentPolicyResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	objectId := d.Id()
	accessPackageAssignmentPolicy, status, err := client.Get(ctx, objectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package assignment policy with Object ID %q was not found - removing from state!", objectId)
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving access package assignment policy with object ID: %q", objectId)
	}

	tf.Set(d, "display_name", accessPackageAssignmentPolicy.DisplayName)
	tf.Set(d, "access_package_id", accessPackageAssignmentPolicy.AccessPackageId)
	tf.Set(d, "description", accessPackageAssignmentPolicy.Description)
	tf.Set(d, "can_extend", accessPackageAssignmentPolicy.CanExtend)
	tf.Set(d, "duration_in_days", accessPackageAssignmentPolicy.DurationInDays)
	if expirationDate := accessPackageAssignmentPolicy.ExpirationDateTime; expirationDate != nil && !expirationDate.IsZero() {
		tf.Set(d, "expiration_date", expirationDate.UTC().Format(time.RFC3339))
	} else {
		tf.Set(d, "expiration_date", "")
	}

	tf.Set(d, "requestor_settings", flattenRequestorSettings(accessPackageAssignmentPolicy.RequestorSettings))
	tf.Set(d, "approval_settings", falttenApprovalSettings(accessPackageAssignmentPolicy.RequestApprovalSettings))
	tf.Set(d, "assignment_review_settings", flattenReviewSettings(accessPackageAssignmentPolicy.AccessReviewSettings))
	tf.Set(d, "question", flattenAssignmentPolicyQuestions(accessPackageAssignmentPolicy.Questions))

	return nil
}

func accessPackageAssignmentPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient
	accessPackageAssignmentPolicyId := d.Id()

	_, status, err := client.Get(ctx, accessPackageAssignmentPolicyId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(fmt.Errorf("Access package assignment policy was not found"), "id", "Retrieving user with object ID %q", accessPackageAssignmentPolicyId)
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving access package assignment policy with object ID %q", accessPackageAssignmentPolicyId)
	}

	status, err = client.Delete(ctx, accessPackageAssignmentPolicyId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting access package assignment policy with object ID %q, got status %d", accessPackageAssignmentPolicyId, status)
	}

	// Wait for user object to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, accessPackageAssignmentPolicyId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for deletion of access package assignment policy with object ID %q", accessPackageAssignmentPolicyId)
	}
	return nil
}

func buildAssignmentPolicyResourceData(ctx context.Context, d *schema.ResourceData, meta interface{}) (msgraph.AccessPackageAssignmentPolicy, error) {
	accessPackageClient := meta.(*clients.Client).IdentityGovernance.AccessPackageClient

	accessPackageId := d.Get("access_package_id").(string)
	_, status, err := accessPackageClient.Get(ctx, accessPackageId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access package with Object ID %q was not found - removing from state!", accessPackageId)
		}
		return msgraph.AccessPackageAssignmentPolicy{}, fmt.Errorf("Error retrieving access package with ID %v: %v", accessPackageId, err)
	}

	properties := msgraph.AccessPackageAssignmentPolicy{
		ID:              utils.String(d.Id()),
		DisplayName:     utils.String(d.Get("display_name").(string)),
		Description:     utils.String(d.Get("description").(string)),
		CanExtend:       utils.Bool(d.Get("can_extend").(bool)),
		DurationInDays:  utils.Int32(int32(d.Get("duration_in_days").(int))),
		Questions:       expandAccessPakcageAssignmentPolicyQuestions(d.Get("question").([]interface{})),
		AccessPackageId: utils.String(d.Get("access_package_id").(string)),
	}

	expirationDateValue := d.Get("expiration_date").(string)
	if expirationDateValue != "" {
		expirationDate, err := time.Parse(time.RFC3339, expirationDateValue)
		if err != nil {
			return properties, fmt.Errorf("Error converting expiration date %v to a valide date", expirationDate)
		}
		properties.ExpirationDateTime = &expirationDate
	}
	properties.RequestorSettings = buildAssignmentPolicyRequestorSettings(d.Get("requestor_settings").([]interface{}))
	properties.RequestApprovalSettings = buildAssignmentPolicyApprovalSettings(d.Get("approval_settings").([]interface{}))
	reviewSettingsStruct, err := buildAssignmentPolicyReviewSettings(d.Get("assignment_review_settings").([]interface{}))
	if err != nil {
		return properties, fmt.Errorf("Error building assignment_review_settings configuration: %v", err)
	}
	properties.AccessReviewSettings = reviewSettingsStruct

	return properties, nil
}

func assignmentPolicyDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if k == "approval_settings.#" && old == "1" && new == "0" {
		return true
	}

	if k == "requestor_settings.#" && old == "1" && new == "0" {
		return true
	}

	if k == "requestor_settings.0.scope_type" && old == msgraph.RequestorSettingsScopeTypeNoSubjects && len(new) == 0 {
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

func assignmentPolicyCustomDiff(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	if reviewSettings := diff.Get("assignment_review_settings").([]interface{}); len(reviewSettings) > 0 {
		reviewSetting := reviewSettings[0].(map[string]interface{})
		if reviewSetting["is_enabled"].(bool) &&
			(reviewSetting["duration_in_days"] == 0 ||
				len(reviewSetting["review_frequency"].(string)) == 0 ||
				len(reviewSetting["access_review_timeout_behavior"].(string)) == 0) {
			return fmt.Errorf("`duration_in_days`, `review_frequency`, `access_review_timeout_behavior` must be set when review is enabled")
		}
	}

	return nil
}
