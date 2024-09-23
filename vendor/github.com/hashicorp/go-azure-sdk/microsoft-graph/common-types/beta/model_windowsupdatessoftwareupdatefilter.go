package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSoftwareUpdateFilter interface {
	WindowsUpdatesContentFilter
	WindowsUpdatesSoftwareUpdateFilter() BaseWindowsUpdatesSoftwareUpdateFilterImpl
}

var _ WindowsUpdatesSoftwareUpdateFilter = BaseWindowsUpdatesSoftwareUpdateFilterImpl{}

type BaseWindowsUpdatesSoftwareUpdateFilterImpl struct {

	// Fields inherited from WindowsUpdatesContentFilter

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesSoftwareUpdateFilterImpl) WindowsUpdatesSoftwareUpdateFilter() BaseWindowsUpdatesSoftwareUpdateFilterImpl {
	return s
}

func (s BaseWindowsUpdatesSoftwareUpdateFilterImpl) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return BaseWindowsUpdatesContentFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsUpdatesSoftwareUpdateFilter = RawWindowsUpdatesSoftwareUpdateFilterImpl{}

// RawWindowsUpdatesSoftwareUpdateFilterImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesSoftwareUpdateFilterImpl struct {
	windowsUpdatesSoftwareUpdateFilter BaseWindowsUpdatesSoftwareUpdateFilterImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawWindowsUpdatesSoftwareUpdateFilterImpl) WindowsUpdatesSoftwareUpdateFilter() BaseWindowsUpdatesSoftwareUpdateFilterImpl {
	return s.windowsUpdatesSoftwareUpdateFilter
}

func (s RawWindowsUpdatesSoftwareUpdateFilterImpl) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return s.windowsUpdatesSoftwareUpdateFilter.WindowsUpdatesContentFilter()
}

var _ json.Marshaler = BaseWindowsUpdatesSoftwareUpdateFilterImpl{}

func (s BaseWindowsUpdatesSoftwareUpdateFilterImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsUpdatesSoftwareUpdateFilterImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsUpdatesSoftwareUpdateFilterImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsUpdatesSoftwareUpdateFilterImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.softwareUpdateFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsUpdatesSoftwareUpdateFilterImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsUpdatesSoftwareUpdateFilterImplementation(input []byte) (WindowsUpdatesSoftwareUpdateFilter, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesSoftwareUpdateFilter into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.windowsUpdateFilter") {
		var out WindowsUpdatesWindowsUpdateFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesWindowsUpdateFilter: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesSoftwareUpdateFilterImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesSoftwareUpdateFilterImpl: %+v", err)
	}

	return RawWindowsUpdatesSoftwareUpdateFilterImpl{
		windowsUpdatesSoftwareUpdateFilter: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
