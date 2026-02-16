package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Acl struct {
	AccessType     *AccessType         `json:"accessType,omitempty"`
	IdentitySource *IdentitySourceType `json:"identitySource,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Type  *AclType `json:"type,omitempty"`
	Value *string  `json:"value,omitempty"`
}
