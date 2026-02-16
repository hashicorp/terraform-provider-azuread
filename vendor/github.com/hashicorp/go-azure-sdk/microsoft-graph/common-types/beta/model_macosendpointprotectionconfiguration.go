package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = MacOSEndpointProtectionConfiguration{}

type MacOSEndpointProtectionConfiguration struct {
	// Possible values of a property
	AdvancedThreatProtectionAutomaticSampleSubmission *Enablement `json:"advancedThreatProtectionAutomaticSampleSubmission,omitempty"`

	// Possible values of a property
	AdvancedThreatProtectionCloudDelivered *Enablement `json:"advancedThreatProtectionCloudDelivered,omitempty"`

	// Possible values of a property
	AdvancedThreatProtectionDiagnosticDataCollection *Enablement `json:"advancedThreatProtectionDiagnosticDataCollection,omitempty"`

	// A list of file extensions to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on
	// macOS.
	AdvancedThreatProtectionExcludedExtensions *[]string `json:"advancedThreatProtectionExcludedExtensions,omitempty"`

	// A list of paths to files to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on
	// macOS.
	AdvancedThreatProtectionExcludedFiles *[]string `json:"advancedThreatProtectionExcludedFiles,omitempty"`

	// A list of paths to folders to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on
	// macOS.
	AdvancedThreatProtectionExcludedFolders *[]string `json:"advancedThreatProtectionExcludedFolders,omitempty"`

	// A list of process names to exclude from antivirus scanning for Microsoft Defender Advanced Threat Protection on
	// macOS.
	AdvancedThreatProtectionExcludedProcesses *[]string `json:"advancedThreatProtectionExcludedProcesses,omitempty"`

	// Possible values of a property
	AdvancedThreatProtectionRealTime *Enablement `json:"advancedThreatProtectionRealTime,omitempty"`

	// Optional. If set to true, the user can defer the enabling of FileVault until they sign out.
	FileVaultAllowDeferralUntilSignOut *bool `json:"fileVaultAllowDeferralUntilSignOut,omitempty"`

	// Optional. When using the Defer option, if set to true, the user is not prompted to enable FileVault at sign-out.
	FileVaultDisablePromptAtSignOut *bool `json:"fileVaultDisablePromptAtSignOut,omitempty"`

	// Whether FileVault should be enabled or not.
	FileVaultEnabled *bool `json:"fileVaultEnabled,omitempty"`

	// Optional. A hidden personal recovery key does not appear on the user's screen during FileVault encryption, reducing
	// the risk of it ending up in the wrong hands.
	FileVaultHidePersonalRecoveryKey *bool `json:"fileVaultHidePersonalRecoveryKey,omitempty"`

	// Required if selected recovery key type(s) include InstitutionalRecoveryKey. The DER Encoded certificate file used to
	// set an institutional recovery key.
	FileVaultInstitutionalRecoveryKeyCertificate nullable.Type[string] `json:"fileVaultInstitutionalRecoveryKeyCertificate,omitempty"`

	// File name of the institutional recovery key certificate to display in UI. (.der).
	FileVaultInstitutionalRecoveryKeyCertificateFileName nullable.Type[string] `json:"fileVaultInstitutionalRecoveryKeyCertificateFileName,omitempty"`

	// Optional. When using the Defer option, this is the maximum number of times the user can ignore prompts to enable
	// FileVault before FileVault will be required for the user to sign in. If set to -1, it will always prompt to enable
	// FileVault until FileVault is enabled, though it will allow the user to bypass enabling FileVault. Setting this to 0
	// will disable the feature.
	FileVaultNumberOfTimesUserCanIgnore nullable.Type[int64] `json:"fileVaultNumberOfTimesUserCanIgnore,omitempty"`

	// Required if selected recovery key type(s) include PersonalRecoveryKey. A short message displayed to the user that
	// explains how they can retrieve their personal recovery key.
	FileVaultPersonalRecoveryKeyHelpMessage nullable.Type[string] `json:"fileVaultPersonalRecoveryKeyHelpMessage,omitempty"`

	// Optional. If selected recovery key type(s) include PersonalRecoveryKey, the frequency to rotate that key, in months.
	FileVaultPersonalRecoveryKeyRotationInMonths nullable.Type[int64] `json:"fileVaultPersonalRecoveryKeyRotationInMonths,omitempty"`

	// Recovery key types for macOS FileVault
	FileVaultSelectedRecoveryKeyTypes *MacOSFileVaultRecoveryKeyTypes `json:"fileVaultSelectedRecoveryKeyTypes,omitempty"`

	// List of applications with firewall settings. Firewall settings for applications not on this list are determined by
	// the user. This collection can contain a maximum of 500 elements.
	FirewallApplications *[]MacOSFirewallApplication `json:"firewallApplications,omitempty"`

	// Corresponds to the 'Block all incoming connections' option.
	FirewallBlockAllIncoming *bool `json:"firewallBlockAllIncoming,omitempty"`

	// Corresponds to 'Enable stealth mode.'
	FirewallEnableStealthMode *bool `json:"firewallEnableStealthMode,omitempty"`

	// Whether the firewall should be enabled or not.
	FirewallEnabled *bool `json:"firewallEnabled,omitempty"`

	// App source options for macOS Gatekeeper.
	GatekeeperAllowedAppSource *MacOSGatekeeperAppSources `json:"gatekeeperAllowedAppSource,omitempty"`

	// If set to true, the user override for Gatekeeper will be disabled.
	GatekeeperBlockOverride *bool `json:"gatekeeperBlockOverride,omitempty"`

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

func (s MacOSEndpointProtectionConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s MacOSEndpointProtectionConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSEndpointProtectionConfiguration{}

func (s MacOSEndpointProtectionConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MacOSEndpointProtectionConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSEndpointProtectionConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSEndpointProtectionConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSEndpointProtectionConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSEndpointProtectionConfiguration: %+v", err)
	}

	return encoded, nil
}
