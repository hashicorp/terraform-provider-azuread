package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiableCredentialSettings struct {
	// The types of verifiable credentials that a requestor must present when requesting an access package that has the
	// policy.
	CredentialTypes *[]VerifiableCredentialType `json:"credentialTypes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
