package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementAlertConfiguration interface {
	Entity
	UnifiedRoleManagementAlertConfiguration() BaseUnifiedRoleManagementAlertConfigurationImpl
}

var _ UnifiedRoleManagementAlertConfiguration = BaseUnifiedRoleManagementAlertConfigurationImpl{}

type BaseUnifiedRoleManagementAlertConfigurationImpl struct {
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

func (s BaseUnifiedRoleManagementAlertConfigurationImpl) UnifiedRoleManagementAlertConfiguration() BaseUnifiedRoleManagementAlertConfigurationImpl {
	return s
}

func (s BaseUnifiedRoleManagementAlertConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ UnifiedRoleManagementAlertConfiguration = RawUnifiedRoleManagementAlertConfigurationImpl{}

// RawUnifiedRoleManagementAlertConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawUnifiedRoleManagementAlertConfigurationImpl struct {
	unifiedRoleManagementAlertConfiguration BaseUnifiedRoleManagementAlertConfigurationImpl
	Type                                    string
	Values                                  map[string]interface{}
}

func (s RawUnifiedRoleManagementAlertConfigurationImpl) UnifiedRoleManagementAlertConfiguration() BaseUnifiedRoleManagementAlertConfigurationImpl {
	return s.unifiedRoleManagementAlertConfiguration
}

func (s RawUnifiedRoleManagementAlertConfigurationImpl) Entity() BaseEntityImpl {
	return s.unifiedRoleManagementAlertConfiguration.Entity()
}

var _ json.Marshaler = BaseUnifiedRoleManagementAlertConfigurationImpl{}

func (s BaseUnifiedRoleManagementAlertConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseUnifiedRoleManagementAlertConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseUnifiedRoleManagementAlertConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseUnifiedRoleManagementAlertConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementAlertConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseUnifiedRoleManagementAlertConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalUnifiedRoleManagementAlertConfigurationImplementation(input []byte) (UnifiedRoleManagementAlertConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementAlertConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.invalidLicenseAlertConfiguration") {
		var out InvalidLicenseAlertConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvalidLicenseAlertConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.noMfaOnRoleActivationAlertConfiguration") {
		var out NoMfaOnRoleActivationAlertConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoMfaOnRoleActivationAlertConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.redundantAssignmentAlertConfiguration") {
		var out RedundantAssignmentAlertConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RedundantAssignmentAlertConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration") {
		var out RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sequentialActivationRenewalsAlertConfiguration") {
		var out SequentialActivationRenewalsAlertConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SequentialActivationRenewalsAlertConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.staleSignInAlertConfiguration") {
		var out StaleSignInAlertConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StaleSignInAlertConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertConfiguration") {
		var out TooManyGlobalAdminsAssignedToTenantAlertConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TooManyGlobalAdminsAssignedToTenantAlertConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseUnifiedRoleManagementAlertConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUnifiedRoleManagementAlertConfigurationImpl: %+v", err)
	}

	return RawUnifiedRoleManagementAlertConfigurationImpl{
		unifiedRoleManagementAlertConfiguration: parent,
		Type:                                    value,
		Values:                                  temp,
	}, nil

}
