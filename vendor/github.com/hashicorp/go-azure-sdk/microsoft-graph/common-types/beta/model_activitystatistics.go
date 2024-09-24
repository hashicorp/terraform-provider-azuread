package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityStatistics interface {
	Entity
	ActivityStatistics() BaseActivityStatisticsImpl
}

var _ ActivityStatistics = BaseActivityStatisticsImpl{}

type BaseActivityStatisticsImpl struct {
	// The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and
	// meeting.
	Activity *AnalyticsActivityType `json:"activity,omitempty"`

	// Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
	Duration *string `json:"duration,omitempty"`

	// Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could
	// be '2019-07-03' that follows the YYYY-MM-DD format.
	EndDate *string `json:"endDate,omitempty"`

	// Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value
	// could be '2019-07-04' that follows the YYYY-MM-DD format.
	StartDate *string `json:"startDate,omitempty"`

	// The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value
	// could be 'Pacific Standard Time.'
	TimeZoneUsed nullable.Type[string] `json:"timeZoneUsed,omitempty"`

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

func (s BaseActivityStatisticsImpl) ActivityStatistics() BaseActivityStatisticsImpl {
	return s
}

func (s BaseActivityStatisticsImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ActivityStatistics = RawActivityStatisticsImpl{}

// RawActivityStatisticsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawActivityStatisticsImpl struct {
	activityStatistics BaseActivityStatisticsImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawActivityStatisticsImpl) ActivityStatistics() BaseActivityStatisticsImpl {
	return s.activityStatistics
}

func (s RawActivityStatisticsImpl) Entity() BaseEntityImpl {
	return s.activityStatistics.Entity()
}

var _ json.Marshaler = BaseActivityStatisticsImpl{}

func (s BaseActivityStatisticsImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseActivityStatisticsImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseActivityStatisticsImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseActivityStatisticsImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.activityStatistics"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseActivityStatisticsImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalActivityStatisticsImplementation(input []byte) (ActivityStatistics, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ActivityStatistics into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.callActivityStatistics") {
		var out CallActivityStatistics
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallActivityStatistics: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatActivityStatistics") {
		var out ChatActivityStatistics
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatActivityStatistics: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.emailActivityStatistics") {
		var out EmailActivityStatistics
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailActivityStatistics: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.focusActivityStatistics") {
		var out FocusActivityStatistics
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FocusActivityStatistics: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingActivityStatistics") {
		var out MeetingActivityStatistics
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingActivityStatistics: %+v", err)
		}
		return out, nil
	}

	var parent BaseActivityStatisticsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseActivityStatisticsImpl: %+v", err)
	}

	return RawActivityStatisticsImpl{
		activityStatistics: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
