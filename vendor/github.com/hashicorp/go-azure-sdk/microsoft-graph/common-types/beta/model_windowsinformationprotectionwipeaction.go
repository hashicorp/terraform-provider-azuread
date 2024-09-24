package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsInformationProtectionWipeAction{}

type WindowsInformationProtectionWipeAction struct {
	// Last checkin time of the device that was targeted by this wipe action.
	LastCheckInDateTime *string `json:"lastCheckInDateTime,omitempty"`

	Status *ActionState `json:"status,omitempty"`

	// Targeted device Mac address.
	TargetedDeviceMacAddress nullable.Type[string] `json:"targetedDeviceMacAddress,omitempty"`

	// Targeted device name.
	TargetedDeviceName nullable.Type[string] `json:"targetedDeviceName,omitempty"`

	// The DeviceRegistrationId being targeted by this wipe action.
	TargetedDeviceRegistrationId nullable.Type[string] `json:"targetedDeviceRegistrationId,omitempty"`

	// The UserId being targeted by this wipe action.
	TargetedUserId nullable.Type[string] `json:"targetedUserId,omitempty"`

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

func (s WindowsInformationProtectionWipeAction) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsInformationProtectionWipeAction{}

func (s WindowsInformationProtectionWipeAction) MarshalJSON() ([]byte, error) {
	type wrapper WindowsInformationProtectionWipeAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsInformationProtectionWipeAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsInformationProtectionWipeAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsInformationProtectionWipeAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsInformationProtectionWipeAction: %+v", err)
	}

	return encoded, nil
}
