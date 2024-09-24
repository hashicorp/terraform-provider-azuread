package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
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

// RawSubjectSetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSubjectSetImpl struct {
	subjectSet BaseSubjectSetImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawSubjectSetImpl) SubjectSet() BaseSubjectSetImpl {
	return s.subjectSet
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

	if strings.EqualFold(value, "#microsoft.graph.attributeRuleMembers") {
		var out AttributeRuleMembers
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttributeRuleMembers: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.connectedOrganizationMembers") {
		var out ConnectedOrganizationMembers
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectedOrganizationMembers: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalSponsors") {
		var out ExternalSponsors
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalSponsors: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupMembers") {
		var out GroupMembers
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupMembers: %+v", err)
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

	if strings.EqualFold(value, "#microsoft.graph.internalSponsors") {
		var out InternalSponsors
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InternalSponsors: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.requestorManager") {
		var out RequestorManager
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RequestorManager: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.singleServicePrincipal") {
		var out SingleServicePrincipal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleServicePrincipal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.singleUser") {
		var out SingleUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetApplicationOwners") {
		var out TargetApplicationOwners
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetApplicationOwners: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetManager") {
		var out TargetManager
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetManager: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetUserSponsors") {
		var out TargetUserSponsors
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetUserSponsors: %+v", err)
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
