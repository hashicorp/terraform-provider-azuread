package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewDecision{}

type AccessReviewDecision struct {
	// The feature- generated recommendation shown to the reviewer, one of: Approve, Deny, NotAvailable.
	AccessRecommendation nullable.Type[string] `json:"accessRecommendation,omitempty"`

	// The feature-generated ID of the access review.
	AccessReviewId nullable.Type[string] `json:"accessReviewId,omitempty"`

	// When the review completes, if the results were manually applied, the user identity of the user who applied the
	// decision. If the review was autoapplied, the userPrincipalName is empty.
	AppliedBy *UserIdentity `json:"appliedBy,omitempty"`

	// The date and time when the review decision was applied.
	AppliedDateTime nullable.Type[string] `json:"appliedDateTime,omitempty"`

	// The outcome of applying the decision, one of: NotApplied, Success, Failed, NotFound, NotSupported.
	ApplyResult nullable.Type[string] `json:"applyResult,omitempty"`

	// The reviewer's business justification, if supplied.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// The result of the review, one of NotReviewed, Deny, DontKnow or Approve.
	ReviewResult nullable.Type[string] `json:"reviewResult,omitempty"`

	// The identity of the reviewer. If the recommendation was used as the review, the userPrincipalName is empty.
	ReviewedBy *UserIdentity `json:"reviewedBy,omitempty"`

	ReviewedDateTime nullable.Type[string] `json:"reviewedDateTime,omitempty"`

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

func (s AccessReviewDecision) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewDecision{}

func (s AccessReviewDecision) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewDecision
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewDecision: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewDecision: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewDecision"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewDecision: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessReviewDecision{}

func (s *AccessReviewDecision) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessRecommendation nullable.Type[string] `json:"accessRecommendation,omitempty"`
		AccessReviewId       nullable.Type[string] `json:"accessReviewId,omitempty"`
		AppliedDateTime      nullable.Type[string] `json:"appliedDateTime,omitempty"`
		ApplyResult          nullable.Type[string] `json:"applyResult,omitempty"`
		Justification        nullable.Type[string] `json:"justification,omitempty"`
		ReviewResult         nullable.Type[string] `json:"reviewResult,omitempty"`
		ReviewedDateTime     nullable.Type[string] `json:"reviewedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessRecommendation = decoded.AccessRecommendation
	s.AccessReviewId = decoded.AccessReviewId
	s.AppliedDateTime = decoded.AppliedDateTime
	s.ApplyResult = decoded.ApplyResult
	s.Justification = decoded.Justification
	s.ReviewResult = decoded.ReviewResult
	s.ReviewedDateTime = decoded.ReviewedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewDecision into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appliedBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppliedBy' for 'AccessReviewDecision': %+v", err)
		}
		s.AppliedBy = &impl
	}

	if v, ok := temp["reviewedBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReviewedBy' for 'AccessReviewDecision': %+v", err)
		}
		s.ReviewedBy = &impl
	}

	return nil
}
