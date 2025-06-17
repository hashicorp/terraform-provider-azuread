package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookRangeFormat{}

type WorkbookRangeFormat struct {
	// Collection of border objects that apply to the overall range selected Read-only.
	Borders *[]WorkbookRangeBorder `json:"borders,omitempty"`

	// Returns the fill object defined on the overall range. Read-only.
	Fill *WorkbookRangeFill `json:"fill,omitempty"`

	// Returns the font object defined on the overall range selected Read-only.
	Font *WorkbookRangeFont `json:"font,omitempty"`

	// The horizontal alignment for the specified object. Possible values are: General, Left, Center, Right, Fill, Justify,
	// CenterAcrossSelection, Distributed.
	HorizontalAlignment nullable.Type[string] `json:"horizontalAlignment,omitempty"`

	// Returns the format protection object for a range. Read-only.
	Protection *WorkbookFormatProtection `json:"protection,omitempty"`

	// The vertical alignment for the specified object. Possible values are: Top, Center, Bottom, Justify, Distributed.
	VerticalAlignment nullable.Type[string] `json:"verticalAlignment,omitempty"`

	// Indicates whether Excel wraps the text in the object. A null value indicates that the entire range doesn't have a
	// uniform wrap setting.
	WrapText nullable.Type[bool] `json:"wrapText,omitempty"`

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

func (s WorkbookRangeFormat) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookRangeFormat{}

func (s WorkbookRangeFormat) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookRangeFormat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookRangeFormat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookRangeFormat: %+v", err)
	}

	delete(decoded, "borders")
	delete(decoded, "fill")
	delete(decoded, "font")
	delete(decoded, "protection")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookRangeFormat"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookRangeFormat: %+v", err)
	}

	return encoded, nil
}
