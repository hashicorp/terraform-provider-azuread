package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomAuthenticationExtension interface {
	Entity
	CustomCalloutExtension
	CustomAuthenticationExtension() BaseCustomAuthenticationExtensionImpl
}

var _ CustomAuthenticationExtension = BaseCustomAuthenticationExtensionImpl{}

type BaseCustomAuthenticationExtensionImpl struct {
	// The behaviour on error for the custom authentication extension.
	BehaviorOnError CustomExtensionBehaviorOnError `json:"behaviorOnError"`

	// Fields inherited from CustomCalloutExtension

	// Configuration for securing the API call to the logic app. For example, using OAuth client credentials flow.
	AuthenticationConfiguration CustomExtensionAuthenticationConfiguration `json:"authenticationConfiguration"`

	// HTTP connection settings that define how long Microsoft Entra ID can wait for a connection to a logic app, how many
	// times you can retry a timed-out connection and the exception scenarios when retries are allowed.
	ClientConfiguration *CustomExtensionClientConfiguration `json:"clientConfiguration,omitempty"`

	// Description for the customCalloutExtension object.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for the customCalloutExtension object.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The type and details for configuring the endpoint to call the logic app's workflow.
	EndpointConfiguration CustomExtensionEndpointConfiguration `json:"endpointConfiguration"`

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

func (s BaseCustomAuthenticationExtensionImpl) CustomAuthenticationExtension() BaseCustomAuthenticationExtensionImpl {
	return s
}

func (s BaseCustomAuthenticationExtensionImpl) CustomCalloutExtension() BaseCustomCalloutExtensionImpl {
	return BaseCustomCalloutExtensionImpl{
		AuthenticationConfiguration: s.AuthenticationConfiguration,
		ClientConfiguration:         s.ClientConfiguration,
		Description:                 s.Description,
		DisplayName:                 s.DisplayName,
		EndpointConfiguration:       s.EndpointConfiguration,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s BaseCustomAuthenticationExtensionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ CustomAuthenticationExtension = RawCustomAuthenticationExtensionImpl{}

// RawCustomAuthenticationExtensionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomAuthenticationExtensionImpl struct {
	customAuthenticationExtension BaseCustomAuthenticationExtensionImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawCustomAuthenticationExtensionImpl) CustomAuthenticationExtension() BaseCustomAuthenticationExtensionImpl {
	return s.customAuthenticationExtension
}

func (s RawCustomAuthenticationExtensionImpl) CustomCalloutExtension() BaseCustomCalloutExtensionImpl {
	return s.customAuthenticationExtension.CustomCalloutExtension()
}

func (s RawCustomAuthenticationExtensionImpl) Entity() BaseEntityImpl {
	return s.customAuthenticationExtension.Entity()
}

var _ json.Marshaler = BaseCustomAuthenticationExtensionImpl{}

func (s BaseCustomAuthenticationExtensionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseCustomAuthenticationExtensionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseCustomAuthenticationExtensionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseCustomAuthenticationExtensionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customAuthenticationExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseCustomAuthenticationExtensionImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseCustomAuthenticationExtensionImpl{}

func (s *BaseCustomAuthenticationExtensionImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ClientConfiguration *CustomExtensionClientConfiguration `json:"clientConfiguration,omitempty"`
		Description         nullable.Type[string]               `json:"description,omitempty"`
		DisplayName         nullable.Type[string]               `json:"displayName,omitempty"`
		Id                  *string                             `json:"id,omitempty"`
		ODataId             *string                             `json:"@odata.id,omitempty"`
		ODataType           *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ClientConfiguration = decoded.ClientConfiguration
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseCustomAuthenticationExtensionImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authenticationConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionAuthenticationConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthenticationConfiguration' for 'BaseCustomAuthenticationExtensionImpl': %+v", err)
		}
		s.AuthenticationConfiguration = impl
	}

	if v, ok := temp["behaviorOnError"]; ok {
		impl, err := UnmarshalCustomExtensionBehaviorOnErrorImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'BehaviorOnError' for 'BaseCustomAuthenticationExtensionImpl': %+v", err)
		}
		s.BehaviorOnError = impl
	}

	if v, ok := temp["endpointConfiguration"]; ok {
		impl, err := UnmarshalCustomExtensionEndpointConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EndpointConfiguration' for 'BaseCustomAuthenticationExtensionImpl': %+v", err)
		}
		s.EndpointConfiguration = impl
	}

	return nil
}

func UnmarshalCustomAuthenticationExtensionImplementation(input []byte) (CustomAuthenticationExtension, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomAuthenticationExtension into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionStartCustomExtension") {
		var out OnAttributeCollectionStartCustomExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionStartCustomExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionSubmitCustomExtension") {
		var out OnAttributeCollectionSubmitCustomExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionSubmitCustomExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onOtpSendCustomExtension") {
		var out OnOtpSendCustomExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnOtpSendCustomExtension: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onTokenIssuanceStartCustomExtension") {
		var out OnTokenIssuanceStartCustomExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnTokenIssuanceStartCustomExtension: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomAuthenticationExtensionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomAuthenticationExtensionImpl: %+v", err)
	}

	return RawCustomAuthenticationExtensionImpl{
		customAuthenticationExtension: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
