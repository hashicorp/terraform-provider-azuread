package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentReviewSettings struct {
	// The default decision to apply if the access is not reviewed. The possible values are: keepAccess, removeAccess,
	// acceptAccessRecommendation, unknownFutureValue.
	ExpirationBehavior *AccessReviewExpirationBehavior `json:"expirationBehavior,omitempty"`

	// This collection specifies the users who will be the fallback reviewers when the primary reviewers don't respond.
	FallbackReviewers *[]SubjectSet `json:"fallbackReviewers,omitempty"`

	// If true, access reviews are required for assignments through this policy.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Specifies whether to display recommendations to the reviewer. The default value is true.
	IsRecommendationEnabled nullable.Type[bool] `json:"isRecommendationEnabled,omitempty"`

	// Specifies whether the reviewer must provide justification for the approval. The default value is true.
	IsReviewerJustificationRequired nullable.Type[bool] `json:"isReviewerJustificationRequired,omitempty"`

	// Specifies whether the principals can review their own assignments.
	IsSelfReview nullable.Type[bool] `json:"isSelfReview,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// This collection specifies the users or group of users who will review the access package assignments.
	PrimaryReviewers *[]SubjectSet `json:"primaryReviewers,omitempty"`

	// When the first review should start and how often it should recur.
	Schedule *EntitlementManagementSchedule `json:"schedule,omitempty"`
}

var _ json.Unmarshaler = &AccessPackageAssignmentReviewSettings{}

func (s *AccessPackageAssignmentReviewSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExpirationBehavior              *AccessReviewExpirationBehavior `json:"expirationBehavior,omitempty"`
		IsEnabled                       nullable.Type[bool]             `json:"isEnabled,omitempty"`
		IsRecommendationEnabled         nullable.Type[bool]             `json:"isRecommendationEnabled,omitempty"`
		IsReviewerJustificationRequired nullable.Type[bool]             `json:"isReviewerJustificationRequired,omitempty"`
		IsSelfReview                    nullable.Type[bool]             `json:"isSelfReview,omitempty"`
		ODataId                         *string                         `json:"@odata.id,omitempty"`
		ODataType                       *string                         `json:"@odata.type,omitempty"`
		Schedule                        *EntitlementManagementSchedule  `json:"schedule,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExpirationBehavior = decoded.ExpirationBehavior
	s.IsEnabled = decoded.IsEnabled
	s.IsRecommendationEnabled = decoded.IsRecommendationEnabled
	s.IsReviewerJustificationRequired = decoded.IsReviewerJustificationRequired
	s.IsSelfReview = decoded.IsSelfReview
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Schedule = decoded.Schedule

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageAssignmentReviewSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["fallbackReviewers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling FallbackReviewers into list []json.RawMessage: %+v", err)
		}

		output := make([]SubjectSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSubjectSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'FallbackReviewers' for 'AccessPackageAssignmentReviewSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.FallbackReviewers = &output
	}

	if v, ok := temp["primaryReviewers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PrimaryReviewers into list []json.RawMessage: %+v", err)
		}

		output := make([]SubjectSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSubjectSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PrimaryReviewers' for 'AccessPackageAssignmentReviewSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PrimaryReviewers = &output
	}

	return nil
}
