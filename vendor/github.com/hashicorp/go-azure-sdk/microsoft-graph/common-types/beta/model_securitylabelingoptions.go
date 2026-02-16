package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLabelingOptions struct {
	AssignmentMethod *SecurityAssignmentMethod `json:"assignmentMethod,omitempty"`

	// The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
	DowngradeJustification *SecurityDowngradeJustification `json:"downgradeJustification,omitempty"`

	// Extended properties will be parsed and returned in the standard Microsoft Purview Information Protection labeled
	// metadata format as part of the label information.
	ExtendedProperties *[]SecurityKeyValuePair `json:"extendedProperties,omitempty"`

	// The GUID of the label that should be applied to the information.
	LabelId *string `json:"labelId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
