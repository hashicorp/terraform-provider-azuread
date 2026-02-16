package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = WindowsIdentityProtectionConfiguration{}

type WindowsIdentityProtectionConfiguration struct {
	// Boolean value used to enable enhanced anti-spoofing for facial feature recognition on Windows Hello face
	// authentication.
	EnhancedAntiSpoofingForFacialFeaturesEnabled *bool `json:"enhancedAntiSpoofingForFacialFeaturesEnabled,omitempty"`

	// Integer value specifies the period (in days) that a PIN can be used before the system requires the user to change it.
	// Valid values are 0 to 730 inclusive. Valid values 0 to 730
	PinExpirationInDays nullable.Type[int64] `json:"pinExpirationInDays,omitempty"`

	// Possible values of the ConfigurationUsage list.
	PinLowercaseCharactersUsage *ConfigurationUsage `json:"pinLowercaseCharactersUsage,omitempty"`

	// Integer value that sets the maximum number of characters allowed for the work PIN. Valid values are 4 to 127
	// inclusive and greater than or equal to the value set for the minimum PIN. Valid values 4 to 127
	PinMaximumLength nullable.Type[int64] `json:"pinMaximumLength,omitempty"`

	// Integer value that sets the minimum number of characters required for the Windows Hello for Business PIN. Valid
	// values are 4 to 127 inclusive and less than or equal to the value set for the maximum PIN. Valid values 4 to 127
	PinMinimumLength nullable.Type[int64] `json:"pinMinimumLength,omitempty"`

	// Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the
	// current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not
	// preserved through a PIN reset. Valid values 0 to 50
	PinPreviousBlockCount nullable.Type[int64] `json:"pinPreviousBlockCount,omitempty"`

	// Boolean value that enables a user to change their PIN by using the Windows Hello for Business PIN recovery service.
	PinRecoveryEnabled *bool `json:"pinRecoveryEnabled,omitempty"`

	// Possible values of the ConfigurationUsage list.
	PinSpecialCharactersUsage *ConfigurationUsage `json:"pinSpecialCharactersUsage,omitempty"`

	// Possible values of the ConfigurationUsage list.
	PinUppercaseCharactersUsage *ConfigurationUsage `json:"pinUppercaseCharactersUsage,omitempty"`

	// Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM
	// provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False,
	// all devices can provision Windows Hello for Business even if there is not a usable TPM.
	SecurityDeviceRequired *bool `json:"securityDeviceRequired,omitempty"`

	// Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for
	// Business PIN. If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in
	// case of failures.
	UnlockWithBiometricsEnabled *bool `json:"unlockWithBiometricsEnabled,omitempty"`

	// Boolean value that enables Windows Hello for Business to use certificates to authenticate on-premise resources.
	UseCertificatesForOnPremisesAuthEnabled *bool `json:"useCertificatesForOnPremisesAuthEnabled,omitempty"`

	// Boolean value used to enable the Windows Hello security key as a logon credential.
	UseSecurityKeyForSignin *bool `json:"useSecurityKeyForSignin,omitempty"`

	// Boolean value that blocks Windows Hello for Business as a method for signing into Windows.
	WindowsHelloForBusinessBlocked nullable.Type[bool] `json:"windowsHelloForBusinessBlocked,omitempty"`

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

func (s WindowsIdentityProtectionConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s WindowsIdentityProtectionConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsIdentityProtectionConfiguration{}

func (s WindowsIdentityProtectionConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsIdentityProtectionConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsIdentityProtectionConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsIdentityProtectionConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsIdentityProtectionConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsIdentityProtectionConfiguration: %+v", err)
	}

	return encoded, nil
}
