package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OutlookTaskGroup{}

type OutlookTaskGroup struct {
	// The version of the task group.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// The unique GUID identifier for the task group.
	GroupKey nullable.Type[string] `json:"groupKey,omitempty"`

	// True if the task group is the default task group.
	IsDefaultGroup nullable.Type[bool] `json:"isDefaultGroup,omitempty"`

	// The name of the task group.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The collection of task folders in the task group. Read-only. Nullable.
	TaskFolders *[]OutlookTaskFolder `json:"taskFolders,omitempty"`

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

func (s OutlookTaskGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OutlookTaskGroup{}

func (s OutlookTaskGroup) MarshalJSON() ([]byte, error) {
	type wrapper OutlookTaskGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OutlookTaskGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OutlookTaskGroup: %+v", err)
	}

	delete(decoded, "taskFolders")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.outlookTaskGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OutlookTaskGroup: %+v", err)
	}

	return encoded, nil
}
