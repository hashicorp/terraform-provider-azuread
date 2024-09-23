package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SettingStateDeviceSummary{}

type SettingStateDeviceSummary struct {
	// Device Compliant count for the setting
	CompliantDeviceCount *int64 `json:"compliantDeviceCount,omitempty"`

	// Device conflict error count for the setting
	ConflictDeviceCount *int64 `json:"conflictDeviceCount,omitempty"`

	// Device error count for the setting
	ErrorDeviceCount *int64 `json:"errorDeviceCount,omitempty"`

	// Name of the InstancePath for the setting
	InstancePath nullable.Type[string] `json:"instancePath,omitempty"`

	// Device NonCompliant count for the setting
	NonCompliantDeviceCount *int64 `json:"nonCompliantDeviceCount,omitempty"`

	// Device Not Applicable count for the setting
	NotApplicableDeviceCount *int64 `json:"notApplicableDeviceCount,omitempty"`

	// Device Compliant count for the setting
	RemediatedDeviceCount *int64 `json:"remediatedDeviceCount,omitempty"`

	// Name of the setting
	SettingName nullable.Type[string] `json:"settingName,omitempty"`

	// Device Unkown count for the setting
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

func (s SettingStateDeviceSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SettingStateDeviceSummary{}

func (s SettingStateDeviceSummary) MarshalJSON() ([]byte, error) {
	type wrapper SettingStateDeviceSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SettingStateDeviceSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SettingStateDeviceSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.settingStateDeviceSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SettingStateDeviceSummary: %+v", err)
	}

	return encoded, nil
}
