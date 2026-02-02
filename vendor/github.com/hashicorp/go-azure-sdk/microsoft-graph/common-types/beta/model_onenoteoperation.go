package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Operation = OnenoteOperation{}

type OnenoteOperation struct {
	// The error returned by the operation.
	Error *OnenoteOperationError `json:"error,omitempty"`

	// The operation percent complete if the operation is still in running status.
	PercentComplete nullable.Type[string] `json:"percentComplete,omitempty"`

	// The resource id.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// The resource URI for the object. For example, the resource URI for a copied page or section.
	ResourceLocation nullable.Type[string] `json:"resourceLocation,omitempty"`

	// Fields inherited from Operation

	// The start time of the operation.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The time of the last action of the operation.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// Possible values are: notStarted, running, completed, failed. Read-only.
	Status *OperationStatus `json:"status,omitempty"`

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

func (s OnenoteOperation) Operation() BaseOperationImpl {
	return BaseOperationImpl{
		CreatedDateTime:    s.CreatedDateTime,
		LastActionDateTime: s.LastActionDateTime,
		Status:             s.Status,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s OnenoteOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnenoteOperation{}

func (s OnenoteOperation) MarshalJSON() ([]byte, error) {
	type wrapper OnenoteOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnenoteOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnenoteOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onenoteOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnenoteOperation: %+v", err)
	}

	return encoded, nil
}
