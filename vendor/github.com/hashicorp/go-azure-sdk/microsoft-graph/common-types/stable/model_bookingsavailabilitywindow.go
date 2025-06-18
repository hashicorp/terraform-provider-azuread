package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BookingsAvailability = BookingsAvailabilityWindow{}

type BookingsAvailabilityWindow struct {
	// End date of the availability window.
	EndDate nullable.Type[string] `json:"endDate,omitempty"`

	// Start date of the availability window.
	StartDate nullable.Type[string] `json:"startDate,omitempty"`

	// Fields inherited from BookingsAvailability

	AvailabilityType *BookingsServiceAvailabilityType `json:"availabilityType,omitempty"`

	// The hours of operation in a week. The business hours value is set to null if the availability type isn't
	// customWeeklyHours.
	BusinessHours *[]BookingWorkHours `json:"businessHours,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BookingsAvailabilityWindow) BookingsAvailability() BaseBookingsAvailabilityImpl {
	return BaseBookingsAvailabilityImpl{
		AvailabilityType: s.AvailabilityType,
		BusinessHours:    s.BusinessHours,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

var _ json.Marshaler = BookingsAvailabilityWindow{}

func (s BookingsAvailabilityWindow) MarshalJSON() ([]byte, error) {
	type wrapper BookingsAvailabilityWindow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookingsAvailabilityWindow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingsAvailabilityWindow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingsAvailabilityWindow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookingsAvailabilityWindow: %+v", err)
	}

	return encoded, nil
}
