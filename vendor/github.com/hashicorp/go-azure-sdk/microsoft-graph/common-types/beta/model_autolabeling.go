package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoLabeling struct {
	// The message displayed to the user when the label is applied automatically.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The list of sensitive information type (SIT) IDs that trigger the automatic application of this label.
	SensitiveTypeIds *[]string `json:"sensitiveTypeIds,omitempty"`
}
