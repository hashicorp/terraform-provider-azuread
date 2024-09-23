package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MonthlyUserInsightMetricsRoot{}

type MonthlyUserInsightMetricsRoot struct {
	// Insights for active users on apps registered in the tenant for a specified period.
	ActiveUsers *[]ActiveUsersMetric `json:"activeUsers,omitempty"`

	// Insights for authentications on apps registered in the tenant for a specified period.
	Authentications *[]AuthenticationsMetric `json:"authentications,omitempty"`

	InactiveUsers              *[]MonthlyInactiveUsersMetric              `json:"inactiveUsers,omitempty"`
	InactiveUsersByApplication *[]MonthlyInactiveUsersByApplicationMetric `json:"inactiveUsersByApplication,omitempty"`

	// Insights for MFA usage on apps registered in the tenant for a specified period.
	MfaCompletions *[]MfaCompletionMetric `json:"mfaCompletions,omitempty"`

	// Insights for all user requests on apps registered in the tenant for a specified period.
	Requests *[]UserRequestsMetric `json:"requests,omitempty"`

	// Total sign-ups on apps registered in the tenant for a specified period.
	SignUps *[]UserSignUpMetric `json:"signUps,omitempty"`

	// Summary of all usage insights on apps registered in the tenant for a specified period.
	Summary *[]InsightSummary `json:"summary,omitempty"`

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

func (s MonthlyUserInsightMetricsRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MonthlyUserInsightMetricsRoot{}

func (s MonthlyUserInsightMetricsRoot) MarshalJSON() ([]byte, error) {
	type wrapper MonthlyUserInsightMetricsRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MonthlyUserInsightMetricsRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MonthlyUserInsightMetricsRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.monthlyUserInsightMetricsRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MonthlyUserInsightMetricsRoot: %+v", err)
	}

	return encoded, nil
}
