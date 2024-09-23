package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectorySizeQuota struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total amount of the directory quota.
	Total nullable.Type[int64] `json:"total,omitempty"`

	// Used amount of the directory quota.
	Used nullable.Type[int64] `json:"used,omitempty"`
}
