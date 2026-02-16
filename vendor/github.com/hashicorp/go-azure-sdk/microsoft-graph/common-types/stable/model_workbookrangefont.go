package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookRangeFont{}

type WorkbookRangeFont struct {
	// Inidicates whether the font is bold.
	Bold nullable.Type[bool] `json:"bold,omitempty"`

	// The HTML color code representation of the text color. For example, #FF0000 represents the color red.
	Color nullable.Type[string] `json:"color,omitempty"`

	// Inidicates whether the font is italic.
	Italic nullable.Type[bool] `json:"italic,omitempty"`

	// The font name. For example, 'Calibri'.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The type of underlining applied to the font. The possible values are: None, Single, Double, SingleAccountant,
	// DoubleAccountant.
	Underline nullable.Type[string] `json:"underline,omitempty"`

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

func (s WorkbookRangeFont) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookRangeFont{}

func (s WorkbookRangeFont) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookRangeFont
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookRangeFont: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookRangeFont: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookRangeFont"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookRangeFont: %+v", err)
	}

	return encoded, nil
}
