package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdateInstallScheduleType = WindowsUpdateScheduledInstall{}

type WindowsUpdateScheduledInstall struct {
	// Possible values for a weekly schedule.
	ScheduledInstallDay *WeeklySchedule `json:"scheduledInstallDay,omitempty"`

	// Scheduled Install Time during day
	ScheduledInstallTime *string `json:"scheduledInstallTime,omitempty"`

	// Fields inherited from WindowsUpdateInstallScheduleType

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdateScheduledInstall) WindowsUpdateInstallScheduleType() BaseWindowsUpdateInstallScheduleTypeImpl {
	return BaseWindowsUpdateInstallScheduleTypeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdateScheduledInstall{}

func (s WindowsUpdateScheduledInstall) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdateScheduledInstall
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdateScheduledInstall: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdateScheduledInstall: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdateScheduledInstall"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdateScheduledInstall: %+v", err)
	}

	return encoded, nil
}
