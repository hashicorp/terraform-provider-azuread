package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewRecommendationInsightSetting interface {
	AccessReviewRecommendationInsightSetting() BaseAccessReviewRecommendationInsightSettingImpl
}

var _ AccessReviewRecommendationInsightSetting = BaseAccessReviewRecommendationInsightSettingImpl{}

type BaseAccessReviewRecommendationInsightSettingImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessReviewRecommendationInsightSettingImpl) AccessReviewRecommendationInsightSetting() BaseAccessReviewRecommendationInsightSettingImpl {
	return s
}

var _ AccessReviewRecommendationInsightSetting = RawAccessReviewRecommendationInsightSettingImpl{}

// RawAccessReviewRecommendationInsightSettingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccessReviewRecommendationInsightSettingImpl struct {
	accessReviewRecommendationInsightSetting BaseAccessReviewRecommendationInsightSettingImpl
	Type                                     string
	Values                                   map[string]interface{}
}

func (s RawAccessReviewRecommendationInsightSettingImpl) AccessReviewRecommendationInsightSetting() BaseAccessReviewRecommendationInsightSettingImpl {
	return s.accessReviewRecommendationInsightSetting
}

func UnmarshalAccessReviewRecommendationInsightSettingImplementation(input []byte) (AccessReviewRecommendationInsightSetting, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewRecommendationInsightSetting into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPeerOutlierRecommendationInsightSettings") {
		var out GroupPeerOutlierRecommendationInsightSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPeerOutlierRecommendationInsightSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userLastSignInRecommendationInsightSetting") {
		var out UserLastSignInRecommendationInsightSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserLastSignInRecommendationInsightSetting: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessReviewRecommendationInsightSettingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessReviewRecommendationInsightSettingImpl: %+v", err)
	}

	return RawAccessReviewRecommendationInsightSettingImpl{
		accessReviewRecommendationInsightSetting: parent,
		Type:                                     value,
		Values:                                   temp,
	}, nil

}
