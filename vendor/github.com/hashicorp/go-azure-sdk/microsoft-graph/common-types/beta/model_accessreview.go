package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReview{}

type AccessReview struct {
	// The business flow template identifier. Required on create. This value is case sensitive.
	BusinessFlowTemplateId *string `json:"businessFlowTemplateId,omitempty"`

	// The user who created this review.
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`

	// The collection of decisions for this access review.
	Decisions *[]AccessReviewDecision `json:"decisions,omitempty"`

	// The description provided by the access review creator, to show to the reviewers.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The access review name. Required on create.
	DisplayName *string `json:"displayName,omitempty"`

	// The DateTime when the review is scheduled to end. This must be at least one day later than the start date. Required
	// on create.
	EndDateTime *string `json:"endDateTime,omitempty"`

	// The collection of access reviews instances past, present, and future, if this object is a recurring access review.
	Instances *[]AccessReview `json:"instances,omitempty"`

	// The collection of decisions for the caller, if the caller is a reviewer.
	MyDecisions *[]AccessReviewDecision `json:"myDecisions,omitempty"`

	// The object for which the access review is reviewing the access rights assignments. This identity can be the group for
	// the review of memberships of users in a group, or the app for a review of assignments of users to an application.
	// Required on create.
	ReviewedEntity Identity `json:"reviewedEntity"`

	// The relationship type of reviewer to the target object, one of: self, delegated, entityOwners. Required on create.
	ReviewerType nullable.Type[string] `json:"reviewerType,omitempty"`

	// The collection of reviewers for an access review, if access review reviewerType is of type delegated.
	Reviewers *[]AccessReviewReviewer `json:"reviewers,omitempty"`

	// The settings of an accessReview, see type definition below.
	Settings AccessReviewSettings `json:"settings"`

	// The date and time when the review is scheduled to be start. This date can be in the future. Required on create.
	StartDateTime *string `json:"startDateTime,omitempty"`

	// This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted,
	// Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
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

func (s AccessReview) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReview{}

func (s AccessReview) MarshalJSON() ([]byte, error) {
	type wrapper AccessReview
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReview: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReview: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReview"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReview: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessReview{}

func (s *AccessReview) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BusinessFlowTemplateId *string                 `json:"businessFlowTemplateId,omitempty"`
		Decisions              *[]AccessReviewDecision `json:"decisions,omitempty"`
		Description            nullable.Type[string]   `json:"description,omitempty"`
		DisplayName            *string                 `json:"displayName,omitempty"`
		EndDateTime            *string                 `json:"endDateTime,omitempty"`
		Instances              *[]AccessReview         `json:"instances,omitempty"`
		MyDecisions            *[]AccessReviewDecision `json:"myDecisions,omitempty"`
		ReviewerType           nullable.Type[string]   `json:"reviewerType,omitempty"`
		Reviewers              *[]AccessReviewReviewer `json:"reviewers,omitempty"`
		StartDateTime          *string                 `json:"startDateTime,omitempty"`
		Status                 nullable.Type[string]   `json:"status,omitempty"`
		Id                     *string                 `json:"id,omitempty"`
		ODataId                *string                 `json:"@odata.id,omitempty"`
		ODataType              *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BusinessFlowTemplateId = decoded.BusinessFlowTemplateId
	s.Decisions = decoded.Decisions
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.EndDateTime = decoded.EndDateTime
	s.Instances = decoded.Instances
	s.MyDecisions = decoded.MyDecisions
	s.ReviewerType = decoded.ReviewerType
	s.Reviewers = decoded.Reviewers
	s.StartDateTime = decoded.StartDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReview into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'AccessReview': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["reviewedEntity"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReviewedEntity' for 'AccessReview': %+v", err)
		}
		s.ReviewedEntity = impl
	}

	if v, ok := temp["settings"]; ok {
		impl, err := UnmarshalAccessReviewSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Settings' for 'AccessReview': %+v", err)
		}
		s.Settings = impl
	}

	return nil
}
