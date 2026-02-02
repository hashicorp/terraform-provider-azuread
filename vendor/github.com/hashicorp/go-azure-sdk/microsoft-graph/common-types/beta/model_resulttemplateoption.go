package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResultTemplateOption struct {
	// Indicates whether search display layouts are enabled. If enabled, the user will get the result template to render the
	// search results content in the resultTemplates property of the response. The result template is based on Adaptive
	// Cards. This property is optional.
	EnableResultTemplate nullable.Type[bool] `json:"enableResultTemplate,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
