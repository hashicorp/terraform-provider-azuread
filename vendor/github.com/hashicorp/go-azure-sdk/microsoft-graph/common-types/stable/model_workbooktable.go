package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookTable{}

type WorkbookTable struct {
	// The list of all the columns in the table. Read-only.
	Columns *[]WorkbookTableColumn `json:"columns,omitempty"`

	// Indicates whether the first column contains special formatting.
	HighlightFirstColumn *bool `json:"highlightFirstColumn,omitempty"`

	// Indicates whether the last column contains special formatting.
	HighlightLastColumn *bool `json:"highlightLastColumn,omitempty"`

	// A legacy identifier used in older Excel clients. The value of the identifier remains the same even when the table is
	// renamed. This property should be interpreted as an opaque string value and shouldn't be parsed to any other type.
	// Read-only.
	LegacyId nullable.Type[string] `json:"legacyId,omitempty"`

	// The name of the table.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The list of all the rows in the table. Read-only.
	Rows *[]WorkbookTableRow `json:"rows,omitempty"`

	// Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones
	// to make reading the table easier.
	ShowBandedColumns *bool `json:"showBandedColumns,omitempty"`

	// Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to
	// make reading the table easier.
	ShowBandedRows *bool `json:"showBandedRows,omitempty"`

	// Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if
	// the table contains a header row.
	ShowFilterButton *bool `json:"showFilterButton,omitempty"`

	// Indicates whether the header row is visible or not. This value can be set to show or remove the header row.
	ShowHeaders *bool `json:"showHeaders,omitempty"`

	// Indicates whether the total row is visible or not. This value can be set to show or remove the total row.
	ShowTotals *bool `json:"showTotals,omitempty"`

	// The sorting for the table. Read-only.
	Sort *WorkbookTableSort `json:"sort,omitempty"`

	// A constant value that represents the Table style. Possible values are: TableStyleLight1 through TableStyleLight21,
	// TableStyleMedium1 through TableStyleMedium28, TableStyleStyleDark1 through TableStyleStyleDark11. A custom
	// user-defined style present in the workbook can also be specified.
	Style nullable.Type[string] `json:"style,omitempty"`

	// The worksheet containing the current table. Read-only.
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

func (s WorkbookTable) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookTable{}

func (s WorkbookTable) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookTable
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookTable: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookTable: %+v", err)
	}

	delete(decoded, "columns")
	delete(decoded, "legacyId")
	delete(decoded, "rows")
	delete(decoded, "sort")
	delete(decoded, "worksheet")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookTable"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookTable: %+v", err)
	}

	return encoded, nil
}
