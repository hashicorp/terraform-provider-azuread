package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = PlatformCredentialAuthenticationMethod{}

type PlatformCredentialAuthenticationMethod struct {
	// The registered device on which this Platform Credential resides. Supports $expand. When you get a user's Platform
	// Credential registration information, this property is returned only on a single GET and when you specify ?$expand.
	// For example, GET
	// /users/admin@contoso.com/authentication/platformCredentialAuthenticationMethod/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
	Device *Device `json:"device,omitempty"`

	// The name of the device on which Platform Credential is registered.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Key strength of this Platform Credential key. Possible values are: normal, weak, unknown.
	KeyStrength *AuthenticationMethodKeyStrength `json:"keyStrength,omitempty"`

	// Platform on which this Platform Credential key is present. Possible values are: unknown, windows, macOS,iOS, android,
	// linux.
	Platform *AuthenticationMethodPlatform `json:"platform,omitempty"`

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

func (s PlatformCredentialAuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s PlatformCredentialAuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlatformCredentialAuthenticationMethod{}

func (s PlatformCredentialAuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper PlatformCredentialAuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlatformCredentialAuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlatformCredentialAuthenticationMethod: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.platformCredentialAuthenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlatformCredentialAuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
