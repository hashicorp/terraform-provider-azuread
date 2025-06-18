package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Identity = TeamworkUserIdentity{}

type TeamworkUserIdentity struct {
	// Type of user. Possible values are: aadUser, onPremiseAadUser, anonymousGuest, federatedUser,
	// personalMicrosoftAccountUser, skypeUser, phoneUser, emailUser and azureCommunicationServicesUser.
	UserIdentityType *TeamworkUserIdentityType `json:"userIdentityType,omitempty"`

	// User principal name (UPN) of the user.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Fields inherited from Identity

	// The display name of the identity. This property is read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identifier of the identity. This property is read-only.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TeamworkUserIdentity) Identity() BaseIdentityImpl {
	return BaseIdentityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ json.Marshaler = TeamworkUserIdentity{}

func (s TeamworkUserIdentity) MarshalJSON() ([]byte, error) {
	type wrapper TeamworkUserIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamworkUserIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkUserIdentity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamworkUserIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamworkUserIdentity: %+v", err)
	}

	return encoded, nil
}
