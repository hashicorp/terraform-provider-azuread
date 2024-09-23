package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfSignedCertificate struct {
	// Custom key identifier.
	CustomKeyIdentifier nullable.Type[string] `json:"customKeyIdentifier,omitempty"`

	// The friendly name for the key.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time at which the credential expires. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The value for the key credential. Should be a base-64 encoded value.
	Key nullable.Type[string] `json:"key,omitempty"`

	// The unique identifier (GUID) for the key.
	KeyId nullable.Type[string] `json:"keyId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date and time at which the credential becomes valid. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The thumbprint value for the key.
	Thumbprint nullable.Type[string] `json:"thumbprint,omitempty"`

	// The type of key credential. 'AsymmetricX509Cert'.
	Type nullable.Type[string] `json:"type,omitempty"`

	// A string that describes the purpose for which the key can be used. For example, 'Verify'.
	Usage nullable.Type[string] `json:"usage,omitempty"`
}
