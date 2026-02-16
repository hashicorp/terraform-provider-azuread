package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CredentialUserRegistrationCount{}

type CredentialUserRegistrationCount struct {
	// Provides the count of users with accountEnabled set to true in the tenant.
	TotalUserCount *int64 `json:"totalUserCount,omitempty"`

	// A collection of registration count and status information for users in your tenant.
	UserRegistrationCounts *[]UserRegistrationCount `json:"userRegistrationCounts,omitempty"`

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

func (s CredentialUserRegistrationCount) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CredentialUserRegistrationCount{}

func (s CredentialUserRegistrationCount) MarshalJSON() ([]byte, error) {
	type wrapper CredentialUserRegistrationCount
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CredentialUserRegistrationCount: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CredentialUserRegistrationCount: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.credentialUserRegistrationCount"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CredentialUserRegistrationCount: %+v", err)
	}

	return encoded, nil
}
