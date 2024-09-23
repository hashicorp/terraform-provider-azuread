package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityRetentionEvent{}

type SecurityRetentionEvent struct {
	// The user who created the retentionEvent.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date time when the retentionEvent was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Optional information about the event.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the event.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Represents the success status of a created event and additional information.
	EventPropagationResults *[]SecurityEventPropagationResult `json:"eventPropagationResults,omitempty"`

	// Represents the workload (SharePoint Online, OneDrive for Business, Exchange Online) and identification information
	// associated with a retention event.
	EventQueries *[]SecurityEventQuery `json:"eventQueries,omitempty"`

	// Status of event propogation to the scoped locations after the event has been created.
	EventStatus *SecurityRetentionEventStatus `json:"eventStatus,omitempty"`

	// Optional time when the event should be triggered.
	EventTriggerDateTime nullable.Type[string] `json:"eventTriggerDateTime,omitempty"`

	// The user who last modified the retentionEvent.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The latest date time when the retentionEvent was modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Last time the status of the event was updated.
	LastStatusUpdateDateTime nullable.Type[string] `json:"lastStatusUpdateDateTime,omitempty"`

	// Specifies the event that will start the retention period for labels that use this event type when an event is
	// created.
	RetentionEventType *SecurityRetentionEventType `json:"retentionEventType,omitempty"`

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

func (s SecurityRetentionEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityRetentionEvent{}

func (s SecurityRetentionEvent) MarshalJSON() ([]byte, error) {
	type wrapper SecurityRetentionEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityRetentionEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityRetentionEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.retentionEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityRetentionEvent: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityRetentionEvent{}

func (s *SecurityRetentionEvent) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime          nullable.Type[string]             `json:"createdDateTime,omitempty"`
		Description              nullable.Type[string]             `json:"description,omitempty"`
		DisplayName              nullable.Type[string]             `json:"displayName,omitempty"`
		EventPropagationResults  *[]SecurityEventPropagationResult `json:"eventPropagationResults,omitempty"`
		EventQueries             *[]SecurityEventQuery             `json:"eventQueries,omitempty"`
		EventStatus              *SecurityRetentionEventStatus     `json:"eventStatus,omitempty"`
		EventTriggerDateTime     nullable.Type[string]             `json:"eventTriggerDateTime,omitempty"`
		LastModifiedDateTime     nullable.Type[string]             `json:"lastModifiedDateTime,omitempty"`
		LastStatusUpdateDateTime nullable.Type[string]             `json:"lastStatusUpdateDateTime,omitempty"`
		RetentionEventType       *SecurityRetentionEventType       `json:"retentionEventType,omitempty"`
		Id                       *string                           `json:"id,omitempty"`
		ODataId                  *string                           `json:"@odata.id,omitempty"`
		ODataType                *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.EventPropagationResults = decoded.EventPropagationResults
	s.EventQueries = decoded.EventQueries
	s.EventStatus = decoded.EventStatus
	s.EventTriggerDateTime = decoded.EventTriggerDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LastStatusUpdateDateTime = decoded.LastStatusUpdateDateTime
	s.RetentionEventType = decoded.RetentionEventType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityRetentionEvent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityRetentionEvent': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SecurityRetentionEvent': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
