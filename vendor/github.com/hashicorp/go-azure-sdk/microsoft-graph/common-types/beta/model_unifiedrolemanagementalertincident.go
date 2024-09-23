package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementAlertIncident interface {
	Entity
	UnifiedRoleManagementAlertIncident() BaseUnifiedRoleManagementAlertIncidentImpl
}

var _ UnifiedRoleManagementAlertIncident = BaseUnifiedRoleManagementAlertIncidentImpl{}

type BaseUnifiedRoleManagementAlertIncidentImpl struct {

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

func (s BaseUnifiedRoleManagementAlertIncidentImpl) UnifiedRoleManagementAlertIncident() BaseUnifiedRoleManagementAlertIncidentImpl {
	return s
}

func (s BaseUnifiedRoleManagementAlertIncidentImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ UnifiedRoleManagementAlertIncident = RawUnifiedRoleManagementAlertIncidentImpl{}

// RawUnifiedRoleManagementAlertIncidentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawUnifiedRoleManagementAlertIncidentImpl struct {
	unifiedRoleManagementAlertIncident BaseUnifiedRoleManagementAlertIncidentImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawUnifiedRoleManagementAlertIncidentImpl) UnifiedRoleManagementAlertIncident() BaseUnifiedRoleManagementAlertIncidentImpl {
	return s.unifiedRoleManagementAlertIncident
}

func (s RawUnifiedRoleManagementAlertIncidentImpl) Entity() BaseEntityImpl {
	return s.unifiedRoleManagementAlertIncident.Entity()
}

var _ json.Marshaler = BaseUnifiedRoleManagementAlertIncidentImpl{}

func (s BaseUnifiedRoleManagementAlertIncidentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseUnifiedRoleManagementAlertIncidentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseUnifiedRoleManagementAlertIncidentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseUnifiedRoleManagementAlertIncidentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementAlertIncident"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseUnifiedRoleManagementAlertIncidentImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalUnifiedRoleManagementAlertIncidentImplementation(input []byte) (UnifiedRoleManagementAlertIncident, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementAlertIncident into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.invalidLicenseAlertIncident") {
		var out InvalidLicenseAlertIncident
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvalidLicenseAlertIncident: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.noMfaOnRoleActivationAlertIncident") {
		var out NoMfaOnRoleActivationAlertIncident
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoMfaOnRoleActivationAlertIncident: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.redundantAssignmentAlertIncident") {
		var out RedundantAssignmentAlertIncident
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RedundantAssignmentAlertIncident: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertIncident") {
		var out RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sequentialActivationRenewalsAlertIncident") {
		var out SequentialActivationRenewalsAlertIncident
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SequentialActivationRenewalsAlertIncident: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.staleSignInAlertIncident") {
		var out StaleSignInAlertIncident
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StaleSignInAlertIncident: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertIncident") {
		var out TooManyGlobalAdminsAssignedToTenantAlertIncident
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TooManyGlobalAdminsAssignedToTenantAlertIncident: %+v", err)
		}
		return out, nil
	}

	var parent BaseUnifiedRoleManagementAlertIncidentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUnifiedRoleManagementAlertIncidentImpl: %+v", err)
	}

	return RawUnifiedRoleManagementAlertIncidentImpl{
		unifiedRoleManagementAlertIncident: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
