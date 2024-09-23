package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ColumnValidation struct {
	// Default BCP 47 language tag for the description.
	DefaultLanguage nullable.Type[string] `json:"defaultLanguage,omitempty"`

	// Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted
	// with this message if validation fails.
	Descriptions *[]DisplayNameLocalization `json:"descriptions,omitempty"`

	// The formula to validate column value. For examples, see Examples of common formulas in lists.
	Formula nullable.Type[string] `json:"formula,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
