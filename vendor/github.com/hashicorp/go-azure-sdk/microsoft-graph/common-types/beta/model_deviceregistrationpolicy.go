package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceRegistrationPolicy{}

type DeviceRegistrationPolicy struct {
	// Specifies the authorization policy for controlling registration of new devices using Microsoft Entra join within your
	// organization. Required. For more information, see What is a device identity?.
	AzureADJoin AzureADJoinPolicy `json:"azureADJoin"`

	// Specifies the authorization policy for controlling registration of new devices using Microsoft Entra registered
	// within your organization. Required. For more information, see What is a device identity?.
	AzureADRegistration AzureADRegistrationPolicy `json:"azureADRegistration"`

	// The description of the device registration policy. It's always set to Tenant-wide policy that manages intial
	// provisioning controls using quota restrictions, additional authentication and authorization checks. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the device registration policy. It's always set to Device Registration Policy. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Specifies the setting for Local Admin Password Solution (LAPS) within your organization.
	LocalAdminPassword *LocalAdminPasswordSettings `json:"localAdminPassword,omitempty"`

	MultiFactorAuthConfiguration *MultiFactorAuthConfiguration `json:"multiFactorAuthConfiguration,omitempty"`

	// Specifies the maximum number of devices that a user can have within your organization before blocking new device
	// registrations. The default value is set to 50. If this property isn't specified during the policy update operation,
	// it's automatically reset to 0 to indicate that users aren't allowed to join any devices.
	UserDeviceQuota *int64 `json:"userDeviceQuota,omitempty"`

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

func (s DeviceRegistrationPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceRegistrationPolicy{}

func (s DeviceRegistrationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper DeviceRegistrationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceRegistrationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceRegistrationPolicy: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "displayName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceRegistrationPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceRegistrationPolicy: %+v", err)
	}

	return encoded, nil
}
