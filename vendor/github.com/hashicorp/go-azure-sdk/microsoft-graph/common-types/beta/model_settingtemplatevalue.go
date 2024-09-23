package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingTemplateValue struct {
	// Default value for the setting. Read-only.
	DefaultValue nullable.Type[string] `json:"defaultValue,omitempty"`

	// Description of the setting. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the setting. Read-only.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of the setting. Read-only.
	Type nullable.Type[string] `json:"type,omitempty"`
}

var _ json.Marshaler = SettingTemplateValue{}

func (s SettingTemplateValue) MarshalJSON() ([]byte, error) {
	type wrapper SettingTemplateValue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SettingTemplateValue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SettingTemplateValue: %+v", err)
	}

	delete(decoded, "defaultValue")
	delete(decoded, "description")
	delete(decoded, "name")
	delete(decoded, "type")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SettingTemplateValue: %+v", err)
	}

	return encoded, nil
}
