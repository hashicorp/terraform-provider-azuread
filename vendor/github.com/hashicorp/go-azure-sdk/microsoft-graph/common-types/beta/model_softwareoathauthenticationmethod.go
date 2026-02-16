package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = SoftwareOathAuthenticationMethod{}

type SoftwareOathAuthenticationMethod struct {
	// The secret key of the method. Always returns null.
	SecretKey nullable.Type[string] `json:"secretKey,omitempty"`

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

func (s SoftwareOathAuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SoftwareOathAuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SoftwareOathAuthenticationMethod{}

func (s SoftwareOathAuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper SoftwareOathAuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SoftwareOathAuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SoftwareOathAuthenticationMethod: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.softwareOathAuthenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SoftwareOathAuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
