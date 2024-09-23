package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileContainedApp interface {
	Entity
	MobileContainedApp() BaseMobileContainedAppImpl
}

var _ MobileContainedApp = BaseMobileContainedAppImpl{}

type BaseMobileContainedAppImpl struct {

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseMobileContainedAppImpl) MobileContainedApp() BaseMobileContainedAppImpl {
	return s
}

func (s BaseMobileContainedAppImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MobileContainedApp = RawMobileContainedAppImpl{}

// RawMobileContainedAppImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMobileContainedAppImpl struct {
	mobileContainedApp BaseMobileContainedAppImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawMobileContainedAppImpl) MobileContainedApp() BaseMobileContainedAppImpl {
	return s.mobileContainedApp
}

func (s RawMobileContainedAppImpl) Entity() BaseEntityImpl {
	return s.mobileContainedApp.Entity()
}

var _ json.Marshaler = BaseMobileContainedAppImpl{}

func (s BaseMobileContainedAppImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMobileContainedAppImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMobileContainedAppImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMobileContainedAppImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileContainedApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMobileContainedAppImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalMobileContainedAppImplementation(input []byte) (MobileContainedApp, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileContainedApp into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftStoreForBusinessContainedApp") {
		var out MicrosoftStoreForBusinessContainedApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftStoreForBusinessContainedApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUniversalAppXContainedApp") {
		var out WindowsUniversalAppXContainedApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUniversalAppXContainedApp: %+v", err)
		}
		return out, nil
	}

	var parent BaseMobileContainedAppImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMobileContainedAppImpl: %+v", err)
	}

	return RawMobileContainedAppImpl{
		mobileContainedApp: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
