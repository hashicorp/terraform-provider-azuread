package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegeEscalation{}

type PrivilegeEscalation struct {
	// The list of actions that the identity could perform.
	Actions *[]AuthorizationSystemTypeAction `json:"actions,omitempty"`

	// A detailed description of the privilege escalation.
	Description *string `json:"description,omitempty"`

	// The name of the policy that defines the escalation
	DisplayName *string `json:"displayName,omitempty"`

	// The list of resources that the identity could perform actions on.
	Resources *[]AuthorizationSystemResource `json:"resources,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PrivilegeEscalation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegeEscalation{}

func (s PrivilegeEscalation) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegeEscalation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegeEscalation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegeEscalation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegeEscalation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegeEscalation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrivilegeEscalation{}

func (s *PrivilegeEscalation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description *string `json:"description,omitempty"`
		DisplayName *string `json:"displayName,omitempty"`
		Id          *string `json:"id,omitempty"`
		ODataId     *string `json:"@odata.id,omitempty"`
		ODataType   *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PrivilegeEscalation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Actions into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthorizationSystemTypeAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthorizationSystemTypeActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Actions' for 'PrivilegeEscalation': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Actions = &output
	}

	if v, ok := temp["resources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Resources into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthorizationSystemResource, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthorizationSystemResourceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Resources' for 'PrivilegeEscalation': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Resources = &output
	}

	return nil
}
