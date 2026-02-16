package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerBucketCreation = PlannerExternalBucketSource{}

type PlannerExternalBucketSource struct {
	// Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format.
	// For example, Contoso company owned application for customer support would have a value like
	// 'com.constoso.customerSupport'.
	ContextScenarioId nullable.Type[string] `json:"contextScenarioId,omitempty"`

	// Nullable. The ID of the external entity's containing entity or context.
	ExternalContextId nullable.Type[string] `json:"externalContextId,omitempty"`

	// Nullable. The ID of the entity that an external service associates with a bucket.
	ExternalObjectId nullable.Type[string] `json:"externalObjectId,omitempty"`

	// Fields inherited from PlannerBucketCreation

	// Specifies what kind of creation source the bucket is created with. The possible values are: external, publication and
	// unknownFutureValue.
	CreationSourceKind *PlannerCreationSourceKind `json:"creationSourceKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PlannerExternalBucketSource) PlannerBucketCreation() BasePlannerBucketCreationImpl {
	return BasePlannerBucketCreationImpl{
		CreationSourceKind: s.CreationSourceKind,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

var _ json.Marshaler = PlannerExternalBucketSource{}

func (s PlannerExternalBucketSource) MarshalJSON() ([]byte, error) {
	type wrapper PlannerExternalBucketSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerExternalBucketSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerExternalBucketSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerExternalBucketSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerExternalBucketSource: %+v", err)
	}

	return encoded, nil
}
