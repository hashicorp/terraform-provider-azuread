package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingTemplateValue struct {
	// Default value for the setting.
	DefaultValue nullable.Type[string] `json:"defaultValue,omitempty"`

	// Description of the setting.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the setting.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of the setting.
	Type nullable.Type[string] `json:"type,omitempty"`
}
