package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DailyUserInsightMetricsRoot{}

type DailyUserInsightMetricsRoot struct {
	// Insights for active users on apps registered in the tenant for a specified period.
	ActiveUsers *[]ActiveUsersMetric `json:"activeUsers,omitempty"`

	// Insights for authentications on apps registered in the tenant for a specified period.
	Authentications *[]AuthenticationsMetric `json:"authentications,omitempty"`

	InactiveUsers              *[]DailyInactiveUsersMetric              `json:"inactiveUsers,omitempty"`
	InactiveUsersByApplication *[]DailyInactiveUsersByApplicationMetric `json:"inactiveUsersByApplication,omitempty"`

	// Insights for MFA usage on apps registered in the tenant for a specified period.
	MfaCompletions *[]MfaCompletionMetric `json:"mfaCompletions,omitempty"`

	MfaTelecomFraud *[]MfaTelecomFraudMetric `json:"mfaTelecomFraud,omitempty"`

	// Total sign-ups on apps registered in the tenant for a specified period.
	SignUps *[]UserSignUpMetric `json:"signUps,omitempty"`

	// Summary of all usage insights on apps registered in the tenant for a specified period.
	Summary *[]InsightSummary `json:"summary,omitempty"`

	// Insights for total users on apps registered in the tenant for a specified period.
	UserCount *[]UserCountMetric `json:"userCount,omitempty"`

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

func (s DailyUserInsightMetricsRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DailyUserInsightMetricsRoot{}

func (s DailyUserInsightMetricsRoot) MarshalJSON() ([]byte, error) {
	type wrapper DailyUserInsightMetricsRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DailyUserInsightMetricsRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DailyUserInsightMetricsRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.dailyUserInsightMetricsRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DailyUserInsightMetricsRoot: %+v", err)
	}

	return encoded, nil
}
