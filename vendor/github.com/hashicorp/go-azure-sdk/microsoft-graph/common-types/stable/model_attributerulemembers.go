package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectSet = AttributeRuleMembers{}

type AttributeRuleMembers struct {
	// A description of the membership rule.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Determines the allowed target users for this policy. For more information about the syntax of the membership rule,
	// see Membership Rules syntax.
	MembershipRule nullable.Type[string] `json:"membershipRule,omitempty"`

	// Fields inherited from SubjectSet

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AttributeRuleMembers) SubjectSet() BaseSubjectSetImpl {
	return BaseSubjectSetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AttributeRuleMembers{}

func (s AttributeRuleMembers) MarshalJSON() ([]byte, error) {
	type wrapper AttributeRuleMembers
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AttributeRuleMembers: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AttributeRuleMembers: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attributeRuleMembers"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AttributeRuleMembers: %+v", err)
	}

	return encoded, nil
}
