package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsDeviceHealthStatus{}

type ManagedTenantsDeviceHealthStatus struct {
	BlueScreenCount     nullable.Type[int64]  `json:"blueScreenCount,omitempty"`
	DeviceId            nullable.Type[string] `json:"deviceId,omitempty"`
	DeviceMake          nullable.Type[string] `json:"deviceMake,omitempty"`
	DeviceModel         nullable.Type[string] `json:"deviceModel,omitempty"`
	DeviceName          nullable.Type[string] `json:"deviceName,omitempty"`
	HealthStatus        nullable.Type[string] `json:"healthStatus,omitempty"`
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
	OsVersion           nullable.Type[string] `json:"osVersion,omitempty"`
	PrimaryDiskType     nullable.Type[string] `json:"primaryDiskType,omitempty"`
	RestartCount        nullable.Type[int64]  `json:"restartCount,omitempty"`
	TenantDisplayName   nullable.Type[string] `json:"tenantDisplayName,omitempty"`
	TenantId            nullable.Type[string] `json:"tenantId,omitempty"`
	TopProcesses        nullable.Type[string] `json:"topProcesses,omitempty"`

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

func (s ManagedTenantsDeviceHealthStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsDeviceHealthStatus{}

func (s ManagedTenantsDeviceHealthStatus) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsDeviceHealthStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsDeviceHealthStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsDeviceHealthStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.deviceHealthStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsDeviceHealthStatus: %+v", err)
	}

	return encoded, nil
}
