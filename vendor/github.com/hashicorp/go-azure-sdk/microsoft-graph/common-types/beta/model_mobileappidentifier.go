package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIdentifier interface {
	MobileAppIdentifier() BaseMobileAppIdentifierImpl
}

var _ MobileAppIdentifier = BaseMobileAppIdentifierImpl{}

type BaseMobileAppIdentifierImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseMobileAppIdentifierImpl) MobileAppIdentifier() BaseMobileAppIdentifierImpl {
	return s
}

var _ MobileAppIdentifier = RawMobileAppIdentifierImpl{}

// RawMobileAppIdentifierImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawMobileAppIdentifierImpl struct {
	mobileAppIdentifier BaseMobileAppIdentifierImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawMobileAppIdentifierImpl) MobileAppIdentifier() BaseMobileAppIdentifierImpl {
	return s.mobileAppIdentifier
}

func (s RawMobileAppIdentifierImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalMobileAppIdentifierImplementation(input []byte) (MobileAppIdentifier, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppIdentifier into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidMobileAppIdentifier") {
		var out AndroidMobileAppIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidMobileAppIdentifier: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosMobileAppIdentifier") {
		var out IosMobileAppIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosMobileAppIdentifier: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macAppIdentifier") {
		var out MacAppIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacAppIdentifier: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAppIdentifier") {
		var out WindowsAppIdentifier
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAppIdentifier: %+v", err)
		}
		return out, nil
	}

	var parent BaseMobileAppIdentifierImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMobileAppIdentifierImpl: %+v", err)
	}

	return RawMobileAppIdentifierImpl{
		mobileAppIdentifier: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
