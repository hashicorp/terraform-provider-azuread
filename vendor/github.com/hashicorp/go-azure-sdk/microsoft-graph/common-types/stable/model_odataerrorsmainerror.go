package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ODataErrorsMainError struct {
	Code    *string                    `json:"code,omitempty"`
	Details *[]ODataErrorsErrorDetails `json:"details,omitempty"`

	// The structure of this object is service-specific
	InnerError *ODataErrorsInnerError `json:"innerError,omitempty"`

	Message *string `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Target nullable.Type[string] `json:"target,omitempty"`
}
