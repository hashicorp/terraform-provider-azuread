package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSingleSignOnExtension interface {
	SingleSignOnExtension
	MacOSSingleSignOnExtension() BaseMacOSSingleSignOnExtensionImpl
}

var _ MacOSSingleSignOnExtension = BaseMacOSSingleSignOnExtensionImpl{}

type BaseMacOSSingleSignOnExtensionImpl struct {

	// Fields inherited from SingleSignOnExtension

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseMacOSSingleSignOnExtensionImpl) MacOSSingleSignOnExtension() BaseMacOSSingleSignOnExtensionImpl {
	return s
}

func (s BaseMacOSSingleSignOnExtensionImpl) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return BaseSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MacOSSingleSignOnExtension = RawMacOSSingleSignOnExtensionImpl{}

// RawMacOSSingleSignOnExtensionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMacOSSingleSignOnExtensionImpl struct {
	macOSSingleSignOnExtension BaseMacOSSingleSignOnExtensionImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawMacOSSingleSignOnExtensionImpl) MacOSSingleSignOnExtension() BaseMacOSSingleSignOnExtensionImpl {
	return s.macOSSingleSignOnExtension
}

func (s RawMacOSSingleSignOnExtensionImpl) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return s.macOSSingleSignOnExtension.SingleSignOnExtension()
}

var _ json.Marshaler = BaseMacOSSingleSignOnExtensionImpl{}

func (s BaseMacOSSingleSignOnExtensionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMacOSSingleSignOnExtensionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMacOSSingleSignOnExtensionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMacOSSingleSignOnExtensionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSSingleSignOnExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMacOSSingleSignOnExtensionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalMacOSSingleSignOnExtensionImplementation(input []byte) (MacOSSingleSignOnExtension, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSSingleSignOnExtension into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSAzureAdSingleSignOnExtension") {
		var out MacOSAzureAdSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSAzureAdSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSCredentialSingleSignOnExtension") {
		var out MacOSCredentialSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSCredentialSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSKerberosSingleSignOnExtension") {
		var out MacOSKerberosSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSKerberosSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSRedirectSingleSignOnExtension") {
		var out MacOSRedirectSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSRedirectSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	var parent BaseMacOSSingleSignOnExtensionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMacOSSingleSignOnExtensionImpl: %+v", err)
	}

	return RawMacOSSingleSignOnExtensionImpl{
		macOSSingleSignOnExtension: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
