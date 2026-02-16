package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemViewpoint struct {
	// Indicates whether the user can perform the described actions on this item.
	AccessOperations *DriveItemAccessOperationsViewpoint `json:"accessOperations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates sharing operations the current user can take on the specified item.
	Sharing *SharingViewpoint `json:"sharing,omitempty"`
}
