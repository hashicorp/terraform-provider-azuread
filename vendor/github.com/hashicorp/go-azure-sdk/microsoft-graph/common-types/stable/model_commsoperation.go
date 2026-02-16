package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommsOperation interface {
	Entity
	CommsOperation() BaseCommsOperationImpl
}

var _ CommsOperation = BaseCommsOperationImpl{}

type BaseCommsOperationImpl struct {
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

func (s BaseCommsOperationImpl) CommsOperation() BaseCommsOperationImpl {
	return s
}

func (s BaseCommsOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ CommsOperation = RawCommsOperationImpl{}

// RawCommsOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCommsOperationImpl struct {
	commsOperation BaseCommsOperationImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawCommsOperationImpl) CommsOperation() BaseCommsOperationImpl {
	return s.commsOperation
}

func (s RawCommsOperationImpl) Entity() BaseEntityImpl {
	return s.commsOperation.Entity()
}

var _ json.Marshaler = BaseCommsOperationImpl{}

func (s BaseCommsOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseCommsOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseCommsOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseCommsOperationImpl: %+v", err)
	}

	delete(decoded, "resultInfo")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.commsOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseCommsOperationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalCommsOperationImplementation(input []byte) (CommsOperation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CommsOperation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.addLargeGalleryViewOperation") {
		var out AddLargeGalleryViewOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddLargeGalleryViewOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cancelMediaProcessingOperation") {
		var out CancelMediaProcessingOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CancelMediaProcessingOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inviteParticipantsOperation") {
		var out InviteParticipantsOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InviteParticipantsOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.muteParticipantOperation") {
		var out MuteParticipantOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MuteParticipantOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.playPromptOperation") {
		var out PlayPromptOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlayPromptOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.recordOperation") {
		var out RecordOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecordOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sendDtmfTonesOperation") {
		var out SendDtmfTonesOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SendDtmfTonesOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.startHoldMusicOperation") {
		var out StartHoldMusicOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StartHoldMusicOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stopHoldMusicOperation") {
		var out StopHoldMusicOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StopHoldMusicOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subscribeToToneOperation") {
		var out SubscribeToToneOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubscribeToToneOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unmuteParticipantOperation") {
		var out UnmuteParticipantOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnmuteParticipantOperation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.updateRecordingStatusOperation") {
		var out UpdateRecordingStatusOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UpdateRecordingStatusOperation: %+v", err)
		}
		return out, nil
	}

	var parent BaseCommsOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCommsOperationImpl: %+v", err)
	}

	return RawCommsOperationImpl{
		commsOperation: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
