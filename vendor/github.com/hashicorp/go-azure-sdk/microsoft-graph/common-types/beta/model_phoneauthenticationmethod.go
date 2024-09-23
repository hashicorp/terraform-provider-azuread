package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = PhoneAuthenticationMethod{}

type PhoneAuthenticationMethod struct {
	// The phone number to text or call for authentication. Phone numbers use the format '+<country code>
	// <number>x<extension>', with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers
	// are rejected when creating/updating if they don't match the required format.
	PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`

	// The type of this phone. Possible values are: mobile, alternateMobile, or office.
	PhoneType *AuthenticationPhoneType `json:"phoneType,omitempty"`

	// Whether a phone is ready to be used for SMS sign-in or not. Possible values are: notSupported, notAllowedByPolicy,
	// notEnabled, phoneNumberNotUnique, ready, or notConfigured, unknownFutureValue.
	SmsSignInState *AuthenticationMethodSignInState `json:"smsSignInState,omitempty"`

	// Fields inherited from AuthenticationMethod

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

func (s PhoneAuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s PhoneAuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PhoneAuthenticationMethod{}

func (s PhoneAuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper PhoneAuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PhoneAuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PhoneAuthenticationMethod: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.phoneAuthenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PhoneAuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
