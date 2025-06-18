package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSecurityAndGovernance = TenantDataSecurityAndGovernance{}

type TenantDataSecurityAndGovernance struct {
	ProtectionScopes *TenantProtectionScopeContainer `json:"protectionScopes,omitempty"`

	// Fields inherited from DataSecurityAndGovernance

	SensitivityLabels *[]SensitivityLabel `json:"sensitivityLabels,omitempty"`

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

func (s TenantDataSecurityAndGovernance) DataSecurityAndGovernance() BaseDataSecurityAndGovernanceImpl {
	return BaseDataSecurityAndGovernanceImpl{
		SensitivityLabels: s.SensitivityLabels,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s TenantDataSecurityAndGovernance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TenantDataSecurityAndGovernance{}

func (s TenantDataSecurityAndGovernance) MarshalJSON() ([]byte, error) {
	type wrapper TenantDataSecurityAndGovernance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TenantDataSecurityAndGovernance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TenantDataSecurityAndGovernance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.tenantDataSecurityAndGovernance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TenantDataSecurityAndGovernance: %+v", err)
	}

	return encoded, nil
}
