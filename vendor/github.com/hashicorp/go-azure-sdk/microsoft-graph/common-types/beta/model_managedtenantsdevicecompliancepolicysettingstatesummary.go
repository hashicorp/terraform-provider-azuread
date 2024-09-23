package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsDeviceCompliancePolicySettingStateSummary{}

type ManagedTenantsDeviceCompliancePolicySettingStateSummary struct {
	// The number of devices in a conflict state. Optional. Read-only.
	ConflictDeviceCount nullable.Type[int64] `json:"conflictDeviceCount,omitempty"`

	// The number of devices in an error state. Optional. Read-only.
	ErrorDeviceCount nullable.Type[int64] `json:"errorDeviceCount,omitempty"`

	// The number of devices in a failed state. Optional. Read-only.
	FailedDeviceCount nullable.Type[int64] `json:"failedDeviceCount,omitempty"`

	// The identifer for the Microsoft Intune account. Required. Read-only.
	IntuneAccountId nullable.Type[string] `json:"intuneAccountId,omitempty"`

	// The identifier for the Intune setting. Optional. Read-only.
	IntuneSettingId nullable.Type[string] `json:"intuneSettingId,omitempty"`

	// Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
	LastRefreshedDateTime nullable.Type[string] `json:"lastRefreshedDateTime,omitempty"`

	// The number of devices in a not applicable state. Optional. Read-only.
	NotApplicableDeviceCount nullable.Type[int64] `json:"notApplicableDeviceCount,omitempty"`

	// The number of devices in a pending state. Optional. Read-only.
	PendingDeviceCount nullable.Type[int64] `json:"pendingDeviceCount,omitempty"`

	// The type for the device compliance policy. Optional. Read-only.
	PolicyType nullable.Type[string] `json:"policyType,omitempty"`

	// The name for the setting within the device compliance policy. Optional. Read-only.
	SettingName nullable.Type[string] `json:"settingName,omitempty"`

	// The number of devices in a succeeded state. Optional. Read-only.
	SucceededDeviceCount nullable.Type[int64] `json:"succeededDeviceCount,omitempty"`

	// The display name for the managed tenant. Required. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Required. Read-only.
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

func (s ManagedTenantsDeviceCompliancePolicySettingStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsDeviceCompliancePolicySettingStateSummary{}

func (s ManagedTenantsDeviceCompliancePolicySettingStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsDeviceCompliancePolicySettingStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsDeviceCompliancePolicySettingStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsDeviceCompliancePolicySettingStateSummary: %+v", err)
	}

	delete(decoded, "conflictDeviceCount")
	delete(decoded, "errorDeviceCount")
	delete(decoded, "failedDeviceCount")
	delete(decoded, "intuneAccountId")
	delete(decoded, "intuneSettingId")
	delete(decoded, "lastRefreshedDateTime")
	delete(decoded, "notApplicableDeviceCount")
	delete(decoded, "pendingDeviceCount")
	delete(decoded, "policyType")
	delete(decoded, "settingName")
	delete(decoded, "succeededDeviceCount")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "tenantId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.deviceCompliancePolicySettingStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsDeviceCompliancePolicySettingStateSummary: %+v", err)
	}

	return encoded, nil
}
