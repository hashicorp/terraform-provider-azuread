package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DelegatedAdminRelationshipOperation{}

type DelegatedAdminRelationshipOperation struct {
	// The time in ISO 8601 format and in UTC time when the long-running operation was created. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The data (payload) for the operation. Read-only.
	Data *string `json:"data,omitempty"`

	// The time in ISO 8601 format and in UTC time when the long-running operation was last modified. Read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	OperationType *DelegatedAdminRelationshipOperationType `json:"operationType,omitempty"`
	Status        *LongRunningOperationStatus              `json:"status,omitempty"`

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

func (s DelegatedAdminRelationshipOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DelegatedAdminRelationshipOperation{}

func (s DelegatedAdminRelationshipOperation) MarshalJSON() ([]byte, error) {
	type wrapper DelegatedAdminRelationshipOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DelegatedAdminRelationshipOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DelegatedAdminRelationshipOperation: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "data")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.delegatedAdminRelationshipOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DelegatedAdminRelationshipOperation: %+v", err)
	}

	return encoded, nil
}
