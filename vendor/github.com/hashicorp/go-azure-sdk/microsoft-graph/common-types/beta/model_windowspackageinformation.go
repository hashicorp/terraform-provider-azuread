package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPackageInformation struct {
	// Contains properties for Windows architecture.
	ApplicableArchitecture *WindowsArchitecture `json:"applicableArchitecture,omitempty"`

	// The Display Name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The Identity Name.
	IdentityName nullable.Type[string] `json:"identityName,omitempty"`

	// The Identity Publisher.
	IdentityPublisher nullable.Type[string] `json:"identityPublisher,omitempty"`

	// The Identity Resource Identifier.
	IdentityResourceIdentifier nullable.Type[string] `json:"identityResourceIdentifier,omitempty"`

	// The Identity Version.
	IdentityVersion nullable.Type[string] `json:"identityVersion,omitempty"`

	// The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
