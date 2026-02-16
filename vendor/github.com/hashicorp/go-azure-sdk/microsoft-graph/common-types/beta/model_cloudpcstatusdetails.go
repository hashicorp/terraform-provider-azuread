package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCStatusDetails struct {
	// Any additional information about the Cloud PC status.
	AdditionalInformation *[]KeyValuePair `json:"additionalInformation,omitempty"`

	// The code associated with the Cloud PC status.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The status message.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
