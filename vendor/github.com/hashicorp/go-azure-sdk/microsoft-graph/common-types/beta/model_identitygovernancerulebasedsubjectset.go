package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectSet = IdentityGovernanceRuleBasedSubjectSet{}

type IdentityGovernanceRuleBasedSubjectSet struct {
	// The rule for the subject set. Lifecycle Workflows supports a rich set of user properties for configuring the rules
	// using $filter query expressions. For more information, see supported user and query parameters.
	Rule *string `json:"rule,omitempty"`

	// Fields inherited from SubjectSet

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceRuleBasedSubjectSet) SubjectSet() BaseSubjectSetImpl {
	return BaseSubjectSetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceRuleBasedSubjectSet{}

func (s IdentityGovernanceRuleBasedSubjectSet) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceRuleBasedSubjectSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceRuleBasedSubjectSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceRuleBasedSubjectSet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.ruleBasedSubjectSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceRuleBasedSubjectSet: %+v", err)
	}

	return encoded, nil
}
