package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookRangeBorder{}

type WorkbookRangeBorder struct {
	// HTML color code representing the color of the border line, of the form #RRGGBB (for example 'FFA500') or as a named
	// HTML color (for example 'orange').
	Color nullable.Type[string] `json:"color,omitempty"`

	// Constant value that indicates the specific side of the border. Possible values are: EdgeTop, EdgeBottom, EdgeLeft,
	// EdgeRight, InsideVertical, InsideHorizontal, DiagonalDown, DiagonalUp. Read-only.
	SideIndex nullable.Type[string] `json:"sideIndex,omitempty"`

	// One of the constants of line style specifying the line style for the border. Possible values are: None, Continuous,
	// Dash, DashDot, DashDotDot, Dot, Double, SlantDashDot.
	Style nullable.Type[string] `json:"style,omitempty"`

	// Specifies the weight of the border around a range. Possible values are: Hairline, Thin, Medium, Thick.
	Weight nullable.Type[string] `json:"weight,omitempty"`

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

func (s WorkbookRangeBorder) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookRangeBorder{}

func (s WorkbookRangeBorder) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookRangeBorder
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookRangeBorder: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookRangeBorder: %+v", err)
	}

	delete(decoded, "sideIndex")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookRangeBorder"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookRangeBorder: %+v", err)
	}

	return encoded, nil
}
