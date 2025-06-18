package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionInputConfiguration struct {
	// The built-in or custom attribute for which a value is being collected.
	Attribute *string `json:"attribute,omitempty"`

	// The default value of the attribute displayed to the end user. The capability to set the default value isn't available
	// through the Microsoft Entra admin center.
	DefaultValue nullable.Type[string] `json:"defaultValue,omitempty"`

	// Defines whether the attribute is editable by the end user.
	Editable nullable.Type[bool] `json:"editable,omitempty"`

	// Defines whether the attribute is displayed to the end user. The capability to hide isn't available through the
	// Microsoft Entra admin center.
	Hidden nullable.Type[bool] `json:"hidden,omitempty"`

	InputType *AuthenticationAttributeCollectionInputType `json:"inputType,omitempty"`

	// The label of the attribute field that's displayed to end user, unless overridden.
	Label *string `json:"label,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The option values for certain multiple-option input types.
	Options *[]AuthenticationAttributeCollectionOptionConfiguration `json:"options,omitempty"`

	// Defines whether the field is required.
	Required nullable.Type[bool] `json:"required,omitempty"`

	// The regex for the value of the field. For more information about the supported regexes, see validationRegEx values
	// for inputType objects. To understand how to specify regexes, see the Regular expressions cheat sheet.
	ValidationRegEx nullable.Type[string] `json:"validationRegEx,omitempty"`

	// Defines whether Microsoft Entra ID stores the value that it collects.
	WriteToDirectory nullable.Type[bool] `json:"writeToDirectory,omitempty"`
}
