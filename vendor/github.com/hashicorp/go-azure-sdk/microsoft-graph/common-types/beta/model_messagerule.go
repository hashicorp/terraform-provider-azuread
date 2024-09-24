package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MessageRule{}

type MessageRule struct {
	// Actions to be taken on a message when the corresponding conditions are fulfilled.
	Actions *MessageRuleActions `json:"actions,omitempty"`

	// Conditions that when fulfilled trigger the corresponding actions for that rule.
	Conditions *MessageRulePredicates `json:"conditions,omitempty"`

	// The display name of the rule.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Exception conditions for the rule.
	Exceptions *MessageRulePredicates `json:"exceptions,omitempty"`

	// Indicates whether the rule is in an error condition. Read-only.
	HasError nullable.Type[bool] `json:"hasError,omitempty"`

	// Indicates whether the rule is enabled to be applied to messages.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
	IsReadOnly nullable.Type[bool] `json:"isReadOnly,omitempty"`

	// Indicates the order in which the rule is executed, among other rules.
	Sequence nullable.Type[int64] `json:"sequence,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MessageRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MessageRule{}

func (s MessageRule) MarshalJSON() ([]byte, error) {
	type wrapper MessageRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MessageRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MessageRule: %+v", err)
	}

	delete(decoded, "hasError")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.messageRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MessageRule: %+v", err)
	}

	return encoded, nil
}
