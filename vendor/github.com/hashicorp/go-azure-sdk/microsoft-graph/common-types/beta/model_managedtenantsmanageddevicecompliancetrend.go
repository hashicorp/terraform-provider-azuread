package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedDeviceComplianceTrend{}

type ManagedTenantsManagedDeviceComplianceTrend struct {
	// The number of devices with a compliant status. Required. Read-only.
	CompliantDeviceCount nullable.Type[int64] `json:"compliantDeviceCount,omitempty"`

	// The number of devices manged by Configuration Manager. Required. Read-only.
	ConfigManagerDeviceCount nullable.Type[int64] `json:"configManagerDeviceCount,omitempty"`

	// The date and time compliance snapshot was performed. Required. Read-only.
	CountDateTime nullable.Type[string] `json:"countDateTime,omitempty"`

	// The number of devices with an error status. Required. Read-only.
	ErrorDeviceCount nullable.Type[int64] `json:"errorDeviceCount,omitempty"`

	// The number of devices that are in a grace period status. Required. Read-only.
	InGracePeriodDeviceCount nullable.Type[int64] `json:"inGracePeriodDeviceCount,omitempty"`

	// The number of devices that are in a non-compliant status. Required. Read-only.
	NoncompliantDeviceCount nullable.Type[int64] `json:"noncompliantDeviceCount,omitempty"`

	// The display name for the managed tenant. Optional. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Optional. Read-only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The number of devices in an unknown status. Required. Read-only.
	UnknownDeviceCount nullable.Type[int64] `json:"unknownDeviceCount,omitempty"`

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

func (s ManagedTenantsManagedDeviceComplianceTrend) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedDeviceComplianceTrend{}

func (s ManagedTenantsManagedDeviceComplianceTrend) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedDeviceComplianceTrend
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedDeviceComplianceTrend: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedDeviceComplianceTrend: %+v", err)
	}

	delete(decoded, "compliantDeviceCount")
	delete(decoded, "configManagerDeviceCount")
	delete(decoded, "countDateTime")
	delete(decoded, "errorDeviceCount")
	delete(decoded, "inGracePeriodDeviceCount")
	delete(decoded, "noncompliantDeviceCount")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "tenantId")
	delete(decoded, "unknownDeviceCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedDeviceComplianceTrend"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedDeviceComplianceTrend: %+v", err)
	}

	return encoded, nil
}
