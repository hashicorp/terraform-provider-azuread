package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SingleSignOnExtension interface {
	SingleSignOnExtension() BaseSingleSignOnExtensionImpl
}

var _ SingleSignOnExtension = BaseSingleSignOnExtensionImpl{}

type BaseSingleSignOnExtensionImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSingleSignOnExtensionImpl) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return s
}

var _ SingleSignOnExtension = RawSingleSignOnExtensionImpl{}

// RawSingleSignOnExtensionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSingleSignOnExtensionImpl struct {
	singleSignOnExtension BaseSingleSignOnExtensionImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawSingleSignOnExtensionImpl) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return s.singleSignOnExtension
}

func UnmarshalSingleSignOnExtensionImplementation(input []byte) (SingleSignOnExtension, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SingleSignOnExtension into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.credentialSingleSignOnExtension") {
		var out CredentialSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CredentialSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosSingleSignOnExtension") {
		var out IosSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.kerberosSingleSignOnExtension") {
		var out KerberosSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KerberosSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSSingleSignOnExtension") {
		var out MacOSSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.redirectSingleSignOnExtension") {
		var out RedirectSingleSignOnExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RedirectSingleSignOnExtension: %+v", err)
		}
		return out, nil
	}

	var parent BaseSingleSignOnExtensionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSingleSignOnExtensionImpl: %+v", err)
	}

	return RawSingleSignOnExtensionImpl{
		singleSignOnExtension: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
