package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeployableContent interface {
	WindowsUpdatesDeployableContent() BaseWindowsUpdatesDeployableContentImpl
}

var _ WindowsUpdatesDeployableContent = BaseWindowsUpdatesDeployableContentImpl{}

type BaseWindowsUpdatesDeployableContentImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesDeployableContentImpl) WindowsUpdatesDeployableContent() BaseWindowsUpdatesDeployableContentImpl {
	return s
}

var _ WindowsUpdatesDeployableContent = RawWindowsUpdatesDeployableContentImpl{}

// RawWindowsUpdatesDeployableContentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesDeployableContentImpl struct {
	windowsUpdatesDeployableContent BaseWindowsUpdatesDeployableContentImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawWindowsUpdatesDeployableContentImpl) WindowsUpdatesDeployableContent() BaseWindowsUpdatesDeployableContentImpl {
	return s.windowsUpdatesDeployableContent
}

func UnmarshalWindowsUpdatesDeployableContentImplementation(input []byte) (WindowsUpdatesDeployableContent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesDeployableContent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.catalogContent") {
		var out WindowsUpdatesCatalogContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesCatalogContent: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesDeployableContentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesDeployableContentImpl: %+v", err)
	}

	return RawWindowsUpdatesDeployableContentImpl{
		windowsUpdatesDeployableContent: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
