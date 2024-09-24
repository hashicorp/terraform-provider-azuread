package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVendorInformation struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specific provider (product/service - not vendor company); for example, WindowsDefenderATP.
	Provider nullable.Type[string] `json:"provider,omitempty"`

	// Version of the provider or subprovider, if it exists, that generated the alert. Required
	ProviderVersion nullable.Type[string] `json:"providerVersion,omitempty"`

	// Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen.
	SubProvider nullable.Type[string] `json:"subProvider,omitempty"`

	// Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required
	Vendor nullable.Type[string] `json:"vendor,omitempty"`
}
