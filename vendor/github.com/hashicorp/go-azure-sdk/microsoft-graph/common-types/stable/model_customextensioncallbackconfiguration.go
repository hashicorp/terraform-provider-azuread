package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCallbackConfiguration interface {
	CustomExtensionCallbackConfiguration() BaseCustomExtensionCallbackConfigurationImpl
}

var _ CustomExtensionCallbackConfiguration = BaseCustomExtensionCallbackConfigurationImpl{}

type BaseCustomExtensionCallbackConfigurationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The maximum duration in ISO 8601 format that Microsoft Entra ID will wait for a resume action for the callout it sent
	// to the logic app. The valid range for custom extensions in lifecycle workflows is five minutes to three hours. The
	// valid range for custom extensions in entitlement management is between 5 minutes and 14 days. For example, PT3H
	// refers to three hours, P3D refers to three days, PT10M refers to ten minutes.
	TimeoutDuration nullable.Type[string] `json:"timeoutDuration,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomExtensionCallbackConfigurationImpl) CustomExtensionCallbackConfiguration() BaseCustomExtensionCallbackConfigurationImpl {
	return s
}

var _ CustomExtensionCallbackConfiguration = RawCustomExtensionCallbackConfigurationImpl{}

// RawCustomExtensionCallbackConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomExtensionCallbackConfigurationImpl struct {
	customExtensionCallbackConfiguration BaseCustomExtensionCallbackConfigurationImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawCustomExtensionCallbackConfigurationImpl) CustomExtensionCallbackConfiguration() BaseCustomExtensionCallbackConfigurationImpl {
	return s.customExtensionCallbackConfiguration
}

func UnmarshalCustomExtensionCallbackConfigurationImplementation(input []byte) (CustomExtensionCallbackConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomExtensionCallbackConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.customTaskExtensionCallbackConfiguration") {
		var out IdentityGovernanceCustomTaskExtensionCallbackConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceCustomTaskExtensionCallbackConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomExtensionCallbackConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomExtensionCallbackConfigurationImpl: %+v", err)
	}

	return RawCustomExtensionCallbackConfigurationImpl{
		customExtensionCallbackConfiguration: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
