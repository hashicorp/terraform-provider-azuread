package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ServicePrincipalSignInActivity{}

type ServicePrincipalSignInActivity struct {
	// The globally unique appId (also called client ID on the Microsoft Entra admin center) of the credentialed resource
	// application.
	AppId *string `json:"appId,omitempty"`

	// The sign-in activity of the application in a app-only authentication flow (app-to-app tokens) where the application
	// acts like a client.
	ApplicationAuthenticationClientSignInActivity *SignInActivity `json:"applicationAuthenticationClientSignInActivity,omitempty"`

	// The sign-in activity of the application in a app-only authentication flow (app-to-app tokens) where the application
	// acts like a resource.
	ApplicationAuthenticationResourceSignInActivity *SignInActivity `json:"applicationAuthenticationResourceSignInActivity,omitempty"`

	// The sign-in activity of the application in a delegated flow (user sign-in) where the application acts like a client.
	DelegatedClientSignInActivity *SignInActivity `json:"delegatedClientSignInActivity,omitempty"`

	// The sign-in activity of the application in a delegated flow (user sign-in) where the application acts like a
	// resource.
	DelegatedResourceSignInActivity *SignInActivity `json:"delegatedResourceSignInActivity,omitempty"`

	// The most recent sign-in activity of the application across delegated or app-only flows where the application is used
	// either as a client or resource.
	LastSignInActivity *SignInActivity `json:"lastSignInActivity,omitempty"`

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

func (s ServicePrincipalSignInActivity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServicePrincipalSignInActivity{}

func (s ServicePrincipalSignInActivity) MarshalJSON() ([]byte, error) {
	type wrapper ServicePrincipalSignInActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServicePrincipalSignInActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServicePrincipalSignInActivity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.servicePrincipalSignInActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServicePrincipalSignInActivity: %+v", err)
	}

	return encoded, nil
}
