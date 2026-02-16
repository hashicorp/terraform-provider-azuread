package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GovernanceCriteria = GroupMembershipGovernanceCriteria{}

type GroupMembershipGovernanceCriteria struct {
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Fields inherited from GovernanceCriteria

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s GroupMembershipGovernanceCriteria) GovernanceCriteria() BaseGovernanceCriteriaImpl {
	return BaseGovernanceCriteriaImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupMembershipGovernanceCriteria{}

func (s GroupMembershipGovernanceCriteria) MarshalJSON() ([]byte, error) {
	type wrapper GroupMembershipGovernanceCriteria
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupMembershipGovernanceCriteria: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupMembershipGovernanceCriteria: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupMembershipGovernanceCriteria"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupMembershipGovernanceCriteria: %+v", err)
	}

	return encoded, nil
}
