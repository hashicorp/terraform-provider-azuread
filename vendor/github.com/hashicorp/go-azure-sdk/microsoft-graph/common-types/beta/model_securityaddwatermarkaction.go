package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityInformationProtectionAction = SecurityAddWatermarkAction{}

type SecurityAddWatermarkAction struct {
	// Color of the font to use for the watermark.
	FontColor nullable.Type[string] `json:"fontColor,omitempty"`

	// Name of the font to use for the watermark.
	FontName nullable.Type[string] `json:"fontName,omitempty"`

	// Font size to use for the watermark.
	FontSize *int64 `json:"fontSize,omitempty"`

	Layout *SecurityWatermarkLayout `json:"layout,omitempty"`

	// The contents of the watermark itself.
	Text nullable.Type[string] `json:"text,omitempty"`

	// The name of the UI element where the watermark should be placed.
	UiElementName nullable.Type[string] `json:"uiElementName,omitempty"`

	// Fields inherited from SecurityInformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityAddWatermarkAction) SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl {
	return BaseSecurityInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityAddWatermarkAction{}

func (s SecurityAddWatermarkAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityAddWatermarkAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityAddWatermarkAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAddWatermarkAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.addWatermarkAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityAddWatermarkAction: %+v", err)
	}

	return encoded, nil
}
