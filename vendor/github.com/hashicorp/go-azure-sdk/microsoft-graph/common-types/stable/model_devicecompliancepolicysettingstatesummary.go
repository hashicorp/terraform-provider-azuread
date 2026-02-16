package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceCompliancePolicySettingStateSummary{}

type DeviceCompliancePolicySettingStateSummary struct {
	// Number of compliant devices
	CompliantDeviceCount *int64 `json:"compliantDeviceCount,omitempty"`

	// Number of conflict devices
	ConflictDeviceCount *int64 `json:"conflictDeviceCount,omitempty"`

	// Not yet documented
	DeviceComplianceSettingStates *[]DeviceComplianceSettingState `json:"deviceComplianceSettingStates,omitempty"`

	// Number of error devices
	ErrorDeviceCount *int64 `json:"errorDeviceCount,omitempty"`

	// Number of NonCompliant devices
	NonCompliantDeviceCount *int64 `json:"nonCompliantDeviceCount,omitempty"`

	// Number of not applicable devices
	NotApplicableDeviceCount *int64 `json:"notApplicableDeviceCount,omitempty"`

	// Supported platform types for policies.
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`

	// Number of remediated devices
	RemediatedDeviceCount *int64 `json:"remediatedDeviceCount,omitempty"`

	// The setting class name and property name.
	Setting nullable.Type[string] `json:"setting,omitempty"`

	// Name of the setting.
	SettingName nullable.Type[string] `json:"settingName,omitempty"`

	// Number of unknown devices
	UnknownDeviceCount *int64 `json:"unknownDeviceCount,omitempty"`

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

func (s DeviceCompliancePolicySettingStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceCompliancePolicySettingStateSummary{}

func (s DeviceCompliancePolicySettingStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper DeviceCompliancePolicySettingStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceCompliancePolicySettingStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceCompliancePolicySettingStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceCompliancePolicySettingStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceCompliancePolicySettingStateSummary: %+v", err)
	}

	return encoded, nil
}
