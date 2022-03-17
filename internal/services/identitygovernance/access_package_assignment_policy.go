package identitygovernance

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

func accessPackageAssignmentPolicyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: accessPackageAssignmentPolicyResourceCreate,
		ReadContext:   accessPackageAssignmentPolicyResourceRead,
		UpdateContext: accessPackageAssignmentPolicyResourceUpdate,
		DeleteContext: accessPackageAssignmentPolicyResourceDelete,

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

			"access_package_id": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.UUID,
			},


			"access_review_settings": {
				Description:      "Description of the accessPackageAssignmentPolicy",
				Type:             schema.TypeList,
				Optional:         true,
				Default: nil,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type: schema.TypeBool,
							Optional: true,
							Default: true,
						},
						"recurrence_type": {
							Type: schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.AccessReviewRecurranceTypeMonthly,
								msgraph.AccessReviewRecurranceTypeQuarterly,
								msgraph.AccessReviewRecurranceTypeHalfYearly,
								msgraph.AccessReviewRecurranceTypeHalfYearly,
							}, false),
						},
						"reviewer_type": {
							Type: schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.AccessReviewReviewerTypeSelf,
								msgraph.AccessReviewReviewerTypeReviewers,
							}, false),
							Default: msgraph.AccessReviewReviewerTypeSelf,
						},
						"start_date_time": {
						Type: schema.TypeString,
						Required: true,
						ValidateFunc: validation.IsRFC3339Time,
						},
						"duration_in_days": {
							Type: schema.TypeInt,
							Required: true,
						},
						"reviewers": {
							Type: schema.TypeList,
							Optional: true,
							//MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_type": {
										Type: schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{
											odata.ShortTypeSingleUser,
											odata.ShortTypeGroupMembers,
											odata.ShortTypeConnectedOrganizationMembers,
											odata.ShortTypeRequestorManager,
											odata.ShortTypeInternalSponsors,
											odata.ShortTypeExternalSponsors,
										}, false),
									},
									"id": {
										Type: schema.TypeString,
										Required: true,
										ValidateFunc: validation.IsUUID,
									},
									"description": {
										Type: schema.TypeString,
										Optional: true,
									},
									"backup": {
										Type: schema.TypeBool,
										Optional: true,
										Default: false,
										Description: "Specify whether user or group is a backup approver",
									},
									"manager_level": {
										Type: schema.TypeInt,
										Optional: true,
										ValidateFunc: validation.IntInSlice([]int{
											1,
											2,
											}),
									},
								},
							},
						},
					},
				},					
			},

			"can_extend": {
				Type:             schema.TypeBool,
				Optional: true,
				Default: false,
			},

			"created_by": {
				Description: "Who created the assignment",
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_date": {
				Description: "Sets if the access package hidden",
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"duration_in_days": {
				Description: "The number of days in which assignments from this policy last until they are expired.",
				Type:     schema.TypeInt,
				Required: true,
			},
			"expiration_date_time": {
				Description: "Sets if the access package hidden",
				Type:     schema.TypeString,
				Optional: true,
			},
			"modified_by": {
				Description: "Sets if the access package hidden",
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_date_time": {
				Description: "Sets if the access package hidden",
				Type:     schema.TypeString,
				Computed: true,
			},
			"request_approval_settings": {
				Description: "Sets if the access package hidden",
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"approval_required": {
							Type: schema.TypeBool,
							Optional: true,
							Default: false,
						},
						"approval_required_for_extension": {
							Type: schema.TypeBool,
							Optional: true,
							Default: false,
						},
						"requestor_justification_required": {
							Type: schema.TypeBool,
							Optional: true,
							RequiredWith: []string{"request_approval_settings.0.approval_required"}, // May bug
						},
						"approval_mode": {
							Type: schema.TypeString,
							Optional: true,
							RequiredWith: []string{"request_approval_settings.0.approval_required"}, // May bug
							ValidateFunc: validation.StringInSlice([]string{
								msgraph.ApprovalModeNoApproval,
								msgraph.ApprovalModeSerial,
								msgraph.ApprovalModeSingleStage,
							}, false),
							// May need a default here
						},
						"approval_stages": {
							Type: schema.TypeList,
							Optional: true,
							RequiredWith: []string{"request_approval_settings.0.approval_mode"},
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeout_in_days": {
										Type: schema.TypeInt,
										Required: true,
									},
									"approver_justification_required": {
										Type: schema.TypeBool,
										Optional: true,
										Default: false,
									},
									"escalation_enabled": {
										Type: schema.TypeBool,
										Optional: true,
										Default: false,
									},
									"escalation_time_in_minutes": {
										Type: schema.TypeInt,
										Optional: true,
										RequiredWith: []string{"request_approval_settings.0.approval_stages.0.escalation_enabled"},
									},
									"primary_approvers": {
										Type: schema.TypeList,
										Required: true,
										//MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"data_type": {
													Type: schema.TypeString,
													Required: true,
													ValidateFunc: validation.StringInSlice([]string{
														odata.ShortTypeSingleUser,
														odata.ShortTypeGroupMembers,
														odata.ShortTypeConnectedOrganizationMembers,
														odata.ShortTypeRequestorManager,
														odata.ShortTypeInternalSponsors,
														odata.ShortTypeExternalSponsors,
													}, false),
												},
												"id": {
													Type: schema.TypeString,
													Required: true,
													ValidateFunc: validation.IsUUID,
												},
												"description": {
													Type: schema.TypeString,
													Optional: true,
												},
												"backup": {
													Type: schema.TypeBool,
													Optional: true,
													Default: false,
													Description: "Specify whether user or group is a backup approver",
												},
												"manager_level": {
													Type: schema.TypeInt,
													Optional: true,
													ValidateFunc: validation.IntInSlice([]int{
														1,
														2,
														}),
												},
											},
										},
									},
									"escalation_approvers": {
										Type: schema.TypeList,
										Optional: true,
										//MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"data_type": {
													Type: schema.TypeString,
													Required: true,
													ValidateFunc: validation.StringInSlice([]string{
														odata.ShortTypeSingleUser,
														odata.ShortTypeGroupMembers,
														odata.ShortTypeConnectedOrganizationMembers,
														odata.ShortTypeRequestorManager,
														odata.ShortTypeInternalSponsors,
														odata.ShortTypeExternalSponsors,
													}, false),
												},
												"id": {
													Type: schema.TypeString,
													Required: true,
													ValidateFunc: validation.IsUUID,
												},
												"description": {
													Type: schema.TypeString,
													Optional: true,
												},
												"backup": {
													Type: schema.TypeBool,
													Optional: true,
													Default: false,
													Description: "Specify whether user or group is a backup approver",
												},
												"manager_level": {
													Type: schema.TypeInt,
													Optional: true,
													ValidateFunc: validation.IntInSlice([]int{
														1,
														2,
														}),
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
			"requestor_settings": {
				Description: "Sets if the access package hidden",
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope_type": {
							Type: schema.TypeString,
							Required: true,
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
						"accept_requests": {
							Type: schema.TypeBool,
							Required: true,
						},
						"allowed_requestors": { 
						Type: schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"data_type": {
									Type: schema.TypeString,
									Required: true,
									ValidateFunc: validation.StringInSlice([]string{
										odata.ShortTypeSingleUser,
										odata.ShortTypeGroupMembers,
										odata.ShortTypeConnectedOrganizationMembers,
										odata.ShortTypeRequestorManager,
										odata.ShortTypeInternalSponsors,
										odata.ShortTypeExternalSponsors,
									}, false),
								},
								"id": {
									Type: schema.TypeString,
									Required: true,
									ValidateFunc: validation.IsUUID,
								},
								"description": {
									Type: schema.TypeString,
									Optional: true,
								},
								"backup": {
									Type: schema.TypeBool,
									Computed: true,
								},
								"manager_level": {
									Type: schema.TypeInt,
									Computed: true,
									Default: nil,
								},
							},
						},
					},
				},
			},
			},
			"questions": {
				Description: "Sets if the access package hidden",
				Type:     schema.TypeList,
				//RequiredWith: []string{"request_approval_settings.0.approval_required"},
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeString,
							//ValidateFunc: validation.IsUUID,
							Computed: true,
						},
						"data_type": {
							Type: schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"TextInputQuestion",
								"MultipleChoiceQuestion",
							}, false),
						},
						"required": {
							Type: schema.TypeBool,
							Optional: true,
							Default: false,
						},
						"sequence": {
							Required: true,
							Type: schema.TypeInt,
						},
						"text": {
							Required: true,
							Type: schema.TypeList,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_text": {
										Type: schema.TypeString,
										Required: true,
									},
									"localized_texts": {
										Required: true,
										Type: schema.TypeList,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"text": {
													Type: schema.TypeString,
													Required: true,
												},
												"language_code": {
													Type: schema.TypeString,
													Required: true,
													ValidateDiagFunc: validate.ISO639Language,
												},
											},
										},
									},
								},
							},
						},
						"single_line_question": {
							Type: schema.TypeBool,
							Optional: true,
							//ConflictsWith: []string{"questions.0.allows_multiple_selection"},
						},
						"allows_multiple_selection": {
							Type: schema.TypeBool,
							Optional: true,
							//ConflictsWith: []string{"questions.0.single_line_question"},
						},
						"choices": {
							Type: schema.TypeList,
							Optional: true,
							//RequiredWith: []string{"questions.0.allows_multiple_selection"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"actual_value": {
										Type: schema.TypeString,
										Required: true,
									},
									"display_value": {
										Type: schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"default_text": {
													Type: schema.TypeString,
													Required: true,
												},
												"localized_texts": {
													Required: true,
													Type: schema.TypeList,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"text": {
																Type: schema.TypeString,
																Required: true,
															},
															"language_code": {
																Type: schema.TypeString,
																Required: true,
																ValidateDiagFunc: validate.ISO639Language,
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
				},
			},
		},
	}
}

