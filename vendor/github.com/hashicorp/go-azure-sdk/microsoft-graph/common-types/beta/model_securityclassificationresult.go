package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityClassificationResult struct {
	// The confidence level, 0 to 100, of the result.
	ConfidenceLevel *int64 `json:"confidenceLevel,omitempty"`

	// The number of instances of the specific information type in the input.
	Count *int64 `json:"count,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The GUID of the discovered sensitive information type.
	SensitiveTypeId *string `json:"sensitiveTypeId,omitempty"`
}
