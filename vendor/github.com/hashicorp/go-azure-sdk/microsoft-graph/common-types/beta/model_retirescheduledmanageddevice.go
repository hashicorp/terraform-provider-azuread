package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDevice struct {
	ComplianceState *ComplianceStatus `json:"complianceState,omitempty"`

	// Device Compliance PolicyId
	DeviceCompliancePolicyId nullable.Type[string] `json:"deviceCompliancePolicyId,omitempty"`

	// Device Compliance Policy Name
	DeviceCompliancePolicyName nullable.Type[string] `json:"deviceCompliancePolicyName,omitempty"`

	// Device type.
	DeviceType *DeviceType `json:"deviceType,omitempty"`

	// Key of the entity.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Managed DeviceId
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// Managed Device Name
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// Management agent type.
	ManagementAgent *ManagementAgentType `json:"managementAgent,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Owner type of device.
	OwnerType *ManagedDeviceOwnerType `json:"ownerType,omitempty"`

	// Managed Device Retire After DateTime
	RetireAfterDateTime *string `json:"retireAfterDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`
}