func expandODataTypeShortToLong(shortType string) (odata.Type, error) {
	switch shortType {
	case odata.ShortTypeSingleUser:
		return odata.TypeSingleUser, nil
	case odata.ShortTypeGroupMembers:
		return odata.TypeGroupMembers, nil
	case odata.ShortTypeConnectedOrganizationMembers:
		return odata.TypeConnectedOrganizationMembers, nil
	case odata.ShortTypeRequestorManager:
		return odata.TypeRequestorManager, nil
	case odata.ShortTypeInternalSponsors:
		return odata.TypeInternalSponsors, nil
	case odata.ShortTypeExternalSponsors:
		return odata.TypeExternalSponsors, nil
	case "TextInputQuestion":
		return odata.TypeAccessPackageTextInputQuestion, nil
	case "MultipleChoiceQuestion":
		return odata.TypeAccessPackageMultipleChoiceQuestion, nil
	}
	return "", fmt.Errorf("No return value from data_type cast %v", utils.String(shortType))
}

func flattenODataTypeLongToShort(longType odata.Type) (string, error) {
	switch longType {
	case odata.TypeSingleUser:
		return odata.ShortTypeSingleUser, nil
	case odata.TypeGroupMembers:
		return odata.ShortTypeGroupMembers, nil
	case odata.TypeConnectedOrganizationMembers:
		return odata.ShortTypeConnectedOrganizationMembers, nil
	case odata.TypeRequestorManager:
		return odata.ShortTypeRequestorManager, nil
	case odata.TypeInternalSponsors:
		return odata.ShortTypeInternalSponsors, nil
	case odata.TypeExternalSponsors:
		return odata.ShortTypeExternalSponsors, nil
	case odata.TypeAccessPackageTextInputQuestion:
		return "TextInputQuestion", nil
	case odata.TypeAccessPackageMultipleChoiceQuestion:
		return "MultipleChoiceQuestion", nil
	}
	return "", fmt.Errorf("No return value from data_type cast %v", utils.String(longType))
}

