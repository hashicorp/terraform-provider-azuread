package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartDataLabelFormat{}

type WorkbookChartDataLabelFormat struct {
	// Represents the fill format of the current chart data label. Read-only.
	Fill *WorkbookChartFill `json:"fill,omitempty"`

	// Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
	Font *WorkbookChartFont `json:"font,omitempty"`

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

func (s WorkbookChartDataLabelFormat) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartDataLabelFormat{}

func (s WorkbookChartDataLabelFormat) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartDataLabelFormat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartDataLabelFormat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartDataLabelFormat: %+v", err)
	}

	delete(decoded, "fill")
	delete(decoded, "font")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartDataLabelFormat"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartDataLabelFormat: %+v", err)
	}

	return encoded, nil
}
