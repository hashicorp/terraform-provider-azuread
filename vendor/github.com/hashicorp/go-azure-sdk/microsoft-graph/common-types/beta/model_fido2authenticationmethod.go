package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = Fido2AuthenticationMethod{}

type Fido2AuthenticationMethod struct {
	// Authenticator Attestation GUID, an identifier that indicates the type (such as make and model) of the authenticator.
	AaGuid nullable.Type[string] `json:"aaGuid,omitempty"`

	// The attestation certificate or certificates attached to this security key.
	AttestationCertificates *[]string `json:"attestationCertificates,omitempty"`

	// The attestation level of this FIDO2 security key. Possible values are: attested, notAttested, unknownFutureValue.
	AttestationLevel *AttestationLevel `json:"attestationLevel,omitempty"`

	// The display name of the key as given by the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The manufacturer-assigned model of the FIDO2 security key.
	Model nullable.Type[string] `json:"model,omitempty"`

	// Contains the WebAuthn public key credential information being registered. Only used for write requests. This property
	// isn't returned on read operations.
	PublicKeyCredential *WebauthnPublicKeyCredential `json:"publicKeyCredential,omitempty"`

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

func (s Fido2AuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Fido2AuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Fido2AuthenticationMethod{}

func (s Fido2AuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper Fido2AuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Fido2AuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Fido2AuthenticationMethod: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.fido2AuthenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Fido2AuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
