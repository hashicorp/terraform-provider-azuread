package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsAcl struct {
	AccessType *ExternalConnectorsAccessType `json:"accessType,omitempty"`

	// The source of identity. Possible values are azureActiveDirectory or external.
	IdentitySource *ExternalConnectorsIdentitySourceType `json:"identitySource,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Type *ExternalConnectorsAclType `json:"type,omitempty"`

	// The unique identifer of the identity. For Microsoft Entra identities, value is set to the object identifier of the
	// user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. For external
	// groups, value is set to the ID of the externalGroup.
	Value *string `json:"value,omitempty"`
}
