package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRemovableDrivePolicy struct {
	// This policy setting determines whether BitLocker protection is required for removable data drives to be writable on a
	// computer.
	BlockCrossOrganizationWriteAccess *bool `json:"blockCrossOrganizationWriteAccess,omitempty"`

	// Select the encryption method for removable drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
	EncryptionMethod *BitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether to block write access to devices configured in another organization. If
	// requireEncryptionForWriteAccess is false, this value does not affect.
	RequireEncryptionForWriteAccess *bool `json:"requireEncryptionForWriteAccess,omitempty"`
}
