package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceNowAuthenticationMethod = ServiceNowOauthSecretAuthentication{}

type ServiceNowOauthSecretAuthentication struct {
	// Tenant appId registered with Azure AD
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Fields inherited from ServiceNowAuthenticationMethod

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ServiceNowOauthSecretAuthentication) ServiceNowAuthenticationMethod() BaseServiceNowAuthenticationMethodImpl {
	return BaseServiceNowAuthenticationMethodImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceNowOauthSecretAuthentication{}

func (s ServiceNowOauthSecretAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper ServiceNowOauthSecretAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceNowOauthSecretAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceNowOauthSecretAuthentication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceNowOauthSecretAuthentication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceNowOauthSecretAuthentication: %+v", err)
	}

	return encoded, nil
}
