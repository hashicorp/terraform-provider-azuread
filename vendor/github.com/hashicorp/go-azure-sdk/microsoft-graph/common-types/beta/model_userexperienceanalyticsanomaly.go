package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAnomaly{}

type UserExperienceAnalyticsAnomaly struct {
	// Indicates the first occurrence date and time for the anomaly.
	AnomalyFirstOccurrenceDateTime *string `json:"anomalyFirstOccurrenceDateTime,omitempty"`

	// The unique identifier of the anomaly.
	AnomalyId nullable.Type[string] `json:"anomalyId,omitempty"`

	// Indicates the latest occurrence date and time for the anomaly.
	AnomalyLatestOccurrenceDateTime *string `json:"anomalyLatestOccurrenceDateTime,omitempty"`

	// The name of the anomaly.
	AnomalyName nullable.Type[string] `json:"anomalyName,omitempty"`

	// Indicates the category of the anomaly. Eg: anomaly type can be device, application, stop error, driver or other.
	AnomalyType *UserExperienceAnalyticsAnomalyType `json:"anomalyType,omitempty"`

	// The name of the application or module that caused the anomaly.
	AssetName nullable.Type[string] `json:"assetName,omitempty"`

	// The publisher of the application or module that caused the anomaly.
	AssetPublisher nullable.Type[string] `json:"assetPublisher,omitempty"`

	// The version of the application or module that caused the anomaly.
	AssetVersion nullable.Type[string] `json:"assetVersion,omitempty"`

	// The unique identifier of the anomaly detection model.
	DetectionModelId nullable.Type[string] `json:"detectionModelId,omitempty"`

	// The number of devices impacted by the anomaly. Valid values -2147483648 to 2147483647
	DeviceImpactedCount *int64 `json:"deviceImpactedCount,omitempty"`

	// The unique identifier of the anomaly detection model.
	IssueId nullable.Type[string] `json:"issueId,omitempty"`

	// Indicates the severity of the anomaly. Eg: anomaly severity can be high, medium, low, informational or other.
	Severity *UserExperienceAnalyticsAnomalySeverity `json:"severity,omitempty"`

	// Indicates the state of the anomaly. Eg: anomaly severity can be new, active, disabled, removed or other.
	State *UserExperienceAnalyticsAnomalyState `json:"state,omitempty"`

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

func (s UserExperienceAnalyticsAnomaly) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAnomaly{}

func (s UserExperienceAnalyticsAnomaly) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAnomaly
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAnomaly: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAnomaly: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAnomaly"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAnomaly: %+v", err)
	}

	return encoded, nil
}
