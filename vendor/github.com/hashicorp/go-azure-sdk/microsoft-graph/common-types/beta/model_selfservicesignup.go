package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SelfServiceSignUp{}

type SelfServiceSignUp struct {
	AppDisplayName         nullable.Type[string]                 `json:"appDisplayName,omitempty"`
	AppId                  nullable.Type[string]                 `json:"appId,omitempty"`
	AppliedEventListeners  *[]AppliedAuthenticationEventListener `json:"appliedEventListeners,omitempty"`
	CorrelationId          *string                               `json:"correlationId,omitempty"`
	CreatedDateTime        *string                               `json:"createdDateTime,omitempty"`
	SignUpIdentity         *SignUpIdentity                       `json:"signUpIdentity,omitempty"`
	SignUpIdentityProvider *string                               `json:"signUpIdentityProvider,omitempty"`
	SignUpStage            *SignUpStage                          `json:"signUpStage,omitempty"`
	Status                 *SignUpStatus                         `json:"status,omitempty"`
	UserSnapshot           *CiamUserSnapshot                     `json:"userSnapshot,omitempty"`

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

func (s SelfServiceSignUp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SelfServiceSignUp{}

func (s SelfServiceSignUp) MarshalJSON() ([]byte, error) {
	type wrapper SelfServiceSignUp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SelfServiceSignUp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SelfServiceSignUp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.selfServiceSignUp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SelfServiceSignUp: %+v", err)
	}

	return encoded, nil
}
