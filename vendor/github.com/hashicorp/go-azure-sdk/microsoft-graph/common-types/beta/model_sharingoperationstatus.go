package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingOperationStatus struct {
	// Provides a description of why this operation is not enabled. Only returned if this operation is not enabled.
	DisabledReason nullable.Type[string] `json:"disabledReason,omitempty"`

	// Indicates whether this operation is enabled.
	Enabled nullable.Type[bool] `json:"enabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
