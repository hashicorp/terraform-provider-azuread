package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HardwareConfigurationUserState{}

type HardwareConfigurationUserState struct {
	// Error device count for specific user.
	ErrorDeviceCount *int64 `json:"errorDeviceCount,omitempty"`

	// Failed device count for specific user.
	FailedDeviceCount *int64 `json:"failedDeviceCount,omitempty"`

	// Last timestamp when the hardware configuration executed
	LastStateUpdateDateTime *string `json:"lastStateUpdateDateTime,omitempty"`

	// Not applicable device count for specific user.
	NotApplicableDeviceCount *int64 `json:"notApplicableDeviceCount,omitempty"`

	// Pending device count for specific user.
	PendingDeviceCount *int64 `json:"pendingDeviceCount,omitempty"`

	// Success device count for specific user.
	SuccessfulDeviceCount *int64 `json:"successfulDeviceCount,omitempty"`

	// Unknown device count for specific user.
	UnknownDeviceCount *int64 `json:"unknownDeviceCount,omitempty"`

	// User Principal Name (UPN).
	Upn nullable.Type[string] `json:"upn,omitempty"`

	// User Email address.
	UserEmail nullable.Type[string] `json:"userEmail,omitempty"`

	// User name
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s HardwareConfigurationUserState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HardwareConfigurationUserState{}

func (s HardwareConfigurationUserState) MarshalJSON() ([]byte, error) {
	type wrapper HardwareConfigurationUserState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HardwareConfigurationUserState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HardwareConfigurationUserState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hardwareConfigurationUserState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HardwareConfigurationUserState: %+v", err)
	}

	return encoded, nil
}
