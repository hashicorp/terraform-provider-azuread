package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RoleManagementAlert{}

type RoleManagementAlert struct {
	// The various configurations of an alert for Microsoft Entra roles. The configurations are predefined and can't be
	// created or deleted, but some of the configurations can be modified.
	AlertConfigurations *[]UnifiedRoleManagementAlertConfiguration `json:"alertConfigurations,omitempty"`

	// Defines an alert, its impact, and measures to mitigate or prevent it.
	AlertDefinitions *[]UnifiedRoleManagementAlertDefinition `json:"alertDefinitions,omitempty"`

	// Represents the alert entity.
	Alerts *[]UnifiedRoleManagementAlert `json:"alerts,omitempty"`

	// Represents operations on resources that take a long time to complete and can run in the background until completion.
	Operations *[]LongRunningOperation `json:"operations,omitempty"`

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

func (s RoleManagementAlert) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RoleManagementAlert{}

func (s RoleManagementAlert) MarshalJSON() ([]byte, error) {
	type wrapper RoleManagementAlert
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RoleManagementAlert: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RoleManagementAlert: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.roleManagementAlert"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RoleManagementAlert: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &RoleManagementAlert{}

func (s *RoleManagementAlert) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AlertDefinitions *[]UnifiedRoleManagementAlertDefinition `json:"alertDefinitions,omitempty"`
		Alerts           *[]UnifiedRoleManagementAlert           `json:"alerts,omitempty"`
		Id               *string                                 `json:"id,omitempty"`
		ODataId          *string                                 `json:"@odata.id,omitempty"`
		ODataType        *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AlertDefinitions = decoded.AlertDefinitions
	s.Alerts = decoded.Alerts
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RoleManagementAlert into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["alertConfigurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AlertConfigurations into list []json.RawMessage: %+v", err)
		}

		output := make([]UnifiedRoleManagementAlertConfiguration, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUnifiedRoleManagementAlertConfigurationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AlertConfigurations' for 'RoleManagementAlert': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AlertConfigurations = &output
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]LongRunningOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLongRunningOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'RoleManagementAlert': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}

	return nil
}
