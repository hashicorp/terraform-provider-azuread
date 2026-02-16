package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSLobChildApp struct {
	// The build number of the app.
	BuildNumber nullable.Type[string] `json:"buildNumber,omitempty"`

	// The bundleId of the app.
	BundleId nullable.Type[string] `json:"bundleId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The version number of the app.
	VersionNumber nullable.Type[string] `json:"versionNumber,omitempty"`
}
