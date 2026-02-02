package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityUserFlowAttribute = IdentityCustomUserFlowAttribute{}

type IdentityCustomUserFlowAttribute struct {

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

func (s IdentityCustomUserFlowAttribute) IdentityUserFlowAttribute() BaseIdentityUserFlowAttributeImpl {
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

func (s IdentityCustomUserFlowAttribute) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityCustomUserFlowAttribute{}

func (s IdentityCustomUserFlowAttribute) MarshalJSON() ([]byte, error) {
	type wrapper IdentityCustomUserFlowAttribute
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityCustomUserFlowAttribute: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityCustomUserFlowAttribute: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityCustomUserFlowAttribute"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityCustomUserFlowAttribute: %+v", err)
	}

	return encoded, nil
}
