package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeZoneBase interface {
	TimeZoneBase() BaseTimeZoneBaseImpl
}

var _ TimeZoneBase = BaseTimeZoneBaseImpl{}

type BaseTimeZoneBaseImpl struct {
	// The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized
	// Time Zone' for a custom time zone.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseTimeZoneBaseImpl) TimeZoneBase() BaseTimeZoneBaseImpl {
	return s
}

var _ TimeZoneBase = RawTimeZoneBaseImpl{}

// RawTimeZoneBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTimeZoneBaseImpl struct {
	timeZoneBase BaseTimeZoneBaseImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawTimeZoneBaseImpl) TimeZoneBase() BaseTimeZoneBaseImpl {
	return s.timeZoneBase
}

func UnmarshalTimeZoneBaseImplementation(input []byte) (TimeZoneBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TimeZoneBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.customTimeZone") {
		var out CustomTimeZone
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomTimeZone: %+v", err)
		}
		return out, nil
	}

	var parent BaseTimeZoneBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTimeZoneBaseImpl: %+v", err)
	}

	return RawTimeZoneBaseImpl{
		timeZoneBase: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
