package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAnomalyCorrelationGroupOverview{}

type UserExperienceAnalyticsAnomalyCorrelationGroupOverview struct {
	// Indicates the number of correlation groups in the anomaly. Valid values -2147483648 to 2147483647
	AnomalyCorrelationGroupCount *int64 `json:"anomalyCorrelationGroupCount,omitempty"`

	// The unique identifier of the anomaly. Anomaly details such as name and type can be found in the
	// UserExperienceAnalyticsAnomalySeverityOverview entity.
	AnomalyId nullable.Type[string] `json:"anomalyId,omitempty"`

	// Indicates the total number of devices affected by the anomaly in the correlation group. Valid values -2147483648 to
	// 2147483647
	CorrelationGroupAnomalousDeviceCount *int64 `json:"correlationGroupAnomalousDeviceCount,omitempty"`

	// Indicates the total number of devices at risk in the correlation group. Valid values -2147483648 to 2147483647
	CorrelationGroupAtRiskDeviceCount *int64 `json:"correlationGroupAtRiskDeviceCount,omitempty"`

	// Indicates the total number of devices in a correlation group. Valid values -2147483648 to 2147483647
	CorrelationGroupDeviceCount *int64 `json:"correlationGroupDeviceCount,omitempty"`

	// Describes the features of a device that are shared between all devices in a correlation group.
	CorrelationGroupFeatures *[]UserExperienceAnalyticsAnomalyCorrelationGroupFeature `json:"correlationGroupFeatures,omitempty"`

	// The unique identifier for the correlation group which will uniquely identify one of the correlation group within an
	// anomaly. The correlation Id can be mapped to the correlation group name by concatinating the correlation group
	// features. Example of correlation group name which is the indicative of concatenated features names are for names,
	// Contoso manufacture 4.4.1 and Windows 11.22621.1485.
	CorrelationGroupId nullable.Type[string] `json:"correlationGroupId,omitempty"`

	// Indicates the level of prevalence of the correlation group features in the anomaly. Possible values are: high, medium
	// or low
	CorrelationGroupPrevalence *UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence `json:"correlationGroupPrevalence,omitempty"`

	// Indicates the total number of devices in the tenant. Valid values -2147483648 to 2147483647
	TotalDeviceCount *int64 `json:"totalDeviceCount,omitempty"`

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

func (s UserExperienceAnalyticsAnomalyCorrelationGroupOverview) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAnomalyCorrelationGroupOverview{}

func (s UserExperienceAnalyticsAnomalyCorrelationGroupOverview) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAnomalyCorrelationGroupOverview
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAnomalyCorrelationGroupOverview: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAnomalyCorrelationGroupOverview: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAnomalyCorrelationGroupOverview"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAnomalyCorrelationGroupOverview: %+v", err)
	}

	return encoded, nil
}
