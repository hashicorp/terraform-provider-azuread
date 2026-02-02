package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyCredential struct {
	// A 40-character binary type that can be used to identify the credential. Optional. When not provided in the payload,
	// defaults to the thumbprint of the certificate.
	CustomKeyIdentifier nullable.Type[string] `json:"customKeyIdentifier,omitempty"`

	// The friendly name for the key, with a maximum length of 90 characters. Longer values are accepted but shortened.
	// Optional.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time at which the credential expires. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Value for the key credential. Should be a Base64 encoded value. Returned only on $select for a single object, that
	// is, GET applications/{applicationId}?$select=keyCredentials or GET
	// servicePrincipals/{servicePrincipalId}?$select=keyCredentials; otherwise, it's always null. From a .cer certificate,
	// you can read the key using the Convert.ToBase64String() method. For more information, see Get the certificate key.
	Key nullable.Type[string] `json:"key,omitempty"`

	// The unique identifier for the key.
	KeyId nullable.Type[string] `json:"keyId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date and time at which the credential becomes valid.The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The type of key credential; for example, Symmetric, AsymmetricX509Cert, or X509CertAndPassword.
	Type nullable.Type[string] `json:"type,omitempty"`

	// A string that describes the purpose for which the key can be used; for example, None​, Verify​,
	// PairwiseIdentifier​, Delegation​, Decrypt​, Encrypt​, HashedIdentifier​, SelfSignedTls, or Sign. If usage
	// is Sign​, the type should be X509CertAndPassword​, and the passwordCredentials​ for signing should be defined.
	Usage nullable.Type[string] `json:"usage,omitempty"`
}
