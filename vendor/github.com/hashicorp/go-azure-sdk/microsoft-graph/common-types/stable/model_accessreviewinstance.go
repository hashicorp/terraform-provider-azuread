package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewInstance{}

type AccessReviewInstance struct {
	// Returns the collection of reviewers who were contacted to complete this review. While the reviewers and
	// fallbackReviewers properties of the accessReviewScheduleDefinition might specify group owners or managers as
	// reviewers, contactedReviewers returns their individual identities. Supports $select. Read-only.
	ContactedReviewers *[]AccessReviewReviewer `json:"contactedReviewers,omitempty"`

	// Each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not
	// yet reviewed.
	Decisions *[]AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`

	// DateTime when review instance is scheduled to end.The DatetimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports
	// $select. Read-only.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be
	// notified to take action if no users are found from the list of reviewers specified. This could occur when either the
	// group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but
	// a user's manager does not exist. Supports $select.
	FallbackReviewers *[]AccessReviewReviewerScope `json:"fallbackReviewers,omitempty"`

	// This collection of access review scopes is used to define who the reviewers are. Supports $select. For examples of
	// options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
	Reviewers *[]AccessReviewReviewerScope `json:"reviewers,omitempty"`

	// Created based on scope and instanceEnumerationScope at the accessReviewScheduleDefinition level. Defines the scope of
	// users reviewed in a group. Supports $select and $filter (contains only). Read-only.
	Scope *AccessReviewScope `json:"scope,omitempty"`

	// If the instance has multiple stages, this returns the collection of stages. A new stage will only be created when the
	// previous stage ends. The existence, number, and settings of stages on a review instance are created based on the
	// accessReviewStageSettings on the parent accessReviewScheduleDefinition.
	Stages *[]AccessReviewStage `json:"stages,omitempty"`

	// DateTime when review instance is scheduled to start. May be in the future. The DateTimeOffset type represents date
	// and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Supports $select. Read-only.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Specifies the status of an accessReview. Possible values: Initializing, NotStarted, Starting, InProgress, Completing,
	// Completed, AutoReviewing, and AutoReviewed. Supports $select, $orderby, and $filter (eq only). Read-only.
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

func (s AccessReviewInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewInstance{}

func (s AccessReviewInstance) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstance: %+v", err)
	}

	delete(decoded, "contactedReviewers")
	delete(decoded, "endDateTime")
	delete(decoded, "scope")
	delete(decoded, "startDateTime")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewInstance: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessReviewInstance{}

func (s *AccessReviewInstance) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ContactedReviewers *[]AccessReviewReviewer             `json:"contactedReviewers,omitempty"`
		Decisions          *[]AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`
		EndDateTime        nullable.Type[string]               `json:"endDateTime,omitempty"`
		FallbackReviewers  *[]AccessReviewReviewerScope        `json:"fallbackReviewers,omitempty"`
		Reviewers          *[]AccessReviewReviewerScope        `json:"reviewers,omitempty"`
		Stages             *[]AccessReviewStage                `json:"stages,omitempty"`
		StartDateTime      nullable.Type[string]               `json:"startDateTime,omitempty"`
		Status             nullable.Type[string]               `json:"status,omitempty"`
		Id                 *string                             `json:"id,omitempty"`
		ODataId            *string                             `json:"@odata.id,omitempty"`
		ODataType          *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ContactedReviewers = decoded.ContactedReviewers
	s.Decisions = decoded.Decisions
	s.EndDateTime = decoded.EndDateTime
	s.FallbackReviewers = decoded.FallbackReviewers
	s.Reviewers = decoded.Reviewers
	s.Stages = decoded.Stages
	s.StartDateTime = decoded.StartDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["scope"]; ok {
		impl, err := UnmarshalAccessReviewScopeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Scope' for 'AccessReviewInstance': %+v", err)
		}
		s.Scope = &impl
	}

	return nil
}
