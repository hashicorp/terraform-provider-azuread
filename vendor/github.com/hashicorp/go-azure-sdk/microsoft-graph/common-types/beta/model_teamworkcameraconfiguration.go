package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkCameraConfiguration struct {
	Cameras *[]TeamworkPeripheral `json:"cameras,omitempty"`

	// The configuration for the content camera.
	ContentCameraConfiguration *TeamworkContentCameraConfiguration `json:"contentCameraConfiguration,omitempty"`

	DefaultContentCamera *TeamworkPeripheral `json:"defaultContentCamera,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
