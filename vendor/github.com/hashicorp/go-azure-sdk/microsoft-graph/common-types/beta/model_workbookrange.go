package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookRange{}

type WorkbookRange struct {
	// Represents the range reference in A1-style. Address value contains the Sheet reference (for example, Sheet1!A1:B4).
	// Read-only.
	Address nullable.Type[string] `json:"address,omitempty"`

	// Represents range reference for the specified range in the language of the user. Read-only.
	AddressLocal nullable.Type[string] `json:"addressLocal,omitempty"`

	// Number of cells in the range. Read-only.
	CellCount *int64 `json:"cellCount,omitempty"`

	// Represents the total number of columns in the range. Read-only.
	ColumnCount *int64 `json:"columnCount,omitempty"`

	// Indicates whether all columns of the current range are hidden.
	ColumnHidden nullable.Type[bool] `json:"columnHidden,omitempty"`

	// Represents the column number of the first cell in the range. Zero-indexed. Read-only.
	ColumnIndex *int64 `json:"columnIndex,omitempty"`

	// Returns a format object, encapsulating the range's font, fill, borders, alignment, and other properties. Read-only.
	Format *WorkbookRangeFormat `json:"format,omitempty"`

	// Represents the formula in A1-style notation.
	Formulas *Json `json:"formulas,omitempty"`

	// Represents the formula in A1-style notation, in the user's language and number-formatting locale. For example, the
	// English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
	FormulasLocal *Json `json:"formulasLocal,omitempty"`

	// Represents the formula in R1C1-style notation.
	FormulasR1C1 *Json `json:"formulasR1C1,omitempty"`

	// Represents if all cells of the current range are hidden. Read-only.
	Hidden nullable.Type[bool] `json:"hidden,omitempty"`

	// Represents Excel's number format code for the given cell.
	NumberFormat *Json `json:"numberFormat,omitempty"`

	// Returns the total number of rows in the range. Read-only.
	RowCount *int64 `json:"rowCount,omitempty"`

	// Indicates whether all rows of the current range are hidden.
	RowHidden nullable.Type[bool] `json:"rowHidden,omitempty"`

	// Returns the row number of the first cell in the range. Zero-indexed. Read-only.
	RowIndex *int64 `json:"rowIndex,omitempty"`

	// The worksheet containing the current range. Read-only.
	Sort *WorkbookRangeSort `json:"sort,omitempty"`

	// Text values of the specified range. The text value doesn't depend on the cell width. The # sign substitution that
	// happens in Excel UI doesn't affect the text value returned by the API. Read-only.
	Text *Json `json:"text,omitempty"`

	// Represents the type of data of each cell. Possible values are: Unknown, Empty, String, Integer, Double, Boolean,
	// Error. Read-only.
	ValueTypes *Json `json:"valueTypes,omitempty"`

	// Represents the raw values of the specified range. The data returned could be of type string, number, or a Boolean.
	// Cell that contains an error returns the error string.
	Values *Json `json:"values,omitempty"`

	// The worksheet containing the current range. Read-only.
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

func (s WorkbookRange) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookRange{}

func (s WorkbookRange) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookRange
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookRange: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookRange: %+v", err)
	}

	delete(decoded, "address")
	delete(decoded, "addressLocal")
	delete(decoded, "cellCount")
	delete(decoded, "columnCount")
	delete(decoded, "columnIndex")
	delete(decoded, "format")
	delete(decoded, "hidden")
	delete(decoded, "rowCount")
	delete(decoded, "rowIndex")
	delete(decoded, "sort")
	delete(decoded, "text")
	delete(decoded, "valueTypes")
	delete(decoded, "worksheet")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookRange"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookRange: %+v", err)
	}

	return encoded, nil
}
