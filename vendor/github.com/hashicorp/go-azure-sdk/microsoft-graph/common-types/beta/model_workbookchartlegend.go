package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartLegend{}

type WorkbookChartLegend struct {
	// Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
	Format *WorkbookChartLegendFormat `json:"format,omitempty"`

	// Indicates whether the chart legend should overlap with the main body of the chart.
	Overlay nullable.Type[bool] `json:"overlay,omitempty"`

	// Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner,
	// Custom.
	Position nullable.Type[string] `json:"position,omitempty"`

	// Indicates whether the chart legend is visible.
	Visible *bool `json:"visible,omitempty"`

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

func (s WorkbookChartLegend) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartLegend{}

func (s WorkbookChartLegend) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartLegend
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartLegend: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartLegend: %+v", err)
	}

	delete(decoded, "format")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartLegend"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartLegend: %+v", err)
	}

	return encoded, nil
}
