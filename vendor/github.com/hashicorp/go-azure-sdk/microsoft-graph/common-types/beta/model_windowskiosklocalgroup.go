package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskUser = WindowsKioskLocalGroup{}

type WindowsKioskLocalGroup struct {
	// The name of the local group that will be locked to this kiosk configuration
	GroupName *string `json:"groupName,omitempty"`

	// Fields inherited from WindowsKioskUser

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsKioskLocalGroup) WindowsKioskUser() BaseWindowsKioskUserImpl {
	return BaseWindowsKioskUserImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsKioskLocalGroup{}

func (s WindowsKioskLocalGroup) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskLocalGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskLocalGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskLocalGroup: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsKioskLocalGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskLocalGroup: %+v", err)
	}

	return encoded, nil
}
