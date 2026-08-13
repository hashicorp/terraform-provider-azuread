package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewScope interface {
	AccessReviewScope() BaseAccessReviewScopeImpl
}

var _ AccessReviewScope = BaseAccessReviewScopeImpl{}

type BaseAccessReviewScopeImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessReviewScopeImpl) AccessReviewScope() BaseAccessReviewScopeImpl {
	return s
}

var _ AccessReviewScope = RawAccessReviewScopeImpl{}

// RawAccessReviewScopeImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAccessReviewScopeImpl struct {
	accessReviewScope BaseAccessReviewScopeImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawAccessReviewScopeImpl) AccessReviewScope() BaseAccessReviewScopeImpl {
	return s.accessReviewScope
}

func (s RawAccessReviewScopeImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAccessReviewScopeImplementation(input []byte) (AccessReviewScope, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewScope into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewQueryScope") {
		var out AccessReviewQueryScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewQueryScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewReviewerScope") {
		var out AccessReviewReviewerScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewReviewerScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.principalResourceMembershipsScope") {
		var out PrincipalResourceMembershipsScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrincipalResourceMembershipsScope: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessReviewScopeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessReviewScopeImpl: %+v", err)
	}

	return RawAccessReviewScopeImpl{
		accessReviewScope: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
