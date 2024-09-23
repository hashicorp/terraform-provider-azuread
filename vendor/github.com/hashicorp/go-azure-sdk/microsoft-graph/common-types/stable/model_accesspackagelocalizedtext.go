package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageLocalizedText struct {
	// The language code that text is in. For example, 'en-us'. The language component follows 2-letter codes as defined in
	// ISO 639-1, and the country component follows 2-letter codes as defined in ISO 3166-1 alpha-2. Required.
	LanguageCode string `json:"languageCode"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The question in the specific language. Required.
	Text nullable.Type[string] `json:"text,omitempty"`
}
