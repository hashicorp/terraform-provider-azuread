package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MarkContent = AddFooter{}

type AddFooter struct {
	Alignment *Alignment           `json:"alignment,omitempty"`
	Margin    nullable.Type[int64] `json:"margin,omitempty"`

	// Fields inherited from MarkContent

	FontColor nullable.Type[string] `json:"fontColor,omitempty"`
	FontSize  nullable.Type[int64]  `json:"fontSize,omitempty"`
	Text      nullable.Type[string] `json:"text,omitempty"`

	// Fields inherited from LabelActionBase

	// The name of the action (for example, 'Encrypt', 'AddHeader').
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AddFooter) MarkContent() BaseMarkContentImpl {
	return BaseMarkContentImpl{
		FontColor: s.FontColor,
		FontSize:  s.FontSize,
		Text:      s.Text,
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AddFooter) LabelActionBase() BaseLabelActionBaseImpl {
	return BaseLabelActionBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AddFooter{}

func (s AddFooter) MarshalJSON() ([]byte, error) {
	type wrapper AddFooter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddFooter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddFooter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.addFooter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddFooter: %+v", err)
	}

	return encoded, nil
}
