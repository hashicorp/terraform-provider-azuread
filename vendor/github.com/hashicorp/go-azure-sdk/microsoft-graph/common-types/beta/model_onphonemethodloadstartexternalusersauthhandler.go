package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnPhoneMethodLoadStartHandler = OnPhoneMethodLoadStartExternalUsersAuthHandler{}

type OnPhoneMethodLoadStartExternalUsersAuthHandler struct {
	SmsOptions   *PhoneOptions `json:"smsOptions,omitempty"`
	VoiceOptions *PhoneOptions `json:"voiceOptions,omitempty"`

	// Fields inherited from OnPhoneMethodLoadStartHandler

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnPhoneMethodLoadStartExternalUsersAuthHandler) OnPhoneMethodLoadStartHandler() BaseOnPhoneMethodLoadStartHandlerImpl {
	return BaseOnPhoneMethodLoadStartHandlerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnPhoneMethodLoadStartExternalUsersAuthHandler{}

func (s OnPhoneMethodLoadStartExternalUsersAuthHandler) MarshalJSON() ([]byte, error) {
	type wrapper OnPhoneMethodLoadStartExternalUsersAuthHandler
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPhoneMethodLoadStartExternalUsersAuthHandler: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPhoneMethodLoadStartExternalUsersAuthHandler: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onPhoneMethodLoadStartExternalUsersAuthHandler"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPhoneMethodLoadStartExternalUsersAuthHandler: %+v", err)
	}

	return encoded, nil
}
