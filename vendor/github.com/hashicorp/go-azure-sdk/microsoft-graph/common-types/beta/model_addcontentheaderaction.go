package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InformationProtectionAction = AddContentHeaderAction{}

type AddContentHeaderAction struct {
	Alignment *ContentAlignment `json:"alignment,omitempty"`

	// Color of the font to use for the header.
	FontColor nullable.Type[string] `json:"fontColor,omitempty"`

	// Name of the font to use for the header.
	FontName nullable.Type[string] `json:"fontName,omitempty"`

	// Font size to use for the header.
	FontSize *int64 `json:"fontSize,omitempty"`

	// The margin of the header from the top of the document.
	Margin *int64 `json:"margin,omitempty"`

	// The contents of the header itself.
	Text nullable.Type[string] `json:"text,omitempty"`

	// The name of the UI element where the header should be placed.
	UiElementName nullable.Type[string] `json:"uiElementName,omitempty"`

	// Fields inherited from InformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AddContentHeaderAction) InformationProtectionAction() BaseInformationProtectionActionImpl {
	return BaseInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AddContentHeaderAction{}

func (s AddContentHeaderAction) MarshalJSON() ([]byte, error) {
	type wrapper AddContentHeaderAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddContentHeaderAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddContentHeaderAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.addContentHeaderAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddContentHeaderAction: %+v", err)
	}

	return encoded, nil
}
