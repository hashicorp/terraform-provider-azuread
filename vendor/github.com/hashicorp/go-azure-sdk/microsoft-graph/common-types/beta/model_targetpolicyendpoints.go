package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetPolicyEndpoints struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Use to filter the notification distribution to a specific platform or platforms. Valid values are Windows, iOS,
	// Android and WebPush. By default, all push endpoint types (Windows, iOS, Android and WebPush) are enabled.
	PlatformTypes *[]string `json:"platformTypes,omitempty"`
}
