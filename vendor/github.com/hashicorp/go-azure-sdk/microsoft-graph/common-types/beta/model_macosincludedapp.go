package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSIncludedApp struct {
	// The bundleId of the app. This maps to the CFBundleIdentifier in the app's bundle configuration.
	BundleId *string `json:"bundleId,omitempty"`

	// The version of the app. This maps to the CFBundleShortVersion in the app's bundle configuration.
	BundleVersion *string `json:"bundleVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
