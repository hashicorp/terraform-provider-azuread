package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = MicrosoftAuthenticatorAuthenticationMethod{}

type MicrosoftAuthenticatorAuthenticationMethod struct {
	// The date and time that this app was registered. This property is null if the device isn't registered for passwordless
	// Phone Sign-In.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The registered device on which Microsoft Authenticator resides. This property is null if the device isn't registered
	// for passwordless Phone Sign-In.
	Device *Device `json:"device,omitempty"`

	// Tags containing app metadata.
	DeviceTag nullable.Type[string] `json:"deviceTag,omitempty"`

	// The name of the device on which this app is registered.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Numerical version of this instance of the Authenticator app.
	PhoneAppVersion nullable.Type[string] `json:"phoneAppVersion,omitempty"`

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

func (s MicrosoftAuthenticatorAuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s MicrosoftAuthenticatorAuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftAuthenticatorAuthenticationMethod{}

func (s MicrosoftAuthenticatorAuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftAuthenticatorAuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftAuthenticatorAuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftAuthenticatorAuthenticationMethod: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftAuthenticatorAuthenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftAuthenticatorAuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
