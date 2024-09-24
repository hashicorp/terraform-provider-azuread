package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OutlookTaskFolder{}

type OutlookTaskFolder struct {
	// The version of the task folder.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// True if the folder is the default task folder.
	IsDefaultFolder nullable.Type[bool] `json:"isDefaultFolder,omitempty"`

	// The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The name of the task folder.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The unique GUID identifier for the task folder's parent group.
	ParentGroupKey nullable.Type[string] `json:"parentGroupKey,omitempty"`

	// The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The tasks in this task folder. Read-only. Nullable.
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

func (s OutlookTaskFolder) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OutlookTaskFolder{}

func (s OutlookTaskFolder) MarshalJSON() ([]byte, error) {
	type wrapper OutlookTaskFolder
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OutlookTaskFolder: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OutlookTaskFolder: %+v", err)
	}

	delete(decoded, "multiValueExtendedProperties")
	delete(decoded, "singleValueExtendedProperties")
	delete(decoded, "tasks")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.outlookTaskFolder"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OutlookTaskFolder: %+v", err)
	}

	return encoded, nil
}
