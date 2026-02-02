package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataAdditionalClassGroupOptions struct {
	// Indicates whether a team should be created for the class group.
	CreateTeam *bool `json:"createTeam,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether the class group display name should be set on create.
	WriteDisplayNameOnCreateOnly *bool `json:"writeDisplayNameOnCreateOnly,omitempty"`
}
