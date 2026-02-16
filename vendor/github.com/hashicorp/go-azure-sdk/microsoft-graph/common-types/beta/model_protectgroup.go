package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LabelActionBase = ProtectGroup{}

type ProtectGroup struct {
	AllowEmailFromGuestUsers nullable.Type[bool] `json:"allowEmailFromGuestUsers,omitempty"`
	AllowGuestUsers          nullable.Type[bool] `json:"allowGuestUsers,omitempty"`
	Privacy                  *GroupPrivacy       `json:"privacy,omitempty"`

	// Fields inherited from LabelActionBase

	// The name of the action (for example, 'Encrypt', 'AddHeader').
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ProtectGroup) LabelActionBase() BaseLabelActionBaseImpl {
	return BaseLabelActionBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ProtectGroup{}

func (s ProtectGroup) MarshalJSON() ([]byte, error) {
	type wrapper ProtectGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProtectGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectGroup: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.protectGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProtectGroup: %+v", err)
	}

	return encoded, nil
}
