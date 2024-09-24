package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartDataLabels{}

type WorkbookChartDataLabels struct {
	// Represents the format of chart data labels, which includes fill and font formatting. Read-only.
	Format *WorkbookChartDataLabelFormat `json:"format,omitempty"`

	// DataLabelPosition value that represents the position of the data label. The possible values are: None, Center,
	// InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
	Position nullable.Type[string] `json:"position,omitempty"`

	// String representing the separator used for the data labels on a chart.
	Separator nullable.Type[string] `json:"separator,omitempty"`

	// Boolean value representing if the data label bubble size is visible or not.
	ShowBubbleSize nullable.Type[bool] `json:"showBubbleSize,omitempty"`

	// Boolean value representing if the data label category name is visible or not.
	ShowCategoryName nullable.Type[bool] `json:"showCategoryName,omitempty"`

	// Boolean value representing if the data label legend key is visible or not.
	ShowLegendKey nullable.Type[bool] `json:"showLegendKey,omitempty"`

	// Boolean value representing if the data label percentage is visible or not.
	ShowPercentage nullable.Type[bool] `json:"showPercentage,omitempty"`

	// Boolean value representing if the data label series name is visible or not.
	ShowSeriesName nullable.Type[bool] `json:"showSeriesName,omitempty"`

	// Boolean value representing if the data label value is visible or not.
	ShowValue nullable.Type[bool] `json:"showValue,omitempty"`

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

func (s WorkbookChartDataLabels) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartDataLabels{}

func (s WorkbookChartDataLabels) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartDataLabels
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartDataLabels: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartDataLabels: %+v", err)
	}

	delete(decoded, "format")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartDataLabels"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartDataLabels: %+v", err)
	}

	return encoded, nil
}
