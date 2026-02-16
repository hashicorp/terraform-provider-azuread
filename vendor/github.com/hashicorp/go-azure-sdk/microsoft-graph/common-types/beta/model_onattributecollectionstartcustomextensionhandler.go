package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnAttributeCollectionStartHandler = OnAttributeCollectionStartCustomExtensionHandler{}

type OnAttributeCollectionStartCustomExtensionHandler struct {
	// Configuration regarding properties of the custom extension that are can be overwritten per event listener.
	Configuration *CustomExtensionOverwriteConfiguration `json:"configuration,omitempty"`

	CustomExtension *OnAttributeCollectionStartCustomExtension `json:"customExtension,omitempty"`

	// Fields inherited from OnAttributeCollectionStartHandler

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnAttributeCollectionStartCustomExtensionHandler) OnAttributeCollectionStartHandler() BaseOnAttributeCollectionStartHandlerImpl {
	return BaseOnAttributeCollectionStartHandlerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnAttributeCollectionStartCustomExtensionHandler{}

func (s OnAttributeCollectionStartCustomExtensionHandler) MarshalJSON() ([]byte, error) {
	type wrapper OnAttributeCollectionStartCustomExtensionHandler
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnAttributeCollectionStartCustomExtensionHandler: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnAttributeCollectionStartCustomExtensionHandler: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onAttributeCollectionStartCustomExtensionHandler"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnAttributeCollectionStartCustomExtensionHandler: %+v", err)
	}

	return encoded, nil
}
