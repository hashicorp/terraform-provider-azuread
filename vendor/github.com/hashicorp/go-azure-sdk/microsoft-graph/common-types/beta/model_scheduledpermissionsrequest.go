package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ScheduledPermissionsRequest{}

type ScheduledPermissionsRequest struct {
	Action *UnifiedRoleScheduleRequestActions `json:"action,omitempty"`

	// Defines when the identity created the request.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The identity's justification for the request.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// Additional context for the permissions request.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	RequestedPermissions PermissionsDefinition `json:"requestedPermissions"`

	// When to assign the requested permissions.
	ScheduleInfo *RequestSchedule `json:"scheduleInfo,omitempty"`

	StatusDetail *StatusDetail `json:"statusDetail,omitempty"`

	// Ticketing-related metadata that you can use to correlate to the request.
	TicketInfo *TicketInfo `json:"ticketInfo,omitempty"`

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

func (s ScheduledPermissionsRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ScheduledPermissionsRequest{}

func (s ScheduledPermissionsRequest) MarshalJSON() ([]byte, error) {
	type wrapper ScheduledPermissionsRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScheduledPermissionsRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScheduledPermissionsRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.scheduledPermissionsRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScheduledPermissionsRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ScheduledPermissionsRequest{}

func (s *ScheduledPermissionsRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action          *UnifiedRoleScheduleRequestActions `json:"action,omitempty"`
		CreatedDateTime *string                            `json:"createdDateTime,omitempty"`
		Justification   nullable.Type[string]              `json:"justification,omitempty"`
		Notes           nullable.Type[string]              `json:"notes,omitempty"`
		ScheduleInfo    *RequestSchedule                   `json:"scheduleInfo,omitempty"`
		StatusDetail    *StatusDetail                      `json:"statusDetail,omitempty"`
		TicketInfo      *TicketInfo                        `json:"ticketInfo,omitempty"`
		Id              *string                            `json:"id,omitempty"`
		ODataId         *string                            `json:"@odata.id,omitempty"`
		ODataType       *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Justification = decoded.Justification
	s.Notes = decoded.Notes
	s.ScheduleInfo = decoded.ScheduleInfo
	s.StatusDetail = decoded.StatusDetail
	s.TicketInfo = decoded.TicketInfo
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ScheduledPermissionsRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["requestedPermissions"]; ok {
		impl, err := UnmarshalPermissionsDefinitionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RequestedPermissions' for 'ScheduledPermissionsRequest': %+v", err)
		}
		s.RequestedPermissions = impl
	}

	return nil
}
