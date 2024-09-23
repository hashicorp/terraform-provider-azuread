package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccountsWithAccess = EnumeratedAccountsWithAccess{}

type EnumeratedAccountsWithAccess struct {
	Accounts *[]AuthorizationSystem `json:"accounts,omitempty"`

	// Fields inherited from AccountsWithAccess

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EnumeratedAccountsWithAccess) AccountsWithAccess() BaseAccountsWithAccessImpl {
	return BaseAccountsWithAccessImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EnumeratedAccountsWithAccess{}

func (s EnumeratedAccountsWithAccess) MarshalJSON() ([]byte, error) {
	type wrapper EnumeratedAccountsWithAccess
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnumeratedAccountsWithAccess: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnumeratedAccountsWithAccess: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enumeratedAccountsWithAccess"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnumeratedAccountsWithAccess: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EnumeratedAccountsWithAccess{}

func (s *EnumeratedAccountsWithAccess) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EnumeratedAccountsWithAccess into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["accounts"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Accounts into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthorizationSystem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthorizationSystemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Accounts' for 'EnumeratedAccountsWithAccess': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Accounts = &output
	}

	return nil
}
