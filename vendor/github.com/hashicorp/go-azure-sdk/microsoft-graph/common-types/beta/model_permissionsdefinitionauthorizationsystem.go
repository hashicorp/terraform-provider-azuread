package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinitionAuthorizationSystem struct {
	// ID of the authorization system retrieved from the customer cloud environment.
	AuthorizationSystemId *string `json:"authorizationSystemId,omitempty"`

	// The type of authorization system.
	AuthorizationSystemType *string `json:"authorizationSystemType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
