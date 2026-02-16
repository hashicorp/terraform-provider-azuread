package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContext struct {
	// Nullable. An app-defined type of association between the plannerPlan and the app. The app can use this information to
	// track different kinds of relationships to the same plannerPlan.
	AssociationType nullable.Type[string] `json:"associationType,omitempty"`

	// Read-only. The date and time when the plannerPlanContext was created. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The segments of the name of the external experience. Segments represent a hierarchical structure that allows other
	// apps to display the relationship.
	DisplayNameSegments *[]string `json:"displayNameSegments,omitempty"`

	// Read-only. Indicates whether the plan is created from the specified context. Auto-generated based on whether the
	// context is specified as part of plan creation.
	IsCreationContext nullable.Type[bool] `json:"isCreationContext,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Read-only. ID of the app that created the plannerPlanContext.
	OwnerAppId nullable.Type[string] `json:"ownerAppId,omitempty"`
}

var _ json.Marshaler = PlannerPlanContext{}

func (s PlannerPlanContext) MarshalJSON() ([]byte, error) {
	type wrapper PlannerPlanContext
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerPlanContext: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerPlanContext: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "isCreationContext")
	delete(decoded, "ownerAppId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerPlanContext: %+v", err)
	}

	return encoded, nil
}
