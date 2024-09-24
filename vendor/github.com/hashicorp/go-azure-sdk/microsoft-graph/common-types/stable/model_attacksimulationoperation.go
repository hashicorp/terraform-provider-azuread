package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LongRunningOperation = AttackSimulationOperation{}

type AttackSimulationOperation struct {
	// Percentage of completion of the respective operation.
	PercentageCompleted *int64 `json:"percentageCompleted,omitempty"`

	// Tenant identifier.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The attack simulation operation type. Possible values are: createSimulation, updateSimulation, unknownFutureValue.
	Type *AttackSimulationOperationType `json:"type,omitempty"`

	// Fields inherited from LongRunningOperation

	// The start time of the operation. The Timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The time of the last action in the operation. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// URI of the resource that the operation is performed on.
	ResourceLocation nullable.Type[string] `json:"resourceLocation,omitempty"`

	// The status of the operation. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue.
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

func (s AttackSimulationOperation) LongRunningOperation() BaseLongRunningOperationImpl {
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

func (s AttackSimulationOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AttackSimulationOperation{}

func (s AttackSimulationOperation) MarshalJSON() ([]byte, error) {
	type wrapper AttackSimulationOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AttackSimulationOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AttackSimulationOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attackSimulationOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AttackSimulationOperation: %+v", err)
	}

	return encoded, nil
}
