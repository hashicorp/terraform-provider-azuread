package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRoleManagementAlert{}

type UnifiedRoleManagementAlert struct {
	// The configuration of the alert in PIM for Microsoft Entra roles. Alert configurations are pre-defined and cannot be
	// created or deleted, but some configurations can be modified. Supports $filter for the isEnabled property and $expand.
	AlertConfiguration *UnifiedRoleManagementAlertConfiguration `json:"alertConfiguration,omitempty"`

	// Contains the description, impact, and measures to mitigate or prevent the security alert from being triggered in your
	// tenant. Supports $expand.
	AlertDefinition *UnifiedRoleManagementAlertDefinition `json:"alertDefinition,omitempty"`

	// The identifier of an alert definition. Supports $filter (eq, ne).
	AlertDefinitionId nullable.Type[string] `json:"alertDefinitionId,omitempty"`

	// Represents the incidents of this type of alert that have been triggered in Privileged Identity Management (PIM) for
	// Microsoft Entra roles in the tenant. Supports $expand.
	AlertIncidents *[]UnifiedRoleManagementAlertIncident `json:"alertIncidents,omitempty"`

	// The number of incidents triggered in the tenant and relating to the alert. Can only be a positive integer.
	IncidentCount nullable.Type[int64] `json:"incidentCount,omitempty"`

	// false by default. true if the alert is active.
	IsActive nullable.Type[bool] `json:"isActive,omitempty"`

	// The date time when the alert configuration was updated or new incidents generated.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The date time when the tenant was last scanned for incidents that trigger this alert.
	LastScannedDateTime nullable.Type[string] `json:"lastScannedDateTime,omitempty"`

	// The identifier of the scope where the alert is related. / is the only supported one for the tenant. Supports $filter
	// (eq, ne).
	ScopeId nullable.Type[string] `json:"scopeId,omitempty"`

	// The type of scope where the alert is created. DirectoryRole is the only currently supported scope type for Microsoft
	// Entra roles.
	ScopeType nullable.Type[string] `json:"scopeType,omitempty"`

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

func (s UnifiedRoleManagementAlert) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleManagementAlert{}

func (s UnifiedRoleManagementAlert) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleManagementAlert
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleManagementAlert: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementAlert: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementAlert"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleManagementAlert: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UnifiedRoleManagementAlert{}

func (s *UnifiedRoleManagementAlert) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AlertDefinition      *UnifiedRoleManagementAlertDefinition `json:"alertDefinition,omitempty"`
		AlertDefinitionId    nullable.Type[string]                 `json:"alertDefinitionId,omitempty"`
		IncidentCount        nullable.Type[int64]                  `json:"incidentCount,omitempty"`
		IsActive             nullable.Type[bool]                   `json:"isActive,omitempty"`
		LastModifiedDateTime nullable.Type[string]                 `json:"lastModifiedDateTime,omitempty"`
		LastScannedDateTime  nullable.Type[string]                 `json:"lastScannedDateTime,omitempty"`
		ScopeId              nullable.Type[string]                 `json:"scopeId,omitempty"`
		ScopeType            nullable.Type[string]                 `json:"scopeType,omitempty"`
		Id                   *string                               `json:"id,omitempty"`
		ODataId              *string                               `json:"@odata.id,omitempty"`
		ODataType            *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AlertDefinition = decoded.AlertDefinition
	s.AlertDefinitionId = decoded.AlertDefinitionId
	s.IncidentCount = decoded.IncidentCount
	s.IsActive = decoded.IsActive
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LastScannedDateTime = decoded.LastScannedDateTime
	s.ScopeId = decoded.ScopeId
	s.ScopeType = decoded.ScopeType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UnifiedRoleManagementAlert into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["alertConfiguration"]; ok {
		impl, err := UnmarshalUnifiedRoleManagementAlertConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AlertConfiguration' for 'UnifiedRoleManagementAlert': %+v", err)
		}
		s.AlertConfiguration = &impl
	}

	if v, ok := temp["alertIncidents"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AlertIncidents into list []json.RawMessage: %+v", err)
		}

		output := make([]UnifiedRoleManagementAlertIncident, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUnifiedRoleManagementAlertIncidentImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AlertIncidents' for 'UnifiedRoleManagementAlert': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AlertIncidents = &output
	}

	return nil
}
