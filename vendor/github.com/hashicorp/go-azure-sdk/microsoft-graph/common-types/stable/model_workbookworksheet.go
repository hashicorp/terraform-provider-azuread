package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookWorksheet{}

type WorkbookWorksheet struct {
	// Returns collection of charts that are part of the worksheet. Read-only.
	Charts *[]WorkbookChart `json:"charts,omitempty"`

	// The display name of the worksheet.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Returns collection of names that are associated with the worksheet. Read-only.
	Names *[]WorkbookNamedItem `json:"names,omitempty"`

	// Collection of PivotTables that are part of the worksheet.
	PivotTables *[]WorkbookPivotTable `json:"pivotTables,omitempty"`

	// The zero-based position of the worksheet within the workbook.
	Position *int64 `json:"position,omitempty"`

	// Returns sheet protection object for a worksheet. Read-only.
	Protection *WorkbookWorksheetProtection `json:"protection,omitempty"`

	// Collection of tables that are part of the worksheet. Read-only.
	Tables *[]WorkbookTable `json:"tables,omitempty"`

	// The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
	Visibility *string `json:"visibility,omitempty"`

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

func (s WorkbookWorksheet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookWorksheet{}

func (s WorkbookWorksheet) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookWorksheet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookWorksheet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookWorksheet: %+v", err)
	}

	delete(decoded, "charts")
	delete(decoded, "names")
	delete(decoded, "protection")
	delete(decoded, "tables")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookWorksheet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookWorksheet: %+v", err)
	}

	return encoded, nil
}
