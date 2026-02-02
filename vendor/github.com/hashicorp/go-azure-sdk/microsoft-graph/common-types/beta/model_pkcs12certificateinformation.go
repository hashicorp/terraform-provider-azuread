package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Pkcs12CertificateInformation struct {
	// Represents whether the certificate is the active certificate to be used for calling the API connector. The active
	// certificate is the most recently uploaded certificate that isn't yet expired but whose notBefore time is in the past.
	IsActive *bool `json:"isActive,omitempty"`

	// The certificate's expiry. This value is a NumericDate as defined in RFC 7519 (A JSON numeric value representing the
	// number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.)
	NotAfter *int64 `json:"notAfter,omitempty"`

	// The certificate's issue time (not before). This value is a NumericDate as defined in RFC 7519 (A JSON numeric value
	// representing the number of seconds from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap
	// seconds.)
	NotBefore *int64 `json:"notBefore,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The certificate thumbprint.
	Thumbprint nullable.Type[string] `json:"thumbprint,omitempty"`
}
