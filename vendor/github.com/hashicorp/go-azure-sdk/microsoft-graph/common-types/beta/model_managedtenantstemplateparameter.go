package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTemplateParameter struct {
	// The description for the template parameter. Optional. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the template parameter. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
	JsonAllowedValues nullable.Type[string] `json:"jsonAllowedValues,omitempty"`

	// The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
	JsonDefaultValue nullable.Type[string] `json:"jsonDefaultValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ValueType *ManagedTenantsManagementParameterValueType `json:"valueType,omitempty"`
}

var _ json.Marshaler = ManagedTenantsTemplateParameter{}

func (s ManagedTenantsTemplateParameter) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsTemplateParameter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsTemplateParameter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsTemplateParameter: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "jsonAllowedValues")
	delete(decoded, "jsonDefaultValue")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsTemplateParameter: %+v", err)
	}

	return encoded, nil
}
