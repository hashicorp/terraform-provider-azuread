package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionBehaviorOnError interface {
	CustomExtensionBehaviorOnError() BaseCustomExtensionBehaviorOnErrorImpl
}

var _ CustomExtensionBehaviorOnError = BaseCustomExtensionBehaviorOnErrorImpl{}

type BaseCustomExtensionBehaviorOnErrorImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomExtensionBehaviorOnErrorImpl) CustomExtensionBehaviorOnError() BaseCustomExtensionBehaviorOnErrorImpl {
	return s
}

var _ CustomExtensionBehaviorOnError = RawCustomExtensionBehaviorOnErrorImpl{}

// RawCustomExtensionBehaviorOnErrorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomExtensionBehaviorOnErrorImpl struct {
	customExtensionBehaviorOnError BaseCustomExtensionBehaviorOnErrorImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawCustomExtensionBehaviorOnErrorImpl) CustomExtensionBehaviorOnError() BaseCustomExtensionBehaviorOnErrorImpl {
	return s.customExtensionBehaviorOnError
}

func UnmarshalCustomExtensionBehaviorOnErrorImplementation(input []byte) (CustomExtensionBehaviorOnError, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomExtensionBehaviorOnError into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.fallbackToMicrosoftProviderOnError") {
		var out FallbackToMicrosoftProviderOnError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FallbackToMicrosoftProviderOnError: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomExtensionBehaviorOnErrorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomExtensionBehaviorOnErrorImpl: %+v", err)
	}

	return RawCustomExtensionBehaviorOnErrorImpl{
		customExtensionBehaviorOnError: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
