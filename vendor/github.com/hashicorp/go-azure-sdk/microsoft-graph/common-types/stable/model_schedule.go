package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Schedule{}

type Schedule struct {
	// The day notes in the schedule.
	DayNotes *[]DayNote `json:"dayNotes,omitempty"`

	// Indicates whether the schedule is enabled for the team. Required.
	Enabled nullable.Type[bool] `json:"enabled,omitempty"`

	// Indicates whether copied shifts include activities from the original shift.
	IsActivitiesIncludedWhenCopyingShiftsEnabled nullable.Type[bool] `json:"isActivitiesIncludedWhenCopyingShiftsEnabled,omitempty"`

	// The offer requests for shifts in the schedule.
	OfferShiftRequests *[]OfferShiftRequest `json:"offerShiftRequests,omitempty"`

	// Indicates whether offer shift requests are enabled for the schedule.
	OfferShiftRequestsEnabled nullable.Type[bool] `json:"offerShiftRequestsEnabled,omitempty"`

	// The open shift requests in the schedule.
	OpenShiftChangeRequests *[]OpenShiftChangeRequest `json:"openShiftChangeRequests,omitempty"`

	// The set of open shifts in a scheduling group in the schedule.
	OpenShifts *[]OpenShift `json:"openShifts,omitempty"`

	// Indicates whether open shifts are enabled for the schedule.
	OpenShiftsEnabled nullable.Type[bool] `json:"openShiftsEnabled,omitempty"`

	// The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
	ProvisionStatus *OperationStatus `json:"provisionStatus,omitempty"`

	// Additional information about why schedule provisioning failed.
	ProvisionStatusCode nullable.Type[string] `json:"provisionStatusCode,omitempty"`

	// The logical grouping of users in the schedule (usually by role).
	SchedulingGroups *[]SchedulingGroup `json:"schedulingGroups,omitempty"`

	// The shifts in the schedule.
	Shifts *[]Shift `json:"shifts,omitempty"`

	// Indicates the start day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday,
	// saturday.
	StartDayOfWeek *DayOfWeek `json:"startDayOfWeek,omitempty"`

	// The swap requests for shifts in the schedule.
	SwapShiftsChangeRequests *[]SwapShiftsChangeRequest `json:"swapShiftsChangeRequests,omitempty"`

	// Indicates whether swap shifts requests are enabled for the schedule.
	SwapShiftsRequestsEnabled nullable.Type[bool] `json:"swapShiftsRequestsEnabled,omitempty"`

	// The time cards in the schedule.
	TimeCards *[]TimeCard `json:"timeCards,omitempty"`

	// Indicates whether time clock is enabled for the schedule.
	TimeClockEnabled nullable.Type[bool] `json:"timeClockEnabled,omitempty"`

	// The time clock location settings for this schedule.
	TimeClockSettings *TimeClockSettings `json:"timeClockSettings,omitempty"`

	// The set of reasons for a time off in the schedule.
	TimeOffReasons *[]TimeOffReason `json:"timeOffReasons,omitempty"`

	// The time off requests in the schedule.
	TimeOffRequests *[]TimeOffRequest `json:"timeOffRequests,omitempty"`

	// Indicates whether time off requests are enabled for the schedule.
	TimeOffRequestsEnabled nullable.Type[bool] `json:"timeOffRequestsEnabled,omitempty"`

	// Indicates the time zone of the schedule team using tz database format. Required.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`

	// The instances of times off in the schedule.
	TimesOff *[]TimeOff `json:"timesOff,omitempty"`

	// The IDs for the workforce integrations associated with this schedule.
	WorkforceIntegrationIds *[]string `json:"workforceIntegrationIds,omitempty"`

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

func (s Schedule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Schedule{}

func (s Schedule) MarshalJSON() ([]byte, error) {
	type wrapper Schedule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Schedule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Schedule: %+v", err)
	}

	delete(decoded, "provisionStatus")
	delete(decoded, "provisionStatusCode")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.schedule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Schedule: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Schedule{}

func (s *Schedule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DayNotes                                     *[]DayNote                 `json:"dayNotes,omitempty"`
		Enabled                                      nullable.Type[bool]        `json:"enabled,omitempty"`
		IsActivitiesIncludedWhenCopyingShiftsEnabled nullable.Type[bool]        `json:"isActivitiesIncludedWhenCopyingShiftsEnabled,omitempty"`
		OfferShiftRequestsEnabled                    nullable.Type[bool]        `json:"offerShiftRequestsEnabled,omitempty"`
		OpenShiftChangeRequests                      *[]OpenShiftChangeRequest  `json:"openShiftChangeRequests,omitempty"`
		OpenShifts                                   *[]OpenShift               `json:"openShifts,omitempty"`
		OpenShiftsEnabled                            nullable.Type[bool]        `json:"openShiftsEnabled,omitempty"`
		ProvisionStatus                              *OperationStatus           `json:"provisionStatus,omitempty"`
		ProvisionStatusCode                          nullable.Type[string]      `json:"provisionStatusCode,omitempty"`
		SchedulingGroups                             *[]SchedulingGroup         `json:"schedulingGroups,omitempty"`
		Shifts                                       *[]Shift                   `json:"shifts,omitempty"`
		StartDayOfWeek                               *DayOfWeek                 `json:"startDayOfWeek,omitempty"`
		SwapShiftsChangeRequests                     *[]SwapShiftsChangeRequest `json:"swapShiftsChangeRequests,omitempty"`
		SwapShiftsRequestsEnabled                    nullable.Type[bool]        `json:"swapShiftsRequestsEnabled,omitempty"`
		TimeCards                                    *[]TimeCard                `json:"timeCards,omitempty"`
		TimeClockEnabled                             nullable.Type[bool]        `json:"timeClockEnabled,omitempty"`
		TimeClockSettings                            *TimeClockSettings         `json:"timeClockSettings,omitempty"`
		TimeOffReasons                               *[]TimeOffReason           `json:"timeOffReasons,omitempty"`
		TimeOffRequests                              *[]TimeOffRequest          `json:"timeOffRequests,omitempty"`
		TimeOffRequestsEnabled                       nullable.Type[bool]        `json:"timeOffRequestsEnabled,omitempty"`
		TimeZone                                     nullable.Type[string]      `json:"timeZone,omitempty"`
		TimesOff                                     *[]TimeOff                 `json:"timesOff,omitempty"`
		WorkforceIntegrationIds                      *[]string                  `json:"workforceIntegrationIds,omitempty"`
		Id                                           *string                    `json:"id,omitempty"`
		ODataId                                      *string                    `json:"@odata.id,omitempty"`
		ODataType                                    *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DayNotes = decoded.DayNotes
	s.Enabled = decoded.Enabled
	s.IsActivitiesIncludedWhenCopyingShiftsEnabled = decoded.IsActivitiesIncludedWhenCopyingShiftsEnabled
	s.OfferShiftRequestsEnabled = decoded.OfferShiftRequestsEnabled
	s.OpenShiftChangeRequests = decoded.OpenShiftChangeRequests
	s.OpenShifts = decoded.OpenShifts
	s.OpenShiftsEnabled = decoded.OpenShiftsEnabled
	s.ProvisionStatus = decoded.ProvisionStatus
	s.ProvisionStatusCode = decoded.ProvisionStatusCode
	s.SchedulingGroups = decoded.SchedulingGroups
	s.Shifts = decoded.Shifts
	s.StartDayOfWeek = decoded.StartDayOfWeek
	s.SwapShiftsChangeRequests = decoded.SwapShiftsChangeRequests
	s.SwapShiftsRequestsEnabled = decoded.SwapShiftsRequestsEnabled
	s.TimeCards = decoded.TimeCards
	s.TimeClockEnabled = decoded.TimeClockEnabled
	s.TimeClockSettings = decoded.TimeClockSettings
	s.TimeOffReasons = decoded.TimeOffReasons
	s.TimeOffRequests = decoded.TimeOffRequests
	s.TimeOffRequestsEnabled = decoded.TimeOffRequestsEnabled
	s.TimeZone = decoded.TimeZone
	s.TimesOff = decoded.TimesOff
	s.WorkforceIntegrationIds = decoded.WorkforceIntegrationIds
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Schedule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["offerShiftRequests"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling OfferShiftRequests into list []json.RawMessage: %+v", err)
		}

		output := make([]OfferShiftRequest, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalOfferShiftRequestImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'OfferShiftRequests' for 'Schedule': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.OfferShiftRequests = &output
	}

	return nil
}
