package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentReviewSettings struct {
	// The default decision to apply if the request isn't reviewed within the period specified in durationInDays. The
	// possible values are: acceptAccessRecommendation, keepAccess, removeAccess, and unknownFutureValue.
	AccessReviewTimeoutBehavior *AccessReviewTimeoutBehavior `json:"accessReviewTimeoutBehavior,omitempty"`

	// The number of days within which reviewers should provide input.
	DurationInDays nullable.Type[int64] `json:"durationInDays,omitempty"`

	// Specifies whether to display recommendations to the reviewer. The default value is true
	IsAccessRecommendationEnabled nullable.Type[bool] `json:"isAccessRecommendationEnabled,omitempty"`

	// Specifies whether the reviewer must provide justification for the approval. The default value is true.
	IsApprovalJustificationRequired nullable.Type[bool] `json:"isApprovalJustificationRequired,omitempty"`

	// If true, access reviews are required for assignments from this policy.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The interval for recurrence, such as monthly or quarterly.
	RecurrenceType nullable.Type[string] `json:"recurrenceType,omitempty"`

	// Who should be asked to do the review, either Self, Reviewers or Manager.
	ReviewerType nullable.Type[string] `json:"reviewerType,omitempty"`

	// If the reviewerType is Reviewers, this collection specifies the users who will be reviewers, either by ID or as
	// members of a group, using a collection of singleUser and groupMembers.
	Reviewers *[]UserSet `json:"reviewers,omitempty"`

	// When the first review should start.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
}

var _ json.Unmarshaler = &AssignmentReviewSettings{}

func (s *AssignmentReviewSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessReviewTimeoutBehavior     *AccessReviewTimeoutBehavior `json:"accessReviewTimeoutBehavior,omitempty"`
		DurationInDays                  nullable.Type[int64]         `json:"durationInDays,omitempty"`
		IsAccessRecommendationEnabled   nullable.Type[bool]          `json:"isAccessRecommendationEnabled,omitempty"`
		IsApprovalJustificationRequired nullable.Type[bool]          `json:"isApprovalJustificationRequired,omitempty"`
		IsEnabled                       nullable.Type[bool]          `json:"isEnabled,omitempty"`
		ODataId                         *string                      `json:"@odata.id,omitempty"`
		ODataType                       *string                      `json:"@odata.type,omitempty"`
		RecurrenceType                  nullable.Type[string]        `json:"recurrenceType,omitempty"`
		ReviewerType                    nullable.Type[string]        `json:"reviewerType,omitempty"`
		StartDateTime                   nullable.Type[string]        `json:"startDateTime,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessReviewTimeoutBehavior = decoded.AccessReviewTimeoutBehavior
	s.DurationInDays = decoded.DurationInDays
	s.IsAccessRecommendationEnabled = decoded.IsAccessRecommendationEnabled
	s.IsApprovalJustificationRequired = decoded.IsApprovalJustificationRequired
	s.IsEnabled = decoded.IsEnabled
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecurrenceType = decoded.RecurrenceType
	s.ReviewerType = decoded.ReviewerType
	s.StartDateTime = decoded.StartDateTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AssignmentReviewSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["reviewers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Reviewers into list []json.RawMessage: %+v", err)
		}

		output := make([]UserSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUserSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Reviewers' for 'AssignmentReviewSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Reviewers = &output
	}

	return nil
}
