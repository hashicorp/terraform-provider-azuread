package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsAutopilotDeviceIdentity{}

type WindowsAutopilotDeviceIdentity struct {
	// Addressable user name.
	AddressableUserName nullable.Type[string] `json:"addressableUserName,omitempty"`

	// AAD Device ID - to be deprecated
	AzureActiveDirectoryDeviceId nullable.Type[string] `json:"azureActiveDirectoryDeviceId,omitempty"`

	// Display Name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	EnrollmentState *EnrollmentState `json:"enrollmentState,omitempty"`

	// Group Tag of the Windows autopilot device.
	GroupTag nullable.Type[string] `json:"groupTag,omitempty"`

	// Intune Last Contacted Date Time of the Windows autopilot device.
	LastContactedDateTime *string `json:"lastContactedDateTime,omitempty"`

	// Managed Device ID
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// Oem manufacturer of the Windows autopilot device.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// Model name of the Windows autopilot device.
	Model nullable.Type[string] `json:"model,omitempty"`

	// Product Key of the Windows autopilot device.
	ProductKey nullable.Type[string] `json:"productKey,omitempty"`

	// Purchase Order Identifier of the Windows autopilot device.
	PurchaseOrderIdentifier nullable.Type[string] `json:"purchaseOrderIdentifier,omitempty"`

	// Resource Name.
	ResourceName nullable.Type[string] `json:"resourceName,omitempty"`

	// Serial number of the Windows autopilot device.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// SKU Number
	SkuNumber nullable.Type[string] `json:"skuNumber,omitempty"`

	// System Family
	SystemFamily nullable.Type[string] `json:"systemFamily,omitempty"`

	// User Principal Name.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s WindowsAutopilotDeviceIdentity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsAutopilotDeviceIdentity{}

func (s WindowsAutopilotDeviceIdentity) MarshalJSON() ([]byte, error) {
	type wrapper WindowsAutopilotDeviceIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsAutopilotDeviceIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsAutopilotDeviceIdentity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsAutopilotDeviceIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsAutopilotDeviceIdentity: %+v", err)
	}

	return encoded, nil
}
