package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditActivityInitiator struct {
	// If the resource initiating the activity is an app, this property indicates all the app related information like appId
	// and name.
	App *AppIdentity `json:"app,omitempty"`

	// A set of linkable claims to link together all the authentication artifacts issued from a single interactive root
	// authentication.
	LinkableIdentifiers *LinkableIdentifiers `json:"linkableIdentifiers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If the resource initiating the activity is a user, this property Indicates all the user related information like user
	// ID and userPrincipalName.
	User *AuditUserIdentity `json:"user,omitempty"`
}
