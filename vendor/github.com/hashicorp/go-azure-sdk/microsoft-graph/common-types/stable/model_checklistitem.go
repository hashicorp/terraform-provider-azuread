package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ChecklistItem{}

type ChecklistItem struct {
	// The date and time when the checklistItem was finished.
	CheckedDateTime nullable.Type[string] `json:"checkedDateTime,omitempty"`

	// The date and time when the checklistItem was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Indicates the title of the checklistItem.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// State that indicates whether the item is checked off or not.
	IsChecked nullable.Type[bool] `json:"isChecked,omitempty"`

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

func (s ChecklistItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChecklistItem{}

func (s ChecklistItem) MarshalJSON() ([]byte, error) {
	type wrapper ChecklistItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChecklistItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChecklistItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.checklistItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChecklistItem: %+v", err)
	}

	return encoded, nil
}
