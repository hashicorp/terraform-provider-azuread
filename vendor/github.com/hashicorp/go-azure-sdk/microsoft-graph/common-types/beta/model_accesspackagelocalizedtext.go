package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageLocalizedText struct {
	// The ISO code for the intended language. Required.
	LanguageCode nullable.Type[string] `json:"languageCode,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The text in the specific language. Required.
	Text nullable.Type[string] `json:"text,omitempty"`
}
