package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAnswerChoice struct {
	// The actual value of the selected choice. This is typically a string value which is understandable by applications.
	// Required.
	ActualValue nullable.Type[string] `json:"actualValue,omitempty"`

	// The localized display values shown to the requestor and approvers. Required.
	DisplayValue AccessPackageLocalizedContent `json:"displayValue"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
