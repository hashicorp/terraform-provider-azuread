package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelAssignment struct {
	AssignmentMethod *SensitivityLabelAssignmentMethod `json:"assignmentMethod,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier for the sensitivity label assigned to the file.
	SensitivityLabelId *string `json:"sensitivityLabelId,omitempty"`

	// The unique identifier for the tenant that hosts the file when this label is applied.
	TenantId *string `json:"tenantId,omitempty"`
}
