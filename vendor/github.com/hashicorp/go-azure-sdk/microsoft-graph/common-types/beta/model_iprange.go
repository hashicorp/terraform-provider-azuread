package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPRange interface {
	IPRange() BaseIPRangeImpl
}

var _ IPRange = BaseIPRangeImpl{}

type BaseIPRangeImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIPRangeImpl) IPRange() BaseIPRangeImpl {
	return s
}

var _ IPRange = RawIPRangeImpl{}

// RawIPRangeImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawIPRangeImpl struct {
	iPRange BaseIPRangeImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawIPRangeImpl) IPRange() BaseIPRangeImpl {
	return s.iPRange
}

func (s RawIPRangeImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalIPRangeImplementation(input []byte) (IPRange, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IPRange into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iPv4CidrRange") {
		var out IPv4CIDRRange
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPv4CIDRRange: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iPv4Range") {
		var out IPv4Range
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPv4Range: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iPv6CidrRange") {
		var out IPv6CIDRRange
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPv6CIDRRange: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iPv6Range") {
		var out IPv6Range
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPv6Range: %+v", err)
		}
		return out, nil
	}

	var parent BaseIPRangeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIPRangeImpl: %+v", err)
	}

	return RawIPRangeImpl{
		iPRange: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
