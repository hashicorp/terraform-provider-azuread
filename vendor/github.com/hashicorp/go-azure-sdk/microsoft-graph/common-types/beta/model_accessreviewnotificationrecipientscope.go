package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewNotificationRecipientScope interface {
	AccessReviewNotificationRecipientScope() BaseAccessReviewNotificationRecipientScopeImpl
}

var _ AccessReviewNotificationRecipientScope = BaseAccessReviewNotificationRecipientScopeImpl{}

type BaseAccessReviewNotificationRecipientScopeImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessReviewNotificationRecipientScopeImpl) AccessReviewNotificationRecipientScope() BaseAccessReviewNotificationRecipientScopeImpl {
	return s
}

var _ AccessReviewNotificationRecipientScope = RawAccessReviewNotificationRecipientScopeImpl{}

// RawAccessReviewNotificationRecipientScopeImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAccessReviewNotificationRecipientScopeImpl struct {
	accessReviewNotificationRecipientScope BaseAccessReviewNotificationRecipientScopeImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawAccessReviewNotificationRecipientScopeImpl) AccessReviewNotificationRecipientScope() BaseAccessReviewNotificationRecipientScopeImpl {
	return s.accessReviewNotificationRecipientScope
}

func (s RawAccessReviewNotificationRecipientScopeImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAccessReviewNotificationRecipientScopeImplementation(input []byte) (AccessReviewNotificationRecipientScope, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewNotificationRecipientScope into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewNotificationRecipientQueryScope") {
		var out AccessReviewNotificationRecipientQueryScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewNotificationRecipientQueryScope: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessReviewNotificationRecipientScopeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessReviewNotificationRecipientScopeImpl: %+v", err)
	}

	return RawAccessReviewNotificationRecipientScopeImpl{
		accessReviewNotificationRecipientScope: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
