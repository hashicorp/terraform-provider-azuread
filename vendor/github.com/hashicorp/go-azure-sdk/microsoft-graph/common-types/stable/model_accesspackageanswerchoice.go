package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAnswerChoice struct {
	// The actual value of the selected choice. This is typically a string value which is understandable by applications.
	// Required.
	ActualValue nullable.Type[string] `json:"actualValue,omitempty"`

	// The text of the answer choice represented in a format for a specific locale.
	Localizations *[]AccessPackageLocalizedText `json:"localizations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The string to display for this answer; if an Accept-Language header is provided, and there is a matching localization
	// in localizations, this string will be the matching localized string; otherwise, this string remains as the default
	// non-localized string. Required.
	Text nullable.Type[string] `json:"text,omitempty"`
}
