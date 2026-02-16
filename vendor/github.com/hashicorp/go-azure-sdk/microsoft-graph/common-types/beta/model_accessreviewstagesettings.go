package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewStageSettings struct {
	// Indicate which decisions will go to the next stage. Can be a subset of Approve, Deny, Recommendation, or NotReviewed.
	// If not provided, all decisions will go to the next stage. Optional.
	DecisionsThatWillMoveToNextStage *[]string `json:"decisionsThatWillMoveToNextStage,omitempty"`

	// Defines the sequential or parallel order of the stages and depends on the stageId. Only sequential stages are
	// currently supported. For example, if stageId is 2, then dependsOn must be 1. If stageId is 1, don't specify
	// dependsOn. Required if stageId isn't 1.
	DependsOn *[]string `json:"dependsOn,omitempty"`

	// The duration of the stage. Required. NOTE: The cumulative value of this property across all stages 1. Will override
	// the instanceDurationInDays setting on the accessReviewScheduleDefinition object. 2. Can't exceed the length of one
	// recurrence. That is, if the review recurs weekly, the cumulative durationInDays can't exceed 7.
	DurationInDays int64 `json:"durationInDays"`

	// If provided, the fallback reviewers are asked to complete a review if the primary reviewers don't exist. For example,
	// if managers are selected as reviewers and a principal under review doesn't have a manager in Microsoft Entra ID, the
	// fallback reviewers are asked to review that principal. NOTE: The value of this property overrides the corresponding
	// setting on the accessReviewScheduleDefinition object.
	FallbackReviewers *[]AccessReviewReviewerScope `json:"fallbackReviewers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RecommendationInsightSettings *[]AccessReviewRecommendationInsightSetting `json:"recommendationInsightSettings,omitempty"`

	// Optional field. Indicates the time period of inactivity (with respect to the start date of the review instance) from
	// which that recommendations will be configured. The recommendation is to deny if the user is inactive during the look
	// back duration. For reviews of groups and Microsoft Entra roles, any duration is accepted. For reviews of
	// applications, 30 days is the maximum duration. If not specified, the duration is 30 days. NOTE: The value of this
	// property overrides the corresponding setting on the accessReviewScheduleDefinition object.
	RecommendationLookBackDuration nullable.Type[string] `json:"recommendationLookBackDuration,omitempty"`

	// Indicates whether showing recommendations to reviewers is enabled. Required. NOTE: The value of this property
	// overrides the corresponding setting on the accessReviewScheduleDefinition object.
	RecommendationsEnabled bool `json:"recommendationsEnabled"`

	// Defines who the reviewers are. If none is specified, the review is a self-review (users review their own access). For
	// examples of options for assigning reviewers, see Assign reviewers to your access review definition using the
	// Microsoft Graph API. NOTE: The value of this property overrides the corresponding setting on the
	// accessReviewScheduleDefinition.
	Reviewers *[]AccessReviewReviewerScope `json:"reviewers,omitempty"`

	// Unique identifier of the accessReviewStageSettings. The stageId is used in dependsOn property to indicate the stage
	// relationship. Required.
	StageId string `json:"stageId"`
}

var _ json.Unmarshaler = &AccessReviewStageSettings{}

func (s *AccessReviewStageSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DecisionsThatWillMoveToNextStage *[]string                    `json:"decisionsThatWillMoveToNextStage,omitempty"`
		DependsOn                        *[]string                    `json:"dependsOn,omitempty"`
		DurationInDays                   int64                        `json:"durationInDays"`
		FallbackReviewers                *[]AccessReviewReviewerScope `json:"fallbackReviewers,omitempty"`
		ODataId                          *string                      `json:"@odata.id,omitempty"`
		ODataType                        *string                      `json:"@odata.type,omitempty"`
		RecommendationLookBackDuration   nullable.Type[string]        `json:"recommendationLookBackDuration,omitempty"`
		RecommendationsEnabled           bool                         `json:"recommendationsEnabled"`
		Reviewers                        *[]AccessReviewReviewerScope `json:"reviewers,omitempty"`
		StageId                          string                       `json:"stageId"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DecisionsThatWillMoveToNextStage = decoded.DecisionsThatWillMoveToNextStage
	s.DependsOn = decoded.DependsOn
	s.DurationInDays = decoded.DurationInDays
	s.FallbackReviewers = decoded.FallbackReviewers
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecommendationLookBackDuration = decoded.RecommendationLookBackDuration
	s.RecommendationsEnabled = decoded.RecommendationsEnabled
	s.Reviewers = decoded.Reviewers
	s.StageId = decoded.StageId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewStageSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["recommendationInsightSettings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RecommendationInsightSettings into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessReviewRecommendationInsightSetting, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessReviewRecommendationInsightSettingImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RecommendationInsightSettings' for 'AccessReviewStageSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RecommendationInsightSettings = &output
	}

	return nil
}
