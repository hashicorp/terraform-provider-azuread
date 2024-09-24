package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditActivityInitiator struct {
	// If the actor initiating the activity is an app, this property indicates all its identification information including
	// appId, displayName, servicePrincipalId, and servicePrincipalName.
	App *AppIdentity `json:"app,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If the actor initiating the activity is a user, this property indicates their identification information including
	// their id, displayName, and userPrincipalName.
	User *AuditUserIdentity `json:"user,omitempty"`
}
