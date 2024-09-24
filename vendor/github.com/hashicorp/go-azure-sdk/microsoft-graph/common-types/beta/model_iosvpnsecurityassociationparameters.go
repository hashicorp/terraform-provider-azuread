package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnSecurityAssociationParameters struct {
	// Lifetime (minutes)
	LifetimeInMinutes nullable.Type[int64] `json:"lifetimeInMinutes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Diffie-Hellman Group
	SecurityDiffieHellmanGroup nullable.Type[int64] `json:"securityDiffieHellmanGroup,omitempty"`

	// Encryption algorithm. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm,
	// chaCha20Poly1305.
	SecurityEncryptionAlgorithm *VpnEncryptionAlgorithmType `json:"securityEncryptionAlgorithm,omitempty"`

	// Integrity algorithm. Possible values are: sha2256, sha196, sha1160, sha2384, sha2_512, md5.
	SecurityIntegrityAlgorithm *VpnIntegrityAlgorithmType `json:"securityIntegrityAlgorithm,omitempty"`
}
