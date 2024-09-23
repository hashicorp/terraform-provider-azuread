package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerFixedDrivePolicy struct {
	// Select the encryption method for fixed drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// This policy setting allows you to control how BitLocker-protected fixed data drives are recovered in the absence of
	// the required credentials. This policy setting is applied when you turn on BitLocker.
	RecoveryOptions *BitLockerRecoveryOptions `json:"recoveryOptions,omitempty"`

	// This policy setting determines whether BitLocker protection is required for fixed data drives to be writable on a
	// computer.
	RequireEncryptionForWriteAccess *bool `json:"requireEncryptionForWriteAccess,omitempty"`
}
