package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TodoTaskList{}

type TodoTaskList struct {
	// The name of the task list.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The collection of open extensions defined for the task list. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// True if the user is owner of the given task list.
	IsOwner *bool `json:"isOwner,omitempty"`

	// True if the task list is shared with other users
	IsShared *bool `json:"isShared,omitempty"`

	// The tasks in this task list. Read-only. Nullable.
	Tasks *[]TodoTask `json:"tasks,omitempty"`

	WellknownListName *WellknownListName `json:"wellknownListName,omitempty"`

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

func (s TodoTaskList) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TodoTaskList{}

func (s TodoTaskList) MarshalJSON() ([]byte, error) {
	type wrapper TodoTaskList
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TodoTaskList: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TodoTaskList: %+v", err)
	}

	delete(decoded, "tasks")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.todoTaskList"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TodoTaskList: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TodoTaskList{}

func (s *TodoTaskList) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName       nullable.Type[string] `json:"displayName,omitempty"`
		IsOwner           *bool                 `json:"isOwner,omitempty"`
		IsShared          *bool                 `json:"isShared,omitempty"`
		Tasks             *[]TodoTask           `json:"tasks,omitempty"`
		WellknownListName *WellknownListName    `json:"wellknownListName,omitempty"`
		Id                *string               `json:"id,omitempty"`
		ODataId           *string               `json:"@odata.id,omitempty"`
		ODataType         *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.IsOwner = decoded.IsOwner
	s.IsShared = decoded.IsShared
	s.Tasks = decoded.Tasks
	s.WellknownListName = decoded.WellknownListName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TodoTaskList into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'TodoTaskList': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	return nil
}
