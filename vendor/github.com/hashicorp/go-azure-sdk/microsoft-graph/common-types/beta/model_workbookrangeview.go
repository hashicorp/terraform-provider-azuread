package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookRangeView{}

type WorkbookRangeView struct {
	// Represents the cell addresses
	CellAddresses *Json `json:"cellAddresses,omitempty"`

	// The number of visible columns. Read-only.
	ColumnCount *int64 `json:"columnCount,omitempty"`

	// The formula in A1-style notation.
	Formulas *Json `json:"formulas,omitempty"`

	// The formula in A1-style notation, in the user's language and number-formatting locale. For example, the English
	// '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
	FormulasLocal *Json `json:"formulasLocal,omitempty"`

	// Represents the formula in R1C1-style notation.
	FormulasR1C1 *Json `json:"formulasR1C1,omitempty"`

	// The index of the range.
	Index *int64 `json:"index,omitempty"`

	// Excel's number format code for the given cell. Read-only.
	NumberFormat *Json `json:"numberFormat,omitempty"`

	// The number of visible rows. Read-only.
	RowCount *int64 `json:"rowCount,omitempty"`

	// The collection of range views associated with the range. Read-only. Read-only.
	Rows *[]WorkbookRangeView `json:"rows,omitempty"`

	// The text values of the specified range. The Text value won't depend on the cell width. The # sign substitution that
	// happens in Excel UI won't affect the text value returned by the API. Read-only.
	Text *Json `json:"text,omitempty"`

	// The type of data of each cell. Read-only. Possible values are: Unknown, Empty, String, Integer, Double, Boolean,
	// Error.
	ValueTypes *Json `json:"valueTypes,omitempty"`

	// The raw values of the specified range view. The data returned could be of type string, number, or a Boolean. Cell
	// that contains an error returns the error string.
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
	delete(decoded, "numberFormat")
	delete(decoded, "rowCount")
	delete(decoded, "rows")
	delete(decoded, "text")
	delete(decoded, "valueTypes")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookRangeView"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookRangeView: %+v", err)
	}

	return encoded, nil
}
