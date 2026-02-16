package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookNamedItem{}

type WorkbookNamedItem struct {
	// The comment associated with this name.
	Comment nullable.Type[string] `json:"comment,omitempty"`

	// The name of the object. Read-only.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only.
	Scope *string `json:"scope,omitempty"`

	// The type of reference is associated with the name. Possible values are: String, Integer, Double, Boolean, Range.
	// Read-only.
	Type nullable.Type[string] `json:"type,omitempty"`

	// The formula that the name is defined to refer to. For example, =Sheet14!$B$2:$H$12 and =4.75. Read-only.
	Value *Json `json:"value,omitempty"`

	// Indicates whether the object is visible.
	Visible *bool `json:"visible,omitempty"`

	// Returns the worksheet to which the named item is scoped. Available only if the item is scoped to the worksheet.
	// Read-only.
	Worksheet *WorkbookWorksheet `json:"worksheet,omitempty"`

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

func (s WorkbookNamedItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookNamedItem{}

func (s WorkbookNamedItem) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookNamedItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookNamedItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookNamedItem: %+v", err)
	}

	delete(decoded, "name")
	delete(decoded, "scope")
	delete(decoded, "type")
	delete(decoded, "value")
	delete(decoded, "worksheet")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookNamedItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookNamedItem: %+v", err)
	}

	return encoded, nil
}
