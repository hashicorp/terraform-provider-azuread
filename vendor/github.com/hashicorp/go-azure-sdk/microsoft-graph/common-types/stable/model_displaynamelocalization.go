package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisplayNameLocalization struct {
	// If present, the value of this field contains the displayName string that has been set for the language present in the
	// languageTag field.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Provides the language culture-code and friendly name of the language that the displayName field has been provided in.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
