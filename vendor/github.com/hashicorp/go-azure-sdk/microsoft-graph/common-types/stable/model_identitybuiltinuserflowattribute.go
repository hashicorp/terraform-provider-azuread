package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityUserFlowAttribute = IdentityBuiltInUserFlowAttribute{}

type IdentityBuiltInUserFlowAttribute struct {

	// Fields inherited from IdentityUserFlowAttribute

	DataType *IdentityUserFlowAttributeDataType `json:"dataType,omitempty"`

	// The description of the user flow attribute that's shown to the user at the time of sign up.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the user flow attribute. Supports $filter (eq, ne).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	UserFlowAttributeType *IdentityUserFlowAttributeType `json:"userFlowAttributeType,omitempty"`

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

func (s IdentityBuiltInUserFlowAttribute) IdentityUserFlowAttribute() BaseIdentityUserFlowAttributeImpl {
	return BaseIdentityUserFlowAttributeImpl{
		DataType:              s.DataType,
		Description:           s.Description,
		DisplayName:           s.DisplayName,
		UserFlowAttributeType: s.UserFlowAttributeType,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s IdentityBuiltInUserFlowAttribute) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityBuiltInUserFlowAttribute{}

func (s IdentityBuiltInUserFlowAttribute) MarshalJSON() ([]byte, error) {
	type wrapper IdentityBuiltInUserFlowAttribute
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityBuiltInUserFlowAttribute: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityBuiltInUserFlowAttribute: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityBuiltInUserFlowAttribute"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityBuiltInUserFlowAttribute: %+v", err)
	}

	return encoded, nil
}
