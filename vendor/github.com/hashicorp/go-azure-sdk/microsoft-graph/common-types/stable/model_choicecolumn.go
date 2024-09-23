package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChoiceColumn struct {
	// If true, allows custom values that aren't in the configured choices.
	AllowTextEntry nullable.Type[bool] `json:"allowTextEntry,omitempty"`

	// The list of values available for this column.
	Choices *[]string `json:"choices,omitempty"`

	// How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons
	DisplayAs nullable.Type[string] `json:"displayAs,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
