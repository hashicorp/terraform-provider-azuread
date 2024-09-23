package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdateInstallScheduleType = WindowsUpdateActiveHoursInstall{}

type WindowsUpdateActiveHoursInstall struct {
	// Active Hours End
	ActiveHoursEnd *string `json:"activeHoursEnd,omitempty"`

	// Active Hours Start
	ActiveHoursStart *string `json:"activeHoursStart,omitempty"`

	// Fields inherited from WindowsUpdateInstallScheduleType

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdateActiveHoursInstall) WindowsUpdateInstallScheduleType() BaseWindowsUpdateInstallScheduleTypeImpl {
	return BaseWindowsUpdateInstallScheduleTypeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdateActiveHoursInstall{}

func (s WindowsUpdateActiveHoursInstall) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdateActiveHoursInstall
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdateActiveHoursInstall: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdateActiveHoursInstall: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdateActiveHoursInstall"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdateActiveHoursInstall: %+v", err)
	}

	return encoded, nil
}
