package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementAlertConfiguration = NoMfaOnRoleActivationAlertConfiguration{}

type NoMfaOnRoleActivationAlertConfiguration struct {

	// Fields inherited from UnifiedRoleManagementAlertConfiguration

	// The definition of the alert that contains its description, impact, and measures to mitigate or prevent it. Supports
	// $expand.
	AlertDefinition *UnifiedRoleManagementAlertDefinition `json:"alertDefinition,omitempty"`

	// The identifier of an alert definition. Supports $filter (eq, ne).
	AlertDefinitionId nullable.Type[string] `json:"alertDefinitionId,omitempty"`

	// true if the alert is enabled. Setting it to false disables PIM scanning the tenant to identify instances that trigger
	// the alert.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The identifier of the scope to which the alert is related. Only / is supported to represent the tenant scope.
	// Supports $filter (eq, ne).
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

func (s NoMfaOnRoleActivationAlertConfiguration) UnifiedRoleManagementAlertConfiguration() BaseUnifiedRoleManagementAlertConfigurationImpl {
	return BaseUnifiedRoleManagementAlertConfigurationImpl{
		AlertDefinition:   s.AlertDefinition,
		AlertDefinitionId: s.AlertDefinitionId,
		IsEnabled:         s.IsEnabled,
		ScopeId:           s.ScopeId,
		ScopeType:         s.ScopeType,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s NoMfaOnRoleActivationAlertConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NoMfaOnRoleActivationAlertConfiguration{}

func (s NoMfaOnRoleActivationAlertConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper NoMfaOnRoleActivationAlertConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NoMfaOnRoleActivationAlertConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NoMfaOnRoleActivationAlertConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.noMfaOnRoleActivationAlertConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NoMfaOnRoleActivationAlertConfiguration: %+v", err)
	}

	return encoded, nil
}
