package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfiguration interface {
	Entity
	ManagedDeviceMobileAppConfiguration() BaseManagedDeviceMobileAppConfigurationImpl
}

var _ ManagedDeviceMobileAppConfiguration = BaseManagedDeviceMobileAppConfigurationImpl{}

type BaseManagedDeviceMobileAppConfigurationImpl struct {
	// The list of group assignemenets for app configration.
	Assignments *[]ManagedDeviceMobileAppConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// App configuration device status summary.
	DeviceStatusSummary *ManagedDeviceMobileAppConfigurationDeviceSummary `json:"deviceStatusSummary,omitempty"`

	// List of ManagedDeviceMobileAppConfigurationDeviceStatus.
	DeviceStatuses *[]ManagedDeviceMobileAppConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// the associated app.
	TargetedMobileApps *[]string `json:"targetedMobileApps,omitempty"`

	// App configuration user status summary.
	UserStatusSummary *ManagedDeviceMobileAppConfigurationUserSummary `json:"userStatusSummary,omitempty"`

	// List of ManagedDeviceMobileAppConfigurationUserStatus.
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

func (s BaseManagedDeviceMobileAppConfigurationImpl) ManagedDeviceMobileAppConfiguration() BaseManagedDeviceMobileAppConfigurationImpl {
	return s
}

func (s BaseManagedDeviceMobileAppConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedDeviceMobileAppConfiguration = RawManagedDeviceMobileAppConfigurationImpl{}

// RawManagedDeviceMobileAppConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedDeviceMobileAppConfigurationImpl struct {
	managedDeviceMobileAppConfiguration BaseManagedDeviceMobileAppConfigurationImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawManagedDeviceMobileAppConfigurationImpl) ManagedDeviceMobileAppConfiguration() BaseManagedDeviceMobileAppConfigurationImpl {
	return s.managedDeviceMobileAppConfiguration
}

func (s RawManagedDeviceMobileAppConfigurationImpl) Entity() BaseEntityImpl {
	return s.managedDeviceMobileAppConfiguration.Entity()
}

var _ json.Marshaler = BaseManagedDeviceMobileAppConfigurationImpl{}

func (s BaseManagedDeviceMobileAppConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedDeviceMobileAppConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedDeviceMobileAppConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedDeviceMobileAppConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDeviceMobileAppConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedDeviceMobileAppConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalManagedDeviceMobileAppConfigurationImplementation(input []byte) (ManagedDeviceMobileAppConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceMobileAppConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosMobileAppConfiguration") {
		var out IosMobileAppConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosMobileAppConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedDeviceMobileAppConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedDeviceMobileAppConfigurationImpl: %+v", err)
	}

	return RawManagedDeviceMobileAppConfigurationImpl{
		managedDeviceMobileAppConfiguration: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
