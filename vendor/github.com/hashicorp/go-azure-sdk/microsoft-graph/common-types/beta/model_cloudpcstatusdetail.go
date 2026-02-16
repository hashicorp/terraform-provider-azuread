package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCStatusDetail struct {
	// More information about the Cloud PC status. For example, 'additionalInformation': ['{'@odata.type':
	// 'microsoft.graph.keyValuePair','name': 'retriable','value': true }] '
	AdditionalInformation *[]KeyValuePair `json:"additionalInformation,omitempty"`

	// The error/warning code associated with the Cloud PC status. Example: 'code': 'internalServerError'.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The status message associated with error code. Example: 'message': 'There was an internal server error. Please
	// contact support xxx.'.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
