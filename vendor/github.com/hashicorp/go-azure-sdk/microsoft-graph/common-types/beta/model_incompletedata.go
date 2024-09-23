package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncompleteData struct {
	// The service doesn't have source data before the specified time.
	MissingDataBeforeDateTime nullable.Type[string] `json:"missingDataBeforeDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Some data wasn't recorded due to excessive activity.
	WasThrottled nullable.Type[bool] `json:"wasThrottled,omitempty"`
}
