package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthenticationModeConfiguration struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Rules are configured in addition to the authentication mode to bind a specific x509CertificateRuleType to an
	// x509CertificateAuthenticationMode. For example, bind the policyOID with identifier 1.32.132.343 to
	// x509CertificateMultiFactor authentication mode.
	Rules *[]X509CertificateRule `json:"rules,omitempty"`

	// The type of strong authentication mode. The possible values are: x509CertificateSingleFactor,
	// x509CertificateMultiFactor, unknownFutureValue.
	X509CertificateAuthenticationDefaultMode *X509CertificateAuthenticationMode `json:"x509CertificateAuthenticationDefaultMode,omitempty"`

	// Determines the default value for the tenant affinity binding level. The possible values are: low, high,
	// unknownFutureValue.
	X509CertificateDefaultRequiredAffinityLevel *X509CertificateAffinityLevel `json:"x509CertificateDefaultRequiredAffinityLevel,omitempty"`
}
