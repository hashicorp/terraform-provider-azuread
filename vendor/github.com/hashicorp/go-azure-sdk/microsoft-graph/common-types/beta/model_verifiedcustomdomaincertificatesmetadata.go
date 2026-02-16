package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedCustomDomainCertificatesMetadata struct {
	// The expiry date of the custom domain certificate. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ExpiryDate nullable.Type[string] `json:"expiryDate,omitempty"`

	// The issue date of the custom domain. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	IssueDate nullable.Type[string] `json:"issueDate,omitempty"`

	// The issuer name of the custom domain certificate.
	IssuerName nullable.Type[string] `json:"issuerName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The subject name of the custom domain certificate.
	SubjectName nullable.Type[string] `json:"subjectName,omitempty"`

	// The thumbprint associated with the custom domain certificate.
	Thumbprint nullable.Type[string] `json:"thumbprint,omitempty"`
}
