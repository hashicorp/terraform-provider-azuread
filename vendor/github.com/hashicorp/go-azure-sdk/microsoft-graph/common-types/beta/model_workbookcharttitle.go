package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartTitle{}

type WorkbookChartTitle struct {
	// The formatting of a chart title, which includes fill and font formatting. Read-only.
	Format *WorkbookChartTitleFormat `json:"format,omitempty"`

	// Indicates whether the chart title will overlay the chart or not.
	Overlay nullable.Type[bool] `json:"overlay,omitempty"`

	// The title text of the chart.
	Text nullable.Type[string] `json:"text,omitempty"`

	// Indicates whether the chart title is visible.
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

func (s WorkbookChartTitle) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartTitle{}

func (s WorkbookChartTitle) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartTitle
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartTitle: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartTitle: %+v", err)
	}

	delete(decoded, "format")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartTitle"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartTitle: %+v", err)
	}

	return encoded, nil
}
