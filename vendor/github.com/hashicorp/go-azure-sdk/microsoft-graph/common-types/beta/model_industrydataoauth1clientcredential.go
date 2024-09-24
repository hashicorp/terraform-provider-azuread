package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataOAuthClientCredential = IndustryDataOAuth1ClientCredential{}

type IndustryDataOAuth1ClientCredential struct {

	// Fields inherited from IndustryDataOAuthClientCredential

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

func (s IndustryDataOAuth1ClientCredential) IndustryDataOAuthClientCredential() BaseIndustryDataOAuthClientCredentialImpl {
	return BaseIndustryDataOAuthClientCredentialImpl{
		ClientId:          s.ClientId,
		ClientSecret:      s.ClientSecret,
		DisplayName:       s.DisplayName,
		IsValid:           s.IsValid,
		LastValidDateTime: s.LastValidDateTime,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s IndustryDataOAuth1ClientCredential) IndustryDataCredential() BaseIndustryDataCredentialImpl {
	return BaseIndustryDataCredentialImpl{
		DisplayName:       s.DisplayName,
		IsValid:           s.IsValid,
		LastValidDateTime: s.LastValidDateTime,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataOAuth1ClientCredential{}

func (s IndustryDataOAuth1ClientCredential) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataOAuth1ClientCredential
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataOAuth1ClientCredential: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataOAuth1ClientCredential: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.oAuth1ClientCredential"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataOAuth1ClientCredential: %+v", err)
	}

	return encoded, nil
}
