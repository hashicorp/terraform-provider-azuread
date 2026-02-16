package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenLicenseSummary struct {
	// The Apple Id associated with the given Apple Volume Purchase Program Token.
	AppleId nullable.Type[string] `json:"appleId,omitempty"`

	// The number of VPP licenses available.
	AvailableLicenseCount *int64 `json:"availableLicenseCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The organization associated with the Apple Volume Purchase Program Token.
	OrganizationName nullable.Type[string] `json:"organizationName,omitempty"`

	// The number of VPP licenses in use.
	UsedLicenseCount *int64 `json:"usedLicenseCount,omitempty"`

	// Identifier of the VPP token.
	VppTokenId nullable.Type[string] `json:"vppTokenId,omitempty"`
}
