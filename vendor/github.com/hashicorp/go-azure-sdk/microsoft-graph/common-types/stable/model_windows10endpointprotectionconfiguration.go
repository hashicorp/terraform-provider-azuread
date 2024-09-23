package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = Windows10EndpointProtectionConfiguration{}

type Windows10EndpointProtectionConfiguration struct {
	// Possible values of AppLocker Application Control Types
	AppLockerApplicationControl *AppLockerApplicationControlType `json:"appLockerApplicationControl,omitempty"`

	// Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
	ApplicationGuardAllowPersistence *bool `json:"applicationGuardAllowPersistence,omitempty"`

	// Allow printing to Local Printers from Container
	ApplicationGuardAllowPrintToLocalPrinters *bool `json:"applicationGuardAllowPrintToLocalPrinters,omitempty"`

	// Allow printing to Network Printers from Container
	ApplicationGuardAllowPrintToNetworkPrinters *bool `json:"applicationGuardAllowPrintToNetworkPrinters,omitempty"`

	// Allow printing to PDF from Container
	ApplicationGuardAllowPrintToPDF *bool `json:"applicationGuardAllowPrintToPDF,omitempty"`

	// Allow printing to XPS from Container
	ApplicationGuardAllowPrintToXPS *bool `json:"applicationGuardAllowPrintToXPS,omitempty"`

	// Possible values for applicationGuardBlockClipboardSharingType
	ApplicationGuardBlockClipboardSharing *ApplicationGuardBlockClipboardSharingType `json:"applicationGuardBlockClipboardSharing,omitempty"`

	// Possible values for applicationGuardBlockFileTransfer
	ApplicationGuardBlockFileTransfer *ApplicationGuardBlockFileTransferType `json:"applicationGuardBlockFileTransfer,omitempty"`

	// Block enterprise sites to load non-enterprise content, such as third party plug-ins
	ApplicationGuardBlockNonEnterpriseContent *bool `json:"applicationGuardBlockNonEnterpriseContent,omitempty"`

	// Enable Windows Defender Application Guard
	ApplicationGuardEnabled *bool `json:"applicationGuardEnabled,omitempty"`

	// Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user
	// login-logoff, use of privilege rights, software installation, system changes, etc.)
	ApplicationGuardForceAuditing *bool `json:"applicationGuardForceAuditing,omitempty"`

	// Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
	BitLockerDisableWarningForOtherDiskEncryption *bool `json:"bitLockerDisableWarningForOtherDiskEncryption,omitempty"`

	// Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
	BitLockerEnableStorageCardEncryptionOnMobile *bool `json:"bitLockerEnableStorageCardEncryptionOnMobile,omitempty"`

	// Allows the admin to require encryption to be turned on using BitLocker.
	BitLockerEncryptDevice *bool `json:"bitLockerEncryptDevice,omitempty"`

	// BitLocker Removable Drive Policy.
	BitLockerRemovableDrivePolicy *BitLockerRemovableDrivePolicy `json:"bitLockerRemovableDrivePolicy,omitempty"`

	// List of folder paths to be added to the list of protected folders
	DefenderAdditionalGuardedFolders *[]string `json:"defenderAdditionalGuardedFolders,omitempty"`

	// List of exe files and folders to be excluded from attack surface reduction rules
	DefenderAttackSurfaceReductionExcludedPaths *[]string `json:"defenderAttackSurfaceReductionExcludedPaths,omitempty"`

	// Xml content containing information regarding exploit protection details.
	DefenderExploitProtectionXml nullable.Type[string] `json:"defenderExploitProtectionXml,omitempty"`

	// Name of the file from which DefenderExploitProtectionXml was obtained.
	DefenderExploitProtectionXmlFileName nullable.Type[string] `json:"defenderExploitProtectionXmlFileName,omitempty"`

	// List of paths to exe that are allowed to access protected folders
	DefenderGuardedFoldersAllowedAppPaths *[]string `json:"defenderGuardedFoldersAllowedAppPaths,omitempty"`

	// Indicates whether or not to block user from overriding Exploit Protection settings.
	DefenderSecurityCenterBlockExploitProtectionOverride *bool `json:"defenderSecurityCenterBlockExploitProtectionOverride,omitempty"`

	// Blocks stateful FTP connections to the device
	FirewallBlockStatefulFTP nullable.Type[bool] `json:"firewallBlockStatefulFTP,omitempty"`

	// Possible values for firewallCertificateRevocationListCheckMethod
	FirewallCertificateRevocationListCheckMethod *FirewallCertificateRevocationListCheckMethodType `json:"firewallCertificateRevocationListCheckMethod,omitempty"`

	// Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
	FirewallIPSecExemptionsAllowDHCP *bool `json:"firewallIPSecExemptionsAllowDHCP,omitempty"`

	// Configures IPSec exemptions to allow ICMP
	FirewallIPSecExemptionsAllowICMP *bool `json:"firewallIPSecExemptionsAllowICMP,omitempty"`

	// Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
	FirewallIPSecExemptionsAllowNeighborDiscovery *bool `json:"firewallIPSecExemptionsAllowNeighborDiscovery,omitempty"`

	// Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
	FirewallIPSecExemptionsAllowRouterDiscovery *bool `json:"firewallIPSecExemptionsAllowRouterDiscovery,omitempty"`

	// Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period
	// after which security associations will expire and be deleted. Valid values 300 to 3600
	FirewallIdleTimeoutForSecurityAssociationInSeconds nullable.Type[int64] `json:"firewallIdleTimeoutForSecurityAssociationInSeconds,omitempty"`

	// If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported
	// authentication suites rather than the entire set
	FirewallMergeKeyingModuleSettings nullable.Type[bool] `json:"firewallMergeKeyingModuleSettings,omitempty"`

	// Possible values for firewallPacketQueueingMethod
	FirewallPacketQueueingMethod *FirewallPacketQueueingMethodType `json:"firewallPacketQueueingMethod,omitempty"`

	// Possible values for firewallPreSharedKeyEncodingMethod
	FirewallPreSharedKeyEncodingMethod *FirewallPreSharedKeyEncodingMethodType `json:"firewallPreSharedKeyEncodingMethod,omitempty"`

	// Configures the firewall profile settings for domain networks
	FirewallProfileDomain *WindowsFirewallNetworkProfile `json:"firewallProfileDomain,omitempty"`

	// Configures the firewall profile settings for private networks
	FirewallProfilePrivate *WindowsFirewallNetworkProfile `json:"firewallProfilePrivate,omitempty"`

	// Configures the firewall profile settings for public networks
	FirewallProfilePublic *WindowsFirewallNetworkProfile `json:"firewallProfilePublic,omitempty"`

	// Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
	SmartScreenBlockOverrideForFiles *bool `json:"smartScreenBlockOverrideForFiles,omitempty"`

	// Allows IT Admins to configure SmartScreen for Windows.
	SmartScreenEnableInShell *bool `json:"smartScreenEnableInShell,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s Windows10EndpointProtectionConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:                 s.Assignments,
		CreatedDateTime:             s.CreatedDateTime,
		Description:                 s.Description,
		DeviceSettingStateSummaries: s.DeviceSettingStateSummaries,
		DeviceStatusOverview:        s.DeviceStatusOverview,
		DeviceStatuses:              s.DeviceStatuses,
		DisplayName:                 s.DisplayName,
		LastModifiedDateTime:        s.LastModifiedDateTime,
		UserStatusOverview:          s.UserStatusOverview,
		UserStatuses:                s.UserStatuses,
		Version:                     s.Version,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s Windows10EndpointProtectionConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10EndpointProtectionConfiguration{}

func (s Windows10EndpointProtectionConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10EndpointProtectionConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10EndpointProtectionConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10EndpointProtectionConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10EndpointProtectionConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10EndpointProtectionConfiguration: %+v", err)
	}

	return encoded, nil
}
