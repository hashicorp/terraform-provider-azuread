package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedDeviceMobileAppConfiguration = IosMobileAppConfiguration{}

type IosMobileAppConfiguration struct {
	// mdm app configuration Base64 binary.
	EncodedSettingXml nullable.Type[string] `json:"encodedSettingXml,omitempty"`

	// app configuration setting items.
	Settings *[]AppConfigurationSettingItem `json:"settings,omitempty"`

	// Fields inherited from ManagedDeviceMobileAppConfiguration

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

func (s IosMobileAppConfiguration) ManagedDeviceMobileAppConfiguration() BaseManagedDeviceMobileAppConfigurationImpl {
	return BaseManagedDeviceMobileAppConfigurationImpl{
		Assignments:          s.Assignments,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DeviceStatusSummary:  s.DeviceStatusSummary,
		DeviceStatuses:       s.DeviceStatuses,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		TargetedMobileApps:   s.TargetedMobileApps,
		UserStatusSummary:    s.UserStatusSummary,
		UserStatuses:         s.UserStatuses,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s IosMobileAppConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosMobileAppConfiguration{}

func (s IosMobileAppConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper IosMobileAppConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosMobileAppConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosMobileAppConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosMobileAppConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosMobileAppConfiguration: %+v", err)
	}

	return encoded, nil
}
