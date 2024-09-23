package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ComanagementEligibleDevice{}

type ComanagementEligibleDevice struct {
	// Device registration status.
	ClientRegistrationStatus *DeviceRegistrationState `json:"clientRegistrationStatus,omitempty"`

	// DeviceName
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Device type.
	DeviceType *DeviceType `json:"deviceType,omitempty"`

	// EntitySource
	EntitySource *int64 `json:"entitySource,omitempty"`

	// Management agent type.
	ManagementAgents *ManagementAgentType `json:"managementAgents,omitempty"`

	// Management state of device in Microsoft Intune.
	ManagementState *ManagementState `json:"managementState,omitempty"`

	// Manufacturer
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// MDMStatus
	MdmStatus nullable.Type[string] `json:"mdmStatus,omitempty"`

	// Model
	Model nullable.Type[string] `json:"model,omitempty"`

	// OSDescription
	OsDescription nullable.Type[string] `json:"osDescription,omitempty"`

	// OSVersion
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Owner type of device.
	OwnerType *OwnerType `json:"ownerType,omitempty"`

	// ReferenceId
	ReferenceId nullable.Type[string] `json:"referenceId,omitempty"`

	// SerialNumber
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	Status *ComanagementEligibleType `json:"status,omitempty"`

	// UPN
	Upn nullable.Type[string] `json:"upn,omitempty"`

	// UserEmail
	UserEmail nullable.Type[string] `json:"userEmail,omitempty"`

	// UserId
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// UserName
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

func (s ComanagementEligibleDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ComanagementEligibleDevice{}

func (s ComanagementEligibleDevice) MarshalJSON() ([]byte, error) {
	type wrapper ComanagementEligibleDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ComanagementEligibleDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ComanagementEligibleDevice: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.comanagementEligibleDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ComanagementEligibleDevice: %+v", err)
	}

	return encoded, nil
}
