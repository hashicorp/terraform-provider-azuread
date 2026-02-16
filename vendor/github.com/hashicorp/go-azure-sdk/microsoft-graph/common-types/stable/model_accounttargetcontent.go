package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountTargetContent interface {
	AccountTargetContent() BaseAccountTargetContentImpl
}

var _ AccountTargetContent = BaseAccountTargetContentImpl{}

type BaseAccountTargetContentImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of account target content. Possible values are: unknown, includeAll, addressBook, unknownFutureValue.
	Type *AccountTargetContentType `json:"type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccountTargetContentImpl) AccountTargetContent() BaseAccountTargetContentImpl {
	return s
}

var _ AccountTargetContent = RawAccountTargetContentImpl{}

// RawAccountTargetContentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccountTargetContentImpl struct {
	accountTargetContent BaseAccountTargetContentImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawAccountTargetContentImpl) AccountTargetContent() BaseAccountTargetContentImpl {
	return s.accountTargetContent
}

func UnmarshalAccountTargetContentImplementation(input []byte) (AccountTargetContent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccountTargetContent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.addressBookAccountTargetContent") {
		var out AddressBookAccountTargetContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddressBookAccountTargetContent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.includeAllAccountTargetContent") {
		var out IncludeAllAccountTargetContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IncludeAllAccountTargetContent: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccountTargetContentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccountTargetContentImpl: %+v", err)
	}

	return RawAccountTargetContentImpl{
		accountTargetContent: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
