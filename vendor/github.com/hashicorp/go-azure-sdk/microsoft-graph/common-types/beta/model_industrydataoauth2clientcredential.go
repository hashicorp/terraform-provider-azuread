package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataOAuthClientCredential = IndustryDataOAuth2ClientCredential{}

type IndustryDataOAuth2ClientCredential struct {
	// The OAuth scope that is provided to the authentication process.
	Scope nullable.Type[string] `json:"scope,omitempty"`

	// The URL with which to retrieve the token after authentication happens.
	TokenUrl *string `json:"tokenUrl,omitempty"`

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

func (s IndustryDataOAuth2ClientCredential) IndustryDataOAuthClientCredential() BaseIndustryDataOAuthClientCredentialImpl {
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

func (s IndustryDataOAuth2ClientCredential) IndustryDataCredential() BaseIndustryDataCredentialImpl {
	return BaseIndustryDataCredentialImpl{
		DisplayName:       s.DisplayName,
		IsValid:           s.IsValid,
		LastValidDateTime: s.LastValidDateTime,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataOAuth2ClientCredential{}

func (s IndustryDataOAuth2ClientCredential) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataOAuth2ClientCredential
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataOAuth2ClientCredential: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataOAuth2ClientCredential: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.oAuth2ClientCredential"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataOAuth2ClientCredential: %+v", err)
	}

	return encoded, nil
}
