package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceOverview{}

type ManagedDeviceOverview struct {
	// Distribution of Exchange Access State in Intune
	DeviceExchangeAccessStateSummary *DeviceExchangeAccessStateSummary `json:"deviceExchangeAccessStateSummary,omitempty"`

	// Device operating system summary.
	DeviceOperatingSystemSummary *DeviceOperatingSystemSummary `json:"deviceOperatingSystemSummary,omitempty"`

	// The number of devices enrolled in both MDM and EAS
	DualEnrolledDeviceCount *int64 `json:"dualEnrolledDeviceCount,omitempty"`

	// Total enrolled device count. Does not include PC devices managed via Intune PC Agent
	EnrolledDeviceCount *int64 `json:"enrolledDeviceCount,omitempty"`

	// The number of devices enrolled in MDM
	MdmEnrolledCount *int64 `json:"mdmEnrolledCount,omitempty"`

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

func (s ManagedDeviceOverview) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceOverview{}

func (s ManagedDeviceOverview) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceOverview
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceOverview: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceOverview: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDeviceOverview"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceOverview: %+v", err)
	}

	return encoded, nil
}
