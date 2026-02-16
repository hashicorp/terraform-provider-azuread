package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LongRunningOperation = RichLongRunningOperation{}

type RichLongRunningOperation struct {
	// Error due to which the operation failed.
	Error *PublicError `json:"error,omitempty"`

	// A value between 0 and 100 that indicates the progress of the operation.
	PercentageComplete nullable.Type[int64] `json:"percentageComplete,omitempty"`

	// A unique identifier for the result.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Type of the operation.
	Type nullable.Type[string] `json:"type,omitempty"`

	// Fields inherited from LongRunningOperation

	// The start time of the operation. The timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The time of the last action in the operation. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// URI of the resource that the operation is performed on.
	ResourceLocation nullable.Type[string] `json:"resourceLocation,omitempty"`

	// The status of the operation. The possible values are: notStarted, running, succeeded, failed, skipped,
	// unknownFutureValue.
	Status *LongRunningOperationStatus `json:"status,omitempty"`

	// Details about the status of the operation.
	StatusDetail nullable.Type[string] `json:"statusDetail,omitempty"`

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

func (s RichLongRunningOperation) LongRunningOperation() BaseLongRunningOperationImpl {
	return BaseLongRunningOperationImpl{
		CreatedDateTime:    s.CreatedDateTime,
		LastActionDateTime: s.LastActionDateTime,
		ResourceLocation:   s.ResourceLocation,
		Status:             s.Status,
		StatusDetail:       s.StatusDetail,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s RichLongRunningOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RichLongRunningOperation{}

func (s RichLongRunningOperation) MarshalJSON() ([]byte, error) {
	type wrapper RichLongRunningOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RichLongRunningOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RichLongRunningOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.richLongRunningOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RichLongRunningOperation: %+v", err)
	}

	return encoded, nil
}
