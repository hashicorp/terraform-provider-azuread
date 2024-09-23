package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewStage{}

type AccessReviewStage struct {
	// Each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet
	// reviewed.
	Decisions *[]AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`

	// The date and time in ISO 8601 format and UTC time when the review stage is scheduled to end. This property is the
	// cumulative total of the durationInDays for all stages. Read-only.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers are
	// notified to take action if no users are found from the list of reviewers specified. This could occur when either the
	// group owner is specified as the reviewer but the group owner doesn't exist, or manager is specified as reviewer but a
	// user's manager doesn't exist.
	FallbackReviewers *[]AccessReviewReviewerScope `json:"fallbackReviewers,omitempty"`

	// This collection of access review scopes is used to define who the reviewers are. For examples of options for
	// assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
	Reviewers *[]AccessReviewReviewerScope `json:"reviewers,omitempty"`

	// The date and time in ISO 8601 format and UTC time when the review stage is scheduled to start. Read-only.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Specifies the status of an accessReviewStage. Possible values: Initializing, NotStarted, Starting, InProgress,
	// Completing, Completed, AutoReviewing, and AutoReviewed. Supports $orderby, and $filter (eq only). Read-only.
	Status nullable.Type[string] `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessReviewStage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewStage{}

func (s AccessReviewStage) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewStage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewStage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewStage: %+v", err)
	}

	delete(decoded, "endDateTime")
	delete(decoded, "startDateTime")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewStage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewStage: %+v", err)
	}

	return encoded, nil
}