// Generated via terrassist github.com/manicminer/hamilton/msgraph AccessPackageAssignmentPolicy
func expandAssignmentReviewSettingsPtr(input []interface{}) *msgraph.AssignmentReviewSettings {
	if len(input) == 0 || input[0] == nil {
			return nil
	}
	b := input[0].(map[string]interface{})
	output := &msgraph.AssignmentReviewSettings{
			AccessReviewTimeoutBehavior:     b["access_review_timeout_behavior"].(string),
			DurationInDays:                  utils.Int32(int32(b["duration_in_days"].(int32))),
			IsAccessRecommendationEnabled:   utils.Bool(b["access_recommendation_enabled"].(bool)),
			IsApprovalJustificationRequired: utils.Bool(b["approval_justification_required"].(bool)),
			IsEnabled:                       utils.Bool(b["enabled"].(bool)),
			RecurrenceType:                  b["recurrence_type"].(string),
			ReviewerType:                    b["reviewer_type"].(string),
			Reviewers:                       expandUserSetSlicePtr(b["reviewers"].([]interface{})),
			StartDateTime:                   expandTimePtr(b["start_date_time"].(string)),
	}
	return output
}
func expandTimePtr(input string) *time.Time {
	if len(input) == 0 {
			return nil
	}
    t, err := time.Parse(time.RFC3339, input)
    if err != nil {
        tf.ErrorDiagF(err, "Cannot convert string date to RFC3389 Time - Is your date string formatted correctly?")
		return nil
    }
	return &t
}
func expandUserSetSlicePtr(input []interface{}) *[]msgraph.UserSet {
	if len(input) == 0 {
			return nil
	}
	output := make([]msgraph.UserSet, 0)
	for _, elem := range input {
			elem := elem.(map[string]interface{})
			ODataType, err := expandODataTypeShortToLong(elem["data_type"].(string))
			if err != nil {
				// Cry about it
			}
			//We need to handle omission when its not applicable as cannot nil int so split out 
			if elem["manager_level"] == nil {
				output = append(output, msgraph.UserSet{
					Description:  utils.String(elem["description"].(string)),
					ID:           utils.String(elem["id"].(string)),
					IsBackup:     utils.Bool(elem["backup"].(bool)),
					ODataType:    &ODataType,
			})
			}
			if int32(elem["manager_level"].(int)) > 0 {
				output = append(output, msgraph.UserSet{
					Description:  utils.String(elem["description"].(string)),
					ID:           utils.String(elem["id"].(string)),
					IsBackup:     utils.Bool(elem["backup"].(bool)),
					ManagerLevel: utils.Int32(int32(elem["manager_level"].(int))),
					ODataType:    &ODataType,
			})
			}

	}
	return &output
}
func expandApprovalSettingsPtr(input []interface{}) *msgraph.ApprovalSettings {
	if len(input) == 0 || input[0] == nil {
			return nil
	}
	b := input[0].(map[string]interface{})
	output := &msgraph.ApprovalSettings{
			ApprovalMode:                     b["approval_mode"].(string),
			ApprovalStages:                   expandApprovalStageSlicePtr(b["approval_stages"].([]interface{})),
			IsApprovalRequired:               utils.Bool(b["approval_required"].(bool)),
			IsApprovalRequiredForExtension:   utils.Bool(b["approval_required_for_extension"].(bool)),
			IsRequestorJustificationRequired: utils.Bool(b["requestor_justification_required"].(bool)),
	}
	return output
}
func expandApprovalStageSlicePtr(input []interface{}) *[]msgraph.ApprovalStage {
	if len(input) == 0 {
			return nil
	}
	output := make([]msgraph.ApprovalStage, 0)
	for _, elem := range input {
			elem := elem.(map[string]interface{})
			output = append(output, msgraph.ApprovalStage{
					ApprovalStageTimeOutInDays:      utils.Int32(int32(elem["approval_stage_time_out_in_days"].(int32))),
					EscalationApprovers:             expandUserSetSlicePtr(elem["escalation_approvers"].([]interface{})),
					EscalationTimeInMinutes:         utils.Int32(int32(elem["escalation_time_in_minutes"].(int32))),
					IsApproverJustificationRequired: utils.Bool(elem["approver_justification_required"].(bool)),
					IsEscalationEnabled:             utils.Bool(elem["escalation_enabled"].(bool)),
					PrimaryApprovers:                expandUserSetSlicePtr(elem["primary_approvers"].([]interface{})),
			})
	}
	return &output
}
func expandRequestorSettingsPtr(input []interface{}) *msgraph.RequestorSettings {
	if len(input) == 0 || input[0] == nil {
			return nil
	}
	b := input[0].(map[string]interface{})
	output := &msgraph.RequestorSettings{
			AcceptRequests:    utils.Bool(b["accept_requests"].(bool)),
			AllowedRequestors: expandUserSetSlicePtr(b["allowed_requestors"].([]interface{})),
			ScopeType:         b["scope_type"].(string),
	}
	return output
}
func expandAccessPackageQuestionSlicePtr(input []interface{}) *[]msgraph.AccessPackageQuestion {
	if len(input) == 0 {
			return nil
	}

	output := make([]msgraph.AccessPackageQuestion, 0)
	for _, elem := range input {
			elem := elem.(map[string]interface{})
			ODataType, err := expandODataTypeShortToLong(elem["data_type"].(string))
			if err != nil {
				//Cry about it
			}
			output = append(output, msgraph.AccessPackageQuestion{
					Choices:              expandAccessPackageMultipleChoiceQuestionsSlicePtr(elem["choices"].([]interface{})),
					ID:                   utils.String(elem["id"].(string)),
					IsRequired:           utils.Bool(elem["required"].(bool)),
					IsSingleLineQuestion: utils.Bool(elem["single_line_question"].(bool)),
					ODataType:            &ODataType,
					Sequence:             utils.Int32(int32(elem["sequence"].(int))),
					Text:                 expandAccessPackageLocalizedContentPtr(elem["text"].([]interface{})),
			})
	}
	return &output
}
func expandAccessPackageLocalizedContentPtr(input []interface{}) *msgraph.AccessPackageLocalizedContent {
	if len(input) == 0 || input[0] == nil {
			return nil
	}
	b := input[0].(map[string]interface{})
	output := &msgraph.AccessPackageLocalizedContent{
			DefaultText:    utils.String(b["default_text"].(string)),
			LocalizedTexts: expandAccessPackageLocalizedTextsSlicePtr(b["localized_texts"].([]interface{})),
	}
	return output
}
func expandAccessPackageLocalizedTextsSlicePtr(input []interface{}) *[]msgraph.AccessPackageLocalizedTexts {
	if len(input) == 0 {
			return nil
	}
	output := make([]msgraph.AccessPackageLocalizedTexts, 0)
	for _, elem := range input {
			elem := elem.(map[string]interface{})
			output = append(output, msgraph.AccessPackageLocalizedTexts{
					LanguageCode: utils.String(elem["language_code"].(string)),
					Text:         utils.String(elem["text"].(string)),
			})
	}
	return &output
}
func expandAccessPackageMultipleChoiceQuestionsSlicePtr(input []interface{}) *[]msgraph.AccessPackageMultipleChoiceQuestions {
	if len(input) == 0 {
			return nil
	}
	output := make([]msgraph.AccessPackageMultipleChoiceQuestions, 0)
	for _, elem := range input {
			elem := elem.(map[string]interface{})
			output = append(output, msgraph.AccessPackageMultipleChoiceQuestions{
					ActualValue:  utils.String(elem["actual_value"].(string)),
					DisplayValue: expandAccessPackageLocalizedContentPtr(elem["display_value"].([]interface{})),
					//ODataType:    utils.String(elem["data_type"].(odata.Type)),
			})
	}
	return &output
}
func flattenAssignmentReviewSettingsPtr(input *msgraph.AssignmentReviewSettings) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	var isEnabled bool
	if input.IsEnabled != nil {
			isEnabled = *input.IsEnabled
	}
	var durationInDays int32
	if input.DurationInDays != nil {
			durationInDays = *input.DurationInDays
	}
	var isAccessRecommendationEnabled bool
	if input.IsAccessRecommendationEnabled != nil {
			isAccessRecommendationEnabled = *input.IsAccessRecommendationEnabled
	}
	var isApprovalJustificationRequired bool
	if input.IsApprovalJustificationRequired != nil {
			isApprovalJustificationRequired = *input.IsApprovalJustificationRequired
	}
	return []interface{}{map[string]interface{}{
			"access_review_timeout_behavior":     input.AccessReviewTimeoutBehavior,
			"duration_in_days":                   durationInDays,
			"access_recommendation_enabled":   isAccessRecommendationEnabled,
			"approval_justification_required": isApprovalJustificationRequired,
			"enabled":                         isEnabled,
			"recurrence_type":                    input.RecurrenceType,
			"reviewer_type":                      input.ReviewerType,
			"reviewers":                          flattenUserSetSlicePtr(input.Reviewers),
			"start_date_time":                    flattenTimePtr(input.StartDateTime),
	}}
}
func flattenTimePtr(input *time.Time) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	return []interface{}{map[string]interface{}{}}
}
func flattenUserSetSlicePtr(input *[]msgraph.UserSet) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
			var oDataType string
			if elem.ODataType != nil {
					oDataType = *elem.ODataType
			}
			var isBackup bool
			if elem.IsBackup != nil {
					isBackup = *elem.IsBackup
			}
			var id string
			if elem.ID != nil {
					id = *elem.ID
			}
			var description string
			if elem.Description != nil {
					description = *elem.Description
			}
			var managerLevel int32
			if elem.ManagerLevel != nil {
					managerLevel = *elem.ManagerLevel
			}
			output = append(output, map[string]interface{}{
					"description":   description,
					"id":            id,
					"backup":     isBackup,
					"manager_level": managerLevel,
					"data_type":   oDataType,
			})
	}
	return output
}
func flattenApprovalSettingsPtr(input *msgraph.ApprovalSettings) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	var isApprovalRequired bool
	if input.IsApprovalRequired != nil {
			isApprovalRequired = *input.IsApprovalRequired
	}
	var isApprovalRequiredForExtension bool
	if input.IsApprovalRequiredForExtension != nil {
			isApprovalRequiredForExtension = *input.IsApprovalRequiredForExtension
	}
	var isRequestorJustificationRequired bool
	if input.IsRequestorJustificationRequired != nil {
			isRequestorJustificationRequired = *input.IsRequestorJustificationRequired
	}
	return []interface{}{map[string]interface{}{
			"approval_mode":                       input.ApprovalMode,
			"approval_stages":                     flattenApprovalStageSlicePtr(input.ApprovalStages),
			"approval_required":                isApprovalRequired,
			"approval_required_for_extension":  isApprovalRequiredForExtension,
			"requestor_justification_required": isRequestorJustificationRequired,
	}}
}
func flattenApprovalStageSlicePtr(input *[]msgraph.ApprovalStage) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
			var approvalStageTimeOutInDays int32
			if elem.ApprovalStageTimeOutInDays != nil {
					approvalStageTimeOutInDays = *elem.ApprovalStageTimeOutInDays
			}
			var isApproverJustificationRequired bool
			if elem.IsApproverJustificationRequired != nil {
					isApproverJustificationRequired = *elem.IsApproverJustificationRequired
			}
			var isEscalationEnabled bool
			if elem.IsEscalationEnabled != nil {
					isEscalationEnabled = *elem.IsEscalationEnabled
			}
			var escalationTimeInMinutes int32
			if elem.EscalationTimeInMinutes != nil {
					escalationTimeInMinutes = *elem.EscalationTimeInMinutes
			}
			output = append(output, map[string]interface{}{
					"approval_stage_time_out_in_days":    approvalStageTimeOutInDays,
					"escalation_approvers":               flattenUserSetSlicePtr(elem.EscalationApprovers),
					"escalation_time_in_minutes":         escalationTimeInMinutes,
					"approver_justification_required": isApproverJustificationRequired,
					"escalation_enabled":              isEscalationEnabled,
					"primary_approvers":                  flattenUserSetSlicePtr(elem.PrimaryApprovers),
			})
	}
	return output
}
func flattenRequestorSettingsPtr(input *msgraph.RequestorSettings) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	var acceptRequests bool
	if input.AcceptRequests != nil {
			acceptRequests = *input.AcceptRequests
	}
	return []interface{}{map[string]interface{}{
			"accept_requests":    acceptRequests,
			"allowed_requestors": flattenUserSetSlicePtr(input.AllowedRequestors),
			"scope_type":         input.ScopeType,
	}}
}
func flattenAccessPackageQuestionSlicePtr(input *[]msgraph.AccessPackageQuestion) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		    var oDataTypeString string
			var err error
			if elem.ODataType != nil {
				oDataTypeString, err = flattenODataTypeLongToShort(*elem.ODataType)
				if err != nil {
					//Cry
				}
			}
			var id string
			if elem.ID != nil {
					id = *elem.ID
			}
			var isRequired bool
			if elem.IsRequired != nil {
					isRequired = *elem.IsRequired
			}
			var sequence int32
			if elem.Sequence != nil {
					sequence = *elem.Sequence
			}
			var isSingleLineQuestion bool
			if elem.IsSingleLineQuestion != nil {
					isSingleLineQuestion = *elem.IsSingleLineQuestion
			}
			output = append(output, map[string]interface{}{
					"choices":                 flattenAccessPackageMultipleChoiceQuestionsSlicePtr(elem.Choices),
					"id":                      id,
					"required":             isRequired,
					"single_line_question": isSingleLineQuestion,
					"data_type":             oDataTypeString,
					"sequence":                sequence,
					"text":                    flattenAccessPackageLocalizedContentPtr(elem.Text),
			})
	}
	return output
}
func flattenAccessPackageLocalizedContentPtr(input *msgraph.AccessPackageLocalizedContent) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	var defaultText string
	if input.DefaultText != nil {
			defaultText = *input.DefaultText
	}
	return []interface{}{map[string]interface{}{
			"default_text":    defaultText,
			"localized_texts": flattenAccessPackageLocalizedTextsSlicePtr(input.LocalizedTexts),
	}}
}
func flattenAccessPackageLocalizedTextsSlicePtr(input *[]msgraph.AccessPackageLocalizedTexts) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
			var text string
			if elem.Text != nil {
					text = *elem.Text
			}
			var languageCode string
			if elem.LanguageCode != nil {
					languageCode = *elem.LanguageCode
			}
			output = append(output, map[string]interface{}{
					"language_code": languageCode,
					"text":          text,
			})
	}
	return output
}
func flattenAccessPackageMultipleChoiceQuestionsSlicePtr(input *[]msgraph.AccessPackageMultipleChoiceQuestions) []interface{} {
	if input == nil {
			return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
			// var oDataType string
			// if elem.ODataType != nil {
			// 		oDataType = *elem.ODataType
			// }
			var actualValue string
			if elem.ActualValue != nil {
					actualValue = *elem.ActualValue
			}
			output = append(output, map[string]interface{}{
					"actual_value":  actualValue,
					"display_value": flattenAccessPackageLocalizedContentPtr(elem.DisplayValue),
					//"data_type":   oDataType,
			})
	}
	return output
}

