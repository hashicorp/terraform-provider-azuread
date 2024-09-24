package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AttributeSet{}

type AttributeSet struct {
	// Description of the attribute set. Can be up to 128 characters long and include Unicode characters. Can be changed
	// later.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Maximum number of custom security attributes that can be defined in this attribute set. Default value is null. If not
	// specified, the administrator can add up to the maximum of 500 active attributes per tenant. Can be changed later.
	MaxAttributesPerSet nullable.Type[int64] `json:"maxAttributesPerSet,omitempty"`

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

func (s AttributeSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AttributeSet{}

func (s AttributeSet) MarshalJSON() ([]byte, error) {
	type wrapper AttributeSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AttributeSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AttributeSet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attributeSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AttributeSet: %+v", err)
	}

	return encoded, nil
}
