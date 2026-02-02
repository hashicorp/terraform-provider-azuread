package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewHistoryDefinition{}

type AccessReviewHistoryDefinition struct {
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`

	// Timestamp when the access review definition was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Determines which review decisions will be included in the fetched review history data if specified. Optional on
	// create. All decisions are included by default if no decisions are provided on create. Possible values are: approve,
	// deny, dontKnow, notReviewed, and notNotified.
	Decisions *[]AccessReviewHistoryDecisionFilter `json:"decisions,omitempty"`

	// Name for the access review history data collection. Required.
	DisplayName string `json:"displayName"`

	DownloadUri       nullable.Type[string] `json:"downloadUri,omitempty"`
	FulfilledDateTime nullable.Type[string] `json:"fulfilledDateTime,omitempty"`

	// If the accessReviewHistoryDefinition is a recurring definition, instances represent each recurrence. A definition
	// that doesn't recur will have exactly one instance.
	Instances *[]AccessReviewHistoryInstance `json:"instances,omitempty"`

	// A timestamp. Reviews ending on or before this date will be included in the fetched history data. Only required if
	// scheduleSettings is not defined.
	ReviewHistoryPeriodEndDateTime *string `json:"reviewHistoryPeriodEndDateTime,omitempty"`

	// A timestamp. Reviews starting on or before this date will be included in the fetched history data. Only required if
	// scheduleSettings isn't defined.
	ReviewHistoryPeriodStartDateTime *string `json:"reviewHistoryPeriodStartDateTime,omitempty"`

	// The settings for a recurring access review history definition series. Only required if
	// reviewHistoryPeriodStartDateTime or reviewHistoryPeriodEndDateTime aren't defined. Not supported yet.
	ScheduleSettings *AccessReviewHistoryScheduleSettings `json:"scheduleSettings,omitempty"`

	// Used to scope what reviews are included in the fetched history data. Fetches reviews whose scope matches with this
	// provided scope. Required.
	Scopes []AccessReviewScope `json:"scopes"`

	// Represents the status of the review history data collection. The possible values are: done, inProgress, error,
	// requested, unknownFutureValue.
	Status *AccessReviewHistoryStatus `json:"status,omitempty"`

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

func (s AccessReviewHistoryDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewHistoryDefinition{}

func (s AccessReviewHistoryDefinition) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewHistoryDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewHistoryDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewHistoryDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewHistoryDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewHistoryDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessReviewHistoryDefinition{}

func (s *AccessReviewHistoryDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime                  *string                              `json:"createdDateTime,omitempty"`
		Decisions                        *[]AccessReviewHistoryDecisionFilter `json:"decisions,omitempty"`
		DisplayName                      string                               `json:"displayName"`
		DownloadUri                      nullable.Type[string]                `json:"downloadUri,omitempty"`
		FulfilledDateTime                nullable.Type[string]                `json:"fulfilledDateTime,omitempty"`
		Instances                        *[]AccessReviewHistoryInstance       `json:"instances,omitempty"`
		ReviewHistoryPeriodEndDateTime   *string                              `json:"reviewHistoryPeriodEndDateTime,omitempty"`
		ReviewHistoryPeriodStartDateTime *string                              `json:"reviewHistoryPeriodStartDateTime,omitempty"`
		ScheduleSettings                 *AccessReviewHistoryScheduleSettings `json:"scheduleSettings,omitempty"`
		Status                           *AccessReviewHistoryStatus           `json:"status,omitempty"`
		Id                               *string                              `json:"id,omitempty"`
		ODataId                          *string                              `json:"@odata.id,omitempty"`
		ODataType                        *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Decisions = decoded.Decisions
	s.DisplayName = decoded.DisplayName
	s.DownloadUri = decoded.DownloadUri
	s.FulfilledDateTime = decoded.FulfilledDateTime
	s.Instances = decoded.Instances
	s.ReviewHistoryPeriodEndDateTime = decoded.ReviewHistoryPeriodEndDateTime
	s.ReviewHistoryPeriodStartDateTime = decoded.ReviewHistoryPeriodStartDateTime
	s.ScheduleSettings = decoded.ScheduleSettings
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewHistoryDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'AccessReviewHistoryDefinition': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["scopes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Scopes into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessReviewScope, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessReviewScopeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Scopes' for 'AccessReviewHistoryDefinition': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Scopes = output
	}

	return nil
}
