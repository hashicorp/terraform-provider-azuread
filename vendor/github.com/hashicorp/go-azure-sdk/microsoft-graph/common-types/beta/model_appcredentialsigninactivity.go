package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AppCredentialSignInActivity{}

type AppCredentialSignInActivity struct {
	// The globally unique appId (also called client ID on the Microsoft Entra admin center) of the credential application.
	AppId *string `json:"appId,omitempty"`

	// The ID of the credential application instance.
	AppObjectId nullable.Type[string] `json:"appObjectId,omitempty"`

	// The date and time when the credential was created. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	CredentialOrigin *ApplicationKeyOrigin `json:"credentialOrigin,omitempty"`

	// The date and time when the credential is set to expire. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The key ID of the credential.
	KeyId *string `json:"keyId,omitempty"`

	// Specifies the key type. The possible values are: clientSecret, certificate, unknownFutureValue.
	KeyType *ApplicationKeyType `json:"keyType,omitempty"`

	// Specifies what the key was used for. The possible values are: sign, verify, unknownFutureValue.
	KeyUsage *ApplicationKeyUsage `json:"keyUsage,omitempty"`

	// The ID of the accessed resource.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// The ID of the service principal.
	ServicePrincipalObjectId nullable.Type[string] `json:"servicePrincipalObjectId,omitempty"`

	SignInActivity *SignInActivity `json:"signInActivity,omitempty"`

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

func (s AppCredentialSignInActivity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AppCredentialSignInActivity{}

func (s AppCredentialSignInActivity) MarshalJSON() ([]byte, error) {
	type wrapper AppCredentialSignInActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppCredentialSignInActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppCredentialSignInActivity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appCredentialSignInActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppCredentialSignInActivity: %+v", err)
	}

	return encoded, nil
}
