package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequiredResourceAccess struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
	ResourceAccess *[]ResourceAccess `json:"resourceAccess,omitempty"`

	// The unique identifier for the resource that the application requires access to. This should be equal to the appId
	// declared on the target resource application.
	ResourceAppId *string `json:"resourceAppId,omitempty"`
}
