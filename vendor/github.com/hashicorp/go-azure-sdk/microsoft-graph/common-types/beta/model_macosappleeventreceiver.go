package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSAppleEventReceiver struct {
	// Allow or block this app from receiving Apple events.
	Allowed *bool `json:"allowed,omitempty"`

	// Code requirement for the app or binary that receives the Apple Event.
	CodeRequirement *string `json:"codeRequirement,omitempty"`

	// Bundle ID of the app or file path of the process or executable that receives the Apple Event.
	Identifier *string `json:"identifier,omitempty"`

	// Process identifier types for MacOS Privacy Preferences
	IdentifierType *MacOSProcessIdentifierType `json:"identifierType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
