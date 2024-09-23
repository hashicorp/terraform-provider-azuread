package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesContentFilter interface {
	WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl
}

var _ WindowsUpdatesContentFilter = BaseWindowsUpdatesContentFilterImpl{}

type BaseWindowsUpdatesContentFilterImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesContentFilterImpl) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return s
}

var _ WindowsUpdatesContentFilter = RawWindowsUpdatesContentFilterImpl{}

// RawWindowsUpdatesContentFilterImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesContentFilterImpl struct {
	windowsUpdatesContentFilter BaseWindowsUpdatesContentFilterImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawWindowsUpdatesContentFilterImpl) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return s.windowsUpdatesContentFilter
}

func UnmarshalWindowsUpdatesContentFilterImplementation(input []byte) (WindowsUpdatesContentFilter, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesContentFilter into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.softwareUpdateFilter") {
		var out WindowsUpdatesSoftwareUpdateFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesSoftwareUpdateFilter: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesContentFilterImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesContentFilterImpl: %+v", err)
	}

	return RawWindowsUpdatesContentFilterImpl{
		windowsUpdatesContentFilter: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
