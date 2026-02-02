package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChart{}

type WorkbookChart struct {
	// Represents chart axes. Read-only.
	Axes *WorkbookChartAxes `json:"axes,omitempty"`

	// Represents the datalabels on the chart. Read-only.
	DataLabels *WorkbookChartDataLabels `json:"dataLabels,omitempty"`

	// Encapsulates the format properties for the chart area. Read-only.
	Format *WorkbookChartAreaFormat `json:"format,omitempty"`

	// Represents the legend for the chart. Read-only.
	Legend *WorkbookChartLegend `json:"legend,omitempty"`

	// Represents the name of a chart object.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Represents either a single series or collection of series in the chart. Read-only.
	Series *[]WorkbookChartSeries `json:"series,omitempty"`

	// Represents the title of the specified chart, including the text, visibility, position and formating of the title.
	// Read-only.
	Title *WorkbookChartTitle `json:"title,omitempty"`

	// The worksheet containing the current chart. Read-only.
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

func (s WorkbookChart) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChart{}

func (s WorkbookChart) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChart
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChart: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChart: %+v", err)
	}

	delete(decoded, "axes")
	delete(decoded, "dataLabels")
	delete(decoded, "format")
	delete(decoded, "legend")
	delete(decoded, "series")
	delete(decoded, "title")
	delete(decoded, "worksheet")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChart"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChart: %+v", err)
	}

	return encoded, nil
}
