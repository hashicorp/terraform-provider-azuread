package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartTitleFormat{}

type WorkbookChartTitleFormat struct {
	// Represents the fill format of an object, which includes background formatting information. Read-only.
	Fill *WorkbookChartFill `json:"fill,omitempty"`

	// Represents the font attributes (font name, font size, color, etc.) for the current object. Read-only.
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

func (s WorkbookChartTitleFormat) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartTitleFormat{}

func (s WorkbookChartTitleFormat) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartTitleFormat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartTitleFormat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartTitleFormat: %+v", err)
	}

	delete(decoded, "fill")
	delete(decoded, "font")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartTitleFormat"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartTitleFormat: %+v", err)
	}

	return encoded, nil
}
