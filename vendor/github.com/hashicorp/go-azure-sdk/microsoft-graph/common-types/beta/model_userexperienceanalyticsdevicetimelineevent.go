package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsDeviceTimelineEvent{}

type UserExperienceAnalyticsDeviceTimelineEvent struct {
	// The id of the device where the event occurred.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The time the event occured.
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// The details provided by the event, format depends on event type.
	EventDetails nullable.Type[string] `json:"eventDetails,omitempty"`

	// Indicates device event level. Possible values are: None, Verbose, Information, Warning, Error, Critical
	EventLevel *DeviceEventLevel `json:"eventLevel,omitempty"`

	// The name of the event. Examples include: BootEvent, LogonEvent, AppCrashEvent, AppHangEvent.
	EventName nullable.Type[string] `json:"eventName,omitempty"`

	// The source of the event. Examples include: Intune, Sccm.
	EventSource nullable.Type[string] `json:"eventSource,omitempty"`

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

func (s UserExperienceAnalyticsDeviceTimelineEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsDeviceTimelineEvent{}

func (s UserExperienceAnalyticsDeviceTimelineEvent) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsDeviceTimelineEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsDeviceTimelineEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsDeviceTimelineEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsDeviceTimelineEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsDeviceTimelineEvent: %+v", err)
	}

	return encoded, nil
}
