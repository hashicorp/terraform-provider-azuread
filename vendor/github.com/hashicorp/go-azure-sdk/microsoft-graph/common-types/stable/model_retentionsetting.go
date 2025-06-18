package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionSetting struct {
	// The frequency of the backup.
	Interval nullable.Type[string] `json:"interval,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The period of time to retain the protected data for a single Microsoft 365 service.
	Period nullable.Type[string] `json:"period,omitempty"`
}
