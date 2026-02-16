package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UserSet = SingleUser{}

type SingleUser struct {
	// The name of the user in Microsoft Entra ID. Read only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The ID of the user in Microsoft Entra ID.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Fields inherited from UserSet

	// For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
	IsBackup nullable.Type[bool] `json:"isBackup,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SingleUser) UserSet() BaseUserSetImpl {
	return BaseUserSetImpl{
		IsBackup:  s.IsBackup,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SingleUser{}

func (s SingleUser) MarshalJSON() ([]byte, error) {
	type wrapper SingleUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SingleUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SingleUser: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.singleUser"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SingleUser: %+v", err)
	}

	return encoded, nil
}
