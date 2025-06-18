package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSecurityAndGovernance interface {
	Entity
	DataSecurityAndGovernance() BaseDataSecurityAndGovernanceImpl
}

var _ DataSecurityAndGovernance = BaseDataSecurityAndGovernanceImpl{}

type BaseDataSecurityAndGovernanceImpl struct {
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

func (s BaseDataSecurityAndGovernanceImpl) DataSecurityAndGovernance() BaseDataSecurityAndGovernanceImpl {
	return s
}

func (s BaseDataSecurityAndGovernanceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DataSecurityAndGovernance = RawDataSecurityAndGovernanceImpl{}

// RawDataSecurityAndGovernanceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataSecurityAndGovernanceImpl struct {
	dataSecurityAndGovernance BaseDataSecurityAndGovernanceImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawDataSecurityAndGovernanceImpl) DataSecurityAndGovernance() BaseDataSecurityAndGovernanceImpl {
	return s.dataSecurityAndGovernance
}

func (s RawDataSecurityAndGovernanceImpl) Entity() BaseEntityImpl {
	return s.dataSecurityAndGovernance.Entity()
}

var _ json.Marshaler = BaseDataSecurityAndGovernanceImpl{}

func (s BaseDataSecurityAndGovernanceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDataSecurityAndGovernanceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDataSecurityAndGovernanceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDataSecurityAndGovernanceImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.dataSecurityAndGovernance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDataSecurityAndGovernanceImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDataSecurityAndGovernanceImplementation(input []byte) (DataSecurityAndGovernance, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataSecurityAndGovernance into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.tenantDataSecurityAndGovernance") {
		var out TenantDataSecurityAndGovernance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TenantDataSecurityAndGovernance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userDataSecurityAndGovernance") {
		var out UserDataSecurityAndGovernance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserDataSecurityAndGovernance: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataSecurityAndGovernanceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataSecurityAndGovernanceImpl: %+v", err)
	}

	return RawDataSecurityAndGovernanceImpl{
		dataSecurityAndGovernance: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
