package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignInIdentity = UserSignIn{}

type UserSignIn struct {
	// TenantId of the guest user as applies to Microsoft Entra B2B scenarios.
	ExternalTenantId nullable.Type[string] `json:"externalTenantId,omitempty"`

	ExternalUserType *ConditionalAccessGuestOrExternalUserTypes `json:"externalUserType,omitempty"`

	// Object ID of the user.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Fields inherited from SignInIdentity

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserSignIn) SignInIdentity() BaseSignInIdentityImpl {
	return BaseSignInIdentityImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserSignIn{}

func (s UserSignIn) MarshalJSON() ([]byte, error) {
	type wrapper UserSignIn
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserSignIn: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserSignIn: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userSignIn"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserSignIn: %+v", err)
	}

	return encoded, nil
}
