package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerBaseApprovalAttachment interface {
	PlannerBaseApprovalAttachment() BasePlannerBaseApprovalAttachmentImpl
}

var _ PlannerBaseApprovalAttachment = BasePlannerBaseApprovalAttachmentImpl{}

type BasePlannerBaseApprovalAttachmentImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Status of the approval. The possible values are: requested, approved, rejected, cancelled, unknownFutureValue.
	// Read-only.
	Status *PlannerApprovalStatus `json:"status,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePlannerBaseApprovalAttachmentImpl) PlannerBaseApprovalAttachment() BasePlannerBaseApprovalAttachmentImpl {
	return s
}

var _ PlannerBaseApprovalAttachment = RawPlannerBaseApprovalAttachmentImpl{}

// RawPlannerBaseApprovalAttachmentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPlannerBaseApprovalAttachmentImpl struct {
	plannerBaseApprovalAttachment BasePlannerBaseApprovalAttachmentImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawPlannerBaseApprovalAttachmentImpl) PlannerBaseApprovalAttachment() BasePlannerBaseApprovalAttachmentImpl {
	return s.plannerBaseApprovalAttachment
}

var _ json.Marshaler = BasePlannerBaseApprovalAttachmentImpl{}

func (s BasePlannerBaseApprovalAttachmentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePlannerBaseApprovalAttachmentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePlannerBaseApprovalAttachmentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePlannerBaseApprovalAttachmentImpl: %+v", err)
	}

	delete(decoded, "status")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePlannerBaseApprovalAttachmentImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPlannerBaseApprovalAttachmentImplementation(input []byte) (PlannerBaseApprovalAttachment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerBaseApprovalAttachment into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerBasicApprovalAttachment") {
		var out PlannerBasicApprovalAttachment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerBasicApprovalAttachment: %+v", err)
		}
		return out, nil
	}

	var parent BasePlannerBaseApprovalAttachmentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlannerBaseApprovalAttachmentImpl: %+v", err)
	}

	return RawPlannerBaseApprovalAttachmentImpl{
		plannerBaseApprovalAttachment: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
