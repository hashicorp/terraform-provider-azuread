package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesWindowsUpdateFilter = WindowsUpdatesDriverUpdateFilter{}

type WindowsUpdatesDriverUpdateFilter struct {

	// Fields inherited from WindowsUpdatesContentFilter

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesDriverUpdateFilter) WindowsUpdatesWindowsUpdateFilter() BaseWindowsUpdatesWindowsUpdateFilterImpl {
	return BaseWindowsUpdatesWindowsUpdateFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesDriverUpdateFilter) WindowsUpdatesSoftwareUpdateFilter() BaseWindowsUpdatesSoftwareUpdateFilterImpl {
	return BaseWindowsUpdatesSoftwareUpdateFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesDriverUpdateFilter) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return BaseWindowsUpdatesContentFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesDriverUpdateFilter{}

func (s WindowsUpdatesDriverUpdateFilter) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesDriverUpdateFilter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesDriverUpdateFilter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesDriverUpdateFilter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.driverUpdateFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesDriverUpdateFilter: %+v", err)
	}

	return encoded, nil
}
