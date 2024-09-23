package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerTaskConfigurationRoleBase = PlannerRelationshipBasedUserType{}

type PlannerRelationshipBasedUserType struct {
	Role *PlannerRelationshipUserRoles `json:"role,omitempty"`

	// Fields inherited from PlannerTaskConfigurationRoleBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RoleKind *PlannerUserRoleKind `json:"roleKind,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PlannerRelationshipBasedUserType) PlannerTaskConfigurationRoleBase() BasePlannerTaskConfigurationRoleBaseImpl {
	return BasePlannerTaskConfigurationRoleBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		RoleKind:  s.RoleKind,
	}
}

var _ json.Marshaler = PlannerRelationshipBasedUserType{}

func (s PlannerRelationshipBasedUserType) MarshalJSON() ([]byte, error) {
	type wrapper PlannerRelationshipBasedUserType
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerRelationshipBasedUserType: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerRelationshipBasedUserType: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerRelationshipBasedUserType"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerRelationshipBasedUserType: %+v", err)
	}

	return encoded, nil
}
