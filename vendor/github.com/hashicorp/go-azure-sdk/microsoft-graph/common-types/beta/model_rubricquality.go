package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricQuality struct {
	// The collection of criteria for this rubric quality.
	Criteria *[]RubricCriterion `json:"criteria,omitempty"`

	// The description of this rubric quality.
	Description *EducationItemBody `json:"description,omitempty"`

	// The name of this rubric quality.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ID of this resource.
	QualityId nullable.Type[string] `json:"qualityId,omitempty"`
}
