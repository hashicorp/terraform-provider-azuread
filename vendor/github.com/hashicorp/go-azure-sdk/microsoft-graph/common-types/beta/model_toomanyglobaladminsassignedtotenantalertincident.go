package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementAlertIncident = TooManyGlobalAdminsAssignedToTenantAlertIncident{}

type TooManyGlobalAdminsAssignedToTenantAlertIncident struct {
	// Display name of the subject that the incident applies to.
	AssigneeDisplayName nullable.Type[string] `json:"assigneeDisplayName,omitempty"`

	// The identifier of the subject that the incident applies to.
	AssigneeId nullable.Type[string] `json:"assigneeId,omitempty"`

	// User principal name of the subject that the incident applies to. Applies to user principals.
	AssigneeUserPrincipalName nullable.Type[string] `json:"assigneeUserPrincipalName,omitempty"`

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

func (s TooManyGlobalAdminsAssignedToTenantAlertIncident) UnifiedRoleManagementAlertIncident() BaseUnifiedRoleManagementAlertIncidentImpl {
	return BaseUnifiedRoleManagementAlertIncidentImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s TooManyGlobalAdminsAssignedToTenantAlertIncident) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TooManyGlobalAdminsAssignedToTenantAlertIncident{}

func (s TooManyGlobalAdminsAssignedToTenantAlertIncident) MarshalJSON() ([]byte, error) {
	type wrapper TooManyGlobalAdminsAssignedToTenantAlertIncident
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TooManyGlobalAdminsAssignedToTenantAlertIncident: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TooManyGlobalAdminsAssignedToTenantAlertIncident: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertIncident"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TooManyGlobalAdminsAssignedToTenantAlertIncident: %+v", err)
	}

	return encoded, nil
}
