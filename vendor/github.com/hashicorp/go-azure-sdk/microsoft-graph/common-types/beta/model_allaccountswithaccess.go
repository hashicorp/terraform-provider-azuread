package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccountsWithAccess = AllAccountsWithAccess{}

type AllAccountsWithAccess struct {

	// Fields inherited from AccountsWithAccess

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AllAccountsWithAccess) AccountsWithAccess() BaseAccountsWithAccessImpl {
	return BaseAccountsWithAccessImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AllAccountsWithAccess{}

func (s AllAccountsWithAccess) MarshalJSON() ([]byte, error) {
	type wrapper AllAccountsWithAccess
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AllAccountsWithAccess: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AllAccountsWithAccess: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.allAccountsWithAccess"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AllAccountsWithAccess: %+v", err)
	}

	return encoded, nil
}
