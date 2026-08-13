package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OmaSetting interface {
	OmaSetting() BaseOmaSettingImpl
}

var _ OmaSetting = BaseOmaSettingImpl{}

type BaseOmaSettingImpl struct {
	// Description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display Name.
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// OMA.
	OmaUri *string `json:"omaUri,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOmaSettingImpl) OmaSetting() BaseOmaSettingImpl {
	return s
}

var _ OmaSetting = RawOmaSettingImpl{}

// RawOmaSettingImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawOmaSettingImpl struct {
	omaSetting BaseOmaSettingImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawOmaSettingImpl) OmaSetting() BaseOmaSettingImpl {
	return s.omaSetting
}

func (s RawOmaSettingImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalOmaSettingImplementation(input []byte) (OmaSetting, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OmaSetting into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.omaSettingBase64") {
		var out OmaSettingBase64
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OmaSettingBase64: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.omaSettingBoolean") {
		var out OmaSettingBoolean
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OmaSettingBoolean: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.omaSettingDateTime") {
		var out OmaSettingDateTime
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OmaSettingDateTime: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.omaSettingFloatingPoint") {
		var out OmaSettingFloatingPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OmaSettingFloatingPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.omaSettingInteger") {
		var out OmaSettingInteger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OmaSettingInteger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.omaSettingString") {
		var out OmaSettingString
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OmaSettingString: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.omaSettingStringXml") {
		var out OmaSettingStringXml
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OmaSettingStringXml: %+v", err)
		}
		return out, nil
	}

	var parent BaseOmaSettingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOmaSettingImpl: %+v", err)
	}

	return RawOmaSettingImpl{
		omaSetting: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
