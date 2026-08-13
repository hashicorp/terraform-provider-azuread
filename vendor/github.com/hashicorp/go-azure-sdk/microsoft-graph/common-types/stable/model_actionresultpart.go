package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionResultPart interface {
	ActionResultPart() BaseActionResultPartImpl
}

var _ ActionResultPart = BaseActionResultPartImpl{}

type BaseActionResultPartImpl struct {
	// The error that occurred, if any, during the bulk operation.
	Error *PublicError `json:"error,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseActionResultPartImpl) ActionResultPart() BaseActionResultPartImpl {
	return s
}

var _ ActionResultPart = RawActionResultPartImpl{}

// RawActionResultPartImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawActionResultPartImpl struct {
	actionResultPart BaseActionResultPartImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawActionResultPartImpl) ActionResultPart() BaseActionResultPartImpl {
	return s.actionResultPart
}

func (s RawActionResultPartImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalActionResultPartImplementation(input []byte) (ActionResultPart, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ActionResultPart into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.aadUserConversationMemberResult") {
		var out AadUserConversationMemberResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AadUserConversationMemberResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseActionResultPartImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseActionResultPartImpl: %+v", err)
	}

	return RawActionResultPartImpl{
		actionResultPart: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
