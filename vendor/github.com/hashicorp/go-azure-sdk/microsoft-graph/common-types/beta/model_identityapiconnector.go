package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityApiConnector{}

type IdentityApiConnector struct {
	// The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client
	// certificate are supported.
	AuthenticationConfiguration ApiAuthenticationConfigurationBase `json:"authenticationConfiguration"`

	// The name of the API connector.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The URL of the API endpoint to call.
	TargetUrl nullable.Type[string] `json:"targetUrl,omitempty"`

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

func (s IdentityApiConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityApiConnector{}

func (s IdentityApiConnector) MarshalJSON() ([]byte, error) {
	type wrapper IdentityApiConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityApiConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityApiConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityApiConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityApiConnector: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IdentityApiConnector{}

func (s *IdentityApiConnector) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName nullable.Type[string] `json:"displayName,omitempty"`
		TargetUrl   nullable.Type[string] `json:"targetUrl,omitempty"`
		Id          *string               `json:"id,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.TargetUrl = decoded.TargetUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IdentityApiConnector into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authenticationConfiguration"]; ok {
		impl, err := UnmarshalApiAuthenticationConfigurationBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthenticationConfiguration' for 'IdentityApiConnector': %+v", err)
		}
		s.AuthenticationConfiguration = impl
	}

	return nil
}
