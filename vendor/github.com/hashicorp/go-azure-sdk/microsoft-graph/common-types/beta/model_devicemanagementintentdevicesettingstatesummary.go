package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementIntentDeviceSettingStateSummary{}

type DeviceManagementIntentDeviceSettingStateSummary struct {
	// Number of compliant devices
	CompliantCount *int64 `json:"compliantCount,omitempty"`

	// Number of devices in conflict
	ConflictCount *int64 `json:"conflictCount,omitempty"`

	// Number of error devices
	ErrorCount *int64 `json:"errorCount,omitempty"`

	// Number of non compliant devices
	NonCompliantCount *int64 `json:"nonCompliantCount,omitempty"`

	// Number of not applicable devices
	NotApplicableCount *int64 `json:"notApplicableCount,omitempty"`

	// Number of remediated devices
	RemediatedCount *int64 `json:"remediatedCount,omitempty"`

	// Name of a setting
	SettingName nullable.Type[string] `json:"settingName,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementIntentDeviceSettingStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementIntentDeviceSettingStateSummary{}

func (s DeviceManagementIntentDeviceSettingStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementIntentDeviceSettingStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementIntentDeviceSettingStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementIntentDeviceSettingStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementIntentDeviceSettingStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementIntentDeviceSettingStateSummary: %+v", err)
	}

	return encoded, nil
}
