package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = PasswordlessMicrosoftAuthenticatorAuthenticationMethod{}

type PasswordlessMicrosoftAuthenticatorAuthenticationMethod struct {
	// The timestamp when this method was registered to the user.
	CreationDateTime nullable.Type[string] `json:"creationDateTime,omitempty"`

	Device *Device `json:"device,omitempty"`

	// The display name of the mobile device as given by the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Fields inherited from AuthenticationMethod

	// The date and time the authentication method was registered to the user. Read-only. Optional. This optional value is
	// null if the authentication method doesn't populate it. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

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

func (s PasswordlessMicrosoftAuthenticatorAuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s PasswordlessMicrosoftAuthenticatorAuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PasswordlessMicrosoftAuthenticatorAuthenticationMethod{}

func (s PasswordlessMicrosoftAuthenticatorAuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper PasswordlessMicrosoftAuthenticatorAuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PasswordlessMicrosoftAuthenticatorAuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PasswordlessMicrosoftAuthenticatorAuthenticationMethod: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.passwordlessMicrosoftAuthenticatorAuthenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PasswordlessMicrosoftAuthenticatorAuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
