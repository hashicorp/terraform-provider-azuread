package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsWorkFromAnywhereDevice{}

type UserExperienceAnalyticsWorkFromAnywhereDevice struct {
	// When TRUE, indicates the intune device's autopilot profile is assigned. When FALSE, indicates it's not Assigned.
	// Supports: $select, $OrderBy. Read-only.
	AutoPilotProfileAssigned nullable.Type[bool] `json:"autoPilotProfileAssigned,omitempty"`

	// When TRUE, indicates the intune device's autopilot is registered. When FALSE, indicates it's not registered.
	// Supports: $select, $OrderBy. Read-only.
	AutoPilotRegistered nullable.Type[bool] `json:"autoPilotRegistered,omitempty"`

	// The Azure Active Directory (Azure AD) device Id. Supports: $select, $OrderBy. Read-only.
	AzureAdDeviceId nullable.Type[string] `json:"azureAdDeviceId,omitempty"`

	// The work from anywhere device's Azure Active Directory (Azure AD) join type. Supports: $select, $OrderBy. Read-only.
	AzureAdJoinType nullable.Type[string] `json:"azureAdJoinType,omitempty"`

	// When TRUE, indicates the device's Azure Active Directory (Azure AD) is registered. When False, indicates it's not
	// registered. Supports: $select, $OrderBy. Read-only.
	AzureAdRegistered nullable.Type[bool] `json:"azureAdRegistered,omitempty"`

	// When TRUE, indicates the device's compliance policy is set to intune. When FALSE, indicates it's not set to intune.
	// Supports: $select, $OrderBy. Read-only.
	CompliancePolicySetToIntune nullable.Type[bool] `json:"compliancePolicySetToIntune,omitempty"`

	// The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`

	// When TRUE, indicates the device's Cloud Management Gateway for Configuration Manager is enabled. When FALSE,
	// indicates it's not enabled. Supports: $select, $OrderBy. Read-only.
	IsCloudManagedGatewayEnabled nullable.Type[bool] `json:"isCloudManagedGatewayEnabled,omitempty"`

	// The management agent of the device. Supports: $select, $OrderBy. Read-only.
	ManagedBy nullable.Type[string] `json:"managedBy,omitempty"`

	// The manufacturer name of the device. Supports: $select, $OrderBy. Read-only.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The model name of the device. Supports: $select, $OrderBy. Read-only.
	Model nullable.Type[string] `json:"model,omitempty"`

	// When TRUE, indicates OS check failed for device to upgrade to the latest version of windows. When FALSE, indicates
	// the check succeeded. Supports: $select, $OrderBy. Read-only.
	OsCheckFailed nullable.Type[bool] `json:"osCheckFailed,omitempty"`

	// The OS description of the device. Supports: $select, $OrderBy. Read-only.
	OsDescription nullable.Type[string] `json:"osDescription,omitempty"`

	// The OS version of the device. Supports: $select, $OrderBy. Read-only.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// When TRUE, indicates the device's other workloads is set to intune. When FALSE, indicates it's not set to intune.
	// Supports: $select, $OrderBy. Read-only.
	OtherWorkloadsSetToIntune nullable.Type[bool] `json:"otherWorkloadsSetToIntune,omitempty"`

	// Ownership of the device. Supports: $select, $OrderBy. Read-only.
	Ownership nullable.Type[string] `json:"ownership,omitempty"`

	// When TRUE, indicates processor hardware 64-bit architecture check failed for device to upgrade to the latest version
	// of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
	Processor64BitCheckFailed nullable.Type[bool] `json:"processor64BitCheckFailed,omitempty"`

	// When TRUE, indicates processor hardware core count check failed for device to upgrade to the latest version of
	// windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
	ProcessorCoreCountCheckFailed nullable.Type[bool] `json:"processorCoreCountCheckFailed,omitempty"`

	// When TRUE, indicates processor hardware family check failed for device to upgrade to the latest version of windows.
	// When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
	ProcessorFamilyCheckFailed nullable.Type[bool] `json:"processorFamilyCheckFailed,omitempty"`

	// When TRUE, indicates processor hardware speed check failed for device to upgrade to the latest version of windows.
	// When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
	ProcessorSpeedCheckFailed nullable.Type[bool] `json:"processorSpeedCheckFailed,omitempty"`

	// When TRUE, indicates RAM hardware check failed for device to upgrade to the latest version of windows. When FALSE,
	// indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
	RamCheckFailed nullable.Type[bool] `json:"ramCheckFailed,omitempty"`

	// When TRUE, indicates secure boot hardware check failed for device to upgrade to the latest version of windows. When
	// FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
	SecureBootCheckFailed nullable.Type[bool] `json:"secureBootCheckFailed,omitempty"`

	// The serial number of the device. Supports: $select, $OrderBy. Read-only.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// When TRUE, indicates storage hardware check failed for device to upgrade to the latest version of windows. When
	// FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
	StorageCheckFailed nullable.Type[bool] `json:"storageCheckFailed,omitempty"`

	// When TRUE, indicates the device is Tenant Attached. When FALSE, indicates it's not Tenant Attached. Supports:
	// $select, $OrderBy. Read-only.
	TenantAttached nullable.Type[bool] `json:"tenantAttached,omitempty"`

	// When TRUE, indicates Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade
	// to windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
	TpmCheckFailed nullable.Type[bool] `json:"tpmCheckFailed,omitempty"`

	// Work From Anywhere windows device upgrade eligibility status.
	UpgradeEligibility *OperatingSystemUpgradeEligibility `json:"upgradeEligibility,omitempty"`

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

func (s UserExperienceAnalyticsWorkFromAnywhereDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsWorkFromAnywhereDevice{}

func (s UserExperienceAnalyticsWorkFromAnywhereDevice) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsWorkFromAnywhereDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsWorkFromAnywhereDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsWorkFromAnywhereDevice: %+v", err)
	}

	delete(decoded, "autoPilotProfileAssigned")
	delete(decoded, "autoPilotRegistered")
	delete(decoded, "azureAdDeviceId")
	delete(decoded, "azureAdJoinType")
	delete(decoded, "azureAdRegistered")
	delete(decoded, "compliancePolicySetToIntune")
	delete(decoded, "deviceId")
	delete(decoded, "deviceName")
	delete(decoded, "isCloudManagedGatewayEnabled")
	delete(decoded, "managedBy")
	delete(decoded, "manufacturer")
	delete(decoded, "model")
	delete(decoded, "osCheckFailed")
	delete(decoded, "osDescription")
	delete(decoded, "osVersion")
	delete(decoded, "otherWorkloadsSetToIntune")
	delete(decoded, "ownership")
	delete(decoded, "processor64BitCheckFailed")
	delete(decoded, "processorCoreCountCheckFailed")
	delete(decoded, "processorFamilyCheckFailed")
	delete(decoded, "processorSpeedCheckFailed")
	delete(decoded, "ramCheckFailed")
	delete(decoded, "secureBootCheckFailed")
	delete(decoded, "serialNumber")
	delete(decoded, "storageCheckFailed")
	delete(decoded, "tenantAttached")
	delete(decoded, "tpmCheckFailed")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsWorkFromAnywhereDevice: %+v", err)
	}

	return encoded, nil
}
