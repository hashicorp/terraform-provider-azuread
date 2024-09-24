package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsDeviceAppPerformance{}

type ManagedTenantsDeviceAppPerformance struct {
	AppFriendlyName     nullable.Type[string] `json:"appFriendlyName,omitempty"`
	AppName             nullable.Type[string] `json:"appName,omitempty"`
	AppPublisher        nullable.Type[string] `json:"appPublisher,omitempty"`
	AppVersion          nullable.Type[string] `json:"appVersion,omitempty"`
	DeviceId            nullable.Type[string] `json:"deviceId,omitempty"`
	DeviceManufacturer  nullable.Type[string] `json:"deviceManufacturer,omitempty"`
	DeviceModel         nullable.Type[string] `json:"deviceModel,omitempty"`
	DeviceName          nullable.Type[string] `json:"deviceName,omitempty"`
	HealthStatus        nullable.Type[string] `json:"healthStatus,omitempty"`
	IsLatestUsedVersion nullable.Type[int64]  `json:"isLatestUsedVersion,omitempty"`
	IsMostUsedVersion   nullable.Type[int64]  `json:"isMostUsedVersion,omitempty"`
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
	TenantDisplayName   nullable.Type[string] `json:"tenantDisplayName,omitempty"`
	TenantId            nullable.Type[string] `json:"tenantId,omitempty"`
	TotalAppCrashCount  nullable.Type[int64]  `json:"totalAppCrashCount,omitempty"`
	TotalAppFreezeCount nullable.Type[int64]  `json:"totalAppFreezeCount,omitempty"`

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

func (s ManagedTenantsDeviceAppPerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsDeviceAppPerformance{}

func (s ManagedTenantsDeviceAppPerformance) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsDeviceAppPerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsDeviceAppPerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsDeviceAppPerformance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.deviceAppPerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsDeviceAppPerformance: %+v", err)
	}

	return encoded, nil
}
