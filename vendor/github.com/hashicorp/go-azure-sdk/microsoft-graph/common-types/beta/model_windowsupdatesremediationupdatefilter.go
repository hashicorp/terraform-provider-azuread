package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesWindowsUpdateFilter = WindowsUpdatesRemediationUpdateFilter{}

type WindowsUpdatesRemediationUpdateFilter struct {
	RemediationType *WindowsUpdatesRemediationType `json:"remediationType,omitempty"`

	// Fields inherited from WindowsUpdatesContentFilter

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesRemediationUpdateFilter) WindowsUpdatesWindowsUpdateFilter() BaseWindowsUpdatesWindowsUpdateFilterImpl {
	return BaseWindowsUpdatesWindowsUpdateFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesRemediationUpdateFilter) WindowsUpdatesSoftwareUpdateFilter() BaseWindowsUpdatesSoftwareUpdateFilterImpl {
	return BaseWindowsUpdatesSoftwareUpdateFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesRemediationUpdateFilter) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return BaseWindowsUpdatesContentFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesRemediationUpdateFilter{}

func (s WindowsUpdatesRemediationUpdateFilter) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesRemediationUpdateFilter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesRemediationUpdateFilter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesRemediationUpdateFilter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.remediationUpdateFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesRemediationUpdateFilter: %+v", err)
	}

	return encoded, nil
}
