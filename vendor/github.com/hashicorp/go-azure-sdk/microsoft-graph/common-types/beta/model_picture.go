package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Picture struct {
	Content     nullable.Type[string] `json:"content,omitempty"`
	ContentType nullable.Type[string] `json:"contentType,omitempty"`
	Height      nullable.Type[int64]  `json:"height,omitempty"`
	Id          *string               `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Width nullable.Type[int64] `json:"width,omitempty"`
}
