package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileCardAnnotation struct {
	// If present, the value of this field is used by the profile card as the default property label in the experience (for
	// example, 'Cost Center').
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Each resource in this collection represents the localized value of the attribute name for a given language, used as
	// the default label for that locale. For example, a user with a nb-NO client gets 'Kostnadssenter' as the attribute
	// label, rather than 'Cost Center.'
	Localizations *[]DisplayNameLocalization `json:"localizations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
