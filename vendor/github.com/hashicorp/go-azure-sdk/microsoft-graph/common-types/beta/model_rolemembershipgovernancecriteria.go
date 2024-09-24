package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GovernanceCriteria = RoleMembershipGovernanceCriteria{}

type RoleMembershipGovernanceCriteria struct {
	RoleId         nullable.Type[string] `json:"roleId,omitempty"`
	RoleTemplateId nullable.Type[string] `json:"roleTemplateId,omitempty"`

	// Fields inherited from GovernanceCriteria

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s RoleMembershipGovernanceCriteria) GovernanceCriteria() BaseGovernanceCriteriaImpl {
	return BaseGovernanceCriteriaImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RoleMembershipGovernanceCriteria{}

func (s RoleMembershipGovernanceCriteria) MarshalJSON() ([]byte, error) {
	type wrapper RoleMembershipGovernanceCriteria
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RoleMembershipGovernanceCriteria: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RoleMembershipGovernanceCriteria: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.roleMembershipGovernanceCriteria"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RoleMembershipGovernanceCriteria: %+v", err)
	}

	return encoded, nil
}
