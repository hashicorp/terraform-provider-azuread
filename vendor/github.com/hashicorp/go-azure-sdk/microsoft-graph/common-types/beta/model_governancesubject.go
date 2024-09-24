package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GovernanceSubject{}

type GovernanceSubject struct {
	// The display name of the subject.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The email address of the user subject. If the subject is in other types, it's empty.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The principal name of the user subject. If the subject is in other types, it's empty.
	PrincipalName nullable.Type[string] `json:"principalName,omitempty"`

	// The type of the subject. The value can be User, Group, and ServicePrincipal.
	Type nullable.Type[string] `json:"type,omitempty"`

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

func (s GovernanceSubject) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GovernanceSubject{}

func (s GovernanceSubject) MarshalJSON() ([]byte, error) {
	type wrapper GovernanceSubject
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GovernanceSubject: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GovernanceSubject: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.governanceSubject"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GovernanceSubject: %+v", err)
	}

	return encoded, nil
}
