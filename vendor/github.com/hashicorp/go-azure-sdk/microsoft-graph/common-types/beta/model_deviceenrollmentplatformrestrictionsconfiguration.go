package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceEnrollmentConfiguration = DeviceEnrollmentPlatformRestrictionsConfiguration{}

type DeviceEnrollmentPlatformRestrictionsConfiguration struct {
	// Android for work restrictions based on platform, platform operating system version, and device ownership
	AndroidForWorkRestriction *DeviceEnrollmentPlatformRestriction `json:"androidForWorkRestriction,omitempty"`

	// Android restrictions based on platform, platform operating system version, and device ownership
	AndroidRestriction *DeviceEnrollmentPlatformRestriction `json:"androidRestriction,omitempty"`

	// Ios restrictions based on platform, platform operating system version, and device ownership
	IosRestriction *DeviceEnrollmentPlatformRestriction `json:"iosRestriction,omitempty"`

	// Mac restrictions based on platform, platform operating system version, and device ownership
	MacOSRestriction *DeviceEnrollmentPlatformRestriction `json:"macOSRestriction,omitempty"`

	// Mac restrictions based on platform, platform operating system version, and device ownership
	MacRestriction *DeviceEnrollmentPlatformRestriction `json:"macRestriction,omitempty"`

	// Windows Home Sku restrictions based on platform, platform operating system version, and device ownership
	WindowsHomeSkuRestriction *DeviceEnrollmentPlatformRestriction `json:"windowsHomeSkuRestriction,omitempty"`

	// Windows mobile restrictions based on platform, platform operating system version, and device ownership
	WindowsMobileRestriction *DeviceEnrollmentPlatformRestriction `json:"windowsMobileRestriction,omitempty"`

	// Windows restrictions based on platform, platform operating system version, and device ownership
	WindowsRestriction *DeviceEnrollmentPlatformRestriction `json:"windowsRestriction,omitempty"`

	// Fields inherited from DeviceEnrollmentConfiguration

	// The list of group assignments for the device configuration profile
	Assignments *[]EnrollmentConfigurationAssignment `json:"assignments,omitempty"`

	// Created date time in UTC of the device enrollment configuration
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the device enrollment configuration
	Description nullable.Type[string] `json:"description,omitempty"`

	// Describes the TemplateFamily for the Template entity
	DeviceEnrollmentConfigurationType *DeviceEnrollmentConfigurationType `json:"deviceEnrollmentConfigurationType,omitempty"`

	// The display name of the device enrollment configuration
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Last modified date time in UTC of the device enrollment configuration
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject
	// only to the configuration with the lowest priority value.
	Priority *int64 `json:"priority,omitempty"`

	// Optional role scope tags for the enrollment restrictions.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The version of the device enrollment configuration
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

func (s DeviceEnrollmentPlatformRestrictionsConfiguration) DeviceEnrollmentConfiguration() BaseDeviceEnrollmentConfigurationImpl {
	return BaseDeviceEnrollmentConfigurationImpl{
		Assignments:                       s.Assignments,
		CreatedDateTime:                   s.CreatedDateTime,
		Description:                       s.Description,
		DeviceEnrollmentConfigurationType: s.DeviceEnrollmentConfigurationType,
		DisplayName:                       s.DisplayName,
		LastModifiedDateTime:              s.LastModifiedDateTime,
		Priority:                          s.Priority,
		RoleScopeTagIds:                   s.RoleScopeTagIds,
		Version:                           s.Version,
		Id:                                s.Id,
		ODataId:                           s.ODataId,
		ODataType:                         s.ODataType,
	}
}

func (s DeviceEnrollmentPlatformRestrictionsConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceEnrollmentPlatformRestrictionsConfiguration{}

func (s DeviceEnrollmentPlatformRestrictionsConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper DeviceEnrollmentPlatformRestrictionsConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceEnrollmentPlatformRestrictionsConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceEnrollmentPlatformRestrictionsConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceEnrollmentPlatformRestrictionsConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceEnrollmentPlatformRestrictionsConfiguration: %+v", err)
	}

	return encoded, nil
}
