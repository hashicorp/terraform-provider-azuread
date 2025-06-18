package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnOtpSendHandler = OnOtpSendCustomExtensionHandler{}

type OnOtpSendCustomExtensionHandler struct {
	// Configuration regarding properties of the custom extension that are can be overwritten for the onEmailOtpSendListener
	// event listener.
	Configuration *CustomExtensionOverwriteConfiguration `json:"configuration,omitempty"`

	CustomExtension *OnOtpSendCustomExtension `json:"customExtension,omitempty"`

	// Fields inherited from OnOtpSendHandler

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnOtpSendCustomExtensionHandler) OnOtpSendHandler() BaseOnOtpSendHandlerImpl {
	return BaseOnOtpSendHandlerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnOtpSendCustomExtensionHandler{}

func (s OnOtpSendCustomExtensionHandler) MarshalJSON() ([]byte, error) {
	type wrapper OnOtpSendCustomExtensionHandler
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnOtpSendCustomExtensionHandler: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnOtpSendCustomExtensionHandler: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onOtpSendCustomExtensionHandler"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnOtpSendCustomExtensionHandler: %+v", err)
	}

	return encoded, nil
}
