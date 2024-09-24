package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataCredential interface {
	IndustryDataCredential() BaseIndustryDataCredentialImpl
}

var _ IndustryDataCredential = BaseIndustryDataCredentialImpl{}

type BaseIndustryDataCredentialImpl struct {
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

func (s BaseIndustryDataCredentialImpl) IndustryDataCredential() BaseIndustryDataCredentialImpl {
	return s
}

var _ IndustryDataCredential = RawIndustryDataCredentialImpl{}

// RawIndustryDataCredentialImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataCredentialImpl struct {
	industryDataCredential BaseIndustryDataCredentialImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawIndustryDataCredentialImpl) IndustryDataCredential() BaseIndustryDataCredentialImpl {
	return s.industryDataCredential
}

var _ json.Marshaler = BaseIndustryDataCredentialImpl{}

func (s BaseIndustryDataCredentialImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataCredentialImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataCredentialImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataCredentialImpl: %+v", err)
	}

	delete(decoded, "isValid")
	delete(decoded, "lastValidDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataCredentialImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIndustryDataCredentialImplementation(input []byte) (IndustryDataCredential, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataCredential into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.oAuthClientCredential") {
		var out IndustryDataOAuthClientCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataOAuthClientCredential: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataCredentialImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataCredentialImpl: %+v", err)
	}

	return RawIndustryDataCredentialImpl{
		industryDataCredential: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
