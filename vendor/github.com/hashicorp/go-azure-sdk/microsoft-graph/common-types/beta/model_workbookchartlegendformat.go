package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartLegendFormat{}

type WorkbookChartLegendFormat struct {
	// Represents the fill format of an object, which includes background formating information. Read-only.
	Fill *WorkbookChartFill `json:"fill,omitempty"`

	// Represents the font attributes such as font name, font size, color, etc. of a chart legend. Read-only.
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

func (s WorkbookChartLegendFormat) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartLegendFormat{}

func (s WorkbookChartLegendFormat) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartLegendFormat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartLegendFormat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartLegendFormat: %+v", err)
	}

	delete(decoded, "fill")
	delete(decoded, "font")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartLegendFormat"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartLegendFormat: %+v", err)
	}

	return encoded, nil
}
