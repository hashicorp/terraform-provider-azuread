package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookTableColumn{}

type WorkbookTableColumn struct {
	// The filter applied to the column. Read-only.
	Filter *WorkbookFilter `json:"filter,omitempty"`

	// The index of the column within the columns collection of the table. Zero-indexed. Read-only.
	Index *int64 `json:"index,omitempty"`

	// The name of the table column.
	Name nullable.Type[string] `json:"name,omitempty"`

	// TRepresents the raw values of the specified range. The data returned could be of type string, number, or a Boolean.
	// Cell that contain an error will return the error string.
	Values *Json `json:"values,omitempty"`

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

func (s WorkbookTableColumn) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookTableColumn{}

func (s WorkbookTableColumn) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookTableColumn
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookTableColumn: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookTableColumn: %+v", err)
	}

	delete(decoded, "filter")
	delete(decoded, "index")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookTableColumn"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookTableColumn: %+v", err)
	}

	return encoded, nil
}
