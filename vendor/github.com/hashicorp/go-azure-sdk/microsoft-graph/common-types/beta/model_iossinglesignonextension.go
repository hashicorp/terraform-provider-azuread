package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosSingleSignOnExtension interface {
	SingleSignOnExtension
	IosSingleSignOnExtension() BaseIosSingleSignOnExtensionImpl
}

var _ IosSingleSignOnExtension = BaseIosSingleSignOnExtensionImpl{}

type BaseIosSingleSignOnExtensionImpl struct {

	// Fields inherited from SingleSignOnExtension

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIosSingleSignOnExtensionImpl) IosSingleSignOnExtension() BaseIosSingleSignOnExtensionImpl {
	return s
}

func (s BaseIosSingleSignOnExtensionImpl) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return BaseSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IosSingleSignOnExtension = RawIosSingleSignOnExtensionImpl{}

// RawIosSingleSignOnExtensionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIosSingleSignOnExtensionImpl struct {
	iosSingleSignOnExtension BaseIosSingleSignOnExtensionImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawIosSingleSignOnExtensionImpl) IosSingleSignOnExtension() BaseIosSingleSignOnExtensionImpl {
	return s.iosSingleSignOnExtension
}

func (s RawIosSingleSignOnExtensionImpl) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return s.iosSingleSignOnExtension.SingleSignOnExtension()
}

var _ json.Marshaler = BaseIosSingleSignOnExtensionImpl{}

func (s BaseIosSingleSignOnExtensionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIosSingleSignOnExtensionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIosSingleSignOnExtensionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIosSingleSignOnExtensionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosSingleSignOnExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIosSingleSignOnExtensionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIosSingleSignOnExtensionImplementation(input []byte) (IosSingleSignOnExtension, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IosSingleSignOnExtension into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosAzureAdSingleSignOnExtension") {
		var out IosAzureAdSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosAzureAdSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosCredentialSingleSignOnExtension") {
		var out IosCredentialSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosCredentialSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosKerberosSingleSignOnExtension") {
		var out IosKerberosSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosKerberosSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosRedirectSingleSignOnExtension") {
		var out IosRedirectSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosRedirectSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	var parent BaseIosSingleSignOnExtensionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIosSingleSignOnExtensionImpl: %+v", err)
	}

	return RawIosSingleSignOnExtensionImpl{
		iosSingleSignOnExtension: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
