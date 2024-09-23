package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HardwareConfigurationRunSummary{}

type HardwareConfigurationRunSummary struct {
	// Number of devices for which hardware configuration state is error
	ErrorDeviceCount *int64 `json:"errorDeviceCount,omitempty"`

	// Number of users for which hardware configuration state is error
	ErrorUserCount *int64 `json:"errorUserCount,omitempty"`

	// Number of devices for which hardware configuration found an issue
	FailedDeviceCount *int64 `json:"failedDeviceCount,omitempty"`

	// Number of users for which hardware configuration found an issue
	FailedUserCount *int64 `json:"failedUserCount,omitempty"`

	// Last run time for the configuration across all devices
	LastRunDateTime nullable.Type[string] `json:"lastRunDateTime,omitempty"`

	// Number of devices for which hardware configuration state is not applicable
	NotApplicableDeviceCount *int64 `json:"notApplicableDeviceCount,omitempty"`

	// Number of users for which hardware configuration state is not applicable
	NotApplicableUserCount *int64 `json:"notApplicableUserCount,omitempty"`

	// Number of devices for which hardware configuration is in pending state
	PendingDeviceCount *int64 `json:"pendingDeviceCount,omitempty"`

	// Number of users for which hardware configuration is in pending state
	PendingUserCount *int64 `json:"pendingUserCount,omitempty"`

	// Number of devices for which hardware configured without any issue
	SuccessfulDeviceCount *int64 `json:"successfulDeviceCount,omitempty"`

	// Number of users for which hardware configured without any issue
	SuccessfulUserCount *int64 `json:"successfulUserCount,omitempty"`

	// Number of devices for which hardware configuration state is unknown
	UnknownDeviceCount *int64 `json:"unknownDeviceCount,omitempty"`

	// Number of users for which hardware configuration state is unknown
	UnknownUserCount *int64 `json:"unknownUserCount,omitempty"`

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

func (s HardwareConfigurationRunSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HardwareConfigurationRunSummary{}

func (s HardwareConfigurationRunSummary) MarshalJSON() ([]byte, error) {
	type wrapper HardwareConfigurationRunSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HardwareConfigurationRunSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HardwareConfigurationRunSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hardwareConfigurationRunSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HardwareConfigurationRunSummary: %+v", err)
	}

	return encoded, nil
}
