package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemIdentitySource interface {
	AuthorizationSystemIdentitySource() BaseAuthorizationSystemIdentitySourceImpl
}

var _ AuthorizationSystemIdentitySource = BaseAuthorizationSystemIdentitySourceImpl{}

type BaseAuthorizationSystemIdentitySourceImpl struct {
	// Type of identity provider. Read-only.
	IdentityProviderType nullable.Type[string] `json:"identityProviderType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAuthorizationSystemIdentitySourceImpl) AuthorizationSystemIdentitySource() BaseAuthorizationSystemIdentitySourceImpl {
	return s
}

var _ AuthorizationSystemIdentitySource = RawAuthorizationSystemIdentitySourceImpl{}

// RawAuthorizationSystemIdentitySourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthorizationSystemIdentitySourceImpl struct {
	authorizationSystemIdentitySource BaseAuthorizationSystemIdentitySourceImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawAuthorizationSystemIdentitySourceImpl) AuthorizationSystemIdentitySource() BaseAuthorizationSystemIdentitySourceImpl {
	return s.authorizationSystemIdentitySource
}

var _ json.Marshaler = BaseAuthorizationSystemIdentitySourceImpl{}

func (s BaseAuthorizationSystemIdentitySourceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthorizationSystemIdentitySourceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthorizationSystemIdentitySourceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthorizationSystemIdentitySourceImpl: %+v", err)
	}

	delete(decoded, "identityProviderType")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthorizationSystemIdentitySourceImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthorizationSystemIdentitySourceImplementation(input []byte) (AuthorizationSystemIdentitySource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthorizationSystemIdentitySource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.aadSource") {
		var out AadSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AadSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsSource") {
		var out AwsSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureSource") {
		var out AzureSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gsuiteSource") {
		var out GsuiteSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GsuiteSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unknownSource") {
		var out UnknownSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnknownSource: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthorizationSystemIdentitySourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthorizationSystemIdentitySourceImpl: %+v", err)
	}

	return RawAuthorizationSystemIdentitySourceImpl{
		authorizationSystemIdentitySource: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
