package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateRule struct {
	// The identifier of the X.509 certificate. Required.
	Identifier nullable.Type[string] `json:"identifier,omitempty"`

	// The identifier of the certificate issuer.
	IssuerSubjectIdentifier nullable.Type[string] `json:"issuerSubjectIdentifier,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identifier of the X.509 certificate policyOID.
	PolicyOidIdentifier nullable.Type[string] `json:"policyOidIdentifier,omitempty"`

	// The type of strong authentication mode. The possible values are: x509CertificateSingleFactor,
	// x509CertificateMultiFactor, unknownFutureValue. Required.
	X509CertificateAuthenticationMode X509CertificateAuthenticationMode `json:"x509CertificateAuthenticationMode"`

	// The possible values are: low, high, unknownFutureValue.
	X509CertificateRequiredAffinityLevel *X509CertificateAffinityLevel `json:"x509CertificateRequiredAffinityLevel,omitempty"`

	// The type of the X.509 certificate mode configuration rule. The possible values are: issuerSubject, policyOID,
	// unknownFutureValue, issuerSubjectAndPolicyOID. Use the Prefer: include-unknown-enum-members request header to get the
	// following values from this evolvable enum: issuerSubjectAndPolicyOID. Required.
	X509CertificateRuleType X509CertificateRuleType `json:"x509CertificateRuleType"`
}
