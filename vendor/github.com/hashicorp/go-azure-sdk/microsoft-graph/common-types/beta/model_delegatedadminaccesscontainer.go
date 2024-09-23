package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminAccessContainer struct {
	// The identifier of the access container (for example, a security group). For 'securityGroup' access containers, this
	// must be a valid ID of a Microsoft Entra security group in the Microsoft partner's tenant.
	AccessContainerId *string `json:"accessContainerId,omitempty"`

	AccessContainerType *DelegatedAdminAccessContainerType `json:"accessContainerType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
