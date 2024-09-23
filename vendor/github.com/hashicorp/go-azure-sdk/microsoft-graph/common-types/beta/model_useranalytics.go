package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserAnalytics{}

type UserAnalytics struct {
	// The collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
	ActivityStatistics *[]ActivityStatistics `json:"activityStatistics,omitempty"`

	// The current settings for a user to use the analytics API.
	Settings *Settings `json:"settings,omitempty"`

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

func (s UserAnalytics) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserAnalytics{}

func (s UserAnalytics) MarshalJSON() ([]byte, error) {
	type wrapper UserAnalytics
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserAnalytics: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserAnalytics: %+v", err)
	}

	delete(decoded, "activityStatistics")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userAnalytics"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserAnalytics: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UserAnalytics{}

func (s *UserAnalytics) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Settings  *Settings `json:"settings,omitempty"`
		Id        *string   `json:"id,omitempty"`
		ODataId   *string   `json:"@odata.id,omitempty"`
		ODataType *string   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Settings = decoded.Settings
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UserAnalytics into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["activityStatistics"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ActivityStatistics into list []json.RawMessage: %+v", err)
		}

		output := make([]ActivityStatistics, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalActivityStatisticsImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ActivityStatistics' for 'UserAnalytics': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ActivityStatistics = &output
	}

	return nil
}
