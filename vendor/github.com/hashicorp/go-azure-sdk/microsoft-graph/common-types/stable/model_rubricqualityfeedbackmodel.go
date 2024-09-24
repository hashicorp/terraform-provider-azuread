package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricQualityFeedbackModel struct {
	// Specific feedback for one quality of this rubric.
	Feedback *EducationItemBody `json:"feedback,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ID of the rubricQuality that this feedback is related to.
	QualityId nullable.Type[string] `json:"qualityId,omitempty"`
}
