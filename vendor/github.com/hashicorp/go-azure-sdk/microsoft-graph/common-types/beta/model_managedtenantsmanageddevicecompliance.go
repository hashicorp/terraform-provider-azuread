package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedDeviceCompliance{}

type ManagedTenantsManagedDeviceCompliance struct {
	// Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant,
	// conflict, error, inGracePeriod, configManager. Optional. Read-only.
	ComplianceStatus nullable.Type[string] `json:"complianceStatus,omitempty"`

	// Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia,
	// windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub,
	// androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
	// Optional. Read-only.
	DeviceType nullable.Type[string] `json:"deviceType,omitempty"`

	// The date and time when the grace period will expire. Optional. Read-only.
	InGracePeriodUntilDateTime nullable.Type[string] `json:"inGracePeriodUntilDateTime,omitempty"`

	// Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
	LastRefreshedDateTime nullable.Type[string] `json:"lastRefreshedDateTime,omitempty"`

	// The date and time that the device last completed a successful sync with Microsoft Endpoint Manager. Optional.
	// Read-only.
	LastSyncDateTime nullable.Type[string] `json:"lastSyncDateTime,omitempty"`

	// The identifier for the managed device in Microsoft Endpoint Manager. Optional. Read-only.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The display name for the managed device. Optional. Read-only.
	ManagedDeviceName nullable.Type[string] `json:"managedDeviceName,omitempty"`

	// The manufacture for the device. Optional. Read-only.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The model for the device. Optional. Read-only.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The description of the operating system for the managed device. Optional. Read-only.
	OsDescription nullable.Type[string] `json:"osDescription,omitempty"`

	// The version of the operating system for the managed device. Optional. Read-only.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// The type of owner for the managed device. Optional. Read-only.
	OwnerType nullable.Type[string] `json:"ownerType,omitempty"`

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

func (s ManagedTenantsManagedDeviceCompliance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedDeviceCompliance{}

func (s ManagedTenantsManagedDeviceCompliance) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedDeviceCompliance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedDeviceCompliance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedDeviceCompliance: %+v", err)
	}

	delete(decoded, "complianceStatus")
	delete(decoded, "deviceType")
	delete(decoded, "inGracePeriodUntilDateTime")
	delete(decoded, "lastRefreshedDateTime")
	delete(decoded, "lastSyncDateTime")
	delete(decoded, "managedDeviceId")
	delete(decoded, "managedDeviceName")
	delete(decoded, "manufacturer")
	delete(decoded, "model")
	delete(decoded, "osDescription")
	delete(decoded, "osVersion")
	delete(decoded, "ownerType")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "tenantId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedDeviceCompliance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedDeviceCompliance: %+v", err)
	}

	return encoded, nil
}
