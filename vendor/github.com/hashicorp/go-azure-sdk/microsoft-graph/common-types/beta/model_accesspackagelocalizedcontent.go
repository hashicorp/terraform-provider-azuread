package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageLocalizedContent struct {
	// The fallback string, which is used when a requested localization isn't available. Required.
	DefaultText nullable.Type[string] `json:"defaultText,omitempty"`

	// Content represented in a format for a specific locale.
	LocalizedTexts *[]AccessPackageLocalizedText `json:"localizedTexts,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
