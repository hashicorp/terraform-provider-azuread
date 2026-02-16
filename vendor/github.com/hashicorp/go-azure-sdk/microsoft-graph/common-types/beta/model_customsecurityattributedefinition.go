package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CustomSecurityAttributeDefinition{}

type CustomSecurityAttributeDefinition struct {
	// Values that are predefined for this custom security attribute. This navigation property is not returned by default
	// and must be specified in an $expand query. For example,
	// /directory/customSecurityAttributeDefinitions?$expand=allowedValues.
	AllowedValues *[]AllowedValue `json:"allowedValues,omitempty"`

	// Name of the attribute set. Case insensitive.
	AttributeSet *string `json:"attributeSet,omitempty"`

	// Description of the custom security attribute. Can be up to 128 characters long and include Unicode characters. Can be
	// changed later.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Indicates whether multiple values can be assigned to the custom security attribute. Cannot be changed later. If type
	// is set to Boolean, isCollection cannot be set to true.
	IsCollection *bool `json:"isCollection,omitempty"`

	// Indicates whether custom security attribute values are indexed for searching on objects that are assigned attribute
	// values. Cannot be changed later.
	IsSearchable nullable.Type[bool] `json:"isSearchable,omitempty"`

	// Name of the custom security attribute. Must be unique within an attribute set. Can be up to 32 characters long and
	// include Unicode characters. Cannot contain spaces or special characters. Cannot be changed later. Case insensitive.
	Name *string `json:"name,omitempty"`

	// Specifies whether the custom security attribute is active or deactivated. Acceptable values are: Available and
	// Deprecated. Can be changed later.
	Status *string `json:"status,omitempty"`

	// Data type for the custom security attribute values. Supported types are: Boolean, Integer, and String. Cannot be
	// changed later.
	Type *string `json:"type,omitempty"`

	// Indicates whether only predefined values can be assigned to the custom security attribute. If set to false, free-form
	// values are allowed. Can later be changed from true to false, but cannot be changed from false to true. If type is set
	// to Boolean, usePreDefinedValuesOnly cannot be set to true.
	UsePreDefinedValuesOnly nullable.Type[bool] `json:"usePreDefinedValuesOnly,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CustomSecurityAttributeDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomSecurityAttributeDefinition{}

func (s CustomSecurityAttributeDefinition) MarshalJSON() ([]byte, error) {
	type wrapper CustomSecurityAttributeDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomSecurityAttributeDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomSecurityAttributeDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customSecurityAttributeDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomSecurityAttributeDefinition: %+v", err)
	}

	return encoded, nil
}
