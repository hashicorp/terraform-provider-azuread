package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectSet interface {
	SubjectSet() BaseSubjectSetImpl
}

var _ SubjectSet = BaseSubjectSetImpl{}

type BaseSubjectSetImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSubjectSetImpl) SubjectSet() BaseSubjectSetImpl {
	return s
}

var _ SubjectSet = RawSubjectSetImpl{}

// RawSubjectSetImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawSubjectSetImpl struct {
	subjectSet BaseSubjectSetImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawSubjectSetImpl) SubjectSet() BaseSubjectSetImpl {
	return s.subjectSet
}

func (s RawSubjectSetImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalSubjectSetImplementation(input []byte) (SubjectSet, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SubjectSet into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.groupBasedSubjectSet") {
		var out IdentityGovernanceGroupBasedSubjectSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceGroupBasedSubjectSet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.ruleBasedSubjectSet") {
		var out IdentityGovernanceRuleBasedSubjectSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceRuleBasedSubjectSet: %+v", err)
		}
		return out, nil
	}

	var parent BaseSubjectSetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSubjectSetImpl: %+v", err)
	}

	return RawSubjectSetImpl{
		subjectSet: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
