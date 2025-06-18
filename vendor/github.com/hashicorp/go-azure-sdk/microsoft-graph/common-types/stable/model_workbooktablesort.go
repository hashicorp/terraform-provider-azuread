package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookTableSort{}

type WorkbookTableSort struct {
	// The list of the current conditions last used to sort the table. Read-only.
	Fields *[]WorkbookSortField `json:"fields,omitempty"`

	// Indicates whether the casing impacted the last sort of the table. Read-only.
	MatchCase *bool `json:"matchCase,omitempty"`

	// The Chinese character ordering method last used to sort the table. The possible values are: PinYin, StrokeCount.
	// Read-only.
	Method *string `json:"method,omitempty"`

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

func (s WorkbookTableSort) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookTableSort{}

func (s WorkbookTableSort) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookTableSort
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookTableSort: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookTableSort: %+v", err)
	}

	delete(decoded, "fields")
	delete(decoded, "matchCase")
	delete(decoded, "method")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookTableSort"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookTableSort: %+v", err)
	}

	return encoded, nil
}
