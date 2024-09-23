package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HasPayloadLinkResultItem struct {
	// Exception information indicates if check for this item was successful or not.Empty string for no error.
	Error nullable.Type[string] `json:"error,omitempty"`

	// Indicate whether a payload has any link or not.
	HasLink nullable.Type[bool] `json:"hasLink,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Key of the Payload, In the format of Guid.
	PayloadId nullable.Type[string] `json:"payloadId,omitempty"`

	// The reason where the link comes from.
	Sources *[]DeviceAndAppManagementAssignmentSource `json:"sources,omitempty"`
}
