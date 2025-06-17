package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateCRLValidationConfiguration struct {
	// Represents the SKIs of CAs that should be excluded from the valid CRL distribution point check. SKI is represented as
	// a hexadecimal string.
	ExemptedCertificateAuthoritiesSubjectKeyIdentifiers *[]string `json:"exemptedCertificateAuthoritiesSubjectKeyIdentifiers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *X509CertificateCRLValidationConfigurationState `json:"state,omitempty"`
}
