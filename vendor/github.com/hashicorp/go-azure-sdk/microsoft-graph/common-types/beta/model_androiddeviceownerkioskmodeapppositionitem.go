package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeAppPositionItem struct {
	// Represents an item on the Android Device Owner Managed Home Screen (application, weblink or folder
	Item AndroidDeviceOwnerKioskModeHomeScreenItem `json:"item"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Position of the item on the grid. Valid values 0 to 9999999
	Position *int64 `json:"position,omitempty"`
}

var _ json.Unmarshaler = &AndroidDeviceOwnerKioskModeAppPositionItem{}

func (s *AndroidDeviceOwnerKioskModeAppPositionItem) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
		Position  *int64  `json:"position,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Position = decoded.Position

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AndroidDeviceOwnerKioskModeAppPositionItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["item"]; ok {
		impl, err := UnmarshalAndroidDeviceOwnerKioskModeHomeScreenItemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Item' for 'AndroidDeviceOwnerKioskModeAppPositionItem': %+v", err)
		}
		s.Item = impl
	}

	return nil
}
