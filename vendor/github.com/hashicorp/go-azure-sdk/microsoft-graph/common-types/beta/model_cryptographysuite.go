package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuite struct {
	// Authentication Transform Constants. Possible values are: md596, sha196, sha256128, aes128Gcm, aes192Gcm, aes256Gcm.
	AuthenticationTransformConstants *AuthenticationTransformConstant `json:"authenticationTransformConstants,omitempty"`

	// Cipher Transform Constants. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192,
	// aes192Gcm, chaCha20Poly1305.
	CipherTransformConstants *VpnEncryptionAlgorithmType `json:"cipherTransformConstants,omitempty"`

	// Diffie Hellman Group. Possible values are: group1, group2, group14, ecp256, ecp384, group24.
	DhGroup *DiffieHellmanGroup `json:"dhGroup,omitempty"`

	// Encryption Method. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm,
	// chaCha20Poly1305.
	EncryptionMethod *VpnEncryptionAlgorithmType `json:"encryptionMethod,omitempty"`

	// Integrity Check Method. Possible values are: sha2256, sha196, sha1160, sha2384, sha2_512, md5.
	IntegrityCheckMethod *VpnIntegrityAlgorithmType `json:"integrityCheckMethod,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Perfect Forward Secrecy Group. Possible values are: pfs1, pfs2, pfs2048, ecp256, ecp384, pfsMM, pfs24.
	PfsGroup *PerfectForwardSecrecyGroup `json:"pfsGroup,omitempty"`
}
