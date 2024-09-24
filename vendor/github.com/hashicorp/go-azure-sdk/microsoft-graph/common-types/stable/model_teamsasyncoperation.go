package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsAsyncOperation{}

type TeamsAsyncOperation struct {
	// Number of times the operation was attempted before being marked successful or failed.
	AttemptsCount *int64 `json:"attemptsCount,omitempty"`

	// Time when the operation was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Any error that causes the async operation to fail.
	Error *OperationError `json:"error,omitempty"`

	// Time when the async operation was last updated.
	LastActionDateTime *string `json:"lastActionDateTime,omitempty"`

	OperationType *TeamsAsyncOperationType   `json:"operationType,omitempty"`
	Status        *TeamsAsyncOperationStatus `json:"status,omitempty"`

	// The ID of the object that's created or modified as result of this async operation, typically a team.
	TargetResourceId nullable.Type[string] `json:"targetResourceId,omitempty"`

	// The location of the object that's created or modified as result of this async operation. This URL should be treated
	// as an opaque value and not parsed into its component paths.
	TargetResourceLocation nullable.Type[string] `json:"targetResourceLocation,omitempty"`

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

func (s TeamsAsyncOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAsyncOperation{}

func (s TeamsAsyncOperation) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAsyncOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAsyncOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAsyncOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAsyncOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAsyncOperation: %+v", err)
	}

	return encoded, nil
}
