package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MarkContent = AddHeader{}

type AddHeader struct {
	Alignment *Alignment           `json:"alignment,omitempty"`
	Margin    nullable.Type[int64] `json:"margin,omitempty"`

	// Fields inherited from MarkContent

	FontColor nullable.Type[string] `json:"fontColor,omitempty"`
	FontSize  nullable.Type[int64]  `json:"fontSize,omitempty"`
	Text      nullable.Type[string] `json:"text,omitempty"`

	// Fields inherited from LabelActionBase

	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AddHeader) MarkContent() BaseMarkContentImpl {
	return BaseMarkContentImpl{
		FontColor: s.FontColor,
		FontSize:  s.FontSize,
		Text:      s.Text,
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AddHeader) LabelActionBase() BaseLabelActionBaseImpl {
	return BaseLabelActionBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AddHeader{}

func (s AddHeader) MarshalJSON() ([]byte, error) {
	type wrapper AddHeader
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddHeader: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddHeader: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.addHeader"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddHeader: %+v", err)
	}

	return encoded, nil
}
