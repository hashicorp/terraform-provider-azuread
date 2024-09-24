package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityInformationProtectionAction = SecurityRemoveWatermarkAction{}

type SecurityRemoveWatermarkAction struct {
	// The name of the UI element of watermark to be removed.
	UiElementNames *[]string `json:"uiElementNames,omitempty"`

	// Fields inherited from SecurityInformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityRemoveWatermarkAction) SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl {
	return BaseSecurityInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityRemoveWatermarkAction{}

func (s SecurityRemoveWatermarkAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityRemoveWatermarkAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityRemoveWatermarkAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityRemoveWatermarkAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.removeWatermarkAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityRemoveWatermarkAction: %+v", err)
	}

	return encoded, nil
}
