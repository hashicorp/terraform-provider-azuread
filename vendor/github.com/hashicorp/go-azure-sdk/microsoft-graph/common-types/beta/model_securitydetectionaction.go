package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDetectionAction struct {
	AlertTemplate *SecurityAlertTemplate `json:"alertTemplate,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Groups to which the custom detection rule applies.
	OrganizationalScope *SecurityOrganizationalScope `json:"organizationalScope,omitempty"`

	// Actions taken on impacted assets as set in the custom detection rule.
	ResponseActions *[]SecurityResponseAction `json:"responseActions,omitempty"`
}

var _ json.Unmarshaler = &SecurityDetectionAction{}

func (s *SecurityDetectionAction) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AlertTemplate       *SecurityAlertTemplate       `json:"alertTemplate,omitempty"`
		ODataId             *string                      `json:"@odata.id,omitempty"`
		ODataType           *string                      `json:"@odata.type,omitempty"`
		OrganizationalScope *SecurityOrganizationalScope `json:"organizationalScope,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AlertTemplate = decoded.AlertTemplate
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OrganizationalScope = decoded.OrganizationalScope

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityDetectionAction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["responseActions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ResponseActions into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityResponseAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityResponseActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ResponseActions' for 'SecurityDetectionAction': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ResponseActions = &output
	}

	return nil
}
