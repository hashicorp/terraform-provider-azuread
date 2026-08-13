package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsInsightValue interface {
	UserExperienceAnalyticsInsightValue() BaseUserExperienceAnalyticsInsightValueImpl
}

var _ UserExperienceAnalyticsInsightValue = BaseUserExperienceAnalyticsInsightValueImpl{}

type BaseUserExperienceAnalyticsInsightValueImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseUserExperienceAnalyticsInsightValueImpl) UserExperienceAnalyticsInsightValue() BaseUserExperienceAnalyticsInsightValueImpl {
	return s
}

var _ UserExperienceAnalyticsInsightValue = RawUserExperienceAnalyticsInsightValueImpl{}

// RawUserExperienceAnalyticsInsightValueImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawUserExperienceAnalyticsInsightValueImpl struct {
	userExperienceAnalyticsInsightValue BaseUserExperienceAnalyticsInsightValueImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawUserExperienceAnalyticsInsightValueImpl) UserExperienceAnalyticsInsightValue() BaseUserExperienceAnalyticsInsightValueImpl {
	return s.userExperienceAnalyticsInsightValue
}

func (s RawUserExperienceAnalyticsInsightValueImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalUserExperienceAnalyticsInsightValueImplementation(input []byte) (UserExperienceAnalyticsInsightValue, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsInsightValue into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.insightValueDouble") {
		var out InsightValueDouble
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InsightValueDouble: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.insightValueInt") {
		var out InsightValueInt
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InsightValueInt: %+v", err)
		}
		return out, nil
	}

	var parent BaseUserExperienceAnalyticsInsightValueImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUserExperienceAnalyticsInsightValueImpl: %+v", err)
	}

	return RawUserExperienceAnalyticsInsightValueImpl{
		userExperienceAnalyticsInsightValue: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
