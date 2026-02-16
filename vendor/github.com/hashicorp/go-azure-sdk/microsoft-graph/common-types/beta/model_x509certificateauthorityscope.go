package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthorityScope struct {
	// A collection of groups that are enabled to be in scope to use certificates issued by specific certificate authority.
	IncludeTargets *[]IncludeTarget `json:"includeTargets,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Public Key Infrastructure container object under which the certificate authorities are stored in the Entra PKI based
	// trust store.
	PublicKeyInfrastructureIdentifier nullable.Type[string] `json:"publicKeyInfrastructureIdentifier,omitempty"`

	// Subject Key Identifier that identifies the certificate authority uniquely.
	SubjectKeyIdentifier nullable.Type[string] `json:"subjectKeyIdentifier,omitempty"`
}
