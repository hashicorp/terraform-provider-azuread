package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataOAuthClientCredential interface {
	IndustryDataCredential
	IndustryDataOAuthClientCredential() BaseIndustryDataOAuthClientCredentialImpl
}

var _ IndustryDataOAuthClientCredential = BaseIndustryDataOAuthClientCredentialImpl{}

type BaseIndustryDataOAuthClientCredentialImpl struct {
	// The client identifier of the application that is authenticating.
	ClientId *string `json:"clientId,omitempty"`

	// The client secret that is used to authenticate (write-only).
	ClientSecret nullable.Type[string] `json:"clientSecret,omitempty"`

	// Fields inherited from IndustryDataCredential

	// The name of the credential.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates whether the credential provided is valid based on the last data connector validate operation.
	IsValid *bool `json:"isValid,omitempty"`

	// The time that the credential was last successfully validated by the data connector validate operation.
	LastValidDateTime nullable.Type[string] `json:"lastValidDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIndustryDataOAuthClientCredentialImpl) IndustryDataOAuthClientCredential() BaseIndustryDataOAuthClientCredentialImpl {
	return s
}

func (s BaseIndustryDataOAuthClientCredentialImpl) IndustryDataCredential() BaseIndustryDataCredentialImpl {
	return BaseIndustryDataCredentialImpl{
		DisplayName:       s.DisplayName,
		IsValid:           s.IsValid,
		LastValidDateTime: s.LastValidDateTime,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

var _ IndustryDataOAuthClientCredential = RawIndustryDataOAuthClientCredentialImpl{}

// RawIndustryDataOAuthClientCredentialImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataOAuthClientCredentialImpl struct {
	industryDataOAuthClientCredential BaseIndustryDataOAuthClientCredentialImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawIndustryDataOAuthClientCredentialImpl) IndustryDataOAuthClientCredential() BaseIndustryDataOAuthClientCredentialImpl {
	return s.industryDataOAuthClientCredential
}

func (s RawIndustryDataOAuthClientCredentialImpl) IndustryDataCredential() BaseIndustryDataCredentialImpl {
	return s.industryDataOAuthClientCredential.IndustryDataCredential()
}

var _ json.Marshaler = BaseIndustryDataOAuthClientCredentialImpl{}

func (s BaseIndustryDataOAuthClientCredentialImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataOAuthClientCredentialImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataOAuthClientCredentialImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataOAuthClientCredentialImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.oAuthClientCredential"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataOAuthClientCredentialImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIndustryDataOAuthClientCredentialImplementation(input []byte) (IndustryDataOAuthClientCredential, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataOAuthClientCredential into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.oAuth1ClientCredential") {
		var out IndustryDataOAuth1ClientCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataOAuth1ClientCredential: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.oAuth2ClientCredential") {
		var out IndustryDataOAuth2ClientCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataOAuth2ClientCredential: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataOAuthClientCredentialImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataOAuthClientCredentialImpl: %+v", err)
	}

	return RawIndustryDataOAuthClientCredentialImpl{
		industryDataOAuthClientCredential: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
