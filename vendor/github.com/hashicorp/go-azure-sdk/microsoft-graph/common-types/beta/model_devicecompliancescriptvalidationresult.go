package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptValidationResult struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Errors in json for the script for rules.
	RuleErrors *[]DeviceComplianceScriptRuleError `json:"ruleErrors,omitempty"`

	// Parsed rules from json.
	Rules *[]DeviceComplianceScriptRule `json:"rules,omitempty"`

	// Errors in json for the script.
	ScriptErrors *[]DeviceComplianceScriptError `json:"scriptErrors,omitempty"`
}

var _ json.Unmarshaler = &DeviceComplianceScriptValidationResult{}

func (s *DeviceComplianceScriptValidationResult) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId    *string                            `json:"@odata.id,omitempty"`
		ODataType  *string                            `json:"@odata.type,omitempty"`
		RuleErrors *[]DeviceComplianceScriptRuleError `json:"ruleErrors,omitempty"`
		Rules      *[]DeviceComplianceScriptRule      `json:"rules,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RuleErrors = decoded.RuleErrors
	s.Rules = decoded.Rules

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceComplianceScriptValidationResult into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["scriptErrors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ScriptErrors into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceComplianceScriptError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceComplianceScriptErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ScriptErrors' for 'DeviceComplianceScriptValidationResult': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ScriptErrors = &output
	}

	return nil
}
