package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsCreepIndex struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// This value represents how much risk an identity poses. This risk range is classified in three buckets: 0-33: low,
	// 34-66: medium, 67-100: high..
	Score *int64 `json:"score,omitempty"`
}
