package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApprovalStage{}

type ApprovalStage struct {
	// Indicates whether the stage is assigned to the calling user to review. Read-only.
	AssignedToMe nullable.Type[bool] `json:"assignedToMe,omitempty"`

	// The label provided by the policy creator to identify an approval stage. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The justification associated with the approval stage decision.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
	ReviewResult nullable.Type[string] `json:"reviewResult,omitempty"`

	// The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed.
	// Read-only.
	ReviewedBy *Identity `json:"reviewedBy,omitempty"`

	// The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in
	// UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ReviewedDateTime nullable.Type[string] `json:"reviewedDateTime,omitempty"`

	// The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
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

func (s ApprovalStage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApprovalStage{}

func (s ApprovalStage) MarshalJSON() ([]byte, error) {
	type wrapper ApprovalStage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApprovalStage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApprovalStage: %+v", err)
	}

	delete(decoded, "assignedToMe")
	delete(decoded, "displayName")
	delete(decoded, "reviewedBy")
	delete(decoded, "reviewedDateTime")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.approvalStage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApprovalStage: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ApprovalStage{}

func (s *ApprovalStage) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedToMe     nullable.Type[bool]   `json:"assignedToMe,omitempty"`
		DisplayName      nullable.Type[string] `json:"displayName,omitempty"`
		Justification    nullable.Type[string] `json:"justification,omitempty"`
		ReviewResult     nullable.Type[string] `json:"reviewResult,omitempty"`
		ReviewedDateTime nullable.Type[string] `json:"reviewedDateTime,omitempty"`
		Status           nullable.Type[string] `json:"status,omitempty"`
		Id               *string               `json:"id,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedToMe = decoded.AssignedToMe
	s.DisplayName = decoded.DisplayName
	s.Justification = decoded.Justification
	s.ReviewResult = decoded.ReviewResult
	s.ReviewedDateTime = decoded.ReviewedDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ApprovalStage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["reviewedBy"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReviewedBy' for 'ApprovalStage': %+v", err)
		}
		s.ReviewedBy = &impl
	}

	return nil
}
