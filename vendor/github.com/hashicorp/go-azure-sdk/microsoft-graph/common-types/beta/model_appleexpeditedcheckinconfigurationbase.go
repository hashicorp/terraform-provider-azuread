package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleExpeditedCheckinConfigurationBase interface {
	Entity
	DeviceConfiguration
	AppleExpeditedCheckinConfigurationBase() BaseAppleExpeditedCheckinConfigurationBaseImpl
}

var _ AppleExpeditedCheckinConfigurationBase = BaseAppleExpeditedCheckinConfigurationBaseImpl{}

type BaseAppleExpeditedCheckinConfigurationBaseImpl struct {
	// Gets or sets whether to enable expedited device check-ins.
	EnableExpeditedCheckin *bool `json:"enableExpeditedCheckin,omitempty"`

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

func (s BaseAppleExpeditedCheckinConfigurationBaseImpl) AppleExpeditedCheckinConfigurationBase() BaseAppleExpeditedCheckinConfigurationBaseImpl {
	return s
}

func (s BaseAppleExpeditedCheckinConfigurationBaseImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s BaseAppleExpeditedCheckinConfigurationBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AppleExpeditedCheckinConfigurationBase = RawAppleExpeditedCheckinConfigurationBaseImpl{}

// RawAppleExpeditedCheckinConfigurationBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAppleExpeditedCheckinConfigurationBaseImpl struct {
	appleExpeditedCheckinConfigurationBase BaseAppleExpeditedCheckinConfigurationBaseImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawAppleExpeditedCheckinConfigurationBaseImpl) AppleExpeditedCheckinConfigurationBase() BaseAppleExpeditedCheckinConfigurationBaseImpl {
	return s.appleExpeditedCheckinConfigurationBase
}

func (s RawAppleExpeditedCheckinConfigurationBaseImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.appleExpeditedCheckinConfigurationBase.DeviceConfiguration()
}

func (s RawAppleExpeditedCheckinConfigurationBaseImpl) Entity() BaseEntityImpl {
	return s.appleExpeditedCheckinConfigurationBase.Entity()
}

var _ json.Marshaler = BaseAppleExpeditedCheckinConfigurationBaseImpl{}

func (s BaseAppleExpeditedCheckinConfigurationBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAppleExpeditedCheckinConfigurationBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAppleExpeditedCheckinConfigurationBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAppleExpeditedCheckinConfigurationBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appleExpeditedCheckinConfigurationBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAppleExpeditedCheckinConfigurationBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAppleExpeditedCheckinConfigurationBaseImplementation(input []byte) (AppleExpeditedCheckinConfigurationBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AppleExpeditedCheckinConfigurationBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosExpeditedCheckinConfiguration") {
		var out IosExpeditedCheckinConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosExpeditedCheckinConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAppleExpeditedCheckinConfigurationBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAppleExpeditedCheckinConfigurationBaseImpl: %+v", err)
	}

	return RawAppleExpeditedCheckinConfigurationBaseImpl{
		appleExpeditedCheckinConfigurationBase: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
