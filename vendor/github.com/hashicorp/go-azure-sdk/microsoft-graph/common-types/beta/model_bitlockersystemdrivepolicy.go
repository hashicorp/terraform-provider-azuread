package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicy struct {
	// Select the encryption method for operating system drives. Possible values are: aesCbc128, aesCbc256, xtsAes128,
	// xtsAes256.
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`

	// Indicates the minimum length of startup pin. Valid values 4 to 20
	MinimumPinLength nullable.Type[int64] `json:"minimumPinLength,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Enable pre-boot recovery message and Url. If requireStartupAuthentication is false, this value does not affect.
	PrebootRecoveryEnableMessageAndUrl *bool `json:"prebootRecoveryEnableMessageAndUrl,omitempty"`

	// Defines a custom recovery message.
	PrebootRecoveryMessage nullable.Type[string] `json:"prebootRecoveryMessage,omitempty"`

	// Defines a custom recovery URL.
	PrebootRecoveryUrl nullable.Type[string] `json:"prebootRecoveryUrl,omitempty"`

	// Allows to recover BitLocker encrypted operating system drives in the absence of the required startup key information.
	// This policy setting is applied when you turn on BitLocker.
	RecoveryOptions *BitLockerRecoveryOptions `json:"recoveryOptions,omitempty"`

	// Indicates whether to allow BitLocker without a compatible TPM (requires a password or a startup key on a USB flash
	// drive).
	StartupAuthenticationBlockWithoutTpmChip *bool `json:"startupAuthenticationBlockWithoutTpmChip,omitempty"`

	// Require additional authentication at startup.
	StartupAuthenticationRequired *bool `json:"startupAuthenticationRequired,omitempty"`

	// Possible values of the ConfigurationUsage list.
	StartupAuthenticationTpmKeyUsage *ConfigurationUsage `json:"startupAuthenticationTpmKeyUsage,omitempty"`

	// Possible values of the ConfigurationUsage list.
	StartupAuthenticationTpmPinAndKeyUsage *ConfigurationUsage `json:"startupAuthenticationTpmPinAndKeyUsage,omitempty"`

	// Possible values of the ConfigurationUsage list.
	StartupAuthenticationTpmPinUsage *ConfigurationUsage `json:"startupAuthenticationTpmPinUsage,omitempty"`

	// Possible values of the ConfigurationUsage list.
	StartupAuthenticationTpmUsage *ConfigurationUsage `json:"startupAuthenticationTpmUsage,omitempty"`
}
