package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementAlertIncident = SequentialActivationRenewalsAlertIncident{}

type SequentialActivationRenewalsAlertIncident struct {
	// The length of sequential activation of the same role.
	ActivationCount nullable.Type[int64] `json:"activationCount,omitempty"`

	// Display name of the subject that the incident applies to.
	AssigneeDisplayName nullable.Type[string] `json:"assigneeDisplayName,omitempty"`

	// The identifier of the subject that the incident applies to.
	AssigneeId nullable.Type[string] `json:"assigneeId,omitempty"`

	// User principal name of the subject that the incident applies to. Applies to user principals.
	AssigneeUserPrincipalName nullable.Type[string] `json:"assigneeUserPrincipalName,omitempty"`

	// The identifier for the directory role definition that's in scope of this incident.
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The display name for the directory role.
	RoleDisplayName nullable.Type[string] `json:"roleDisplayName,omitempty"`

	// The globally unique identifier for the directory role.
	RoleTemplateId nullable.Type[string] `json:"roleTemplateId,omitempty"`

	// End date time of the sequential activation event.
	SequenceEndDateTime nullable.Type[string] `json:"sequenceEndDateTime,omitempty"`

	// Start date time of the sequential activation event.
	SequenceStartDateTime nullable.Type[string] `json:"sequenceStartDateTime,omitempty"`

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

func (s SequentialActivationRenewalsAlertIncident) UnifiedRoleManagementAlertIncident() BaseUnifiedRoleManagementAlertIncidentImpl {
	return BaseUnifiedRoleManagementAlertIncidentImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SequentialActivationRenewalsAlertIncident) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SequentialActivationRenewalsAlertIncident{}

func (s SequentialActivationRenewalsAlertIncident) MarshalJSON() ([]byte, error) {
	type wrapper SequentialActivationRenewalsAlertIncident
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SequentialActivationRenewalsAlertIncident: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SequentialActivationRenewalsAlertIncident: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sequentialActivationRenewalsAlertIncident"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SequentialActivationRenewalsAlertIncident: %+v", err)
	}

	return encoded, nil
}
