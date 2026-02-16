package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceNowAuthenticationMethod interface {
	ServiceNowAuthenticationMethod() BaseServiceNowAuthenticationMethodImpl
}

var _ ServiceNowAuthenticationMethod = BaseServiceNowAuthenticationMethodImpl{}

type BaseServiceNowAuthenticationMethodImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseServiceNowAuthenticationMethodImpl) ServiceNowAuthenticationMethod() BaseServiceNowAuthenticationMethodImpl {
	return s
}

var _ ServiceNowAuthenticationMethod = RawServiceNowAuthenticationMethodImpl{}

// RawServiceNowAuthenticationMethodImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceNowAuthenticationMethodImpl struct {
	serviceNowAuthenticationMethod BaseServiceNowAuthenticationMethodImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawServiceNowAuthenticationMethodImpl) ServiceNowAuthenticationMethod() BaseServiceNowAuthenticationMethodImpl {
	return s.serviceNowAuthenticationMethod
}

func UnmarshalServiceNowAuthenticationMethodImplementation(input []byte) (ServiceNowAuthenticationMethod, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceNowAuthenticationMethod into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceNowOauthSecretAuthentication") {
		var out ServiceNowOauthSecretAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceNowOauthSecretAuthentication: %+v", err)
		}
		return out, nil
	}

	var parent BaseServiceNowAuthenticationMethodImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServiceNowAuthenticationMethodImpl: %+v", err)
	}

	return RawServiceNowAuthenticationMethodImpl{
		serviceNowAuthenticationMethod: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
