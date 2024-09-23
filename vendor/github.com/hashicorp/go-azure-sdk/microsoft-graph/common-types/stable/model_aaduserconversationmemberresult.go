package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ActionResultPart = AadUserConversationMemberResult{}

type AadUserConversationMemberResult struct {
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s AadUserConversationMemberResult) ActionResultPart() BaseActionResultPartImpl {
	return BaseActionResultPartImpl{
		Error:     s.Error,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AadUserConversationMemberResult{}

func (s AadUserConversationMemberResult) MarshalJSON() ([]byte, error) {
	type wrapper AadUserConversationMemberResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AadUserConversationMemberResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AadUserConversationMemberResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aadUserConversationMemberResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AadUserConversationMemberResult: %+v", err)
	}

	return encoded, nil
}
