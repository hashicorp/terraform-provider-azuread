package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNetworkUsageRule struct {
	// If set to true, corresponding managed apps will not be allowed to use cellular data when roaming.
	CellularDataBlockWhenRoaming *bool `json:"cellularDataBlockWhenRoaming,omitempty"`

	// If set to true, corresponding managed apps will not be allowed to use cellular data at any time.
	CellularDataBlocked *bool `json:"cellularDataBlocked,omitempty"`

	// Information about the managed apps that this rule is going to apply to. This collection can contain a maximum of 500
	// elements.
	ManagedApps *[]AppListItem `json:"managedApps,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &IosNetworkUsageRule{}

func (s *IosNetworkUsageRule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CellularDataBlockWhenRoaming *bool   `json:"cellularDataBlockWhenRoaming,omitempty"`
		CellularDataBlocked          *bool   `json:"cellularDataBlocked,omitempty"`
		ODataId                      *string `json:"@odata.id,omitempty"`
		ODataType                    *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CellularDataBlockWhenRoaming = decoded.CellularDataBlockWhenRoaming
	s.CellularDataBlocked = decoded.CellularDataBlocked
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosNetworkUsageRule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["managedApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ManagedApps into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ManagedApps' for 'IosNetworkUsageRule': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ManagedApps = &output
	}

	return nil
}
