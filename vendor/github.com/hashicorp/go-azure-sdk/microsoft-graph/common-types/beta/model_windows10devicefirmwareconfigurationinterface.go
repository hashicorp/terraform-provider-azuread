package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = Windows10DeviceFirmwareConfigurationInterface{}

type Windows10DeviceFirmwareConfigurationInterface struct {
	// Possible values of a property
	Bluetooth *Enablement `json:"bluetooth,omitempty"`

	// Possible values of a property
	BootFromBuiltInNetworkAdapters *Enablement `json:"bootFromBuiltInNetworkAdapters,omitempty"`

	// Possible values of a property
	BootFromExternalMedia *Enablement `json:"bootFromExternalMedia,omitempty"`

	// Possible values of a property
	Cameras *Enablement `json:"cameras,omitempty"`

	// Defines the permission level granted to users to enable them change Uefi settings
	ChangeUefiSettingsPermission *ChangeUefiSettingsPermission `json:"changeUefiSettingsPermission,omitempty"`

	// Possible values of a property
	FrontCamera *Enablement `json:"frontCamera,omitempty"`

	// Possible values of a property
	InfraredCamera *Enablement `json:"infraredCamera,omitempty"`

	// Possible values of a property
	Microphone *Enablement `json:"microphone,omitempty"`

	// Possible values of a property
	MicrophonesAndSpeakers *Enablement `json:"microphonesAndSpeakers,omitempty"`

	// Possible values of a property
	NearFieldCommunication *Enablement `json:"nearFieldCommunication,omitempty"`

	// Possible values of a property
	Radios *Enablement `json:"radios,omitempty"`

	// Possible values of a property
	RearCamera *Enablement `json:"rearCamera,omitempty"`

	// Possible values of a property
	SdCard *Enablement `json:"sdCard,omitempty"`

	// Possible values of a property
	SimultaneousMultiThreading *Enablement `json:"simultaneousMultiThreading,omitempty"`

	// Possible values of a property
	UsbTypeAPort *Enablement `json:"usbTypeAPort,omitempty"`

	// Possible values of a property
	VirtualizationOfCpuAndIO *Enablement `json:"virtualizationOfCpuAndIO,omitempty"`

	// Possible values of a property
	WakeOnLAN *Enablement `json:"wakeOnLAN,omitempty"`

	// Possible values of a property
	WakeOnPower *Enablement `json:"wakeOnPower,omitempty"`

	// Possible values of a property
	WiFi *Enablement `json:"wiFi,omitempty"`

	// Possible values of a property
	WindowsPlatformBinaryTable *Enablement `json:"windowsPlatformBinaryTable,omitempty"`

	// Possible values of a property
	WirelessWideAreaNetwork *Enablement `json:"wirelessWideAreaNetwork,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

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

func (s Windows10DeviceFirmwareConfigurationInterface) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s Windows10DeviceFirmwareConfigurationInterface) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10DeviceFirmwareConfigurationInterface{}

func (s Windows10DeviceFirmwareConfigurationInterface) MarshalJSON() ([]byte, error) {
	type wrapper Windows10DeviceFirmwareConfigurationInterface
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10DeviceFirmwareConfigurationInterface: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10DeviceFirmwareConfigurationInterface: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10DeviceFirmwareConfigurationInterface"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10DeviceFirmwareConfigurationInterface: %+v", err)
	}

	return encoded, nil
}
