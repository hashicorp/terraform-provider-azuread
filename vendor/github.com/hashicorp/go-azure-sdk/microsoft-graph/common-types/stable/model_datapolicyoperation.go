package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DataPolicyOperation{}

type DataPolicyOperation struct {
	// Represents when the request for this data policy operation was completed, in UTC time, using the ISO 8601 format. For
	// example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Null until the operation completes.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// Possible values are: notStarted, running, complete, failed, unknownFutureValue.
	Status *DataPolicyOperationStatus `json:"status,omitempty"`

	// The URL location to where data is being exported for export requests.
	StorageLocation nullable.Type[string] `json:"storageLocation,omitempty"`

	// Represents when the request for this data operation was submitted, in UTC time, using the ISO 8601 format. For
	// example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	SubmittedDateTime *string `json:"submittedDateTime,omitempty"`

	// The id for the user on whom the operation is performed.
	UserId *string `json:"userId,omitempty"`

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

func (s DataPolicyOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DataPolicyOperation{}

func (s DataPolicyOperation) MarshalJSON() ([]byte, error) {
	type wrapper DataPolicyOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataPolicyOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataPolicyOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.dataPolicyOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataPolicyOperation: %+v", err)
	}

	return encoded, nil
}
