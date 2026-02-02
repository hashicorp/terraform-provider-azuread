package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ActionResultPart = ForwardToChatResult{}

type ForwardToChatResult struct {
	// The chatMessage ID generated after a message is successfully forwarded to the target chat ID.
	ForwardedMessageId nullable.Type[string] `json:"forwardedMessageId,omitempty"`

	// The target chat ID where the message was forwarded.
	TargetChatId nullable.Type[string] `json:"targetChatId,omitempty"`

	// Fields inherited from ActionResultPart

	// The error that occurred, if any, during the course of the bulk operation.
	Error *PublicError `json:"error,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ForwardToChatResult) ActionResultPart() BaseActionResultPartImpl {
	return BaseActionResultPartImpl{
		Error:     s.Error,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ForwardToChatResult{}

func (s ForwardToChatResult) MarshalJSON() ([]byte, error) {
	type wrapper ForwardToChatResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ForwardToChatResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ForwardToChatResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.forwardToChatResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ForwardToChatResult: %+v", err)
	}

	return encoded, nil
}
