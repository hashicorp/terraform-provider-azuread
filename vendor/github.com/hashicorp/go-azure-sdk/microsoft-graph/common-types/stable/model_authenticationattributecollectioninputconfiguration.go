package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionInputConfiguration struct {
	// The built-in or custom attribute for which a value is being collected.
	Attribute *string `json:"attribute,omitempty"`

	// The default value of the attribute displayed to the end user.
	DefaultValue nullable.Type[string] `json:"defaultValue,omitempty"`

	// Whether the attribute is editable by the end user.
	Editable nullable.Type[bool] `json:"editable,omitempty"`

	// Whether the attribute is displayed to the end user.
	Hidden nullable.Type[bool] `json:"hidden,omitempty"`

	InputType *AuthenticationAttributeCollectionInputType `json:"inputType,omitempty"`

	// The label of the attribute field that is displayed to end user, unless overridden.
	Label *string `json:"label,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The option values for certain multiple-option input types.
	Options *[]AuthenticationAttributeCollectionOptionConfiguration `json:"options,omitempty"`

	// Whether the field is required.
	Required nullable.Type[bool] `json:"required,omitempty"`

	// The regex for the value of the field.
	ValidationRegEx nullable.Type[string] `json:"validationRegEx,omitempty"`

	// Whether the value collected is stored.
	WriteToDirectory nullable.Type[bool] `json:"writeToDirectory,omitempty"`
}
