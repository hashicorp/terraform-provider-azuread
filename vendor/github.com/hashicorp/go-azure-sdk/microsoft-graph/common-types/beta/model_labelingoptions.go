package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelingOptions struct {
	AssignmentMethod *AssignmentMethod `json:"assignmentMethod,omitempty"`

	// The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
	DowngradeJustification *DowngradeJustification `json:"downgradeJustification,omitempty"`

	// Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label
	// information.
	ExtendedProperties *[]KeyValuePair `json:"extendedProperties,omitempty"`

	// The GUID of the label that should be applied to the information.
	LabelId *string `json:"labelId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
