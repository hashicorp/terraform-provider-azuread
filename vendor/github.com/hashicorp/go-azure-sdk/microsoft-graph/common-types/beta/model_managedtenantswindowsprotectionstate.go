package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsWindowsProtectionState{}

type ManagedTenantsWindowsProtectionState struct {
	// The anti-malware version for the managed device. Optional. Read-only.
	AntiMalwareVersion nullable.Type[string] `json:"antiMalwareVersion,omitempty"`

	// A flag indicating whether attention is required for the managed device. Optional. Read-only.
	AttentionRequired nullable.Type[bool] `json:"attentionRequired,omitempty"`

	// A flag indicating whether the managed device has been deleted. Optional. Read-only.
	DeviceDeleted nullable.Type[bool] `json:"deviceDeleted,omitempty"`

	// The date and time the device property has been refreshed. Optional. Read-only.
	DevicePropertyRefreshDateTime nullable.Type[string] `json:"devicePropertyRefreshDateTime,omitempty"`

	// The anti-virus engine version for the managed device. Optional. Read-only.
	EngineVersion nullable.Type[string] `json:"engineVersion,omitempty"`

	// A flag indicating whether quick scan is overdue for the managed device. Optional. Read-only.
	FullScanOverdue nullable.Type[bool] `json:"fullScanOverdue,omitempty"`

	// A flag indicating whether full scan is overdue for the managed device. Optional. Read-only.
	FullScanRequired nullable.Type[bool] `json:"fullScanRequired,omitempty"`

	// The date and time a full scan was completed. Optional. Read-only.
	LastFullScanDateTime nullable.Type[string] `json:"lastFullScanDateTime,omitempty"`

	// The version anti-malware version used to perform the last full scan. Optional. Read-only.
	LastFullScanSignatureVersion nullable.Type[string] `json:"lastFullScanSignatureVersion,omitempty"`

	// The date and time a quick scan was completed. Optional. Read-only.
	LastQuickScanDateTime nullable.Type[string] `json:"lastQuickScanDateTime,omitempty"`

	// The version anti-malware version used to perform the last full scan. Optional. Read-only.
	LastQuickScanSignatureVersion nullable.Type[string] `json:"lastQuickScanSignatureVersion,omitempty"`

	// Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
	LastRefreshedDateTime nullable.Type[string] `json:"lastRefreshedDateTime,omitempty"`

	// The date and time the protection state was last reported for the managed device. Optional. Read-only.
	LastReportedDateTime nullable.Type[string] `json:"lastReportedDateTime,omitempty"`

	// A flag indicating whether malware protection is enabled for the managed device. Optional. Read-only.
	MalwareProtectionEnabled nullable.Type[bool] `json:"malwareProtectionEnabled,omitempty"`

	// The health state for the managed device. Optional. Read-only.
	ManagedDeviceHealthState nullable.Type[string] `json:"managedDeviceHealthState,omitempty"`

	// The unique identifier for the managed device. Optional. Read-only.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The display name for the managed device. Optional. Read-only.
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// A flag indicating whether the network inspection system is enabled. Optional. Read-only.
	NetworkInspectionSystemEnabled nullable.Type[bool] `json:"networkInspectionSystemEnabled,omitempty"`

	// A flag indicating weather a quick scan is overdue. Optional. Read-only.
	QuickScanOverdue nullable.Type[bool] `json:"quickScanOverdue,omitempty"`

	// A flag indicating whether real time protection is enabled. Optional. Read-only.
	RealTimeProtectionEnabled nullable.Type[bool] `json:"realTimeProtectionEnabled,omitempty"`

	// A flag indicating whether a reboot is required. Optional. Read-only.
	RebootRequired nullable.Type[bool] `json:"rebootRequired,omitempty"`

	// A flag indicating whether an signature update is overdue. Optional. Read-only.
	SignatureUpdateOverdue nullable.Type[bool] `json:"signatureUpdateOverdue,omitempty"`

	// The signature version for the managed device. Optional. Read-only.
	SignatureVersion nullable.Type[string] `json:"signatureVersion,omitempty"`

	// The display name for the managed tenant. Optional. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Optional. Read-only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

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

func (s ManagedTenantsWindowsProtectionState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsWindowsProtectionState{}

func (s ManagedTenantsWindowsProtectionState) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsWindowsProtectionState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsWindowsProtectionState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsWindowsProtectionState: %+v", err)
	}

	delete(decoded, "antiMalwareVersion")
	delete(decoded, "attentionRequired")
	delete(decoded, "deviceDeleted")
	delete(decoded, "devicePropertyRefreshDateTime")
	delete(decoded, "engineVersion")
	delete(decoded, "fullScanOverdue")
	delete(decoded, "fullScanRequired")
	delete(decoded, "lastFullScanDateTime")
	delete(decoded, "lastFullScanSignatureVersion")
	delete(decoded, "lastQuickScanDateTime")
	delete(decoded, "lastQuickScanSignatureVersion")
	delete(decoded, "lastRefreshedDateTime")
	delete(decoded, "lastReportedDateTime")
	delete(decoded, "malwareProtectionEnabled")
	delete(decoded, "managedDeviceHealthState")
	delete(decoded, "managedDeviceId")
	delete(decoded, "managedDeviceName")
	delete(decoded, "networkInspectionSystemEnabled")
	delete(decoded, "quickScanOverdue")
	delete(decoded, "realTimeProtectionEnabled")
	delete(decoded, "rebootRequired")
	delete(decoded, "signatureUpdateOverdue")
	delete(decoded, "signatureVersion")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "tenantId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.windowsProtectionState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsWindowsProtectionState: %+v", err)
	}

	return encoded, nil
}
