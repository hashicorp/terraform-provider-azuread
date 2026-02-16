package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CommsOperation = RecordOperation{}

type RecordOperation struct {
	// The access token required to retrieve the recording.
	RecordingAccessToken nullable.Type[string] `json:"recordingAccessToken,omitempty"`

	// The location where the recording is located.
	RecordingLocation nullable.Type[string] `json:"recordingLocation,omitempty"`

	// Fields inherited from CommsOperation

	// Unique Client Context string. Max limit is 256 chars.
	ClientContext nullable.Type[string] `json:"clientContext,omitempty"`

	// The result information. Read-only.
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`

	Status *OperationStatus `json:"status,omitempty"`

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

func (s RecordOperation) CommsOperation() BaseCommsOperationImpl {
	return BaseCommsOperationImpl{
		ClientContext: s.ClientContext,
		ResultInfo:    s.ResultInfo,
		Status:        s.Status,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s RecordOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RecordOperation{}

func (s RecordOperation) MarshalJSON() ([]byte, error) {
	type wrapper RecordOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RecordOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RecordOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.recordOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RecordOperation: %+v", err)
	}

	return encoded, nil
}
