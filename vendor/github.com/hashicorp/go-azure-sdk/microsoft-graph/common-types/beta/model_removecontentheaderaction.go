package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InformationProtectionAction = RemoveContentHeaderAction{}

type RemoveContentHeaderAction struct {
	// The name of the UI element of the header to be removed.
	UiElementNames *[]string `json:"uiElementNames,omitempty"`

	// Fields inherited from InformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s RemoveContentHeaderAction) InformationProtectionAction() BaseInformationProtectionActionImpl {
	return BaseInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RemoveContentHeaderAction{}

func (s RemoveContentHeaderAction) MarshalJSON() ([]byte, error) {
	type wrapper RemoveContentHeaderAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RemoveContentHeaderAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RemoveContentHeaderAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.removeContentHeaderAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RemoveContentHeaderAction: %+v", err)
	}

	return encoded, nil
}
