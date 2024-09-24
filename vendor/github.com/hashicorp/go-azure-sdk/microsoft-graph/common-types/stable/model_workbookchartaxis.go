package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookChartAxis{}

type WorkbookChartAxis struct {
	// Represents the formatting of a chart object, which includes line and font formatting. Read-only.
	Format *WorkbookChartAxisFormat `json:"format,omitempty"`

	// Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
	MajorGridlines *WorkbookChartGridlines `json:"majorGridlines,omitempty"`

	// Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
	MinorGridlines *WorkbookChartGridlines `json:"minorGridlines,omitempty"`

	// Represents the axis title. Read-only.
	Title *WorkbookChartAxisTitle `json:"title,omitempty"`

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

func (s WorkbookChartAxis) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookChartAxis{}

func (s WorkbookChartAxis) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookChartAxis
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookChartAxis: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookChartAxis: %+v", err)
	}

	delete(decoded, "format")
	delete(decoded, "majorGridlines")
	delete(decoded, "minorGridlines")
	delete(decoded, "title")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookChartAxis"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookChartAxis: %+v", err)
	}

	return encoded, nil
}
