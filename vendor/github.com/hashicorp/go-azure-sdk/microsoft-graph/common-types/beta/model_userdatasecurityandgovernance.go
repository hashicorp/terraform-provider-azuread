package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSecurityAndGovernance = UserDataSecurityAndGovernance{}

type UserDataSecurityAndGovernance struct {
	// Container for activity logs (content processing and audit) related to this user. ContainsTarget: true.
	Activities *ActivitiesContainer `json:"activities,omitempty"`

	ProtectionScopes *UserProtectionScopeContainer `json:"protectionScopes,omitempty"`

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

func (s UserDataSecurityAndGovernance) DataSecurityAndGovernance() BaseDataSecurityAndGovernanceImpl {
	return BaseDataSecurityAndGovernanceImpl{
		SensitivityLabels: s.SensitivityLabels,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s UserDataSecurityAndGovernance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserDataSecurityAndGovernance{}

func (s UserDataSecurityAndGovernance) MarshalJSON() ([]byte, error) {
	type wrapper UserDataSecurityAndGovernance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserDataSecurityAndGovernance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserDataSecurityAndGovernance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userDataSecurityAndGovernance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserDataSecurityAndGovernance: %+v", err)
	}

	return encoded, nil
}
