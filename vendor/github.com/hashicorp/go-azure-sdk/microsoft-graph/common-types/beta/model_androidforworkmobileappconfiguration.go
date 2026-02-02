package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedDeviceMobileAppConfiguration = AndroidForWorkMobileAppConfiguration{}

type AndroidForWorkMobileAppConfiguration struct {
	// Setting to specify whether to allow ConnectedApps experience for this app.
	ConnectedAppsEnabled *bool `json:"connectedAppsEnabled,omitempty"`

	// Android For Work app configuration package id.
	PackageId nullable.Type[string] `json:"packageId,omitempty"`

	// Android For Work app configuration JSON payload.
	PayloadJson nullable.Type[string] `json:"payloadJson,omitempty"`

	// List of Android app permissions and corresponding permission actions.
	PermissionActions *[]AndroidPermissionAction `json:"permissionActions,omitempty"`

	// Android profile applicability
	ProfileApplicability *AndroidProfileApplicability `json:"profileApplicability,omitempty"`

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

	// List of Scope Tags for this App configuration entity.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s AndroidForWorkMobileAppConfiguration) ManagedDeviceMobileAppConfiguration() BaseManagedDeviceMobileAppConfigurationImpl {
	return BaseManagedDeviceMobileAppConfigurationImpl{
		Assignments:          s.Assignments,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DeviceStatusSummary:  s.DeviceStatusSummary,
		DeviceStatuses:       s.DeviceStatuses,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RoleScopeTagIds:      s.RoleScopeTagIds,
		TargetedMobileApps:   s.TargetedMobileApps,
		UserStatusSummary:    s.UserStatusSummary,
		UserStatuses:         s.UserStatuses,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s AndroidForWorkMobileAppConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidForWorkMobileAppConfiguration{}

func (s AndroidForWorkMobileAppConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AndroidForWorkMobileAppConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidForWorkMobileAppConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidForWorkMobileAppConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidForWorkMobileAppConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidForWorkMobileAppConfiguration: %+v", err)
	}

	return encoded, nil
}
