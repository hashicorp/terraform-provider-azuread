package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSKernelExtension struct {
	// Bundle ID of the kernel extension.
	BundleId *string `json:"bundleId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The team identifier that was used to sign the kernel extension.
	TeamIdentifier nullable.Type[string] `json:"teamIdentifier,omitempty"`
}
