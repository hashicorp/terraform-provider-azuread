package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationInnerError struct {
	// The activity ID associated with the request that generated the error.
	ActivityId nullable.Type[string] `json:"activityId,omitempty"`

	// The client request ID, if provided by the caller.
	ClientRequestId nullable.Type[string] `json:"clientRequestId,omitempty"`

	// A more specific, potentially internal, error code string.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The date and time the inner error occurred.
	ErrorDateTime nullable.Type[string] `json:"errorDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
