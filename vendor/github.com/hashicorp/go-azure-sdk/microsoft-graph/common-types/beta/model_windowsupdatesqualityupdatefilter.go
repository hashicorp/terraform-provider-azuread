package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesWindowsUpdateFilter = WindowsUpdatesQualityUpdateFilter{}

type WindowsUpdatesQualityUpdateFilter struct {
	// Specifies the cadence for publishing quality updates of the filter. The possible values are: monthly, outOfBand,
	// unknownFutureValue.
	Cadence *WindowsUpdatesQualityUpdateCadence `json:"cadence,omitempty"`

	// Specifies the quality update classification of the filter. The possible values are: all, security, nonSecurity,
	// unknownFutureValue.
	Classification *WindowsUpdatesQualityUpdateClassification `json:"classification,omitempty"`

	// Fields inherited from WindowsUpdatesContentFilter

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesQualityUpdateFilter) WindowsUpdatesWindowsUpdateFilter() BaseWindowsUpdatesWindowsUpdateFilterImpl {
	return BaseWindowsUpdatesWindowsUpdateFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesQualityUpdateFilter) WindowsUpdatesSoftwareUpdateFilter() BaseWindowsUpdatesSoftwareUpdateFilterImpl {
	return BaseWindowsUpdatesSoftwareUpdateFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesQualityUpdateFilter) WindowsUpdatesContentFilter() BaseWindowsUpdatesContentFilterImpl {
	return BaseWindowsUpdatesContentFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesQualityUpdateFilter{}

func (s WindowsUpdatesQualityUpdateFilter) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesQualityUpdateFilter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesQualityUpdateFilter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesQualityUpdateFilter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.qualityUpdateFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesQualityUpdateFilter: %+v", err)
	}

	return encoded, nil
}
