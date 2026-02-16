package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApprovalOperation{}

type ApprovalOperation struct {
	// The date and time when the operation was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The error if the operation failed.
	Error *PublicError `json:"error,omitempty"`

	// The date and time when this operation was most recently updated.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// The URL for the resource that was newly created or acted upon.
	ResourceLocation nullable.Type[string] `json:"resourceLocation,omitempty"`

	// The status of the operation. The possible values are: scheduled, inProgress, succeeded, failed, timeout,
	// unknownFutureValue.
	Status *ApprovalOperationStatus `json:"status,omitempty"`

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

func (s ApprovalOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApprovalOperation{}

func (s ApprovalOperation) MarshalJSON() ([]byte, error) {
	type wrapper ApprovalOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApprovalOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApprovalOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.approvalOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApprovalOperation: %+v", err)
	}

	return encoded, nil
}