// End terrassist

func accessPackageAssignmentPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	properties := msgraph.AccessPackageAssignmentPolicy{
		AccessPackageId:         utils.String(d.Get("access_package_id").(string)),
		AccessReviewSettings:    expandAssignmentReviewSettingsPtr(d.Get("access_review_settings").([]interface{})),
		CanExtend:               utils.Bool(d.Get("can_extend").(bool)),
		CreatedBy:               utils.String(d.Get("created_by").(string)),
		//CreatedDateTime:         expandTimePtr(d.Get("created_date").(string)),
		Description:             utils.String(d.Get("description").(string)),
		DisplayName:             utils.String(d.Get("display_name").(string)),
		DurationInDays:          utils.Int32(int32(d.Get("duration_in_days").(int))),
		ExpirationDateTime:      expandTimePtr(d.Get("expiration_date_time").(string)),
		ModifiedBy:              utils.String(d.Get("modified_by").(string)),
		//ModifiedDateTime:        expandTimePtr(d.Get("modified_date_time").(string)),
		Questions:               expandAccessPackageQuestionSlicePtr(d.Get("questions").([]interface{})),
		RequestApprovalSettings: expandApprovalSettingsPtr(d.Get("request_approval_settings").([]interface{})),
		RequestorSettings:       expandRequestorSettingsPtr(d.Get("requestor_settings").([]interface{})),
	}	

	accessPackageAssignmentPolicy, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create accessPackageAssignmentPolicy")
	}

	if accessPackageAssignmentPolicy.ID == nil || *accessPackageAssignmentPolicy.ID == "" {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Object ID returned for accessPackageAssignmentPolicy is nil/empty")
	}

	d.SetId(*accessPackageAssignmentPolicy.ID)

	return accessPackageAssignmentPolicyResourceRead(ctx, d, meta)
}

func accessPackageAssignmentPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	properties := msgraph.AccessPackageAssignmentPolicy{
		ID:						 utils.String(d.Id()),
		AccessPackageId:         utils.String(d.Get("access_package_id").(string)),
		AccessReviewSettings:    expandAssignmentReviewSettingsPtr(d.Get("access_review_settings").([]interface{})),
		CanExtend:               utils.Bool(d.Get("can_extend").(bool)),
		CreatedBy:               utils.String(d.Get("created_by").(string)),
		//CreatedDateTime:         expandTimePtr(d.Get("created_date").(string)),
		Description:             utils.String(d.Get("description").(string)),
		DisplayName:             utils.String(d.Get("display_name").(string)),
		DurationInDays:          utils.Int32(int32(d.Get("duration_in_days").(int))),
		ExpirationDateTime:      expandTimePtr(d.Get("expiration_date_time").(string)),
		ModifiedBy:              utils.String(d.Get("modified_by").(string)),
		//ModifiedDateTime:        expandTimePtr(d.Get("modified_date_time").(string)),
		Questions:               expandAccessPackageQuestionSlicePtr(d.Get("questions").([]interface{})),
		RequestApprovalSettings: expandApprovalSettingsPtr(d.Get("request_approval_settings").([]interface{})),
		RequestorSettings:       expandRequestorSettingsPtr(d.Get("requestor_settings").([]interface{})),
	}	

	if _, err := client.Update(ctx, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update accessPackageAssignmentPolicy with ID: %q", d.Id())
	}

	return nil
}

func accessPackageAssignmentPolicyResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	accessPackageAssignmentPolicy, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] AP Catalog with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving AP Catalog with object ID %q", d.Id())
	}

	var accessPackageId string
	if accessPackageAssignmentPolicy.AccessPackageId != nil {
			accessPackageId = *accessPackageAssignmentPolicy.AccessPackageId
	}
	var canExtend bool
	if accessPackageAssignmentPolicy.CanExtend != nil {
			canExtend = *accessPackageAssignmentPolicy.CanExtend
	}
	var createdBy string
	if accessPackageAssignmentPolicy.CreatedBy != nil {
			createdBy = *accessPackageAssignmentPolicy.CreatedBy
	}
	var description string
	if accessPackageAssignmentPolicy.Description != nil {
			description = *accessPackageAssignmentPolicy.Description
	}
	var displayName string
	if accessPackageAssignmentPolicy.DisplayName != nil {
			displayName = *accessPackageAssignmentPolicy.DisplayName
	}
	var durationInDays int32
	if accessPackageAssignmentPolicy.DurationInDays != nil {
			durationInDays = *accessPackageAssignmentPolicy.DurationInDays
	}
	var modifiedBy string
	if accessPackageAssignmentPolicy.ModifiedBy != nil {
			modifiedBy = *accessPackageAssignmentPolicy.ModifiedBy
	}

	tf.Set(d, "access_package_id", accessPackageId)
	tf.Set(d, "access_review_settings", flattenAssignmentReviewSettingsPtr(accessPackageAssignmentPolicy.AccessReviewSettings))
	tf.Set(d, "can_extend", canExtend)
	tf.Set(d, "created_by", createdBy)
	tf.Set(d, "created_date", flattenTimePtr(accessPackageAssignmentPolicy.CreatedDateTime))
	tf.Set(d, "description", description)
	tf.Set(d, "display_name", displayName)
	tf.Set(d, "duration_in_days", durationInDays)
	tf.Set(d, "expiration_date_time",  flattenTimePtr(accessPackageAssignmentPolicy.ExpirationDateTime))
	tf.Set(d, "modified_by", modifiedBy)
	tf.Set(d, "modified_date_time", flattenTimePtr(accessPackageAssignmentPolicy.ModifiedDateTime))
	tf.Set(d, "questions", flattenAccessPackageQuestionSlicePtr(accessPackageAssignmentPolicy.Questions))
	tf.Set(d, "request_approval_settings", flattenApprovalSettingsPtr(accessPackageAssignmentPolicy.RequestApprovalSettings))
	tf.Set(d, "requestor_settings", flattenRequestorSettingsPtr(accessPackageAssignmentPolicy.RequestorSettings))

	return nil
}

func accessPackageAssignmentPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).IdentityGovernance.AccessPackageAssignmentPolicyClient

	_, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Access Pacakge with ID %q already deleted", d.Id())
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving AP Catalog with ID %q", d.Id())
	}

	status, err = client.Delete(ctx, d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting AP Catalog with ID %q, got status %d", d.Id(), status)
	}

	return nil
}	
