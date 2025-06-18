package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesWindowsUpdateFilter interface {
	WindowsUpdatesContentFilter
	WindowsUpdatesSoftwareUpdateFilter
	WindowsUpdatesWindowsUpdateFilter() BaseWindowsUpdatesWindowsUpdateFilterImpl
}

var _ WindowsUpdatesWindowsUpdateFilter = BaseWindowsUpdatesWindowsUpdateFilterImpl{}

type BaseWindowsUpdatesWindowsUpdateFilterImpl struct {

	// Fields inherited from WindowsUpdatesContentFilter

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesWindowsUpdateFilterImpl) WindowsUpdatesWindowsUpdateFilter() BaseWindowsUpdatesWindowsUpdateFilterImpl {
	return s
}

func (s BaseWindowsUpdatesWindowsUpdateFilterImpl) WindowsUpdatesSoftwareUpdateFilter() BaseWindowsUpdatesSoftwareUpdateFilterImpl {
	return BaseWindowsUpdatesSoftwareUpdateFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s BaseWindowsUpdatesWindowsUpdateFilterImpl) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return BaseWindowsUpdatesContentFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsUpdatesWindowsUpdateFilter = RawWindowsUpdatesWindowsUpdateFilterImpl{}

// RawWindowsUpdatesWindowsUpdateFilterImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesWindowsUpdateFilterImpl struct {
	windowsUpdatesWindowsUpdateFilter BaseWindowsUpdatesWindowsUpdateFilterImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawWindowsUpdatesWindowsUpdateFilterImpl) WindowsUpdatesWindowsUpdateFilter() BaseWindowsUpdatesWindowsUpdateFilterImpl {
	return s.windowsUpdatesWindowsUpdateFilter
}

func (s RawWindowsUpdatesWindowsUpdateFilterImpl) WindowsUpdatesSoftwareUpdateFilter() BaseWindowsUpdatesSoftwareUpdateFilterImpl {
	return s.windowsUpdatesWindowsUpdateFilter.WindowsUpdatesSoftwareUpdateFilter()
}

func (s RawWindowsUpdatesWindowsUpdateFilterImpl) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return s.windowsUpdatesWindowsUpdateFilter.WindowsUpdatesContentFilter()
}

var _ json.Marshaler = BaseWindowsUpdatesWindowsUpdateFilterImpl{}

func (s BaseWindowsUpdatesWindowsUpdateFilterImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsUpdatesWindowsUpdateFilterImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsUpdatesWindowsUpdateFilterImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsUpdatesWindowsUpdateFilterImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.windowsUpdateFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsUpdatesWindowsUpdateFilterImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsUpdatesWindowsUpdateFilterImplementation(input []byte) (WindowsUpdatesWindowsUpdateFilter, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesWindowsUpdateFilter into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.driverUpdateFilter") {
		var out WindowsUpdatesDriverUpdateFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesDriverUpdateFilter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.qualityUpdateFilter") {
		var out WindowsUpdatesQualityUpdateFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesQualityUpdateFilter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.remediationUpdateFilter") {
		var out WindowsUpdatesRemediationUpdateFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesRemediationUpdateFilter: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesWindowsUpdateFilterImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesWindowsUpdateFilterImpl: %+v", err)
	}

	return RawWindowsUpdatesWindowsUpdateFilterImpl{
		windowsUpdatesWindowsUpdateFilter: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
