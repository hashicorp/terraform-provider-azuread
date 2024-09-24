package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookRangeView{}

type WorkbookRangeView struct {
	// Returns the number of visible columns. Read-only.
	ColumnCount *int64 `json:"columnCount,omitempty"`

	// Index of the range.
	Index *int64 `json:"index,omitempty"`

	// Returns the number of visible rows. Read-only.
	RowCount *int64 `json:"rowCount,omitempty"`

	// Represents a collection of range views associated with the range. Read-only. Read-only.
	Rows *[]WorkbookRangeView `json:"rows,omitempty"`

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

func (s WorkbookRangeView) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookRangeView{}

func (s WorkbookRangeView) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookRangeView
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookRangeView: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookRangeView: %+v", err)
	}

	delete(decoded, "columnCount")
	delete(decoded, "rowCount")
	delete(decoded, "rows")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookRangeView"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookRangeView: %+v", err)
	}

	return encoded, nil
}
