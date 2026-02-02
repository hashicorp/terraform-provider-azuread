package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerPlanCreation = PlannerExternalPlanSource{}

type PlannerExternalPlanSource struct {
	// Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format.
	// For example, Contoso company owned application for customer support would have a value like
	// 'com.constoso.customerSupport'.
	ContextScenarioId nullable.Type[string] `json:"contextScenarioId,omitempty"`

	// Nullable. The ID of the external entity's containing entity or context.
	ExternalContextId nullable.Type[string] `json:"externalContextId,omitempty"`

	// Nullable. The ID of the entity that an external service associates with a plan.
	ExternalObjectId nullable.Type[string] `json:"externalObjectId,omitempty"`

	// Fields inherited from PlannerPlanCreation

	// Specifies what kind of creation source the plan is created with. The possible values are: external, publication and
	// unknownFutureValue.
	CreationSourceKind *PlannerCreationSourceKind `json:"creationSourceKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PlannerExternalPlanSource) PlannerPlanCreation() BasePlannerPlanCreationImpl {
	return BasePlannerPlanCreationImpl{
		CreationSourceKind: s.CreationSourceKind,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

var _ json.Marshaler = PlannerExternalPlanSource{}

func (s PlannerExternalPlanSource) MarshalJSON() ([]byte, error) {
	type wrapper PlannerExternalPlanSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerExternalPlanSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerExternalPlanSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerExternalPlanSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerExternalPlanSource: %+v", err)
	}

	return encoded, nil
}
