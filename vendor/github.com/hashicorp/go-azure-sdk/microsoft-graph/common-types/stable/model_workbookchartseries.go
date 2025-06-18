package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartSeries{}

type WorkbookChartSeries struct {
	// The formatting of a chart series, which includes fill and line formatting. Read-only.
	Format *WorkbookChartSeriesFormat `json:"format,omitempty"`

	// The name of a series in a chart.
	Name nullable.Type[string] `json:"name,omitempty"`

	// A collection of all points in the series. Read-only.
	Points *[]WorkbookChartPoint `json:"points,omitempty"`

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

func (s WorkbookChartSeries) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartSeries{}

func (s WorkbookChartSeries) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartSeries
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartSeries: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartSeries: %+v", err)
	}

	delete(decoded, "format")
	delete(decoded, "points")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartSeries"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartSeries: %+v", err)
	}

	return encoded, nil
}
