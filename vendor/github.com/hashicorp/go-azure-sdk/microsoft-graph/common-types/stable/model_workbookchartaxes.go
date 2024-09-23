package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartAxes{}

type WorkbookChartAxes struct {
	// Represents the category axis in a chart. Read-only.
	CategoryAxis *WorkbookChartAxis `json:"categoryAxis,omitempty"`

	// Represents the series axis of a 3-dimensional chart. Read-only.
	SeriesAxis *WorkbookChartAxis `json:"seriesAxis,omitempty"`

	// Represents the value axis in an axis. Read-only.
	ValueAxis *WorkbookChartAxis `json:"valueAxis,omitempty"`

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

func (s WorkbookChartAxes) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartAxes{}

func (s WorkbookChartAxes) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartAxes
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartAxes: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartAxes: %+v", err)
	}

	delete(decoded, "categoryAxis")
	delete(decoded, "seriesAxis")
	delete(decoded, "valueAxis")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartAxes"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartAxes: %+v", err)
	}

	return encoded, nil
}
