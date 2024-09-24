package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationConnectionSettings interface {
	EducationSynchronizationConnectionSettings() BaseEducationSynchronizationConnectionSettingsImpl
}

var _ EducationSynchronizationConnectionSettings = BaseEducationSynchronizationConnectionSettingsImpl{}

type BaseEducationSynchronizationConnectionSettingsImpl struct {
	// Client ID used to connect to the provider.
	ClientId *string `json:"clientId,omitempty"`

	// Client secret to authenticate the connection to the provider.
	ClientSecret nullable.Type[string] `json:"clientSecret,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEducationSynchronizationConnectionSettingsImpl) EducationSynchronizationConnectionSettings() BaseEducationSynchronizationConnectionSettingsImpl {
	return s
}

var _ EducationSynchronizationConnectionSettings = RawEducationSynchronizationConnectionSettingsImpl{}

// RawEducationSynchronizationConnectionSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEducationSynchronizationConnectionSettingsImpl struct {
	educationSynchronizationConnectionSettings BaseEducationSynchronizationConnectionSettingsImpl
	Type                                       string
	Values                                     map[string]interface{}
}

func (s RawEducationSynchronizationConnectionSettingsImpl) EducationSynchronizationConnectionSettings() BaseEducationSynchronizationConnectionSettingsImpl {
	return s.educationSynchronizationConnectionSettings
}

func UnmarshalEducationSynchronizationConnectionSettingsImplementation(input []byte) (EducationSynchronizationConnectionSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSynchronizationConnectionSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSynchronizationOAuth1ConnectionSettings") {
		var out EducationSynchronizationOAuth1ConnectionSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSynchronizationOAuth1ConnectionSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSynchronizationOAuth2ClientCredentialsConnectionSettings") {
		var out EducationSynchronizationOAuth2ClientCredentialsConnectionSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSynchronizationOAuth2ClientCredentialsConnectionSettings: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationSynchronizationConnectionSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationSynchronizationConnectionSettingsImpl: %+v", err)
	}

	return RawEducationSynchronizationConnectionSettingsImpl{
		educationSynchronizationConnectionSettings: parent,
		Type:   value,
		Values: temp,
	}, nil

}
