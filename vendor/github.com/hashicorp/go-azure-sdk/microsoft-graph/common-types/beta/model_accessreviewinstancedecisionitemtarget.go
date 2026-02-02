package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItemTarget interface {
	AccessReviewInstanceDecisionItemTarget() BaseAccessReviewInstanceDecisionItemTargetImpl
}

var _ AccessReviewInstanceDecisionItemTarget = BaseAccessReviewInstanceDecisionItemTargetImpl{}

type BaseAccessReviewInstanceDecisionItemTargetImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessReviewInstanceDecisionItemTargetImpl) AccessReviewInstanceDecisionItemTarget() BaseAccessReviewInstanceDecisionItemTargetImpl {
	return s
}

var _ AccessReviewInstanceDecisionItemTarget = RawAccessReviewInstanceDecisionItemTargetImpl{}

// RawAccessReviewInstanceDecisionItemTargetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccessReviewInstanceDecisionItemTargetImpl struct {
	accessReviewInstanceDecisionItemTarget BaseAccessReviewInstanceDecisionItemTargetImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawAccessReviewInstanceDecisionItemTargetImpl) AccessReviewInstanceDecisionItemTarget() BaseAccessReviewInstanceDecisionItemTargetImpl {
	return s.accessReviewInstanceDecisionItemTarget
}

func UnmarshalAccessReviewInstanceDecisionItemTargetImplementation(input []byte) (AccessReviewInstanceDecisionItemTarget, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItemTarget into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstanceDecisionItemServicePrincipalTarget") {
		var out AccessReviewInstanceDecisionItemServicePrincipalTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstanceDecisionItemServicePrincipalTarget: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInstanceDecisionItemUserTarget") {
		var out AccessReviewInstanceDecisionItemUserTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInstanceDecisionItemUserTarget: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessReviewInstanceDecisionItemTargetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessReviewInstanceDecisionItemTargetImpl: %+v", err)
	}

	return RawAccessReviewInstanceDecisionItemTargetImpl{
		accessReviewInstanceDecisionItemTarget: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
