package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustFrameworkKey struct {
	// RSA Key - private exponent. The field isn't readable.
	D nullable.Type[string] `json:"d,omitempty"`

	// RSA Key - first exponent. The field isn't readable.
	Dp nullable.Type[string] `json:"dp,omitempty"`

	// RSA Key - second exponent. The field isn't readable.
	Dq nullable.Type[string] `json:"dq,omitempty"`

	// RSA Key - public exponent.
	E nullable.Type[string] `json:"e,omitempty"`

	// This value is a NumericDate as defined in RFC 7519. That is, a JSON numeric value representing the number of seconds
	// from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.
	Exp nullable.Type[int64] `json:"exp,omitempty"`

	// Symmetric Key for oct key type. The field isn't readable.
	K nullable.Type[string] `json:"k,omitempty"`

	// The unique identifier for the key.
	Kid nullable.Type[string] `json:"kid,omitempty"`

	// The kty (key type) parameter identifies the cryptographic algorithm family used with the key. The valid values are
	// rsa, oct.
	Kty nullable.Type[string] `json:"kty,omitempty"`

	// RSA Key - modulus.
	N nullable.Type[string] `json:"n,omitempty"`

	// This value is a NumericDate as defined in RFC 7519. That is, a JSON numeric value representing the number of seconds
	// from 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap seconds.
	Nbf nullable.Type[int64] `json:"nbf,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// RSA Key - first prime. The field isn't readable.
	P nullable.Type[string] `json:"p,omitempty"`

	// RSA Key - second prime. The field isn't readable.
	Q nullable.Type[string] `json:"q,omitempty"`

	// RSA Key - Coefficient. The field isn't readable.
	Qi nullable.Type[string] `json:"qi,omitempty"`

	// Status of the key. The possible values are: enabled, disabled, unknownFutureValue.
	Status *TrustFrameworkKeyStatus `json:"status,omitempty"`

	// The use (public key use) parameter identifies the intended use of the public key. The use parameter is employed to
	// indicate whether a public key is used for encrypting data or verifying the signature on data. Possible values are:
	// sig (signature), enc (encryption).
	Use nullable.Type[string] `json:"use,omitempty"`

	// The x5c (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates. For more information,
	// see RFC 5280.
	X5c *[]string `json:"x5c,omitempty"`

	// The x5t (X.509 certificate SHA-1 thumbprint) parameter is a base64url-encoded SHA-1 thumbprint (also known as digest)
	// of the DER encoding of an X.509 certificate. For more information, see RFC 5280.
	X5t nullable.Type[string] `json:"x5t,omitempty"`
}
