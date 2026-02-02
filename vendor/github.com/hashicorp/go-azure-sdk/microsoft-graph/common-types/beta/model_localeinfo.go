package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocaleInfo struct {
	// A name representing the user's locale in natural language, for example, 'English (United States)'.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A locale representation for the user, which includes the user's preferred language and country/region. For example,
	// 'en-us'. The language component follows 2-letter codes as defined in ISO 639-1, and the country component follows
	// 2-letter codes as defined in ISO 3166-1 alpha-2.
	Locale nullable.Type[string] `json:"locale,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
