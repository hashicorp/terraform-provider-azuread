package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IosLobAppProvisioningConfiguration{}

type IosLobAppProvisioningConfiguration struct {
	// The associated group assignments for IosLobAppProvisioningConfiguration, this determines which devices/users the IOS
	// LOB app provisioning conifguration will be targeted to.
	Assignments *[]IosLobAppProvisioningConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The list of device installation states for this mobile app configuration.
	DeviceStatuses *[]ManagedDeviceMobileAppConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional profile expiration date and time. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Returned by default.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The associated group assignments.
	GroupAssignments *[]MobileAppProvisioningConfigGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Payload. (UTF8 encoded byte array)
	Payload *string `json:"payload,omitempty"`

	// Payload file name (.mobileprovision
	PayloadFileName *string `json:"payloadFileName,omitempty"`

	// List of Scope Tags for this iOS LOB app provisioning configuration entity.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The list of user installation states for this mobile app configuration.
	UserStatuses *[]ManagedDeviceMobileAppConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

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

func (s IosLobAppProvisioningConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosLobAppProvisioningConfiguration{}

func (s IosLobAppProvisioningConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper IosLobAppProvisioningConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosLobAppProvisioningConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosLobAppProvisioningConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosLobAppProvisioningConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosLobAppProvisioningConfiguration: %+v", err)
	}

	return encoded, nil
}
