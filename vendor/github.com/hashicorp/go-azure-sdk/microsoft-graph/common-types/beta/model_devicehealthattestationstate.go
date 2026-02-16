package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationState struct {
	// TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key
	// (EK) certificate.
	AttestationIdentityKey nullable.Type[string] `json:"attestationIdentityKey,omitempty"`

	// On or Off of BitLocker Drive Encryption
	BitLockerStatus nullable.Type[string] `json:"bitLockerStatus,omitempty"`

	// The security version number of the Boot Application
	BootAppSecurityVersion nullable.Type[string] `json:"bootAppSecurityVersion,omitempty"`

	// When bootDebugging is enabled, the device is used in development and testing
	BootDebugging nullable.Type[string] `json:"bootDebugging,omitempty"`

	// The security version number of the Boot Application
	BootManagerSecurityVersion nullable.Type[string] `json:"bootManagerSecurityVersion,omitempty"`

	// The version of the Boot Manager
	BootManagerVersion nullable.Type[string] `json:"bootManagerVersion,omitempty"`

	// The Boot Revision List that was loaded during initial boot on the attested device
	BootRevisionListInfo nullable.Type[string] `json:"bootRevisionListInfo,omitempty"`

	// When code integrity is enabled, code execution is restricted to integrity verified code
	CodeIntegrity nullable.Type[string] `json:"codeIntegrity,omitempty"`

	// The version of the Boot Manager
	CodeIntegrityCheckVersion nullable.Type[string] `json:"codeIntegrityCheckVersion,omitempty"`

	// The Code Integrity policy that is controlling the security of the boot environment
	CodeIntegrityPolicy nullable.Type[string] `json:"codeIntegrityPolicy,omitempty"`

	// The DHA report version. (Namespace version)
	ContentNamespaceUrl nullable.Type[string] `json:"contentNamespaceUrl,omitempty"`

	// The HealthAttestation state schema version
	ContentVersion nullable.Type[string] `json:"contentVersion,omitempty"`

	// DEP Policy defines a set of hardware and software technologies that perform additional checks on memory
	DataExcutionPolicy nullable.Type[string] `json:"dataExcutionPolicy,omitempty"`

	// The DHA report version. (Namespace version)
	DeviceHealthAttestationStatus nullable.Type[string] `json:"deviceHealthAttestationStatus,omitempty"`

	// ELAM provides protection for the computers in your network when they start up
	EarlyLaunchAntiMalwareDriverProtection nullable.Type[string] `json:"earlyLaunchAntiMalwareDriverProtection,omitempty"`

	// A list of possible Firmware protection type for a device. Firmware protection is a set of features that helps to
	// ensure attackers can't get your device to start with untrusted or malicious firmware. Firmware protection type is
	// determined by report sent from Microsoft Azure Attestation service. Only Windows 11 devices will have values
	// "systemGuardSecureLaunch" or "firmwareAttackSurfaceReduction" or "disabled". Windows 10 devices will have value
	// "notApplicable".
	FirmwareProtection *FirmwareProtectionType `json:"firmwareProtection,omitempty"`

	// This attribute indicates if DHA is supported for the device
	HealthAttestationSupportedStatus nullable.Type[string] `json:"healthAttestationSupportedStatus,omitempty"`

	// This attribute appears if DHA-Service detects an integrity issue
	HealthStatusMismatchInfo nullable.Type[string] `json:"healthStatusMismatchInfo,omitempty"`

	// The DateTime when device was evaluated or issued to MDM
	IssuedDateTime *string `json:"issuedDateTime,omitempty"`

	// The Timestamp of the last update.
	LastUpdateDateTime nullable.Type[string] `json:"lastUpdateDateTime,omitempty"`

	// A list of possible Azure Attestation states for a device. Azure Attestation setting status is determined by report
	// sent from Microsoft Azure Attestation service. Only Windows 11 devices will have values "enabled" or "disabled".
	// Windows 10 devices will have value "notApplicable".
	MemoryAccessProtection *AzureAttestationSettingStatus `json:"memoryAccessProtection,omitempty"`

	// A list of possible Azure Attestation states for a device. Azure Attestation setting status is determined by report
	// sent from Microsoft Azure Attestation service. Only Windows 11 devices will have values "enabled" or "disabled".
	// Windows 10 devices will have value "notApplicable".
	MemoryIntegrityProtection *AzureAttestationSettingStatus `json:"memoryIntegrityProtection,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// When operatingSystemKernelDebugging is enabled, the device is used in development and testing
	OperatingSystemKernelDebugging nullable.Type[string] `json:"operatingSystemKernelDebugging,omitempty"`

	// The Operating System Revision List that was loaded during initial boot on the attested device
	OperatingSystemRevListInfo nullable.Type[string] `json:"operatingSystemRevListInfo,omitempty"`

	// The measurement that is captured in PCR[0]
	Pcr0 nullable.Type[string] `json:"pcr0,omitempty"`

	// Informational attribute that identifies the HASH algorithm that was used by TPM
	PcrHashAlgorithm nullable.Type[string] `json:"pcrHashAlgorithm,omitempty"`

	// The number of times a PC device has hibernated or resumed
	ResetCount *int64 `json:"resetCount,omitempty"`

	// The number of times a PC device has rebooted
	RestartCount *int64 `json:"restartCount,omitempty"`

	// Safe mode is a troubleshooting option for Windows that starts your computer in a limited state
	SafeMode nullable.Type[string] `json:"safeMode,omitempty"`

	// When Secure Boot is enabled, the core components must have the correct cryptographic signatures
	SecureBoot nullable.Type[string] `json:"secureBoot,omitempty"`

	// Fingerprint of the Custom Secure Boot Configuration Policy
	SecureBootConfigurationPolicyFingerPrint nullable.Type[string] `json:"secureBootConfigurationPolicyFingerPrint,omitempty"`

	// A list of possible Azure Attestation states for a device. Azure Attestation setting status is determined by report
	// sent from Microsoft Azure Attestation service. Only Windows 11 devices will have values "enabled" or "disabled".
	// Windows 10 devices will have value "notApplicable".
	SecuredCorePC *AzureAttestationSettingStatus `json:"securedCorePC,omitempty"`

	// A list of possible System Management Mode levels for a device. System Management Mode levels is determined by report
	// sent from Microsoft Azure Attestation service. Only specific hardwares support System Management Mode. Windows 11
	// devices will have values "notApplicable", "level1", "level2" or "level3". Windows 10 devices will have value
	// "notApplicable".
	SystemManagementMode *SystemManagementModeLevel `json:"systemManagementMode,omitempty"`

	// When test signing is allowed, the device does not enforce signature validation during boot
	TestSigning nullable.Type[string] `json:"testSigning,omitempty"`

	// The security version number of the Boot Application
	TpmVersion nullable.Type[string] `json:"tpmVersion,omitempty"`

	// Indicates whether the device has Virtual Secure Mode (VSM) enabled. Virtual Secure Mode (VSM) is a container that
	// protects high value assets from a compromised kernel. This property will be deprecated in beta from August 2023.
	// Support for this property will end in August 2025 for v1.0 API. A new property virtualizationBasedSecurity is added
	// and used instead. The value used for virtualSecureMode will be passed by virtualizationBasedSecurity during the
	// deprecation process. Possible values are 'enabled', 'disabled' and 'notApplicable'. 'enabled' indicates Virtual
	// Secure Mode (VSM) is enabled. 'disabled' indicates Virtual Secure Mode (VSM) is disabled. 'notApplicable' indicates
	// the device is not a Windows 11 device. Default value is 'notApplicable'.
	VirtualSecureMode nullable.Type[string] `json:"virtualSecureMode,omitempty"`

	// A list of possible Azure Attestation states for a device. Azure Attestation setting status is determined by report
	// sent from Microsoft Azure Attestation service. Only Windows 11 devices will have values "enabled" or "disabled".
	// Windows 10 devices will have value "notApplicable".
	VirtualizationBasedSecurity *AzureAttestationSettingStatus `json:"virtualizationBasedSecurity,omitempty"`

	// Operating system running with limited services that is used to prepare a computer for Windows
	WindowsPE nullable.Type[string] `json:"windowsPE,omitempty"`
}
