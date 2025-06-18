package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = TemporaryAccessPassAuthenticationMethod{}

type TemporaryAccessPassAuthenticationMethod struct {
	// Determines whether the pass is limited to a one-time use. If true, the pass can be used once; if false, the pass can
	// be used multiple times within the Temporary Access Pass lifetime.
	IsUsableOnce nullable.Type[bool] `json:"isUsableOnce,omitempty"`

	// The lifetime of the Temporary Access Pass in minutes starting at startDateTime. Must be between 10 and 43200
	// inclusive (equivalent to 30 days).
	LifetimeInMinutes nullable.Type[int64] `json:"lifetimeInMinutes,omitempty"`

	// The date and time when the Temporary Access Pass becomes available to use and when isUsable is true is enforced.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The Temporary Access Pass used to authenticate. Returned only on creation of a new
	// temporaryAccessPassAuthenticationMethod object; Hidden in subsequent read operations and returned as null with GET.
	TemporaryAccessPass nullable.Type[string] `json:"temporaryAccessPass,omitempty"`

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

func (s TemporaryAccessPassAuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s TemporaryAccessPassAuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TemporaryAccessPassAuthenticationMethod{}

func (s TemporaryAccessPassAuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper TemporaryAccessPassAuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TemporaryAccessPassAuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TemporaryAccessPassAuthenticationMethod: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.temporaryAccessPassAuthenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TemporaryAccessPassAuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
