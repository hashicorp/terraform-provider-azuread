package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerTaskCreation = PlannerExternalTaskSource{}

type PlannerExternalTaskSource struct {
	// Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format.
	// For example, Contoso company owned application for customer support would have a value like
	// 'com.constoso.customerSupport'.
	ContextScenarioId nullable.Type[string] `json:"contextScenarioId,omitempty"`

	// Specifies how an application should display the link to the associated plannerExternalTaskSource. The possible values
	// are: none, default.
	DisplayLinkType *PlannerExternalTaskSourceDisplayType `json:"displayLinkType,omitempty"`

	// The segments of the name of the external experience. Segments represent a hierarchical structure that allows other
	// apps to display the relationship.
	DisplayNameSegments *[]string `json:"displayNameSegments,omitempty"`

	// Nullable. The id of the external entity's containing entity or context.
	ExternalContextId nullable.Type[string] `json:"externalContextId,omitempty"`

	// Nullable. The id of the entity that an external service associates with a task.
	ExternalObjectId nullable.Type[string] `json:"externalObjectId,omitempty"`

	// Nullable. The external Item Version for the object specified by the externalObjectId.
	ExternalObjectVersion nullable.Type[string] `json:"externalObjectVersion,omitempty"`

	// Nullable. URL of the user experience represented by the associated plannerExternalTaskSource.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

	// Fields inherited from PlannerTaskCreation

	// Specifies what kind of creation source the task is created with. The possible values are: external, publication and
	// unknownFutureValue.
	CreationSourceKind *PlannerCreationSourceKind `json:"creationSourceKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Information about the publication process that created this task. This field is deprecated and clients should move to
	// using the new inheritance model.
	TeamsPublicationInfo *PlannerTeamsPublicationInfo `json:"teamsPublicationInfo,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PlannerExternalTaskSource) PlannerTaskCreation() BasePlannerTaskCreationImpl {
	return BasePlannerTaskCreationImpl{
		CreationSourceKind:   s.CreationSourceKind,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
		TeamsPublicationInfo: s.TeamsPublicationInfo,
	}
}

var _ json.Marshaler = PlannerExternalTaskSource{}

func (s PlannerExternalTaskSource) MarshalJSON() ([]byte, error) {
	type wrapper PlannerExternalTaskSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerExternalTaskSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerExternalTaskSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerExternalTaskSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerExternalTaskSource: %+v", err)
	}

	return encoded, nil
}
