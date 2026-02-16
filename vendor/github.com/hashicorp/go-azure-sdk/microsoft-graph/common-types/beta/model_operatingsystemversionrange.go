package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperatingSystemVersionRange struct {
	// The description of this range (e.g. Valid 1702 builds)
	Description *string `json:"description,omitempty"`

	// The highest inclusive version that this range contains.
	HighestVersion *string `json:"highestVersion,omitempty"`

	// The lowest inclusive version that this range contains.
	LowestVersion *string `json:"lowestVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
