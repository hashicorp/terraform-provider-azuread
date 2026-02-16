package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptions struct {
	// Indicates whether to block certificate-based data recovery agent.
	BlockDataRecoveryAgent *bool `json:"blockDataRecoveryAgent,omitempty"`

	// Indicates whether or not to enable BitLocker until recovery information is stored in AD DS.
	EnableBitLockerAfterRecoveryInformationToStore *bool `json:"enableBitLockerAfterRecoveryInformationToStore,omitempty"`

	// Indicates whether or not to allow BitLocker recovery information to store in AD DS.
	EnableRecoveryInformationSaveToStore *bool `json:"enableRecoveryInformationSaveToStore,omitempty"`

	// Indicates whether or not to allow showing recovery options in BitLocker Setup Wizard for fixed or system disk.
	HideRecoveryOptions *bool `json:"hideRecoveryOptions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// BitLockerRecoveryInformationType types
	RecoveryInformationToStore *BitLockerRecoveryInformationType `json:"recoveryInformationToStore,omitempty"`

	// Possible values of the ConfigurationUsage list.
	RecoveryKeyUsage *ConfigurationUsage `json:"recoveryKeyUsage,omitempty"`

	// Possible values of the ConfigurationUsage list.
	RecoveryPasswordUsage *ConfigurationUsage `json:"recoveryPasswordUsage,omitempty"`
}
