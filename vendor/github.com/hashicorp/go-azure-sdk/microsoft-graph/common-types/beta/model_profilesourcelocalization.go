package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileSourceLocalization struct {
	// Localized display name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Language locale.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Localized profile source URL.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}
