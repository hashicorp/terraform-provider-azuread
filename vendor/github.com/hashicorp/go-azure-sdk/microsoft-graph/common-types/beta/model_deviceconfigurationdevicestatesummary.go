package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceConfigurationDeviceStateSummary{}

type DeviceConfigurationDeviceStateSummary struct {
	// Number of compliant devices
	CompliantDeviceCount *int64 `json:"compliantDeviceCount,omitempty"`

	// Number of conflict devices
	ConflictDeviceCount *int64 `json:"conflictDeviceCount,omitempty"`

	// Number of error devices
	ErrorDeviceCount *int64 `json:"errorDeviceCount,omitempty"`

	// Number of NonCompliant devices
	NonCompliantDeviceCount *int64 `json:"nonCompliantDeviceCount,omitempty"`

	// Number of not applicable devices
	NotApplicableDeviceCount *int64 `json:"notApplicableDeviceCount,omitempty"`

	// Number of remediated devices
	RemediatedDeviceCount *int64 `json:"remediatedDeviceCount,omitempty"`

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

func (s DeviceConfigurationDeviceStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceConfigurationDeviceStateSummary{}

func (s DeviceConfigurationDeviceStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper DeviceConfigurationDeviceStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceConfigurationDeviceStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceConfigurationDeviceStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceConfigurationDeviceStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceConfigurationDeviceStateSummary: %+v", err)
	}

	return encoded, nil
}
