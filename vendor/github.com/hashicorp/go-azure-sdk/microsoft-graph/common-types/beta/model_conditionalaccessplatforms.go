package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPlatforms struct {
	// Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue, linux.
	ExcludePlatforms *[]ConditionalAccessDevicePlatform `json:"excludePlatforms,omitempty"`

	// Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue,linux.
	IncludePlatforms *[]ConditionalAccessDevicePlatform `json:"includePlatforms,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
