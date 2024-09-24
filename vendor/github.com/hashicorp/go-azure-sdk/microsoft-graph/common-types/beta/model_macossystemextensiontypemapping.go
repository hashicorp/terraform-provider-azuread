package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSystemExtensionTypeMapping struct {
	// Flag enum representing the allowed macOS system extension types.
	AllowedTypes *MacOSSystemExtensionType `json:"allowedTypes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Gets or sets the team identifier used to sign the system extension.
	TeamIdentifier *string `json:"teamIdentifier,omitempty"`
}
