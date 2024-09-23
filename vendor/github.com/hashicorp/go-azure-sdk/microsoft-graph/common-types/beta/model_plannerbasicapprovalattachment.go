package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerBaseApprovalAttachment = PlannerBasicApprovalAttachment{}

type PlannerBasicApprovalAttachment struct {
	// Read-only. The identifier of the approval in the approval service.
	ApprovalId nullable.Type[string] `json:"approvalId,omitempty"`

	// Fields inherited from PlannerBaseApprovalAttachment

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

func (s PlannerBasicApprovalAttachment) PlannerBaseApprovalAttachment() BasePlannerBaseApprovalAttachmentImpl {
	return BasePlannerBaseApprovalAttachmentImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Status:    s.Status,
	}
}

var _ json.Marshaler = PlannerBasicApprovalAttachment{}

func (s PlannerBasicApprovalAttachment) MarshalJSON() ([]byte, error) {
	type wrapper PlannerBasicApprovalAttachment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerBasicApprovalAttachment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerBasicApprovalAttachment: %+v", err)
	}

	delete(decoded, "approvalId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerBasicApprovalAttachment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerBasicApprovalAttachment: %+v", err)
	}

	return encoded, nil
}
