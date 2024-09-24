package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitedUserMessageInfo struct {
	// Additional recipients the invitation message should be sent to. Currently only one additional recipient is supported.
	CcRecipients *[]Recipient `json:"ccRecipients,omitempty"`

	// Customized message body you want to send if you don't want the default message.
	CustomizedMessageBody nullable.Type[string] `json:"customizedMessageBody,omitempty"`

	// The language you want to send the default message in. If the customizedMessageBody is specified, this property is
	// ignored, and the message is sent using the customizedMessageBody. The language format should be in ISO 639. The
	// default is en-US.
	MessageLanguage nullable.Type[string] `json:"messageLanguage,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &InvitedUserMessageInfo{}

func (s *InvitedUserMessageInfo) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CustomizedMessageBody nullable.Type[string] `json:"customizedMessageBody,omitempty"`
		MessageLanguage       nullable.Type[string] `json:"messageLanguage,omitempty"`
		ODataId               *string               `json:"@odata.id,omitempty"`
		ODataType             *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CustomizedMessageBody = decoded.CustomizedMessageBody
	s.MessageLanguage = decoded.MessageLanguage
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling InvitedUserMessageInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["ccRecipients"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CcRecipients into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CcRecipients' for 'InvitedUserMessageInfo': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CcRecipients = &output
	}

	return nil
}
