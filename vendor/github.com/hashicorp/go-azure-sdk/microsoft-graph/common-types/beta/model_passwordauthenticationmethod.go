package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = PasswordAuthenticationMethod{}

type PasswordAuthenticationMethod struct {
	// For security, the password is always returned as null from a LIST or GET operation.
	Password nullable.Type[string] `json:"password,omitempty"`

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

func (s PasswordAuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s PasswordAuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PasswordAuthenticationMethod{}

func (s PasswordAuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper PasswordAuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PasswordAuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PasswordAuthenticationMethod: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.passwordAuthenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PasswordAuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
