package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OutlookUser{}

type OutlookUser struct {
	// A list of categories defined for the user.
	MasterCategories *[]OutlookCategory `json:"masterCategories,omitempty"`

	// The user's Outlook task folders. Read-only. Nullable.
	TaskFolders *[]OutlookTaskFolder `json:"taskFolders,omitempty"`

	// The user's Outlook task groups. Read-only. Nullable.
	TaskGroups *[]OutlookTaskGroup `json:"taskGroups,omitempty"`

	// The user's Outlook tasks. Read-only. Nullable.
	Tasks *[]OutlookTask `json:"tasks,omitempty"`

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

func (s OutlookUser) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OutlookUser{}

func (s OutlookUser) MarshalJSON() ([]byte, error) {
	type wrapper OutlookUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OutlookUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OutlookUser: %+v", err)
	}

	delete(decoded, "taskFolders")
	delete(decoded, "taskGroups")
	delete(decoded, "tasks")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.outlookUser"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OutlookUser: %+v", err)
	}

	return encoded, nil
}
