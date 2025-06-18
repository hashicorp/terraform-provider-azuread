package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeRightsAndInheritanceResult struct {
	ContentRights  *[]LabelContentRight `json:"contentRights,omitempty"`
	InheritedLabel *SensitivityLabel    `json:"inheritedLabel,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SensitivityLabels *[]SensitivityLabel `json:"sensitivityLabels,omitempty"`
}
