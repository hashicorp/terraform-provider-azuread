package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceCriteria interface {
	GovernanceCriteria() BaseGovernanceCriteriaImpl
}

var _ GovernanceCriteria = BaseGovernanceCriteriaImpl{}

type BaseGovernanceCriteriaImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseGovernanceCriteriaImpl) GovernanceCriteria() BaseGovernanceCriteriaImpl {
	return s
}

var _ GovernanceCriteria = RawGovernanceCriteriaImpl{}

// RawGovernanceCriteriaImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGovernanceCriteriaImpl struct {
	governanceCriteria BaseGovernanceCriteriaImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawGovernanceCriteriaImpl) GovernanceCriteria() BaseGovernanceCriteriaImpl {
	return s.governanceCriteria
}

func UnmarshalGovernanceCriteriaImplementation(input []byte) (GovernanceCriteria, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GovernanceCriteria into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupMembershipGovernanceCriteria") {
		var out GroupMembershipGovernanceCriteria
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupMembershipGovernanceCriteria: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.roleMembershipGovernanceCriteria") {
		var out RoleMembershipGovernanceCriteria
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleMembershipGovernanceCriteria: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userGovernanceCriteria") {
		var out UserGovernanceCriteria
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserGovernanceCriteria: %+v", err)
		}
		return out, nil
	}

	var parent BaseGovernanceCriteriaImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGovernanceCriteriaImpl: %+v", err)
	}

	return RawGovernanceCriteriaImpl{
		governanceCriteria: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
