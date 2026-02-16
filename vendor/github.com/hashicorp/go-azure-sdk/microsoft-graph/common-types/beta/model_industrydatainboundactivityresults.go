package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataIndustryDataActivityStatistics = IndustryDataInboundActivityResults{}

type IndustryDataInboundActivityResults struct {
	// Number of errors encountered while processing the inbound flow.
	Errors nullable.Type[int64] `json:"errors,omitempty"`

	// Counts of active and inactive groups processed by the inbound flow.
	Groups *IndustryDataIndustryDataRunEntityCountMetric `json:"groups,omitempty"`

	// Number of people matched to a Microsoft Entra user, by role.
	MatchedPeopleByRole *[]IndustryDataIndustryDataRunRoleCountMetric `json:"matchedPeopleByRole,omitempty"`

	// Counts of active and inactive memberships processed by the inbound flow.
	Memberships *IndustryDataIndustryDataRunEntityCountMetric `json:"memberships,omitempty"`

	// Counts of active and inactive organizations processed by the inbound flow.
	Organizations *IndustryDataIndustryDataRunEntityCountMetric `json:"organizations,omitempty"`

	// Counts of active and inactive people processed by the inbound flow.
	People *IndustryDataIndustryDataRunEntityCountMetric `json:"people,omitempty"`

	// Number of people not matched to a Microsoft Entra user, by role.
	UnmatchedPeopleByRole *[]IndustryDataIndustryDataRunRoleCountMetric `json:"unmatchedPeopleByRole,omitempty"`

	// Number of warnings encountered while processing the inbound flow.
	Warnings nullable.Type[int64] `json:"warnings,omitempty"`

	// Fields inherited from IndustryDataIndustryDataActivityStatistics

	// The identifier for the activity that is being reported on.
	ActivityId *string `json:"activityId,omitempty"`

	// The display name of the underlying flow.
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Status *IndustryDataIndustryDataActivityStatus `json:"status,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IndustryDataInboundActivityResults) IndustryDataIndustryDataActivityStatistics() BaseIndustryDataIndustryDataActivityStatisticsImpl {
	return BaseIndustryDataIndustryDataActivityStatisticsImpl{
		ActivityId:  s.ActivityId,
		DisplayName: s.DisplayName,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		Status:      s.Status,
	}
}

var _ json.Marshaler = IndustryDataInboundActivityResults{}

func (s IndustryDataInboundActivityResults) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataInboundActivityResults
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataInboundActivityResults: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataInboundActivityResults: %+v", err)
	}

	delete(decoded, "errors")
	delete(decoded, "groups")
	delete(decoded, "matchedPeopleByRole")
	delete(decoded, "memberships")
	delete(decoded, "organizations")
	delete(decoded, "people")
	delete(decoded, "unmatchedPeopleByRole")
	delete(decoded, "warnings")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.inboundActivityResults"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataInboundActivityResults: %+v", err)
	}

	return encoded, nil
}
