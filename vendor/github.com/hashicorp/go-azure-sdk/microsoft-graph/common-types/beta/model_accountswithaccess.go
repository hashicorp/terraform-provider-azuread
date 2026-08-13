package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountsWithAccess interface {
	AccountsWithAccess() BaseAccountsWithAccessImpl
}

var _ AccountsWithAccess = BaseAccountsWithAccessImpl{}

type BaseAccountsWithAccessImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccountsWithAccessImpl) AccountsWithAccess() BaseAccountsWithAccessImpl {
	return s
}

var _ AccountsWithAccess = RawAccountsWithAccessImpl{}

// RawAccountsWithAccessImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAccountsWithAccessImpl struct {
	accountsWithAccess BaseAccountsWithAccessImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawAccountsWithAccessImpl) AccountsWithAccess() BaseAccountsWithAccessImpl {
	return s.accountsWithAccess
}

func (s RawAccountsWithAccessImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAccountsWithAccessImplementation(input []byte) (AccountsWithAccess, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccountsWithAccess into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.allAccountsWithAccess") {
		var out AllAccountsWithAccess
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllAccountsWithAccess: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enumeratedAccountsWithAccess") {
		var out EnumeratedAccountsWithAccess
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnumeratedAccountsWithAccess: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccountsWithAccessImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccountsWithAccessImpl: %+v", err)
	}

	return RawAccountsWithAccessImpl{
		accountsWithAccess: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
