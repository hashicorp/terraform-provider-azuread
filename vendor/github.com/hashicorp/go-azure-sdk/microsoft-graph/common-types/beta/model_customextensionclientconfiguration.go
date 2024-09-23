package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionClientConfiguration struct {
	// The max number of retries that Microsoft Entra ID makes to the external API. Values of 0 or 1 are supported. If null,
	// the default for the service applies.
	MaximumRetries nullable.Type[int64] `json:"maximumRetries,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The max duration in milliseconds that Microsoft Entra ID waits for a response from the external app before it shuts
	// down the connection. The valid range is between 200 and 2000 milliseconds. If null, the default for the service
	// applies.
	TimeoutInMilliseconds nullable.Type[int64] `json:"timeoutInMilliseconds,omitempty"`
}
