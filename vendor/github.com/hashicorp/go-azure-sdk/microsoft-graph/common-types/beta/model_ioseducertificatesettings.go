package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEduCertificateSettings struct {
	// File name to display in UI.
	CertFileName nullable.Type[string] `json:"certFileName,omitempty"`

	// PKCS Certificate Template Name.
	CertificateTemplateName *string `json:"certificateTemplateName,omitempty"`

	// Certificate Validity Period Options.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`

	// Value for the Certificate Validity Period.
	CertificateValidityPeriodValue *int64 `json:"certificateValidityPeriodValue,omitempty"`

	// PKCS Certification Authority.
	CertificationAuthority *string `json:"certificationAuthority,omitempty"`

	// PKCS Certification Authority Name.
	CertificationAuthorityName *string `json:"certificationAuthorityName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Certificate renewal threshold percentage. Valid values 1 to 99
	RenewalThresholdPercentage *int64 `json:"renewalThresholdPercentage,omitempty"`

	// Trusted Root Certificate.
	TrustedRootCertificate *string `json:"trustedRootCertificate,omitempty"`
}
