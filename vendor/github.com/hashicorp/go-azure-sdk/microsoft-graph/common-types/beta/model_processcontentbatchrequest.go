package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessContentBatchRequest struct {
	ContentToProcess *ProcessContentRequest `json:"contentToProcess,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A unique identifier provided by the client to correlate this specific request item within the batch.
	RequestId nullable.Type[string] `json:"requestId,omitempty"`

	// The unique identifier (Object ID or UPN) of the user in whose context the content should be processed.
	UserId nullable.Type[string] `json:"userId,omitempty"`
}
