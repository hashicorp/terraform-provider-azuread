package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPlatforms struct {
	// Possible values are: android, iOS, windows, windowsPhone, macOS, linux, all, unknownFutureValue.
	ExcludePlatforms *[]ConditionalAccessDevicePlatform `json:"excludePlatforms,omitempty"`

	// Possible values are: android, iOS, windows, windowsPhone, macOS, linux, all, unknownFutureValue.
	IncludePlatforms *[]ConditionalAccessDevicePlatform `json:"includePlatforms,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
