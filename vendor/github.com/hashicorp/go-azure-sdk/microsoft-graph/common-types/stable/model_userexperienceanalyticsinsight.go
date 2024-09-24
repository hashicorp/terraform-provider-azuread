package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsInsight struct {
	// The unique identifier of the user experience analytics insight.
	InsightId nullable.Type[string] `json:"insightId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates severity of insights. Possible values are: None, Informational, Warning, Error.
	Severity *UserExperienceAnalyticsInsightSeverity `json:"severity,omitempty"`

	// The unique identifier of the user experience analytics metric.
	UserExperienceAnalyticsMetricId nullable.Type[string] `json:"userExperienceAnalyticsMetricId,omitempty"`

	// The value of the user experience analytics insight.
	Values *[]UserExperienceAnalyticsInsightValue `json:"values,omitempty"`
}

var _ json.Unmarshaler = &UserExperienceAnalyticsInsight{}

func (s *UserExperienceAnalyticsInsight) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		InsightId                       nullable.Type[string]                   `json:"insightId,omitempty"`
		ODataId                         *string                                 `json:"@odata.id,omitempty"`
		ODataType                       *string                                 `json:"@odata.type,omitempty"`
		Severity                        *UserExperienceAnalyticsInsightSeverity `json:"severity,omitempty"`
		UserExperienceAnalyticsMetricId nullable.Type[string]                   `json:"userExperienceAnalyticsMetricId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.InsightId = decoded.InsightId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Severity = decoded.Severity
	s.UserExperienceAnalyticsMetricId = decoded.UserExperienceAnalyticsMetricId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UserExperienceAnalyticsInsight into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["values"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Values into list []json.RawMessage: %+v", err)
		}

		output := make([]UserExperienceAnalyticsInsightValue, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUserExperienceAnalyticsInsightValueImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Values' for 'UserExperienceAnalyticsInsight': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Values = &output
	}

	return nil
}
