package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGlobalProxy interface {
	AndroidDeviceOwnerGlobalProxy() BaseAndroidDeviceOwnerGlobalProxyImpl
}

var _ AndroidDeviceOwnerGlobalProxy = BaseAndroidDeviceOwnerGlobalProxyImpl{}

type BaseAndroidDeviceOwnerGlobalProxyImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAndroidDeviceOwnerGlobalProxyImpl) AndroidDeviceOwnerGlobalProxy() BaseAndroidDeviceOwnerGlobalProxyImpl {
	return s
}

var _ AndroidDeviceOwnerGlobalProxy = RawAndroidDeviceOwnerGlobalProxyImpl{}

// RawAndroidDeviceOwnerGlobalProxyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAndroidDeviceOwnerGlobalProxyImpl struct {
	androidDeviceOwnerGlobalProxy BaseAndroidDeviceOwnerGlobalProxyImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawAndroidDeviceOwnerGlobalProxyImpl) AndroidDeviceOwnerGlobalProxy() BaseAndroidDeviceOwnerGlobalProxyImpl {
	return s.androidDeviceOwnerGlobalProxy
}

func UnmarshalAndroidDeviceOwnerGlobalProxyImplementation(input []byte) (AndroidDeviceOwnerGlobalProxy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerGlobalProxy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerGlobalProxyAutoConfig") {
		var out AndroidDeviceOwnerGlobalProxyAutoConfig
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerGlobalProxyAutoConfig: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerGlobalProxyDirect") {
		var out AndroidDeviceOwnerGlobalProxyDirect
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerGlobalProxyDirect: %+v", err)
		}
		return out, nil
	}

	var parent BaseAndroidDeviceOwnerGlobalProxyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAndroidDeviceOwnerGlobalProxyImpl: %+v", err)
	}

	return RawAndroidDeviceOwnerGlobalProxyImpl{
		androidDeviceOwnerGlobalProxy: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
