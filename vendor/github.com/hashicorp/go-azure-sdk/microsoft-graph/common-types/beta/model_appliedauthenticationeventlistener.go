package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedAuthenticationEventListener struct {
	// The type of authentication event that triggered the custom authentication extension request. The possible values are:
	// tokenIssuanceStart, pageRenderStart, unknownFutureValue, attributeCollectionStart, attributeCollectionSubmit,
	// emailOtpSend. Use the Prefer: include-unknown-enum-members request header to get the following values in this
	// evolvable enum: attributeCollectionStart, attributeCollectionSubmit, emailOtpSend.
	EventType *AuthenticationEventType `json:"eventType,omitempty"`

	// ID of the authentication event listener that was executed.
	ExecutedListenerId nullable.Type[string] `json:"executedListenerId,omitempty"`

	// The result from the listening client, such as an Azure Logic App and Azure Functions, of this authentication event.
	HandlerResult AuthenticationEventHandlerResult `json:"handlerResult"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &AppliedAuthenticationEventListener{}

func (s *AppliedAuthenticationEventListener) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EventType          *AuthenticationEventType `json:"eventType,omitempty"`
		ExecutedListenerId nullable.Type[string]    `json:"executedListenerId,omitempty"`
		ODataId            *string                  `json:"@odata.id,omitempty"`
		ODataType          *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EventType = decoded.EventType
	s.ExecutedListenerId = decoded.ExecutedListenerId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AppliedAuthenticationEventListener into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["handlerResult"]; ok {
		impl, err := UnmarshalAuthenticationEventHandlerResultImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'HandlerResult' for 'AppliedAuthenticationEventListener': %+v", err)
		}
		s.HandlerResult = impl
	}

	return nil
}
