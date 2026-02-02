package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingDlpRule struct {
	Actions           *[]DlpActionInfo    `json:"actions,omitempty"`
	IsMostRestrictive nullable.Type[bool] `json:"isMostRestrictive,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PolicyId   nullable.Type[string] `json:"policyId,omitempty"`
	PolicyName nullable.Type[string] `json:"policyName,omitempty"`
	Priority   nullable.Type[int64]  `json:"priority,omitempty"`
	RuleId     nullable.Type[string] `json:"ruleId,omitempty"`
	RuleMode   *RuleMode             `json:"ruleMode,omitempty"`
	RuleName   nullable.Type[string] `json:"ruleName,omitempty"`
}

var _ json.Unmarshaler = &MatchingDlpRule{}

func (s *MatchingDlpRule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsMostRestrictive nullable.Type[bool]   `json:"isMostRestrictive,omitempty"`
		ODataId           *string               `json:"@odata.id,omitempty"`
		ODataType         *string               `json:"@odata.type,omitempty"`
		PolicyId          nullable.Type[string] `json:"policyId,omitempty"`
		PolicyName        nullable.Type[string] `json:"policyName,omitempty"`
		Priority          nullable.Type[int64]  `json:"priority,omitempty"`
		RuleId            nullable.Type[string] `json:"ruleId,omitempty"`
		RuleMode          *RuleMode             `json:"ruleMode,omitempty"`
		RuleName          nullable.Type[string] `json:"ruleName,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsMostRestrictive = decoded.IsMostRestrictive
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PolicyId = decoded.PolicyId
	s.PolicyName = decoded.PolicyName
	s.Priority = decoded.Priority
	s.RuleId = decoded.RuleId
	s.RuleMode = decoded.RuleMode
	s.RuleName = decoded.RuleName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MatchingDlpRule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Actions into list []json.RawMessage: %+v", err)
		}

		output := make([]DlpActionInfo, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDlpActionInfoImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Actions' for 'MatchingDlpRule': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Actions = &output
	}

	return nil
}
