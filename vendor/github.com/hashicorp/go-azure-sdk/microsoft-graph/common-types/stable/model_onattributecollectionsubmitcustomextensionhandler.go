package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnAttributeCollectionSubmitHandler = OnAttributeCollectionSubmitCustomExtensionHandler{}

type OnAttributeCollectionSubmitCustomExtensionHandler struct {
	// Configuration regarding properties of the custom extension that can be overwritten per event listener.
	Configuration *CustomExtensionOverwriteConfiguration `json:"configuration,omitempty"`

	CustomExtension *OnAttributeCollectionSubmitCustomExtension `json:"customExtension,omitempty"`

	// Fields inherited from OnAttributeCollectionSubmitHandler

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnAttributeCollectionSubmitCustomExtensionHandler) OnAttributeCollectionSubmitHandler() BaseOnAttributeCollectionSubmitHandlerImpl {
	return BaseOnAttributeCollectionSubmitHandlerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnAttributeCollectionSubmitCustomExtensionHandler{}

func (s OnAttributeCollectionSubmitCustomExtensionHandler) MarshalJSON() ([]byte, error) {
	type wrapper OnAttributeCollectionSubmitCustomExtensionHandler
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnAttributeCollectionSubmitCustomExtensionHandler: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnAttributeCollectionSubmitCustomExtensionHandler: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onAttributeCollectionSubmitCustomExtensionHandler"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnAttributeCollectionSubmitCustomExtensionHandler: %+v", err)
	}

	return encoded, nil
}
