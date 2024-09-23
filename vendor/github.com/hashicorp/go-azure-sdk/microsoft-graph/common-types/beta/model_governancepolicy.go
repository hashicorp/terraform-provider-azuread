package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernancePolicy struct {
	DecisionMakerCriteria *[]GovernanceCriteria         `json:"decisionMakerCriteria,omitempty"`
	NotificationPolicy    *GovernanceNotificationPolicy `json:"notificationPolicy,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &GovernancePolicy{}

func (s *GovernancePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		NotificationPolicy *GovernanceNotificationPolicy `json:"notificationPolicy,omitempty"`
		ODataId            *string                       `json:"@odata.id,omitempty"`
		ODataType          *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.NotificationPolicy = decoded.NotificationPolicy
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GovernancePolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["decisionMakerCriteria"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DecisionMakerCriteria into list []json.RawMessage: %+v", err)
		}

		output := make([]GovernanceCriteria, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalGovernanceCriteriaImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DecisionMakerCriteria' for 'GovernancePolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DecisionMakerCriteria = &output
	}

	return nil
}
