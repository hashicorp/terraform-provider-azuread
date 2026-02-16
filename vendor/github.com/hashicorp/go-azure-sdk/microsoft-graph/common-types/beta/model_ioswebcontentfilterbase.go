package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosWebContentFilterBase interface {
	IosWebContentFilterBase() BaseIosWebContentFilterBaseImpl
}

var _ IosWebContentFilterBase = BaseIosWebContentFilterBaseImpl{}

type BaseIosWebContentFilterBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIosWebContentFilterBaseImpl) IosWebContentFilterBase() BaseIosWebContentFilterBaseImpl {
	return s
}

var _ IosWebContentFilterBase = RawIosWebContentFilterBaseImpl{}

// RawIosWebContentFilterBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIosWebContentFilterBaseImpl struct {
	iosWebContentFilterBase BaseIosWebContentFilterBaseImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawIosWebContentFilterBaseImpl) IosWebContentFilterBase() BaseIosWebContentFilterBaseImpl {
	return s.iosWebContentFilterBase
}

func UnmarshalIosWebContentFilterBaseImplementation(input []byte) (IosWebContentFilterBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IosWebContentFilterBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosWebContentFilterAutoFilter") {
		var out IosWebContentFilterAutoFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosWebContentFilterAutoFilter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess") {
		var out IosWebContentFilterSpecificWebsitesAccess
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosWebContentFilterSpecificWebsitesAccess: %+v", err)
		}
		return out, nil
	}

	var parent BaseIosWebContentFilterBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIosWebContentFilterBaseImpl: %+v", err)
	}

	return RawIosWebContentFilterBaseImpl{
		iosWebContentFilterBase: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
