package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnTokenIssuanceStartHandler = OnTokenIssuanceStartCustomExtensionHandler{}

type OnTokenIssuanceStartCustomExtensionHandler struct {
	Configuration   *CustomExtensionOverwriteConfiguration `json:"configuration,omitempty"`
	CustomExtension *OnTokenIssuanceStartCustomExtension   `json:"customExtension,omitempty"`

	// Fields inherited from OnTokenIssuanceStartHandler

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnTokenIssuanceStartCustomExtensionHandler) OnTokenIssuanceStartHandler() BaseOnTokenIssuanceStartHandlerImpl {
	return BaseOnTokenIssuanceStartHandlerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnTokenIssuanceStartCustomExtensionHandler{}

func (s OnTokenIssuanceStartCustomExtensionHandler) MarshalJSON() ([]byte, error) {
	type wrapper OnTokenIssuanceStartCustomExtensionHandler
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnTokenIssuanceStartCustomExtensionHandler: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnTokenIssuanceStartCustomExtensionHandler: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onTokenIssuanceStartCustomExtensionHandler"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnTokenIssuanceStartCustomExtensionHandler: %+v", err)
	}

	return encoded, nil
}
