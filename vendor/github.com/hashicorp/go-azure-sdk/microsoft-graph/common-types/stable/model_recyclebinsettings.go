package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecycleBinSettings struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Recycle bin retention period override in days for deleted content. The default value is 93; the value range is 7 to
	// 180. The setting applies to newly deleted content only. Setting this property to null reverts to its default value.
	// Read-write.
	RetentionPeriodOverrideDays nullable.Type[int64] `json:"retentionPeriodOverrideDays,omitempty"`
}
