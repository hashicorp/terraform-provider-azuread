package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Workbook{}

type Workbook struct {
	Application *WorkbookApplication `json:"application,omitempty"`
	Comments    *[]WorkbookComment   `json:"comments,omitempty"`
	Functions   *WorkbookFunctions   `json:"functions,omitempty"`

	// Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
	Names *[]WorkbookNamedItem `json:"names,omitempty"`

	// The status of Workbook operations. Getting an operation collection is not supported, but you can get the status of a
	// long-running operation if the Location header is returned in the response. Read-only. Nullable.
	Operations *[]WorkbookOperation `json:"operations,omitempty"`

	// Represents a collection of tables associated with the workbook. Read-only.
	Tables *[]WorkbookTable `json:"tables,omitempty"`

	// Represents a collection of worksheets associated with the workbook. Read-only.
	Worksheets *[]WorkbookWorksheet `json:"worksheets,omitempty"`

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

func (s Workbook) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Workbook{}

func (s Workbook) MarshalJSON() ([]byte, error) {
	type wrapper Workbook
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Workbook: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Workbook: %+v", err)
	}

	delete(decoded, "names")
	delete(decoded, "operations")
	delete(decoded, "tables")
	delete(decoded, "worksheets")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbook"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Workbook: %+v", err)
	}

	return encoded, nil
}
